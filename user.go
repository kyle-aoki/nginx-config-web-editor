package main

import (
	"log"
	"net/http"
)

// ############################################################################
// ############################################################################

type User struct {
	Username string
	Password string
}

var UserSet = map[string]bool{}
var UserSessionSet = map[string]string{}

func FormatUser(user User) string {
	return f("%s:%s", user.Username, user.Password)
}

func LoadUsers() {
	for i := 0; i < len(CFG.Users); i++ {
		UserSet[FormatUser(CFG.Users[i])] = true
	}
}

// ############################################################################
// ############################################################################

type LogInInput struct{ Username, Password string }
type LogInOutput struct{ Session string }

func LogIn(w http.ResponseWriter, r *http.Request) {
	pl := fromJson[LogInInput](r)
	_, exists := UserSet[FormatUser(User{pl.Username, pl.Password})]
	if !exists {
		log.Println("found no users")
		w.WriteHeader(401)
		return
	}
	log.Println("found user")
	session := randomString(64)
	UserSessionSet[session] = pl.Username
	w.Write(toJson(LogInOutput{session}))
}

// ############################################################################
// ############################################################################

type FindUserInput struct{}
type FindUserOutput struct{ Username string }

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write(toJson(FindUserOutput{r.Header.Get("username")}))
}

// ############################################################################
// ############################################################################

type AuthMw struct {
	handler http.Handler
}

func (am AuthMw) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if CFG.DisableAuth {
		log.Println("skipping auth")
		r.Header.Add("username", "skipped")
		am.handler.ServeHTTP(w, r)
		return
	}
	session := r.Header.Get("session")
	username, found := UserSessionSet[session]
	if !found {
		w.WriteHeader(401)
		return
	}
	r.Header.Add("username", username)
	am.handler.ServeHTTP(w, r)
}

func ApplyAuthMiddleware(handler http.Handler) http.Handler {
	return AuthMw{handler}
}

// ############################################################################
// ############################################################################

type HttpMw struct {
	handler http.Handler
}

func (hm HttpMw) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			w.WriteHeader(500)
		}
	}()
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		return
	}
	log.Println(r.URL.Path)
	hm.handler.ServeHTTP(w, r)
}

func ApplyHttpMiddleware(handler http.Handler) http.Handler {
	return HttpMw{handler}
}
