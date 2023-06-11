package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ############################################################################
// ############################################################################

var idx = 1

func next() string {
	idx++
	return os.Args[idx]
}

// ############################################################################
// ############################################################################

var cliCmds = map[string]func(){
	"reload":     func() { fmt.Println(NginxReload(&NginxReloadInput{next()})) },
	"clone":      func() { NginxClone(&NginxCloneInput{next()}) },
	"rename":     func() { NginxRename(&NginxRenameInput{next(), next()}) },
	"list":       func() { fmt.Println(NginxList()) },
	"initialize": func() { NginxInitialize() },
	"read":       func() { fmt.Println(NginxRead(&NginxReadInput{next()})) },
	"save":       func() { NginxSave(&NginxSaveInput{next(), next()}) },
	"delete":     func() { NginxDelete(&NginxDeleteInput{next()}) },
}

// go run . cli rename a.conf b.conf
func CLI() { cliCmds[next()]() }

// ############################################################################
// ############################################################################

func rq(v any) {
	req := must(http.NewRequest("POST", host+path, bytes.NewReader(toJson(v))))
	req.Header.Add("session", session)
	resp := must(http.DefaultClient.Do(req))
	b := string(must(io.ReadAll(resp.Body)))
	if b == "" {
		fmt.Println(resp.Status)
	} else {
		fmt.Println(b)
	}
}

var httpCmds = map[string]func(){
	"/log-in":       func() { rq(&LogInInput{next(), next()}) },
	"/user/find":    func() { rq(nil) },
	"/nginx/reload": func() { rq(&NginxReloadInput{next()}) },
	"/nginx/clone":  func() { rq(&NginxCloneInput{next()}) },
	"/nginx/rename": func() { rq(&NginxRenameInput{next(), next()}) },
	"/nginx/list":   func() { rq(nil) },
	"/nginx/read":   func() { rq(&NginxReadInput{next()}) },
	"/nginx/save":   func() { rq(&NginxSaveInput{next(), next()}) },
	"/nginx/delete": func() { rq(&NginxDeleteInput{next()}) },
}

var (
	host    = ""
	session = ""
	path    = ""
)

// go run . http http://0.0.0.0:9040 LW41C5E1SC4P /nginx/rename a.conf b.conf
func HTTP() {
	host, session, path = next(), next(), next()
	httpCmds[path]()
}
