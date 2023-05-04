// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/memeticofficial/pepecoingo/snow"
	"github.com/memeticofficial/pepecoingo/snow/uptime"
	"github.com/memeticofficial/pepecoingo/utils"
	"github.com/memeticofficial/pepecoingo/utils/timer/mockable"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/config"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/fx"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/reward"
	"github.com/memeticofficial/pepecoingo/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Manager
	Rewards      reward.Calculator
	Bootstrapped *utils.Atomic[bool]
}
