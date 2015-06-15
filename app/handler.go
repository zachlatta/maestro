package app

import (
	"go/build"
	"log"
	"net/http"
	"os"
	"path/filepath"
	
	"github.com/gorilla/mux"
	"github.com/hackedu/maestro/router"
)

var (
	// StaticDir is the directory containing static assets.
	StaticDir = filepath.Join(defaultBase("github.com/hackedu/maestro/app"), "static")
)

func Handler() *mux.Router {
	m := router.App()
	m.Get(router.AppStatic).Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))
	return m
}

func defaultBase(path string) string {
	p, err := build.Default.Import(path, "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p.Dir, err = filepath.Rel(cwd, p.Dir)
	if err != nil {
		log.Fatal(err)
	}

	return p.Dir
}