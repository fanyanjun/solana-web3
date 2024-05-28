// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package spl_token_2022

import (
	common "github.com/donutnomad/solana-web3/common"
	binary "github.com/gagliardetto/binary"
)

// Mint Struct
// Mint data.
type Mint struct {
	// Optional authority used to mint new tokens. The mint authority may only
	// be provided during mint creation. If no mint authority is present
	// then the mint has a fixed supply and no further tokens may be
	// minted.
	MintAuthority *common.PublicKey `bin:"optional"`
	// Total supply of tokens.
	Supply uint64
	// Number of base 10 digits to the right of the decimal place.
	Decimals uint8
	// Is `true` if this structure has been initialized
	IsInitialized bool
	// Optional authority to freeze token accounts.
	FreezeAuthority *common.PublicKey `bin:"optional"`
}

const MINT_SIZE = 82

func (obj *Mint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.WriteUint32(btou32(obj.MintAuthority != nil), binary.LE); err != nil {
		return err
	}
	if obj.MintAuthority != nil {
		if err = encoder.Encode(obj.MintAuthority); err != nil {
			return err
		}
	} else {
		var tmp common.PublicKey
		if err = encoder.Encode(tmp); err != nil {
			return err
		}
	}
	if err = encoder.Encode(&obj.Supply); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Decimals); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.IsInitialized); err != nil {
		return err
	}
	if err = encoder.WriteUint32(btou32(obj.FreezeAuthority != nil), binary.LE); err != nil {
		return err
	}
	if obj.FreezeAuthority != nil {
		if err = encoder.Encode(obj.FreezeAuthority); err != nil {
			return err
		}
	} else {
		var tmp common.PublicKey
		if err = encoder.Encode(tmp); err != nil {
			return err
		}
	}
	return nil
}

func (obj *Mint) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if v, err := decoder.ReadBytes(4); err != nil {
		return err
	} else {
		if err = decoder.Decode(&obj.MintAuthority); err != nil {
			return err
		}
		if v[0] == 0 {
			obj.MintAuthority = nil
		}
	}
	if err = decoder.Decode(&obj.Supply); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Decimals); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.IsInitialized); err != nil {
		return err
	}
	if v, err := decoder.ReadBytes(4); err != nil {
		return err
	} else {
		if err = decoder.Decode(&obj.FreezeAuthority); err != nil {
			return err
		}
		if v[0] == 0 {
			obj.FreezeAuthority = nil
		}
	}
	return nil
}

// Account Struct
// Account data.
type Account struct {
	// The mint associated with this account
	Mint common.PublicKey
	// The owner of this account.
	Owner common.PublicKey
	// The amount of tokens this account holds.
	Amount uint64
	// If `delegate` is `Some` then `delegated_amount` represents
	// the amount authorized by the delegate
	Delegate *common.PublicKey `bin:"optional"`
	// The account's state
	State AccountState
	// If is_some, this is a native token, and the value logs the rent-exempt
	// reserve. An Account is required to be rent-exempt, so the value is
	// used by the Processor to ensure that wrapped SOL accounts do not
	// drop below this threshold.
	IsNative *uint64 `bin:"optional"`
	// The amount delegated
	DelegatedAmount uint64
	// Optional authority to close the account.
	CloseAuthority *common.PublicKey `bin:"optional"`
}

const ACCOUNT_SIZE = 165

func (obj *Account) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Mint); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Owner); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Amount); err != nil {
		return err
	}
	if err = encoder.WriteUint32(btou32(obj.Delegate != nil), binary.LE); err != nil {
		return err
	}
	if obj.Delegate != nil {
		if err = encoder.Encode(obj.Delegate); err != nil {
			return err
		}
	} else {
		var tmp common.PublicKey
		if err = encoder.Encode(tmp); err != nil {
			return err
		}
	}
	if err = encoder.Encode(&obj.State); err != nil {
		return err
	}
	if err = encoder.WriteUint32(btou32(obj.IsNative != nil), binary.LE); err != nil {
		return err
	}
	if obj.IsNative != nil {
		if err = encoder.Encode(obj.IsNative); err != nil {
			return err
		}
	} else {
		var tmp uint64
		if err = encoder.Encode(tmp); err != nil {
			return err
		}
	}
	if err = encoder.Encode(&obj.DelegatedAmount); err != nil {
		return err
	}
	if err = encoder.WriteUint32(btou32(obj.CloseAuthority != nil), binary.LE); err != nil {
		return err
	}
	if obj.CloseAuthority != nil {
		if err = encoder.Encode(obj.CloseAuthority); err != nil {
			return err
		}
	} else {
		var tmp common.PublicKey
		if err = encoder.Encode(tmp); err != nil {
			return err
		}
	}
	return nil
}

func (obj *Account) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Mint); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Owner); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Amount); err != nil {
		return err
	}
	if v, err := decoder.ReadBytes(4); err != nil {
		return err
	} else {
		if err = decoder.Decode(&obj.Delegate); err != nil {
			return err
		}
		if v[0] == 0 {
			obj.Delegate = nil
		}
	}
	if err = decoder.Decode(&obj.State); err != nil {
		return err
	}
	if v, err := decoder.ReadBytes(4); err != nil {
		return err
	} else {
		if err = decoder.Decode(&obj.IsNative); err != nil {
			return err
		}
		if v[0] == 0 {
			obj.IsNative = nil
		}
	}
	if err = decoder.Decode(&obj.DelegatedAmount); err != nil {
		return err
	}
	if v, err := decoder.ReadBytes(4); err != nil {
		return err
	} else {
		if err = decoder.Decode(&obj.CloseAuthority); err != nil {
			return err
		}
		if v[0] == 0 {
			obj.CloseAuthority = nil
		}
	}
	return nil
}

// Multisig Struct
// Multisignature data.
type Multisig struct {
	// Number of signers required
	M uint8
	// Number of valid signers
	N uint8
	// Is `true` if this structure has been initialized
	IsInitialized bool
	// Signer public keys
	Signers [11]common.PublicKey
}

func (obj *Multisig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.M); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.N); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.IsInitialized); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Signers); err != nil {
		return err
	}
	return nil
}

func (obj *Multisig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.M); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.N); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.IsInitialized); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Signers); err != nil {
		return err
	}
	return nil
}
