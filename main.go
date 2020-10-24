package main

import (
	"flag"
	"fmt"
	"github.com/yrshiben/baidu-translation-workflow/model"
	"unicode"
)

var ErrItem = model.AlfredItem{Title: "翻译失败", Subtitle: ""}

var api *TransApi

func main() {
	query, appId, securityKey := GetInputStr()
	api = NewTransApi(appId, securityKey)
	// 默认 英译汉
	var from, to = "en", "zh"
	if IsChinese(query) { // 输入的是汉文，则 汉译英
		from, to = to, from
	}
	var items []model.AlfredItem
	if result, err := api.GetTransResult(query, from, to); err != nil || len(result.TransResult) == 0 {
		items = []model.AlfredItem{ErrItem}
	} else {
		data := result.TransResult
		items = make([]model.AlfredItem, 0, len(data))
		for _, val := range data {
			items = append(items, model.AlfredItem{Title: val.Dst, Subtitle: val.Src, Arg: val.Dst})
		}
	}
	alfredList := model.AlfredList{Items: items}
	fmt.Println(alfredList.ToJson())
}

func GetInputStr() (q, appId, securityKey string) {
	flag.StringVar(&q, "q", "", "要翻译的内容")
	flag.StringVar(&appId, "appId", "", "APP ID")
	flag.StringVar(&securityKey, "securityKey", "", "秘钥")
	// 这行很关键，一定要有
	flag.Parse()
	return
}

// 判断字符串中是否包含中文
func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
