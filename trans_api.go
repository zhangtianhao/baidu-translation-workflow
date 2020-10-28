package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/yrshiben/baidu-translation-workflow/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const transApiHost string = "https://fanyi-api.baidu.com/api/trans/vip/translate"

type TransApi struct {
	AppId       string
	SecurityKey string
}

func NewTransApi(appId string, securityKey string) *TransApi {
	return &TransApi{AppId: appId, SecurityKey: securityKey}
}

// from/to: en or zh
func (api *TransApi) GetTransResult(query, from, to string) (*model.Result, error) {
	params := api.buildParams(query, from, to)
	return get(params)
}

func (api *TransApi) buildParams(query, from, to string) string {
	salt := strconv.FormatInt(time.Now().Unix(), 10)
	sign := Md5(api.AppId + query + salt + api.SecurityKey)
	// ?appid=appid&salt=salt&from=from&to=to&sign=sign&q=query
	queryParams := "?appid=" + api.AppId + "&salt=" + salt + "&from=" + from + "&to=" + to + "&sign=" + sign + "&q=" + url.PathEscape(query)
	return queryParams
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func get(query string) (*model.Result, error) {
	var queryResult *model.Result
	resp, err := http.Get(transApiHost + query)
	if err != nil {
		return queryResult, err
	}
	defer resp.Body.Close()
	if result, err := ioutil.ReadAll(resp.Body); err != nil {
		return queryResult, err
	} else {
		err = json.Unmarshal(result, &queryResult)
		return queryResult, err
	}
}
