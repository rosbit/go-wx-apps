package main

import (
	"github.com/rosbit/go-wx-api/v2/tools"
	"github.com/rosbit/go-wx-api/v2"
	"os"
	"fmt"
	"encoding/json"
)

const (
	service = "test"
)

type ParamsT struct {
	Token      string
	AppId      string
	Secret     string
	TokenCache string `json:"TokenCachePath"`
	MenuJson   string `json:"MenuJsonFile"`
}

func usage() {
	fmt.Printf("Usage: %s <command> <params-json-file>\n", os.Args[0])
	fmt.Printf("Where <command> could be:\n")
	fmt.Printf("  create  ---  create wx menu\n")
	fmt.Printf("  query   ---  query wx menu\n")
	fmt.Printf("  delete  ---  delete wx menu\n")
	fmt.Printf("  info    ---  current information of self-menu\n")
}

func createMenu(paramsFile string) {
	params, err := loadConf(paramsFile);
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	if err = wxtools.CreateMenu(service, params.MenuJson); err != nil {
		fmt.Printf("failed to create menu: %v\n", err)
		return
	}
	fmt.Printf("OK\n")
}

func queryMenu(paramsFile string) {
	if _, err := loadConf(paramsFile); err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.QueryMenu(service)
	if err != nil {
		fmt.Printf("failed to query menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %v\n", resp)
}

func deleteMenu(paramsFile string) {
	if _, err := loadConf(paramsFile); err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	if err := wxtools.DeleteMenu(service); err != nil {
		fmt.Printf("failed to delete menu: %v\n", err)
		return
	}
	fmt.Printf("OK\n")
}

func infoMenu(paramsFile string) {
	if _, err := loadConf(paramsFile); err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.CurrentSelfmenuInfo(service)
	if err != nil {
		fmt.Printf("failed to get information of current self-menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %v\n", resp)
}

func loadConf(paramsFile string) (*ParamsT, error) {
	fp, err := os.Open(paramsFile)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var params ParamsT
	if err = json.NewDecoder(fp).Decode(&params); err != nil {
		return nil, err
	}

	wxapi.SetWxParams(service, params.Token, params.AppId, params.Secret, "")
	wxapi.InitWx(params.TokenCache)

	return &params, nil
}

var _commands = map[string]func(string) {
	"create": createMenu,
	"query":  queryMenu,
	"delete": deleteMenu,
	"info":   infoMenu,
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	command, paramsFile := os.Args[1], os.Args[2]

	fn, ok := _commands[command]
	if !ok {
		fmt.Printf("Uknown command: %s\n", command)
		return
	}

	fn(paramsFile)
}
