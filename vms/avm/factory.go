// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/memeticofficial/pepecoingo/utils/logging"
	"github.com/memeticofficial/pepecoingo/vms"
	"github.com/memeticofficial/pepecoingo/vms/avm/config"
)

var _ vms.Factory = (*Factory)(nil)

type Factory struct {
	config.Config
}

func (f *Factory) New(logging.Logger) (interface{}, error) {
	return &VM{Config: f.Config}, nil
}
