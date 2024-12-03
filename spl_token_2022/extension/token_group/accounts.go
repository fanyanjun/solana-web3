// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package token_group

import (
	"bytes"
	"fmt"
	common "github.com/fanyanjun/solana-web3/common"
	binary "github.com/gagliardetto/binary"
)

// TokenGroup Struct
// Data struct for a `TokenGroup`
// DETERMINANT: spl_token_group_interface:group
type TokenGroup struct {
	// The authority that can sign to update the group
	UpdateAuthority common.PublicKey
	// The associated mint, used to counter spoofing to be sure that group
	// belongs to a particular mint
	Mint common.PublicKey
	// The current number of group members
	Size uint32
	// The maximum number of group members
	MaxSize uint32
}

const TOKEN_GROUP_SIZE = 72

// TokenGroupDiscriminator DETERMINANT: spl_token_group_interface:group
var TokenGroupDiscriminator = [8]byte{214, 15, 63, 132, 49, 119, 209, 40}

func (obj *TokenGroup) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenGroupDiscriminator[:], false)
	if err != nil {
		return err
	}
	if err = encoder.Encode(&obj.UpdateAuthority); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Mint); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Size); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.MaxSize); err != nil {
		return err
	}
	return nil
}

func (obj *TokenGroup) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadNBytes(8)
		if err != nil {
			return err
		}
		if !bytes.Equal(discriminator, TokenGroupDiscriminator[:]) {
			return fmt.Errorf("wrong discriminator: wanted %s, got %s", fmt.Sprint(TokenGroupDiscriminator[:]), fmt.Sprint(discriminator[:]))
		}
	}
	if err = decoder.Decode(&obj.UpdateAuthority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Mint); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Size); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MaxSize); err != nil {
		return err
	}
	return nil
}

// TokenGroupMember Struct
// Data struct for a `TokenGroupMember`
// DETERMINANT: spl_token_group_interface:member
type TokenGroupMember struct {
	// The associated mint, used to counter spoofing to be sure that member
	// belongs to a particular mint
	Mint common.PublicKey
	// The pubkey of the `TokenGroup`
	Group common.PublicKey
	// The member number
	MemberNumber uint32
}

const TOKEN_GROUP_MEMBER_SIZE = 68

// TokenGroupMemberDiscriminator DETERMINANT: spl_token_group_interface:member
var TokenGroupMemberDiscriminator = [8]byte{254, 50, 168, 134, 88, 126, 100, 186}

func (obj *TokenGroupMember) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenGroupMemberDiscriminator[:], false)
	if err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Mint); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Group); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.MemberNumber); err != nil {
		return err
	}
	return nil
}

func (obj *TokenGroupMember) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadNBytes(8)
		if err != nil {
			return err
		}
		if !bytes.Equal(discriminator, TokenGroupMemberDiscriminator[:]) {
			return fmt.Errorf("wrong discriminator: wanted %s, got %s", fmt.Sprint(TokenGroupMemberDiscriminator[:]), fmt.Sprint(discriminator[:]))
		}
	}
	if err = decoder.Decode(&obj.Mint); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Group); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MemberNumber); err != nil {
		return err
	}
	return nil
}
