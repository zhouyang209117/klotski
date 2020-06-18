package controller

import (
	"klotski/model"
	"klotski/constant"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/Workiva/go-datastructures/set"
	"fmt"
	"container/list"
	log2 "log"
)

func select_append(states []model.State, new_states []model.State) []model.State {
	if len(new_states) > 0 {
		states = append(states, new_states...)
	}
	return states
}

func nextStates(state model.State) []model.State {
	states := make([]model.State, 0)
	states = select_append(states, Pos2_2Next(state, constant.CAO_CAO))
	states = select_append(states, Pos2_1Next(state, constant.ZHANG_FEI))
	states = select_append(states, Pos2_1Next(state, constant.MA_CHAO))
	states = select_append(states, Pos2_1Next(state, constant.ZHAO_YUN))
	states = select_append(states, Pos2_1Next(state, constant.HUANG_ZHONG))
	states = select_append(states, Pos1_2Next(state, constant.GUAN_YU))
	states = select_append(states, Pos1_1Next(state, constant.BING1))
	states = select_append(states, Pos1_1Next(state, constant.BING2))
	states = select_append(states, Pos1_1Next(state, constant.BING3))
	states = select_append(states, Pos1_1Next(state, constant.BING4))
	return states
}

func success(state model.State) bool{
	x, y := state.GetCoord(constant.CAO_CAO)
	return x == 3 && y == 1
}

func Klotski(s model.State) (bool, *list.List) {
	steps := list.New()
	q := queue.New(10)
	if success(s) {
		fmt.Println("成功")
		steps.PushFront(s.Step)
		return true, steps
	}
	visited := set.New()
	visited.Add(s.GetKey())
	q.Put(s)
	finish := false
	finalState := model.State{}
	try_state_num := 1
	diff_state_num := 1
	for !q.Empty() && !finish {
		a, err := q.Get(1)
		if err != nil {
			panic("Get queue element error")
		}
		currentS := (a[0]).(model.State)
		newStates := nextStates(currentS)
		for _, newS := range newStates {
			if success(newS) {
				finalState = newS
				finish = true
				break
			}
			key := newS.GetKey()
			if !visited.Exists(key) {
				q.Put(newS)
				visited.Add(key)
				diff_state_num += 1
			}
			try_state_num += 1
		}
	}
	log2.Println(fmt.Sprintf("Total:%d states are tested.", try_state_num))
	log2.Println(fmt.Sprintf("Total:%d states are tested.", diff_state_num))
	if !finish {
		fmt.Println("无解")
		return false, steps
	}
	steps.PushFront(finalState.Step)
	for finalState.Pre != nil {
		steps.PushFront(finalState.Pre.Step)
		finalState = *(finalState.Pre)
	}
	return true, steps
}

