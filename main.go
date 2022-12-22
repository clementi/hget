package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
)

var displayProgress = true

func main() {
	var err error

	cli.AppHelpTemplate = strings.Replace(cli.AppHelpTemplate, "[arguments...]", "[URL]", -1)

	app := cli.App{
		Name:  "hget",
		Usage: "Multipart resumable downloads",
		Action: func(ctx *cli.Context) error {
			if !ctx.Args().Present() {
				fmt.Println("URL required")
				fmt.Println()
				cli.ShowAppHelpAndExit(ctx, 1)
			}

			url := ctx.Args().First()

			Execute(url, nil, int(ctx.Uint("connections")), ctx.Bool("skip-tls"))
			return nil
		},
		Authors: []*cli.Author{
			{
				Name: "huydx (https://github.com/huydx)",
			},
			{
				Name: "clementi (https://github.com/clementi)",
			},
		},
		Version:         "2.0.0-beta1",
		HideHelpCommand: true,
		Commands: []*cli.Command{
			{
				Name:    "tasks",
				Aliases: []string{"t"},
				Usage:   "manage current tasks",
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"ls"},
						Usage:   "list tasks",
						Action: func(ctx *cli.Context) error {
							return TaskPrint()
						},
					},
					{
						Name:    "delete",
						Aliases: []string{"del", "d", "remove", "rm"},
						Usage:   "delete task",
						Action: func(ctx *cli.Context) error {
							if !ctx.Args().Present() {
								fmt.Println("task name required")
								os.Exit(2)
							}
							task := ctx.Args().First()
							if err := Delete(task); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:    "resume",
						Aliases: []string{"r"},
						Usage:   "resume task",
						Action: func(ctx *cli.Context) error {
							if !ctx.Args().Present() {
								fmt.Println("task name required")
								os.Exit(2)
							}
							task := ctx.Args().First()
							state, err := Resume(task)
							if err != nil {
								return err
							}
							Execute(state.Url, state, int(ctx.Uint("connections")), ctx.Bool("skip-tls"))
							return nil
						},
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:     "connections",
				Value:    4,
				Required: false,
				Usage:    "number of connections",
				Aliases:  []string{"n"},
			},
			&cli.BoolFlag{
				Name:     "skip-tls",
				Value:    false,
				Required: false,
				Usage:    "do not verify certificate for HTTPS",
				Aliases:  []string{"s"},
			},
		},
	}

	if err = app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Execute(url string, state *State, conn int, skiptls bool) {
	var err error

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// set up parallel

	var files = make([]string, 0)
	var parts = make([]Part, 0)
	var isInterrupted = false

	doneChan := make(chan bool, conn)
	fileChan := make(chan string, conn)
	errorChan := make(chan error, 1)
	stateChan := make(chan Part, 1)
	interruptChan := make(chan bool, conn)

	var downloader *HttpDownloader
	if state == nil {
		downloader = NewHttpDownloader(url, conn, skiptls)
	} else {
		downloader = &HttpDownloader{url: state.Url, file: filepath.Base(state.Url), par: int64(len(state.Parts)), parts: state.Parts, resumable: true}
	}
	go downloader.Do(doneChan, fileChan, errorChan, interruptChan, stateChan)

	for {
		select {
		case <-signal_chan:
			//send par number of interrupt for each routine
			isInterrupted = true
			for i := 0; i < conn; i++ {
				interruptChan <- true
			}
		case file := <-fileChan:
			files = append(files, file)
		case err := <-errorChan:
			log.Fatalf("%v", err)
			panic(err) //maybe need better style
		case part := <-stateChan:
			parts = append(parts, part)
		case <-doneChan:
			if isInterrupted {
				if downloader.resumable {
					log.Printf("Interrupted, saving state ... \n")
					s := &State{Url: url, Parts: parts}
					err := s.Save()
					if err != nil {
						log.Fatalf("%v\n", err)
					}
					return
				} else {
					log.Printf("Interrupted, but downloading url is not resumable, silently die")
					return
				}
			} else {
				err = JoinFile(files, filepath.Base(url))
				FatalCheck(err)
				err = os.RemoveAll(FolderOf(url))
				FatalCheck(err)
				return
			}
		}
	}
}
