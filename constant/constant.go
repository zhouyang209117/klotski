package constant

const (
	ZHANG_FEI = 0
	CAO_CAO = 1
	MA_CHAO = 2
	ZHAO_YUN = 3
	GUAN_YU = 4
	HUANG_ZHONG = 5
	BING1 = 6
	BING2 = 7
	BING3 = 8
	BING4 = 9
	FREE1 = 10
	FREE2 = 11

	ROW = 5
	COL = 4

	UP = 0
	RIGHT = 1
	DOWN = 2
	LEFT = 3

	BIT_X = 3
	BIT_Y = 2
	BIT_SUM = 60

	BIT_NAME = 4
	BIT_DIRECTION = 2
/*
 算key的时候会用到,注意虽然有4种棋子,但是要用3bit表示一种棋子
 因为0是默认值,可以认为是free
*/
	TYPE_1_1 = 1
	TYPE_1_2 = 2
	TYPE_2_1 = 3
	TYPE_2_2 = 4
	BIT_TYPE = 3
	BIT_SUM_TYPE = BIT_TYPE * COL * ROW

	ZF_EN_NAME = "zhangfei"
	CC_EN_NAME = "caocao"
	MC_EN_NAME = "machao"
	ZY_EN_NAME = "zhaoyun"
	GY_EN_NAME = "guoyu"
	HZ_EN_NAME = "huangzhong"
	B1_EN_NAME = "bing1"
	B2_EN_NAME = "bing2"
	B3_EN_NAME = "bing3"
	B4_EN_NAME = "bing4"

	ZF_CN_NAME = "张飞"
	CC_CN_NAME = "曹操"
	MC_CN_NAME = "马超"
	ZY_CN_NAME = "赵云"
	GY_CN_NAME = "关羽"
	HZ_CN_NAME = "黄忠"
	B1_CN_NAME = "兵1"
	B2_CN_NAME = "兵2"
	B3_CN_NAME = "兵3"
	B4_CN_NAME = "兵4"

	UP_EN_NAME = "up"
	LEFT_EN_NAME = "left"
	DOWN_EN_NAME = "down"
	RIGHT_EN_NAME = "right"

	UP_CN_NAME = "上"
	LEFT_CN_NAME = "左"
	DOWN_CN_NAME = "下"
	RIGHT_CN_NAME = "右"

)
