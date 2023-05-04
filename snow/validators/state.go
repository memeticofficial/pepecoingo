// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"context"
	"sync"

	"github.com/memeticofficial/pepecoingo/ids"
)

var _ State = (*lockedState)(nil)

// State allows the lookup of validator sets on specified subnets at the
// requested P-chain height.
type State interface {
	// GetMinimumHeight returns the minimum height of the block still in the
	// proposal window.
	GetMinimumHeight(context.Context) (uint64, error)
	// GetCurrentHeight returns the current height of the P-chain.
	GetCurrentHeight(context.Context) (uint64, error)

	// GetSubnetID returns the subnetID of the provided chain.
	GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error)

	// GetValidatorSet returns the validators of the provided subnet at the
	// requested P-chain height.
	// The returned map should not be modified.
	GetValidatorSet(
		ctx context.Context,
		height uint64,
		subnetID ids.ID,
	) (map[ids.NodeID]*GetValidatorOutput, error)
}

type lockedState struct {
	lock sync.Locker
	s    State
}

func NewLockedState(lock sync.Locker, s State) State {
	return &lockedState{
		lock: lock,
		s:    s,
	}
}

func (s *lockedState) GetMinimumHeight(ctx context.Context) (uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetMinimumHeight(ctx)
}

func (s *lockedState) GetCurrentHeight(ctx context.Context) (uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetCurrentHeight(ctx)
}

func (s *lockedState) GetSubnetID(ctx context.Context, chainID ids.ID) (ids.ID, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetSubnetID(ctx, chainID)
}

func (s *lockedState) GetValidatorSet(
	ctx context.Context,
	height uint64,
	subnetID ids.ID,
) (map[ids.NodeID]*GetValidatorOutput, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.s.GetValidatorSet(ctx, height, subnetID)
}

type noValidators struct {
	State
}

func NewNoValidatorsState(state State) State {
	return &noValidators{
		State: state,
	}
}

func (*noValidators) GetValidatorSet(context.Context, uint64, ids.ID) (map[ids.NodeID]*GetValidatorOutput, error) {
	return nil, nil
}
