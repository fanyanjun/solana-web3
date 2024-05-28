// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package group_member_pointer

import (
	"errors"
	common "github.com/donutnomad/solana-web3/common"
	binary "github.com/gagliardetto/binary"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Initialize Instruction
type Initialize struct {
	// The public key for the account that can update the group address
	Authority *common.PublicKey
	// The account address that holds the member
	MemberAddress *common.PublicKey
	// [0] = [WRITE] mint `The mint to initialize.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	return &Initialize{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewInitializeInstruction
//
// Parameters:
//
//	authority: The public key for the account that can update the group address
//	memberAddress: The account address that holds the member
//	mint: The mint to initialize.
func NewInitializeInstruction(
	authority common.PublicKey,
	memberAddress common.PublicKey,
	mint common.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetAuthority(authority).
		SetMemberAddress(memberAddress).
		SetMintAccount(mint)
}

// SetAuthority sets the "authority" parameter.
func (obj *Initialize) SetAuthority(authority common.PublicKey) *Initialize {
	obj.Authority = &authority
	return obj
}

// SetMemberAddress sets the "memberAddress" parameter.
func (obj *Initialize) SetMemberAddress(memberAddress common.PublicKey) *Initialize {
	obj.MemberAddress = &memberAddress
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint to initialize.
func (obj *Initialize) SetMintAccount(mint common.PublicKey, multiSigners ...common.PublicKey) *Initialize {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[0] = common.Meta(mint)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	}
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint to initialize.
func (obj *Initialize) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *Initialize) SetProgramId(programId *common.PublicKey) *Initialize {
	obj._programId = programId
	return obj
}

func (obj *Initialize) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Initialize}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Initialize) Validate() error {
	if obj.Authority == nil {
		return errors.New("[Initialize] authority param is not set")
	}
	if obj.MemberAddress == nil {
		return errors.New("[Initialize] memberAddress param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Initialize] accounts.mint is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Initialize) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Authority); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.MemberAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Authority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MemberAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Initialize")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("    Authority", *obj.Authority))
						paramsBranch.Child(format.Param("MemberAddress", *obj.MemberAddress))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("mint", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}

// Update Instruction
type Update struct {
	// The new account address that holds the group
	MemberAddress *common.PublicKey
	// [0] = [WRITE] mint `The mint.`
	// [1] = [SIGNER] authority `The group member pointer authority.,..2+M `[signer]` M signer accounts.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewUpdateInstructionBuilder creates a new `Update` instruction builder.
func NewUpdateInstructionBuilder() *Update {
	return &Update{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewUpdateInstruction
//
// Parameters:
//   memberAddress: The new account address that holds the group
//   mint: The mint.
/*
  authority: The group member pointer authority.
  ..2+M `[signer]` M signer accounts.
*/
//
func NewUpdateInstruction(
	memberAddress common.PublicKey,
	mint common.PublicKey,
	authority common.PublicKey,
) *Update {
	return NewUpdateInstructionBuilder().
		SetMemberAddress(memberAddress).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}

// SetMemberAddress sets the "memberAddress" parameter.
func (obj *Update) SetMemberAddress(memberAddress common.PublicKey) *Update {
	obj.MemberAddress = &memberAddress
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint.
func (obj *Update) SetMintAccount(mint common.PublicKey) *Update {
	obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint.
func (obj *Update) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" parameter.
// The group member pointer authority.
// ..2+M `[signer]` M signer accounts.
func (obj *Update) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *Update {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(authority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(authority).SIGNER()
	}
	return obj
}

// GetAuthorityAccount gets the "authority" parameter.
// The group member pointer authority.
// ..2+M `[signer]` M signer accounts.
func (obj *Update) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *Update) SetProgramId(programId *common.PublicKey) *Update {
	obj._programId = programId
	return obj
}

func (obj *Update) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Update}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Update) Validate() error {
	if obj.MemberAddress == nil {
		return errors.New("[Update] memberAddress param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Update] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Update] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Update) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Update) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.MemberAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Update) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.MemberAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Update) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Update")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("MemberAddress", *obj.MemberAddress))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("authority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}
