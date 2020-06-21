package model

type ResultStep struct {
	PersonType string `json:"person_type"`
	Direction  string `json:"direction"`
}

type Result struct {
	Success bool   `json:"success"`
	StepNum int   `json:"step_num"`
	Steps   []ResultStep `json:"steps"`
}

type ClientState struct {
	ZhangFei   Coord `json:"zhangfei"`
	Caocao     Coord `json:"caocao"`
	MaCho      Coord `json:"machao"`
	ZhaoYun    Coord `json:"zhaoyun"`
	GuangYu    Coord `json:"guanyu"`
	HuangZhong Coord `json:"huangzhong"`
	Bing1      Coord `json:"bing1"`
	Bing2      Coord `json:"bing2"`
	Bing3      Coord `json:"bing3"`
	Bing4      Coord `json:"bing4"`
	Free1      Coord `json:"free1"`
	Free2      Coord `json:"free2"`
}