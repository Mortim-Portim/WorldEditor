package wrldedit

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mortim-portim/GraphEng/GE"
)

type UpdateAble interface {
	GE.UpdateAble
	Update(frame int)
	Draw(screen *ebiten.Image)
}

func GetGroup(members ...UpdateAble) *Group {
	return &Group{members}
}

type Group struct {
	Members []UpdateAble
}

func (g *Group) Init(screen *ebiten.Image, data interface{}) (GE.UpdateFunc, GE.DrawFunc) {
	return g.Update, g.Draw
}

func (g *Group) Start(screen *ebiten.Image, data interface{}) {}
func (g *Group) Stop(screen *ebiten.Image, data interface{})  {}

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
