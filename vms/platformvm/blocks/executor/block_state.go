// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"time"

	"github.com/memeticofficial/pepecoingo/chains/atomic"
	"github.com/memeticofficial/pepecoingo/ids"
	"github.com/memeticofficial/pepecoingo/utils/set"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/blocks"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/state"
)

type standardBlockState struct {
	onAcceptFunc func()
	inputs       set.Set[ids.ID]
}

type proposalBlockState struct {
	initiallyPreferCommit bool
	onCommitState         state.Diff
	onAbortState          state.Diff
}

// The state of a block.
// Note that not all fields will be set for a given block.
type blockState struct {
	standardBlockState
	proposalBlockState
	statelessBlock blocks.Block
	onAcceptState  state.Diff

	timestamp      time.Time
	atomicRequests map[ids.ID]*atomic.Requests
}
