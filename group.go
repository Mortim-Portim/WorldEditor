package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/mortim-portim/GraphEng/GE"
)

type UpdateAble interface {
	GE.UpdateAble
	Update(frame int)
	Draw(screen *ebiten.Image)
}

func GetGroup() *Group {
	return &Group{make([]UpdateAble, 0)}
}

type Group struct {
	Members []UpdateAble
}

func (g *Group) Update(frame int) {
	for _, member := range g.Members {
		member.Update(frame)
	}
}

func (g *Group) Draw(screen *ebiten.Image) {
	for _, member := range g.Members {
		member.Draw(screen)
	}
}

func (g *Group) Add(uas ...UpdateAble) {
	g.Members = append(g.Members, uas...)
}
