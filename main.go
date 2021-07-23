package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yrshiben/baidu-translation-workflow/model"
	"os"
	"unicode"
)

var ErrItem = model.AlfredItem{Title: "翻译失败", Subtitle: ""}

var api *TransApi

func main() {
	var query, appId, securityKey string
	app := &cli.App{
		Name:      "百度翻译工具",
		Usage:     "baidu-translation",
		UsageText: "baidu-translation -q <要翻译的内容> --appId <appId> --securityKey <securityKey>",
		Version:   "1.0.2",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "query", Aliases: []string{"q"}, Usage: "要翻译的内容", Required: true, Destination: &query},
			&cli.StringFlag{Name: "appId", Usage: "在 https://fanyi-api.baidu.com/api/trans/product/desktop?req=developer 申请的 appId", Required: true, Destination: &appId},
			&cli.StringFlag{Name: "securityKey", Usage: "在 https://fanyi-api.baidu.com/api/trans/product/desktop?req=developer 申请的 securityKey", Destination: &securityKey},
		},
		Action: func(c *cli.Context) error {
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
			return nil
		},
	}
	_ = app.Run(os.Args)
}

// IsChinese 判断字符串中是否包含中文
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
