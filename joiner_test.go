package main

import (
	"os"
	"testing"
)

func TestJoiner(t *testing.T) {
	displayProgress = false

	prepare()

	files := [2]string{"file1", "file2"}
	JoinFile(files[:], "join")
	content, err := os.ReadFile("join")
	if err != nil {
		t.Fatalf("err should be nil")
	}
	if string(content) != "file1file2" {
		t.Fatalf("join content should be file1file2")
	}

	clean()
}

func prepare() {
	os.WriteFile("file1", []byte("file1"), 0600)
	os.WriteFile("file2", []byte("file2"), 0600)
}

func clean() {
	os.Remove("file1")
	os.Remove("file2")
	os.Remove("join")
}
