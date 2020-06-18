package tests

import (
	"testing"
	"klotski/model"
	"klotski/constant"
	"klotski/util"
)

func TestCoord1(t *testing.T) {
	s := model.State {
		ChessPos: uint64(0),
	}
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
	x, y := s.GetCoord(constant.ZHANG_FEI)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 0)
	x, y = s.GetCoord(constant.CAO_CAO)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 1)
	x, y = s.GetCoord(constant.MA_CHAO)
	util.AssertUint8(t, x, 0)
	util.AssertUint8(t, y, 3)
	x, y = s.GetCoord(constant.ZHAO_YUN)
	util.AssertUint8(t, x, 2)
	util.AssertUint8(t, y, 0)
	x, y = s.GetCoord(constant.GUAN_YU)
	util.AssertUint8(t, x, 2)
	util.AssertUint8(t, y, 1)
	x, y = s.GetCoord(constant.HUANG_ZHONG)
	util.AssertUint8(t, x, 2)
	util.AssertUint8(t, y, 3)
	x, y = s.GetCoord(constant.BING1)
	util.AssertUint8(t, x, 3)
	util.AssertUint8(t, y, 1)
	x, y = s.GetCoord(constant.BING2)
	util.AssertUint8(t, x, 3)
	util.AssertUint8(t, y, 2)
	x, y = s.GetCoord(constant.BING3)
	util.AssertUint8(t, x, 4)
	util.AssertUint8(t, y, 0)
	x, y = s.GetCoord(constant.BING4)
	util.AssertUint8(t, x, 4)
	util.AssertUint8(t, y, 3)
	x, y = s.GetCoord(constant.FREE1)
	util.AssertUint8(t, x, 4)
	util.AssertUint8(t, y, 1)
	x, y = s.GetCoord(constant.FREE2)
	util.AssertUint8(t, x, 4)
	util.AssertUint8(t, y, 2)
	s.SetStep(constant.CAO_CAO, constant.UP)
	nameIndex, direction := s.GetStep()
	util.AssertUint8(t, nameIndex, constant.CAO_CAO)
	util.AssertUint8(t, direction, constant.UP)
	s.SetStep(constant.GUAN_YU, constant.DOWN)
	nameIndex, direction = s.GetStep()
	util.AssertUint8(t, nameIndex, constant.GUAN_YU)
	util.AssertUint8(t, direction, constant.DOWN)
	s.SetStep(constant.ZHANG_FEI, constant.LEFT)
	nameIndex, direction = s.GetStep()
	util.AssertUint8(t, nameIndex, constant.ZHANG_FEI)
	util.AssertUint8(t, direction, constant.LEFT)
	s.SetStep(constant.ZHAO_YUN, constant.RIGHT)
	nameIndex, direction = s.GetStep()
	util.AssertUint8(t, nameIndex, constant.ZHAO_YUN)
	util.AssertUint8(t, direction, constant.RIGHT)
}

func TestCoord2(t *testing.T) {
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
}