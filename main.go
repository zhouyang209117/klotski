package main

import (
	"net/http"
	"encoding/json"
	"klotski/model"
	"log"
	"klotski/constant"
	"klotski/controller"
	"klotski/util"
	"github.com/astaxie/beego/config"
	"fmt"
	"io/ioutil"
)

func getStartState(s model.ClientState) model.State {
	state := model.State {}
	state.SetCoord(constant.ZHANG_FEI, s.ZhangFei.X, s.ZhangFei.Y)
	state.SetCoord(constant.CAO_CAO, s.Caocao.X, s.Caocao.Y)
	state.SetCoord(constant.MA_CHAO, s.MaoCho.X, s.MaoCho.Y)
	state.SetCoord(constant.ZHAO_YUN, s.ZhaoYun.X, s.ZhaoYun.Y)
	state.SetCoord(constant.GUAN_YU, s.GuangYu.X, s.GuangYu.Y)
	state.SetCoord(constant.HUANG_ZHONG, s.HuangZhong.X, s.HuangZhong.Y)
	state.SetCoord(constant.BING1, s.Bing1.X, s.Bing1.Y)
	state.SetCoord(constant.BING2, s.Bing2.X, s.Bing2.Y)
	state.SetCoord(constant.BING3, s.Bing3.X, s.Bing3.Y)
	state.SetCoord(constant.BING4, s.Bing4.X, s.Bing4.Y)
	state.SetCoord(constant.FREE1, s.Free1.X, s.Free1.Y)
	state.SetCoord(constant.FREE2, s.Free2.X, s.Free2.Y)
	return state
}

func returnRequest(w http.ResponseWriter, result model.Result) {
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func handler(w http.ResponseWriter, req *http.Request) {
	log.Println("call handler")
	s := model.ClientState{}
	result := model.Result{}
	result.Success = false
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("readAll error", err.Error())
		returnRequest(w, result)
		return
	}
	log.Println("body content:", string(body))
	err = json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal("Unmarshal error", err.Error())
		returnRequest(w, result)
		return
	}
	bytesClientStates, _ := json.Marshal(s)
	log.Println(util.PrettyFormat(bytesClientStates))
	startState := getStartState(s)
	success, steps := controller.Klotski(startState)
	if success {
		c := steps.Front()
		if c == nil {
			log.Fatal("状态异常")
			result.Success = false
		} else {
			resultSteps := make([]model.ResultStep, 0)
			m1 := util.GetPersonENMap()
			m2 := util.GetDirectionENMap()
			cnt := 1
			for c = c.Next(); c != nil; c = c.Next() {
				nameIndex, direction := util.GetStep((c.Value).(uint8))
				resultStep := model.ResultStep{
					PersonType: m1[nameIndex], Direction: m2[direction]}
				resultSteps = append(resultSteps, resultStep)
				cnt += 1
			}
			result.Steps = resultSteps
			result.StepNum = len(resultSteps)
			result.Success = true
		}

	}
	returnRequest(w, result)
}

func main()  {
	http.HandleFunc("/", handler)
	conf, _ := config.NewConfig("ini", "conf/app.conf")
	port, err := conf.Int64("port")
	if err != nil {
		log.Println("read port error")
		port = 80
	}
	log.Println("server is running, port:", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("start server error")
	}
}
