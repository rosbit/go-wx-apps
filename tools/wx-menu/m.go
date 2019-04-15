package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/rosbit/go-wx-api/conf"
	"github.com/rosbit/go-wx-api/auth"
	"github.com/rosbit/go-wx-api/tools"
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
	params, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.CreateMenu(accessToken, params.MenuJson)
	if err != nil {
		fmt.Printf("failed to create menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %s\n", string(resp))
}

func queryMenu(paramsFile string) {
	_, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.QueryMenu(accessToken)
	if err != nil {
		fmt.Printf("failed to query menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %s\n", string(resp))
}

func deleteMenu(paramsFile string) {
	_, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.DeleteMenu(accessToken)
	if err != nil {
		fmt.Printf("failed to delete menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %s\n", string(resp))
}

func infoMenu(paramsFile string) {
	_, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}
	resp, err := wxtools.CurrentSelfmenuInfo(accessToken)
	if err != nil {
		fmt.Printf("failed to get information of current self-menu: %v\n", err)
		return
	}
	fmt.Printf("resp: %s\n", string(resp))
}

func loadConf(paramsFile string) (*ParamsT, string, error) {
	paramsContent, err := ioutil.ReadFile(paramsFile)
	if err != nil {
		return nil, "", err
	}

	var params ParamsT
	if err = json.Unmarshal(paramsContent, &params); err != nil {
		return nil, "", err
	}

	wxconf.SetParams(params.Token, params.AppId, params.Secret, "")
	wxconf.TokenStorePath = params.TokenCache

	accessToken, err := wxauth.NewAccessToken().Get()
	if err != nil {
		return nil, "", err
	}
	return &params, accessToken, nil
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
