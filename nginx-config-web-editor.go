package main

import (
	"log"
	"net/http"
	"os"
)

// ############################################################################
// ############################################################################

func main() {
	if len(os.Args) > 1 && os.Args[1] == "cli" {
		CLI()
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "http" {
		HTTP()
		return
	}
	programConfigure()
	log.Println("nginx-config-web-editor server started")
	if !CFG.SkipInitialization {
		NginxInitialize()
	}
	server()
}

// ############################################################################
// ############################################################################

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello from ncwe")) })
	http.HandleFunc("/log-in", LogIn)

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("/find", FindUser)

	nginxRouter := createNginxRouter()

	http.DefaultServeMux.Handle("/user/", ApplyAuthMiddleware(http.StripPrefix("/user", userRouter)))
	http.DefaultServeMux.Handle("/nginx/", ApplyAuthMiddleware(http.StripPrefix("/nginx", nginxRouter)))

	check(http.ListenAndServe(f(":%s", CFG.Port), ApplyHttpMiddleware(http.DefaultServeMux)))
}

// ############################################################################
// ############################################################################

func createNginxRouter() *http.ServeMux {
	nginxRouter := http.NewServeMux()
	nginxRouter.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		w.Write(toJson(NginxReload(fromJson[NginxReloadInput](r))))
	})
	nginxRouter.HandleFunc("/clone", func(w http.ResponseWriter, r *http.Request) {
		w.Write(toJson(NginxClone(fromJson[NginxCloneInput](r))))
	})
	nginxRouter.HandleFunc("/rename", func(w http.ResponseWriter, r *http.Request) {
		NginxRename(fromJson[NginxRenameInput](r))
	})
	nginxRouter.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		w.Write(toJson(NginxList()))
	})
	nginxRouter.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		w.Write(toJson(NginxRead(fromJson[NginxReadInput](r))))
	})
	nginxRouter.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		NginxSave(fromJson[NginxSaveInput](r))
	})
	nginxRouter.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		NginxDelete(fromJson[NginxDeleteInput](r))
	})
	nginxRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write(toJson(NginxTest(fromJson[NginxTestInput](r))))
	})
	return nginxRouter
}
