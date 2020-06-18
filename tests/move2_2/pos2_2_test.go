package move1_1

import (
	"testing"
	"klotski/model"
	"klotski/constant"
	"klotski/controller"
	"klotski/util"
)

func TestUp(t *testing.T) {
	s := model.State{}
	s.SetCoord(constant.ZHANG_FEI, 0, 0)
	s.SetCoord(constant.CAO_CAO, 1, 1)
	s.SetCoord(constant.MA_CHAO, 0, 3)
	s.SetCoord(constant.ZHAO_YUN, 2, 0)
	s.SetCoord(constant.GUAN_YU, 3, 1)
	s.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s.SetCoord(constant.BING1, 4, 0)
	s.SetCoord(constant.BING2, 4, 1)
	s.SetCoord(constant.BING3, 4, 2)
	s.SetCoord(constant.BING4, 4, 3)
	s.SetCoord(constant.FREE1, 0, 1)
	s.SetCoord(constant.FREE2, 0, 2)
	new_states := controller.Pos2_2Next(s, constant.CAO_CAO)
	util.AssertUint8(t, uint8(len(new_states)), 1)
	step_index_name, direction := new_states[0].GetStep()
	util.AssertUint8(t, step_index_name, constant.CAO_CAO)
	util.AssertUint8(t, direction, constant.UP)
	x, y := new_states[0].GetCoord(constant.CAO_CAO)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 1)
	util.AssertDeep(t, new_states[0].Pre, &s)
	x1, y1 := new_states[0].GetCoord(constant.FREE1)
	x2, y2 := new_states[0].GetCoord(constant.FREE2)
	util.AssertUint8(t, x1, 2)
	util.AssertUint8(t, x2, 2)
	if y1 == 1 {
		util.AssertUint8(t, y2, 2)
	} else {
		util.AssertUint8(t, y1, 2)
		util.AssertUint8(t, y2, 1)
	}
}