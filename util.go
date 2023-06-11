package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	check(err)
	return t
}

func fromJson[T any](r *http.Request) *T {
	var t T
	json.Unmarshal(must(io.ReadAll(r.Body)), &t)
	return &t
}

func toJson(v any) []byte {
	return must(json.Marshal(v))
}

func randomString(length int) string {
	const pool = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var str []byte
	for i := 0; i < length; i++ {
		n := must(rand.Int(rand.Reader, big.NewInt(int64(len(pool)))))
		str = append(str, pool[n.Int64()])
	}
	return string(str)
}

func bash(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	stdout := must(cmd.StdoutPipe())
	stderr := must(cmd.StderrPipe())
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	sout := cmdText(stdout)
	serr := cmdText(stderr)
	if len(serr) != 0 {
		return sout, errors.New(serr)
	}
	return sout, nil
}

func cmdText(txt io.ReadCloser) string {
	if txt == nil {
		return ""
	}
	scanner := bufio.NewScanner(txt)
	var strs []string
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	return strings.Join(strs, "\n")
}

func f(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func fsObjectExists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	check(err)
	return true
}
