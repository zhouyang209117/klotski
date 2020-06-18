package tests

import (
	"testing"
	"klotski/model"
	"klotski/constant"
	"klotski/controller"
	"fmt"
	"klotski/util"
)

/*
		---------------------
        |zf   |cc       | mc |
        |     |         |    |
		|     |         |    |
        |     |         |    |
		|--------------------|
        | zy  |gy       | hz |
        |     |         |    |
        |     |---------|    |
		|     |b1  |b2  |    |
        |     |    |    |    |
		|--------------------|
        |b3   |    |    |b4  |
        |     |    |    |    |
        ----------------------
 */
func getHengDaoLiMa() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 0, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 3, 2)
	s.SetCoord(constant.BING3, 4, 0)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 4, 1)
	s.SetCoord(constant.FREE2, 4, 2)
	return s
}


func getSimple1() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 1, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 3)
	s.SetCoord(constant.BING1, 4, 0)
	s.SetCoord(constant.BING2, 0, 1)
	s.SetCoord(constant.BING3, 0, 2)
	s.SetCoord(constant.BING4, 2, 3)
	s.SetCoord(constant.FREE1, 4, 1)
	s.SetCoord(constant.FREE2, 4, 2)
	return s
}

func getSimple2() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 1, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 3)
	s.SetCoord(constant.BING1, 4, 1)
	s.SetCoord(constant.BING2, 0, 1)
	s.SetCoord(constant.BING3, 0, 2)
	s.SetCoord(constant.BING4, 2, 3)
	s.SetCoord(constant.FREE1, 4, 0)
	s.SetCoord(constant.FREE2, 4, 2)
	return s
}

func getSimple3() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 3, 0)
	s.SetCoord(constant.GUAN_YU, 1, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 3)
	s.SetCoord(constant.BING1, 4, 1)
	s.SetCoord(constant.BING2, 0, 1)
	s.SetCoord(constant.BING3, 0, 2)
	s.SetCoord(constant.BING4, 2, 3)
	s.SetCoord(constant.FREE1, 2, 0)
	s.SetCoord(constant.FREE2, 4, 2)
	return s
}

func getSimple4() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 1, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 3)
	s.SetCoord(constant.BING1, 4, 1)
	s.SetCoord(constant.BING2, 0, 1)
	s.SetCoord(constant.BING3, 0, 2)
	s.SetCoord(constant.BING4, 4, 2)
	s.SetCoord(constant.FREE1, 4, 0)
	s.SetCoord(constant.FREE2, 2, 3)
	return s
}

func getSimple5() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 1, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 1, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 3)
	s.SetCoord(constant.BING1, 4, 0)
	s.SetCoord(constant.BING2, 0, 2)
	s.SetCoord(constant.BING3, 0, 3)
	s.SetCoord(constant.BING4, 4, 2)
	s.SetCoord(constant.FREE1, 0, 1)
	s.SetCoord(constant.FREE2, 4, 1)
	return s
}

/*
		---------------------
        |zf   |cc       | mc |
        |     |         |    |
		|     |         |    |
        |     |         |    |
		|--------------------|
        | zy  |gy       | hz |
        |     |         |    |
        |     |---------|    |
		|     |b1  |b2  |    |
        |     |    |    |    |
		|--------------------|
        |b3   |    |    |b4  |
        |     |    |    |    |
        ----------------------
 */

func getSimple6() model.State {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 2, 1)
	s.SetCoord(constant.MA_CHAO, 0, 1)
	s.SetCoord(constant.ZHAO_YUN, 3, 0)
	s.SetCoord(constant.GUAN_YU, 0, 2)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 1, 2)
	s.SetCoord(constant.BING2, 1, 3)
	s.SetCoord(constant.BING3, 4, 1)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 2, 0)
	s.SetCoord(constant.FREE2, 4, 2)
	return s
}

func TestAlgo(t *testing.T) {
	s := getHengDaoLiMa()
	result, steps := controller.Klotski(s)
	if result {
		c := steps.Front()
		if c == nil {
			fmt.Println("状态异常")
			return
		} else {
			m1 := util.GetPersonCNMap()
			m2 := util.GetDirectionCNMap()
			cnt := 1
			for c = c.Next(); c != nil; c = c.Next() {
				nameIndex, direction := util.GetStep((c.Value).(uint8))
				fmt.Println(fmt.Sprintf("%d.%s%s移", cnt, m1[nameIndex], m2[direction]))
				cnt += 1
			}
		}
	}
}
