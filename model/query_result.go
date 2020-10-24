package model

// {"from":"en","to":"zh","trans_result":[{"src":"hello","dst":"你好"}]}
type Result struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []Data `json:"trans_result"`
}
type Data struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
