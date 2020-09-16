package kComponent

import (
	"github.com/whenfitrd/KServer/kInterface"
)

type Component struct {
	active bool
}

func (component *Component) Active() {
	logger.Info("Active component...")
	component.active = true
}

func (component *Component) Freeze() {
	logger.Info("Freeze component...")
	component.active = false
}

func (component *Component) GetComponent() kInterface.IComponent {
	return component
}