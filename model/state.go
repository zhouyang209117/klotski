package model
import
(
	"klotski/constant"
	"klotski/util"
)
/*
12个棋子(2个空位置也算),横坐标长度为4,需要2bit表示,纵坐标长度是5,需要3bit表示
也就是一个棋子的位置要5bit,12个棋子一共是12*5=60bit,一个int64就够用了.
*/
type State struct {
	ChessPos uint64 //各个棋子的位置
	Step uint8       //哪个棋子怎么移动得到现在位置
	Pre *State
}

func(s *State) SetCoord(nameIndex, xUint8, yUint8 uint8) {
	x := uint64(xUint8)
	y := uint64(yUint8)
	offset := (constant.BIT_X + constant.BIT_Y) * nameIndex
	middle := (x << (offset + constant.BIT_Y)) | (y << offset)
	low_all_1 := (uint64(1) << offset) - 1
	mid_low := middle | (s.ChessPos & low_all_1)
	high_num := constant.BIT_SUM - (offset + constant.BIT_X + constant.BIT_Y)
	high_1_mid_low_0 := ((uint64(1) << high_num) - 1) << (offset + constant.BIT_X + constant.BIT_Y)
	s.ChessPos = s.ChessPos & high_1_mid_low_0 | mid_low
}

func(s *State) GetCoord(nameIndex uint8) (uint8, uint8) {
	offset := (constant.BIT_X + constant.BIT_Y) * nameIndex
	mid_low_all_1 := (uint64(1) << (offset + constant.BIT_X + constant.BIT_Y)) - 1
	mid_low := s.ChessPos & mid_low_all_1
	x_y := mid_low >> offset
	x := mid_low >> (offset + constant.BIT_Y)
	y := x_y & ((1 << constant.BIT_Y) - 1)
	return uint8(x), uint8(y)
}

func(s *State) SetStep(nameIndex, direction uint8) {
	s.Step = nameIndex | (direction << constant.BIT_NAME)
}

func(s *State) GetStep() (uint8, uint8) {
	return util.GetStep(s.Step)
}

func(s *State) GetKey() uint64 {
	sk := uint64(0)
	x, y := s.GetCoord(constant.CAO_CAO)
	util.SetKeyPosValue2_2(&sk, x, y)
	x, y = s.GetCoord(constant.ZHANG_FEI)
	util.SetKeyPosValue2_1(&sk, x, y)
	x, y = s.GetCoord(constant.MA_CHAO)
	util.SetKeyPosValue2_1(&sk, x, y)
	x, y = s.GetCoord(constant.ZHAO_YUN)
	util.SetKeyPosValue2_1(&sk, x, y)
	x, y = s.GetCoord(constant.HUANG_ZHONG)
	util.SetKeyPosValue2_1(&sk, x, y)
	x, y = s.GetCoord(constant.GUAN_YU)
	util.SetKeyPosValue1_2(&sk, x, y)
	x, y = s.GetCoord(constant.BING1)
	util.SetKeyPosValue1_1(&sk, x, y)
	x, y = s.GetCoord(constant.BING2)
	util.SetKeyPosValue1_1(&sk, x, y)
	x, y = s.GetCoord(constant.BING3)
	util.SetKeyPosValue1_1(&sk, x, y)
	x, y = s.GetCoord(constant.BING4)
	util.SetKeyPosValue1_1(&sk, x, y)
	return sk
}
