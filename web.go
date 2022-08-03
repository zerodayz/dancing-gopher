// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
)

type podInfo struct {
	DisplayPodName string
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	podName, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	podNameBuf := bytes.NewBuffer(nil)
	podNameBuf.Write([]byte(podName))

	p := podInfo{}
	p.DisplayPodName = podNameBuf.String()

	t := template.Must(template.ParseFiles("dashboard.html"))

	err = t.ExecuteTemplate(w, "dashboard.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/dashboard", DashboardHandler)
	http.HandleFunc("/", RootHandler)
	log.Println("Starting dancing-gopher webserver at :8080")
	http.ListenAndServe(":8080", nil)
}
