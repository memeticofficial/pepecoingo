// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"context"
	"errors"
	"testing"

	"github.com/memeticofficial/pepecoingo/ids"
	"github.com/memeticofficial/pepecoingo/snow/consensus/pepecoin"
)

var (
	errBuild = errors.New("unexpectedly called Build")

	_ Builder = (*TestBuilder)(nil)
)

type TestBuilder struct {
	T             *testing.T
	CantBuildVtx  bool
	BuildStopVtxF func(ctx context.Context, parentIDs []ids.ID) (pepecoin.Vertex, error)
}

func (b *TestBuilder) Default(cant bool) {
	b.CantBuildVtx = cant
}

func (b *TestBuilder) BuildStopVtx(ctx context.Context, parentIDs []ids.ID) (pepecoin.Vertex, error) {
	if b.BuildStopVtxF != nil {
		return b.BuildStopVtxF(ctx, parentIDs)
	}
	if b.CantBuildVtx && b.T != nil {
		b.T.Fatal(errBuild)
	}
	return nil, errBuild
}
