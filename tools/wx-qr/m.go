package main

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
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

func showResult(ticketURL2ShowQrCode, urlIncluedInQrcode string, err error) {
	if err != nil {
		fmt.Printf("failed to create Qrcode: %v\n", err)
		return
	}
	fmt.Printf("ticketURL2ShowQrCode: %s\n", ticketURL2ShowQrCode)
	fmt.Printf("urlIncluedInQrcode: %s\n", urlIncluedInQrcode)
}

func createTempQrInt(accessToken string, sceneId interface{}, expireTime int) {
	showResult(wxtools.CreateTempQrIntScene(accessToken, sceneId.(int), expireTime))
}

func createTempQrStr(accessToken string, sceneId interface{}, expireTime int) {
	showResult(wxtools.CreateTempQrStrScene(accessToken, sceneId.(string), expireTime))
}

func createQrInt(accessToken string, sceneId interface{}, expireTime int) {
	showResult(wxtools.CreateQrIntScene(accessToken, sceneId.(int)))
}

func createQrStr(accessToken string, sceneId interface{}, expireTime int) {
	showResult(wxtools.CreateQrStrScene(accessToken, sceneId.(string)))
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

const (
	TEMP_INT = "temp-int"
	TEMP_STR = "temp-str"
	FOREVER_INT = "forever-int"
	FOREVER_STR = "forever-str"
)

var _commands = map[string]func(string, interface{}, int) {
	TEMP_INT:    createTempQrInt,
	TEMP_STR:    createTempQrStr,
	FOREVER_INT: createQrInt,
	FOREVER_STR: createQrStr,
}

var _needExpireTime = map[string]bool {
	TEMP_INT:    true,
	TEMP_STR:    true,
	FOREVER_INT: false,
	FOREVER_STR: false,
}

var _intSceneId = map[string]bool {
	TEMP_INT:    true,
	TEMP_STR:    false,
	FOREVER_INT: true,
	FOREVER_STR: false,
}

func usage() {
	fmt.Printf("Usage: %s <command> <params-json-file> <sceneId> [ <expireTimeInSeconds>]\n", os.Args[0])
	fmt.Printf("Where <command> could be:\n")
	fmt.Printf("  temp-int    ---  create qrcode with an integer id and expired time, default 30s, max 30 days\n")
	fmt.Printf("  temp-str    ---  create qrcode with a string id and expired time, default 30s, max 30 days\n")
	fmt.Printf("  forever-int ---  create qrcode with an integer id\n")
	fmt.Printf("  forever-str ---  create qrcode with a string id\n")
}

func main() {
	if len(os.Args) < 4 {
		usage()
		return
	}
	command, paramsFile, sceneId := os.Args[1], os.Args[2], os.Args[3]
	intSceneId, ok := _intSceneId[command]
	if !ok {
		fmt.Printf("Unknown command: %s\n", command)
		return
	}
	var iSceneId interface{}
	if intSceneId {
		id, err := strconv.Atoi(sceneId)
		if err != nil {
			fmt.Printf("command %s need an integer sceneId: %v\n", command, err)
			return
		}
		iSceneId = id
	} else {
		iSceneId = sceneId
	}

	needExpireTime, _ := _needExpireTime[command]
	expireTime := 0
	var err error
	if needExpireTime {
		if len(os.Args) < 5 {
			fmt.Printf("command %s need params <expireTimeInSeconds>\n", command)
			usage()
			return
		}
		expireTime, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Printf("<expireTimeInSeconds> error: %v\n", err)
			return
		}
	}

	_, accessToken, err := loadConf(paramsFile)
	if err != nil {
		fmt.Printf("failed to load conf: %v\n", err)
		return
	}

	fn, _ := _commands[command]
	fn(accessToken, iSceneId, expireTime)
}
