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

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <params-json-file> <long-url>\n", os.Args[0])
		return
	}

	paramsFile, longUrl := os.Args[1], os.Args[2]
	_, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	shortUrl, err := wxtools.MakeShorturl(accessToken, longUrl)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n", shortUrl)
}
