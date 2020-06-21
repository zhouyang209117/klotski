package tests

import (
	"testing"
	"net/http"
	"bytes"
	"io/ioutil"
	"klotski/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	util2 "klotski/util"
	"github.com/astaxie/beego/config"
)

func getHenDaoLiMa() model.ClientState {
	r := model.ClientState{}
	r.ZhangFei = model.Coord{X: 0, Y: 0}
	r.Caocao = model.Coord{X: 0, Y: 1}
	r.MaCho = model.Coord{X: 0, Y: 3}
	r.ZhaoYun = model.Coord{X: 2, Y: 0}
	r.GuangYu = model.Coord{X: 2, Y: 1}
	r.HuangZhong = model.Coord{X: 2, Y: 3}
	r.Bing1 = model.Coord{X: 3, Y: 1}
	r.Bing2 = model.Coord{X: 3, Y: 2}
	r.Bing3 = model.Coord{X: 4, Y: 0}
	r.Bing4 = model.Coord{X: 4, Y: 3}
	r.Free1 = model.Coord{X: 4, Y: 1}
	r.Free2 = model.Coord{X: 4, Y: 2}
	return r
}

func getSimple1() model.ClientState {
	r := model.ClientState{}
	r.ZhangFei = model.Coord{X: 0, Y: 0}
	r.Caocao = model.Coord{X: 2, Y: 1}
	r.MaCho = model.Coord{X: 0, Y: 3}
	r.ZhaoYun = model.Coord{X: 2, Y: 0}
	r.GuangYu = model.Coord{X: 1, Y: 1}
	r.HuangZhong = model.Coord{X: 3, Y: 3}
	r.Bing1 = model.Coord{X: 4, Y: 0}
	r.Bing2 = model.Coord{X: 0, Y: 1}
	r.Bing3 = model.Coord{X: 0, Y: 2}
	r.Bing4 = model.Coord{X: 2, Y: 3}
	r.Free1 = model.Coord{X: 4, Y: 1}
	r.Free2 = model.Coord{X: 4, Y: 2}
	return r
}
func TestHttp(t *testing.T) {
	conf, _ := config.NewConfig("ini", "conf/app.conf")
	url := conf.String("url")
	//r := getSimple1()
	r := getHenDaoLiMa()
	bytesR, err := json.Marshal(r)
	prettyResult1 := util2.PrettyFormat(bytesR)
	log.Println(prettyResult1)
	if err != nil {
		assert.Equal(t, 1, 2)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesR))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	prettyResult := util2.PrettyFormat(body)
	log.Println(prettyResult)
}
