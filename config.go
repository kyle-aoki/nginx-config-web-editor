package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ProgramConfiguration struct {
	Port               string
	Users              []User
	DisableAuth        bool
	SkipInitialization bool
}

var CFG = &ProgramConfiguration{}

func programConfigure() {
	if len(os.Args) == 1 {
		baseConfig := &ProgramConfiguration{
			Port: "9040",
			Users: []User{{
				Username: "admin", Password: randomString(32),
			}},
			SkipInitialization: false,
		}
		fmt.Println("# nginx-config-web-editor base configuration file")
		fmt.Println("# to run program: nginx-config-web-editor <path-to-config-file>")
		fmt.Print(string(must(yaml.Marshal(baseConfig))))
		os.Exit(0)
	}
	check(yaml.Unmarshal(must(os.ReadFile(os.Args[1])), CFG))
	LoadUsers()
}
