package model

type StateKey struct {
	ChessPos uint64 //各个棋子的位置
	Step uint8       //哪个棋子怎么移动得到现在位置
	Pre *State
}
