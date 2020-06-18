package controller

import (
	"klotski/model"
	"klotski/constant"
)

func posFree2(d1, d2 model.Coord, s model.State) bool {
	f1_x, f1_y := s.GetCoord(constant.FREE1)
	f2_x, f2_y := s.GetCoord(constant.FREE2)
	return (d1.X == f1_x && d1.Y == f1_y && d2.X == f2_x && d2.Y == f2_y) ||
	(d1.X == f2_x && d1.Y == f2_y && d2.X == f1_x && d2.Y == f1_y)
}

func posFree1(dest model.Coord, s model.State) (bool, uint8){
	f1_x, f1_y := s.GetCoord(constant.FREE1)
	f2_x, f2_y := s.GetCoord(constant.FREE2)
	if dest.X == f1_x && dest.Y == f1_y {
		return true, constant.FREE1
	}
	if dest.X == f2_x && dest.Y == f2_y {
		return true, constant.FREE2
	}
	return false, 0
}

func Pos2_1Next(s model.State, nameIndex uint8) []model.State {
	states := make([]model.State, 0)
	x, y := s.GetCoord(nameIndex)
	if x > 0 {
		r, freeIndex := posFree1(model.Coord{X: x - 1, Y: y}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x - 1, y)
			newS.SetCoord(freeIndex, x + 1, y)
			newS.SetStep(nameIndex, constant.UP)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if y + 1 < constant.COL {
		r := posFree2(model.Coord{X: x, Y: y + 1}, model.Coord{X: x + 1, Y: y + 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y + 1)
			newS.SetCoord(constant.FREE1, x, y)
			newS.SetCoord(constant.FREE2, x + 1, y)
			newS.SetStep(nameIndex, constant.RIGHT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if x + 2 < constant.ROW && len(states) < 2 {
		r, free_index := posFree1(model.Coord{X: x + 2, Y: y}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x + 1, y)
			newS.SetCoord(free_index, x, y)
			newS.SetStep(nameIndex, constant.DOWN)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if y > 0 && len(states) < 2 {
		r := posFree2(model.Coord{X: x, Y: y - 1}, model.Coord{X: x + 1, Y: y - 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y - 1)
			newS.SetCoord(constant.FREE1, x, y)
			newS.SetCoord(constant.FREE2, x + 1, y)
			newS.SetStep(nameIndex, constant.LEFT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	return states
}

func Pos2_2Next(s model.State, nameIndex uint8) []model.State {
	states := make([]model.State, 0)
	x, y := s.GetCoord(nameIndex)
	new_state := model.State{ChessPos: s.ChessPos}
	update := false
	if x > 0 {
		r := posFree2(model.Coord{X: x - 1, Y: y}, model.Coord{X: x - 1, Y: y + 1}, s)
		if r {
			new_state.SetCoord(nameIndex, x - 1, y)
			new_state.SetCoord(constant.FREE1, x + 1, y)
			new_state.SetCoord(constant.FREE2, x + 1, y + 1)
			new_state.SetStep(nameIndex, constant.UP)
			update = true
		}
	}
	if y + 2 < constant.COL && !update {
		r := posFree2(model.Coord{X: x, Y: y + 2}, model.Coord{X: x + 1, Y: y + 2}, s)
		if r {
			new_state.SetCoord(nameIndex, x, y + 1)
			new_state.SetCoord(constant.FREE1, x, y)
			new_state.SetCoord(constant.FREE2, x + 1, y)
			new_state.SetStep(nameIndex, constant.RIGHT)
			update = true
		}
	}
	if x + 2 < constant.ROW && !update {
		r := posFree2(model.Coord{X: x + 2, Y: y}, model.Coord{X: x + 2, Y: y + 1}, s)
		if r {
			new_state.SetCoord(nameIndex, x + 1, y)
			new_state.SetCoord(constant.FREE1, x, y)
			new_state.SetCoord(constant.FREE2, x, y + 1)
			new_state.SetStep(nameIndex, constant.DOWN)
			update = true
		}
	}
	if y - 1 >= 0 && !update {
		r := posFree2(model.Coord{X: x, Y: y - 1}, model.Coord{X: x + 1, Y: y - 1}, s)
		if r {
			new_state.SetCoord(nameIndex, x, y - 1)
			new_state.SetCoord(constant.FREE1, x, y + 1)
			new_state.SetCoord(constant.FREE2, x + 1, y + 1)
			new_state.SetStep(nameIndex, constant.LEFT)
			update = true
		}
	}
	if update {
		new_state.Pre = &s
		states = append(states, new_state)
	}
	return states
}

func Pos1_2Next(s model.State, nameIndex uint8) []model.State {
	states := make([]model.State, 0)
	x, y := s.GetCoord(nameIndex)
	if x > 0 {
		r := posFree2(model.Coord{X: x - 1, Y: y}, model.Coord{X: x - 1, Y: y + 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x - 1, y)
			newS.SetCoord(constant.FREE1, x, y)
			newS.SetCoord(constant.FREE2, x, y + 1)
			newS.SetStep(nameIndex, constant.UP)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if y + 2 < constant.COL {
		r, free_index := posFree1(model.Coord{X: x, Y: y + 2}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y + 1)
			newS.SetCoord(free_index, x, y)
			newS.SetStep(nameIndex, constant.RIGHT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if x + 1 < constant.ROW && len(states) < 2 {
		r := posFree2(model.Coord{X: x + 1, Y: y}, model.Coord{X: x + 1, Y: y + 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x + 1, y)
			newS.SetCoord(constant.FREE1, x, y)
			newS.SetCoord(constant.FREE2, x, y + 1)
			newS.SetStep(nameIndex, constant.DOWN)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if y - 1 >= 0 && len(states) < 2 {
		r, free_index := posFree1(model.Coord{X: x, Y: y - 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y - 1)
			newS.SetCoord(free_index, x, y + 1)
			newS.SetStep(nameIndex, constant.LEFT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	return states
}

func Pos1_1Next(s model.State, nameIndex uint8) []model.State {
	states := make([]model.State, 0)
	x, y := s.GetCoord(nameIndex)
	if x > 0 {
		r, free_index := posFree1(model.Coord{X: x - 1, Y: y}, s)
		if r {
			newState := model.State{ChessPos: s.ChessPos}
			newState.SetCoord(nameIndex, x - 1, y)
			newState.SetCoord(free_index, x, y)
			newState.SetStep(nameIndex, constant.UP)
			newState.Pre = &s
			states = append(states, newState)
		}
	}
	if y + 1 < constant.COL {
		r, free_index := posFree1(model.Coord{X: x, Y: y + 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y + 1)
			newS.SetCoord(free_index, x, y)
			newS.SetStep(nameIndex, constant.RIGHT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if x + 1 < constant.ROW && len(states) < 2 {
		r, free_index := posFree1(model.Coord{ X: x + 1, Y: y}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x + 1, y)
			newS.SetCoord(free_index, x, y)
			newS.SetStep(nameIndex, constant.DOWN)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	if y - 1 >= 0 && len(states) < 2 {
		r, free_index := posFree1(model.Coord{X: x, Y: y - 1}, s)
		if r {
			newS := model.State{ChessPos: s.ChessPos}
			newS.SetCoord(nameIndex, x, y - 1)
			newS.SetCoord(free_index, x, y)
			newS.SetStep(nameIndex, constant.LEFT)
			newS.Pre = &s
			states = append(states, newS)
		}
	}
	return states
}
