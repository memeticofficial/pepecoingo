// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"fmt"

	"github.com/memeticofficial/pepecoingo/codec"
	"github.com/memeticofficial/pepecoingo/ids"
	"github.com/memeticofficial/pepecoingo/snow"
	"github.com/memeticofficial/pepecoingo/utils/crypto/secp256k1"
	"github.com/memeticofficial/pepecoingo/utils/hashing"
	"github.com/memeticofficial/pepecoingo/utils/set"
	"github.com/memeticofficial/pepecoingo/vms/avm/fxs"
	"github.com/memeticofficial/pepecoingo/vms/components/avax"
	"github.com/memeticofficial/pepecoingo/vms/nftfx"
	"github.com/memeticofficial/pepecoingo/vms/propertyfx"
	"github.com/memeticofficial/pepecoingo/vms/secp256k1fx"
)

type UnsignedTx interface {
	snow.ContextInitializable

	SetBytes(unsignedBytes []byte)
	Bytes() []byte

	InputIDs() set.Set[ids.ID]

	ConsumedAssetIDs() set.Set[ids.ID]
	AssetIDs() set.Set[ids.ID]

	NumCredentials() int
	// TODO: deprecate after x-chain linearization
	InputUTXOs() []*avax.UTXOID

	// Visit calls [visitor] with this transaction's concrete type
	Visit(visitor Visitor) error
}

// Tx is the core operation that can be performed. The tx uses the UTXO model.
// Specifically, a txs inputs will consume previous txs outputs. A tx will be
// valid if the inputs have the authority to consume the outputs they are
// attempting to consume and the inputs consume sufficient state to produce the
// outputs.
type Tx struct {
	Unsigned UnsignedTx          `serialize:"true" json:"unsignedTx"`
	Creds    []*fxs.FxCredential `serialize:"true" json:"credentials"` // The credentials of this transaction

	TxID  ids.ID `json:"id"`
	bytes []byte
}

func (t *Tx) Initialize(c codec.Manager) error {
	signedBytes, err := c.Marshal(CodecVersion, t)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}

	unsignedBytesLen, err := c.Size(CodecVersion, &t.Unsigned)
	if err != nil {
		return fmt.Errorf("couldn't calculate UnsignedTx marshal length: %w", err)
	}

	unsignedBytes := signedBytes[:unsignedBytesLen]
	t.SetBytes(unsignedBytes, signedBytes)
	return nil
}

func (t *Tx) SetBytes(unsignedBytes, signedBytes []byte) {
	t.TxID = hashing.ComputeHash256Array(signedBytes)
	t.bytes = signedBytes
	t.Unsigned.SetBytes(unsignedBytes)
}

// ID returns the unique ID of this tx
func (t *Tx) ID() ids.ID {
	return t.TxID
}

// Bytes returns the binary representation of this tx
func (t *Tx) Bytes() []byte {
	return t.bytes
}

// UTXOs returns the UTXOs transaction is producing.
func (t *Tx) UTXOs() []*avax.UTXO {
	u := utxoGetter{tx: t}
	// The visit error is explicitly dropped here because no error is ever
	// returned from the utxoGetter.
	_ = t.Unsigned.Visit(&u)
	return u.utxos
}

func (t *Tx) SignSECP256K1Fx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	unsignedBytes, err := c.Marshal(CodecVersion, &t.Unsigned)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}

	hash := hashing.ComputeHash256(unsignedBytes)
	for _, keys := range signers {
		cred := &secp256k1fx.Credential{
			Sigs: make([][secp256k1.SignatureLen]byte, len(keys)),
		}
		for i, key := range keys {
			sig, err := key.SignHash(hash)
			if err != nil {
				return fmt.Errorf("problem creating transaction: %w", err)
			}
			copy(cred.Sigs[i][:], sig)
		}
		t.Creds = append(t.Creds, &fxs.FxCredential{Verifiable: cred})
	}

	signedBytes, err := c.Marshal(CodecVersion, t)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}
	t.SetBytes(unsignedBytes, signedBytes)
	return nil
}

func (t *Tx) SignPropertyFx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	unsignedBytes, err := c.Marshal(CodecVersion, &t.Unsigned)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}

	hash := hashing.ComputeHash256(unsignedBytes)
	for _, keys := range signers {
		cred := &propertyfx.Credential{Credential: secp256k1fx.Credential{
			Sigs: make([][secp256k1.SignatureLen]byte, len(keys)),
		}}
		for i, key := range keys {
			sig, err := key.SignHash(hash)
			if err != nil {
				return fmt.Errorf("problem creating transaction: %w", err)
			}
			copy(cred.Sigs[i][:], sig)
		}
		t.Creds = append(t.Creds, &fxs.FxCredential{Verifiable: cred})
	}

	signedBytes, err := c.Marshal(CodecVersion, t)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}
	t.SetBytes(unsignedBytes, signedBytes)
	return nil
}

func (t *Tx) SignNFTFx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	unsignedBytes, err := c.Marshal(CodecVersion, &t.Unsigned)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}

	hash := hashing.ComputeHash256(unsignedBytes)
	for _, keys := range signers {
		cred := &nftfx.Credential{Credential: secp256k1fx.Credential{
			Sigs: make([][secp256k1.SignatureLen]byte, len(keys)),
		}}
		for i, key := range keys {
			sig, err := key.SignHash(hash)
			if err != nil {
				return fmt.Errorf("problem creating transaction: %w", err)
			}
			copy(cred.Sigs[i][:], sig)
		}
		t.Creds = append(t.Creds, &fxs.FxCredential{Verifiable: cred})
	}

	signedBytes, err := c.Marshal(CodecVersion, t)
	if err != nil {
		return fmt.Errorf("problem creating transaction: %w", err)
	}
	t.SetBytes(unsignedBytes, signedBytes)
	return nil
}
