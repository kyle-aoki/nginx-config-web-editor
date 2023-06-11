package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ############################################################################
// ############################################################################

const (
	EtcNginx                              = "/etc/nginx"
	EtcNginxNginxConf                     = "/etc/nginx/nginx.conf"
	EtcNginxNginxConfigWebEditor          = "/etc/nginx/nginx-config-web-editor"
	EtcNginxNginxConfigWebEditorNginxConf = "/etc/nginx/nginx-config-web-editor/default-nginx.conf"
)

// ############################################################################
// ############################################################################

func NginxInitialize() {
	log.Println("initializing nginx-config-web-editor")
	if !fsObjectExists(EtcNginx) {
		log.Println("nginx is not installed, could not find", EtcNginx)
		os.Exit(1)
	}
	if !fsObjectExists(EtcNginxNginxConfigWebEditor) {
		log.Println("creating", EtcNginxNginxConfigWebEditor)
		check(os.Mkdir(EtcNginxNginxConfigWebEditor, 0644))
	}
	des := must(os.ReadDir(EtcNginxNginxConfigWebEditor))
	if len(des) == 0 {
		if !fsObjectExists(EtcNginxNginxConf) {
			log.Fatalln("default nginx config file does not exist at", EtcNginxNginxConf)
			os.Exit(1)
		}
		b := must(os.ReadFile(EtcNginxNginxConf))
		check(os.WriteFile(EtcNginxNginxConfigWebEditorNginxConf, b, 0644))
	}
}

// ############################################################################
// ############################################################################

type NginxReloadInput struct{ Name string }
type NginxReloadOutput struct{ Result, Error string }

func NginxReload(nri *NginxReloadInput) *NginxReloadOutput {
	log.Println("reloading", EtcNginxNginxConf, "with", nri.Name)
	nro := NginxRead(&NginxReadInput{nri.Name})
	check(os.WriteFile(EtcNginxNginxConf, []byte(nro.Value), 0644))
	str, err := bash(`nginx -s reload`)
	if err != nil {
		return &NginxReloadOutput{str, err.Error()}
	}
	return &NginxReloadOutput{str, ""}
}

// ############################################################################
// ############################################################################

type NginxCloneInput struct{ Name string }
type NginxCloneOutput struct{ NewName string }

func NginxClone(nci *NginxCloneInput) *NginxCloneOutput {
	log.Println("cloning", nci.Name)
	nro := NginxRead(&NginxReadInput{nci.Name})
	newName := f("%s-copy", nci.Name)
	NginxSave(&NginxSaveInput{newName, nro.Value})
	return &NginxCloneOutput{newName}
}

// ############################################################################
// ############################################################################

type NginxRenameInput struct{ Name, NewName string }

func NginxRename(nri *NginxRenameInput) {
	log.Println("renaming", nri.Name, "to", nri.NewName)
	nro := NginxRead(&NginxReadInput{nri.Name})
	NginxSave(&NginxSaveInput{nri.NewName, nro.Value})
	NginxDelete(&NginxDeleteInput{nri.Name})
}

// ############################################################################
// ############################################################################

type NginxListOutput struct{ Files []string }

func NginxList() *NginxListOutput {
	log.Println("listing config files")
	var confFiles []string
	de := must(os.ReadDir(EtcNginxNginxConfigWebEditor))
	for i := 0; i < len(de); i++ {
		confFiles = append(confFiles, de[i].Name())
	}
	if len(confFiles) == 0 {
		NginxInitialize()
	}
	return &NginxListOutput{confFiles}
}

// ############################################################################
// ############################################################################

type NginxReadInput struct{ Name string }
type NginxReadOutput struct{ Value string }

func NginxRead(nri *NginxReadInput) *NginxReadOutput {
	log.Println("reading", nri.Name)
	path := filepath.Join(EtcNginxNginxConfigWebEditor, nri.Name)
	return &NginxReadOutput{string(must(os.ReadFile(path)))}
}

// ############################################################################
// ############################################################################

type NginxSaveInput struct{ Name, Value string }

func NginxSave(nsi *NginxSaveInput) {
	log.Println("saving", nsi.Name)
	check(os.WriteFile(filepath.Join(EtcNginxNginxConfigWebEditor, nsi.Name), []byte(nsi.Value), 0644))
}

// ############################################################################
// ############################################################################

type NginxDeleteInput struct{ Name string }

func NginxDelete(ndi *NginxDeleteInput) {
	log.Println("deleting", ndi.Name)
	check(os.Remove(filepath.Join(EtcNginxNginxConfigWebEditor, ndi.Name)))
}

// ############################################################################
// ############################################################################

type NginxTestInput struct{ Name string }
type NginxTestOutput struct{ Result, Error string }

func NginxTest(nri *NginxTestInput) *NginxTestOutput {
	log.Println("testing", nri.Name)
	path := filepath.Join(EtcNginxNginxConfigWebEditor, nri.Name)
	cmd := fmt.Sprintf(`nginx -t -c "%s"`, path)
	res, err := bash(cmd)
	if err != nil {
		return &NginxTestOutput{res, err.Error()}
	}
	return &NginxTestOutput{res, ""}
}
