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
}

func loadConf(paramsFile string) (error) {
	fp, err := os.Open(paramsFile)
	if err != nil {
		return err
	}
	defer fp.Close()

	var params ParamsT
	if err = json.NewDecoder(fp).Decode(&params); err != nil {
		return err
	}

	wxapi.SetWxParams(service, params.Token, params.AppId, params.Secret, "")
	wxapi.InitWx(params.TokenCache)

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <params-json-file> <openId>\n", os.Args[0])
		return
	}

	paramsFile, openId := os.Args[1], os.Args[2]
	if err := loadConf(paramsFile); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	res, err := wxtools.GetUserInfo(service, openId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for k, v := range res {
		switch v.(type) {
		case float64:
			fmt.Printf("%s => %d\n", k, int64(v.(float64)))
		default:
			fmt.Printf("%s => %v\n", k, v)
		}
	}
}
