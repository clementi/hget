package main

import (
	"path/filepath"
	"testing"
)

func TestFilterIpV4(t *testing.T) {
}

// func TestFolderOfPanic1(t *testing.T) {
// 	url := "http://foo.bar/.."
// 	shouldPanic(t, func() { FolderOf(url) })
// }

func TestFolderOfPanic2(t *testing.T) {
	url := "http://foo.bar/../../../foobar"
	u := FolderOf(url)
	if filepath.Base(u) != "foobar" {
		t.Fatal("URL of return incorrect value")
	}
}

func TestFolderOfNormal(t *testing.T) {
	url := "http://foo.bar/file"
	u := FolderOf(url)
	if filepath.Base(u) != "file" {
		t.Fatal("URL of return incorrect value")
	}
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Error("should have panicked")
}
