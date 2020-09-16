package mBlock

import (
	"github.com/whenfitrd/KServer/minterface"
)

type Block struct {
	active bool
}

func (block *Block) Active() {
	logger.Info("Active block...")
	block.active = true
}

func (block *Block) Freeze() {
	logger.Info("Freeze block...")
	block.active = false
}

func (block *Block) GetBlock() minterface.IBlock {
	return block
}
