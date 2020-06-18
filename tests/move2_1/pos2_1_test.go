package move2_1

import (
	"testing"
	"klotski/model"
	"klotski/constant"
	"klotski/controller"
	"klotski/util"
)

func TestUp(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 1, 0)
	s.SetCoord(constant.CAO_CAO, 0, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 3, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 3, 2)
	s.SetCoord(constant.BING3, 4, 1)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 0, 0)
	s.SetCoord(constant.FREE2, 4, 2)
	new_states := controller.Pos2_1Next(s, constant.ZHANG_FEI)
	util.AssertUint8(t, uint8(len(new_states)), 1)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.ZHANG_FEI)
	util.AssertUint8(t, direction, constant.UP)
	x, y := new_states[0].GetCoord(constant.ZHANG_FEI)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 0)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	if x1 == 2 {
		util.AssertUint8(t, y1, 0)
		util.AssertUint8(t, x2, 4)
		util.AssertUint8(t, y2, 2)
	} else {
		util.AssertUint8(t, x1, 4)
		util.AssertUint8(t, y1, 2)
		util.AssertUint8(t, x2, 2)
		util.AssertUint8(t, y2, 0)
	}
}

func TestRight(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 0, 2)
	s.SetCoord(constant.MA_CHAO, 3, 2)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 4, 0)
	s.SetCoord(constant.BING3, 4, 1)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 0, 1)
	s.SetCoord(constant.FREE2, 1, 1)
	new_states := controller.Pos2_1Next(s, constant.ZHANG_FEI)
	util.AssertUint8(t, uint8(len(new_states)), 1)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.ZHANG_FEI)
	util.AssertUint8(t, direction, constant.RIGHT)
	x, y := new_states[0].GetCoord(constant.ZHANG_FEI)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 1)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	if x1 == 0 {
		util.AssertUint8(t, y1, 0)
		util.AssertUint8(t, x2, 1)
		util.AssertUint8(t, y2, 0)
	} else {
		util.AssertUint8(t, x1, 1)
		util.AssertUint8(t, y1, 0)
		util.AssertUint8(t, x2, 0)
		util.AssertUint8(t, y2, 0)
	}
}

func TestDown(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 0, 2)
	s.SetCoord(constant.MA_CHAO, 3, 2)
	s.SetCoord(constant.ZHAO_YUN, 3, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 3, 2)
	s.SetCoord(constant.BING3, 4, 1)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 2, 0)
	s.SetCoord(constant.FREE2, 4, 2)
	new_states := controller.Pos2_1Next(s, constant.ZHANG_FEI)
	util.AssertUint8(t, uint8(len(new_states)), 1)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.ZHANG_FEI)
	util.AssertUint8(t, direction, constant.DOWN)
	x, y := new_states[0].GetCoord(constant.ZHANG_FEI)
	util.AssertUint8(t, x, 1)
	util.AssertUint8(t, y, 0)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	if x1 == 0 {
		util.AssertUint8(t, y1, 0)
		util.AssertUint8(t, x2, 4)
		util.AssertUint8(t, y2, 2)
	} else {
		util.AssertUint8(t, x1, 4)
		util.AssertUint8(t, y1, 2)
		util.AssertUint8(t, x2, 0)
		util.AssertUint8(t, y2, 0)
	}
}

func TestLeft(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 1)
	s.SetCoord(constant.CAO_CAO, 0, 2)
	s.SetCoord(constant.MA_CHAO, 3, 2)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 4, 1)
	s.SetCoord(constant.BING3, 4, 0)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 0, 0)
	s.SetCoord(constant.FREE2, 1, 0)
	new_states := controller.Pos2_1Next(s, constant.ZHANG_FEI)
	util.AssertUint8(t, uint8(len(new_states)), 1)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.ZHANG_FEI)
	util.AssertUint8(t, direction, constant.LEFT)
	x, y := new_states[0].GetCoord(constant.ZHANG_FEI)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 0)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	if x1 == 0 {
		util.AssertUint8(t, y1, 1)
		util.AssertUint8(t, x2, 1)
		util.AssertUint8(t, y2, 1)
	} else {
		util.AssertUint8(t, x1, 1)
		util.AssertUint8(t, y1, 1)
		util.AssertUint8(t, x2, 0)
		util.AssertUint8(t, y2, 1)
	}
}

func TestUpDown(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 0, 1)
	s.SetCoord(constant.MA_CHAO, 1, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 2, 1)
	s.SetCoord(constant.HUANG_ZHONG, 3, 2)
	s.SetCoord(constant.BING1, 3, 1)
	s.SetCoord(constant.BING2, 4, 1)
	s.SetCoord(constant.BING3, 4, 0)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 0, 3)
	s.SetCoord(constant.FREE2, 3, 3)
	new_states := controller.Pos2_1Next(s, constant.MA_CHAO)
	util.AssertUint8(t, uint8(len(new_states)), 2)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.MA_CHAO)
	util.AssertUint8(t, direction, constant.UP)
	x, y := new_states[0].GetCoord(constant.MA_CHAO)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 3)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	if x1 == 2 {
		util.AssertUint8(t, y1, 3)
		util.AssertUint8(t, x2, 3)
		util.AssertUint8(t, y2, 3)
	} else {
		util.AssertUint8(t, x1, 3)
		util.AssertUint8(t, y1, 3)
		util.AssertUint8(t, x2, 2)
		util.AssertUint8(t, y2, 3)
	}
	x, y = new_states[1].GetCoord(constant.MA_CHAO)
	util.AssertUint8(t, x, 2)
	util.AssertUint8(t, y, 3)
	step_index_name, direction = new_states[1].GetStep()
	util.AssertUint8(t, step_index_name, constant.MA_CHAO)
	util.AssertUint8(t, direction, constant.DOWN)
	util.AssertDeep(t, new_states[1].Pre, &s)
	x1, y1 = new_states[1].GetCoord(constant.FREE1)
	x2, y2 = new_states[1].GetCoord(constant.FREE2)
	if x1 == 0 {
		util.AssertUint8(t, y1, 3)
		util.AssertUint8(t, x2, 1)
		util.AssertUint8(t, y2, 3)
	} else {
		util.AssertUint8(t, x1, 1)
		util.AssertUint8(t, y1, 3)
		util.AssertUint8(t, x2, 0)
		util.AssertUint8(t, y2, 3)
	}
}
