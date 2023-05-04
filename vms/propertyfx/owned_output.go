// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"github.com/memeticofficial/pepecoingo/vms/secp256k1fx"
)

type OwnedOutput struct {
	secp256k1fx.OutputOwners `serialize:"true"`
}
