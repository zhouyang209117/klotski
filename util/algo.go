package util

import (
	"klotski/constant"
	"encoding/json"
	"bytes"
	"fmt"
)

func SetKeyType(key *uint64, x, y, keyType uint8) {
	tmpKey := *key
	offset := (x * constant.COL + y) * constant.BIT_TYPE
	mid := uint64(keyType) << offset
	low := ((1 << offset) - 1) & tmpKey
	mid_low := mid | low
	f := (((uint64(1) << constant.BIT_SUM_TYPE) - 1) >> (offset + constant.BIT_TYPE)) <<
		(offset + constant.BIT_TYPE)
	*key = (f & tmpKey) | mid_low
}

func GetKeyType(key uint64, x, y uint8) uint8 {
	offset := (x * constant.COL + y) * constant.BIT_TYPE
	mid_low := ((1 << (offset + constant.BIT_TYPE)) - 1) & key
	return uint8(mid_low >> offset)
}

func SetKeyPosValue2_2(key *uint64, x, y uint8) {
	SetKeyType(key, x, y, constant.TYPE_2_2)
	SetKeyType(key, x, y + 1, constant.TYPE_2_2)
	SetKeyType(key, x + 1, y, constant.TYPE_2_2)
	SetKeyType(key, x + 1, y + 1, constant.TYPE_2_2)
}

func SetKeyPosValue2_1(key *uint64, x, y uint8) {
	SetKeyType(key, x, y, constant.TYPE_2_1)
	SetKeyType(key, x + 1, y, constant.TYPE_2_1)
}

func SetKeyPosValue1_2(key *uint64, x, y uint8) {
	SetKeyType(key, x, y, constant.TYPE_1_2)
	SetKeyType(key, x, y + 1, constant.TYPE_1_2)
}

func SetKeyPosValue1_1(key *uint64, x, y uint8) {
	SetKeyType(key, x, y, constant.TYPE_1_1)
}

func GetStep(step uint8) (uint8, uint8) {
	direction := step >> constant.BIT_NAME
	nameIndex := ((uint8(1) << constant.BIT_NAME) - uint8(1)) & step
	return nameIndex, direction
}

func GetPersonCNMap() map[uint8]string {
	m := make(map[uint8]string)
	m[constant.ZHANG_FEI] = constant.ZF_CN_NAME
	m[constant.CAO_CAO] = constant.CC_CN_NAME
	m[constant.MA_CHAO] = constant.MC_CN_NAME
	m[constant.ZHAO_YUN] = constant.ZY_CN_NAME
	m[constant.GUAN_YU] = constant.GY_CN_NAME
	m[constant.HUANG_ZHONG] = constant.HZ_CN_NAME
	m[constant.BING1] = constant.B1_CN_NAME
	m[constant.BING2] = constant.B2_CN_NAME
	m[constant.BING3] = constant.B3_CN_NAME
	m[constant.BING4] = constant.B4_CN_NAME
	return m
}

func GetDirectionCNMap() map[uint8]string {
	m := make(map[uint8]string)
	m[constant.UP] = constant.UP_CN_NAME
	m[constant.RIGHT] = constant.RIGHT_CN_NAME
	m[constant.DOWN] = constant.DOWN_CN_NAME
	m[constant.LEFT] = constant.LEFT_CN_NAME
	return m
}

func GetPersonENMap() map[uint8]string {
	m := make(map[uint8]string)
	m[constant.ZHANG_FEI] = constant.ZF_EN_NAME
	m[constant.CAO_CAO] = constant.CC_EN_NAME
	m[constant.MA_CHAO] = constant.MC_EN_NAME
	m[constant.ZHAO_YUN] = constant.ZY_EN_NAME
	m[constant.GUAN_YU] = constant.GY_EN_NAME
	m[constant.HUANG_ZHONG] = constant.HZ_EN_NAME
	m[constant.BING1] = constant.B1_EN_NAME
	m[constant.BING2] = constant.B2_EN_NAME
	m[constant.BING3] = constant.B3_EN_NAME
	m[constant.BING4] = constant.B4_EN_NAME
	return m
}

func GetDirectionENMap() map[uint8]string {
	m := make(map[uint8]string)
	m[constant.UP] = constant.UP_EN_NAME
	m[constant.RIGHT] = constant.RIGHT_EN_NAME
	m[constant.DOWN] = constant.DOWN_EN_NAME
	m[constant.LEFT] = constant.LEFT_EN_NAME
	return m
}

func PrettyFormat(b []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")
	if err != nil {
		fmt.Print(err.Error())
		fmt.Print(string(b))
	}
	return string(out.Bytes())
}