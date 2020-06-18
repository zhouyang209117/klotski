package tests

import (
	"testing"
	"klotski/constant"
	"klotski/util"
	"klotski/model"
)

func TestStateKey1(t *testing.T) {
	sk := uint64(0)
	util.SetKeyType(&sk, 0, 0, constant.TYPE_1_1)
	util.AssertUint8(t, util.GetKeyType(sk, 0, 0), constant.TYPE_1_1)
	util.SetKeyType(&sk, 4, 3, constant.TYPE_2_2)
	util.SetKeyType(&sk, 2, 2, constant.TYPE_2_1)
	util.SetKeyType(&sk, 0, 0, constant.TYPE_1_1)
	util.SetKeyType(&sk, 4, 3, constant.TYPE_2_2)
}

func TestStateKey2(t *testing.T) {
	s1 := model.State{}
	s1.SetCoord(constant.ZHANG_FEI, 0, 0)
	s1.SetCoord(constant.CAO_CAO, 0, 1)
	s1.SetCoord(constant.MA_CHAO, 0, 3)
	s1.SetCoord(constant.ZHAO_YUN, 2, 0)
	s1.SetCoord(constant.GUAN_YU, 4, 1)
	s1.SetCoord(constant.HUANG_ZHONG, 2, 3)
	s1.SetCoord(constant.BING1, 2, 1)
	s1.SetCoord(constant.BING2, 3, 2)
	s1.SetCoord(constant.BING3, 4, 0)
	s1.SetCoord(constant.BING4, 4, 3)
	s1.SetCoord(constant.FREE1, 2, 2)
	s1.SetCoord(constant.FREE2, 3, 1)
	s2 := model.State{}
	s2.SetCoord(constant.ZHANG_FEI, 2, 3)
	s2.SetCoord(constant.CAO_CAO, 0, 1)
	s2.SetCoord(constant.MA_CHAO, 0, 3)
	s2.SetCoord(constant.ZHAO_YUN, 2, 0)
	s2.SetCoord(constant.GUAN_YU, 4, 1)
	s2.SetCoord(constant.HUANG_ZHONG, 0, 0)
	s2.SetCoord(constant.BING1, 2, 1)
	s2.SetCoord(constant.BING2, 3, 2)
	s2.SetCoord(constant.BING3, 4, 0)
	s2.SetCoord(constant.BING4, 4, 3)
	s2.SetCoord(constant.FREE1, 2, 2)
	s2.SetCoord(constant.FREE2, 3, 1)
	k1 := s1.GetKey()
	k2 := s2.GetKey()
	util.AssertUint64(t, k1, k2)
}
