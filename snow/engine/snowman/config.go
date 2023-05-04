// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/memeticofficial/pepecoingo/snow"
	"github.com/memeticofficial/pepecoingo/snow/consensus/snowball"
	"github.com/memeticofficial/pepecoingo/snow/consensus/snowman"
	"github.com/memeticofficial/pepecoingo/snow/engine/common"
	"github.com/memeticofficial/pepecoingo/snow/engine/snowman/block"
	"github.com/memeticofficial/pepecoingo/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
