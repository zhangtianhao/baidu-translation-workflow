package model

import "encoding/json"

// 想让 alfred 显示列表，我们需要返回对应格式的数据
type AlfredList struct {
	Items []AlfredItem `json:"items"`
}
type AlfredItem struct {
	Title    string `json:"title"`
	Icon     Icon   `json:"icon"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}
type Icon struct {
	Path string `json:"path"`
}

// 将 AlfredList 数据转换为 json 字符串
func (alfredList *AlfredList) ToJson() string {
	bytes, _ := json.Marshal(alfredList)
	return string(bytes)
}
