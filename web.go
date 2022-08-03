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

var (
	key       string = "/certs/tls.key"
	cert      string = "/certs/tls.crt"
	httpPort  string = ":8080"
	httpsPort string = ":8443"
)

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

	log.Println("Serving SSL Key:", key, "and SSL Cert:", cert)
	log.Println("Starting dancing-gopher server at", httpPort)
	go http.ListenAndServe(httpPort, nil)
	log.Println("Starting dancing-gopher server at", httpsPort)
	log.Fatal(http.ListenAndServeTLS(httpsPort, cert, key, nil))
}
