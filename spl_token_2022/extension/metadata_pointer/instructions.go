// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package metadata_pointer

import (
	"errors"
	common "github.com/donutnomad/solana-web3/common"
	binary "github.com/gagliardetto/binary"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Initialize Instruction
type Initialize struct {
	Authority       *common.PublicKey
	MetadataAddress *common.PublicKey
	// [0] = [WRITE] mint `The mint to initialize`
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
//	authority:
//	metadataAddress:
//	mint: The mint to initialize
func NewInitializeInstruction(
	authority common.PublicKey,
	metadataAddress common.PublicKey,
	mint common.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetAuthority(authority).
		SetMetadataAddress(metadataAddress).
		SetMintAccount(mint)
}

// SetAuthority sets the "authority" parameter.
func (obj *Initialize) SetAuthority(authority common.PublicKey) *Initialize {
	obj.Authority = &authority
	return obj
}

// SetMetadataAddress sets the "metadataAddress" parameter.
func (obj *Initialize) SetMetadataAddress(metadataAddress common.PublicKey) *Initialize {
	obj.MetadataAddress = &metadataAddress
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint to initialize
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
// The mint to initialize
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
	if obj.MetadataAddress == nil {
		return errors.New("[Initialize] metadataAddress param is not set")
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
	if err = encoder.Encode(&obj.MetadataAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Authority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MetadataAddress); err != nil {
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
						paramsBranch.Child(format.Param("      Authority", *obj.Authority))
						paramsBranch.Child(format.Param("MetadataAddress", *obj.MetadataAddress))
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
	MetadataAddress *common.PublicKey
	// [0] = [WRITE] mint `The mint.`
	// [1] = [SIGNER] owner `The metadata pointer authority.`
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
//
//	metadataAddress:
//	mint: The mint.
//	owner: The metadata pointer authority.
func NewUpdateInstruction(
	metadataAddress common.PublicKey,
	mint common.PublicKey,
	owner common.PublicKey,
) *Update {
	return NewUpdateInstructionBuilder().
		SetMetadataAddress(metadataAddress).
		SetMintAccount(mint).
		SetOwnerAccount(owner)
}

// SetMetadataAddress sets the "metadataAddress" parameter.
func (obj *Update) SetMetadataAddress(metadataAddress common.PublicKey) *Update {
	obj.MetadataAddress = &metadataAddress
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

// SetOwnerAccount sets the "owner" parameter.
// The metadata pointer authority.
func (obj *Update) SetOwnerAccount(owner common.PublicKey, multiSigners ...common.PublicKey) *Update {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(owner)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(owner).SIGNER()
	}
	return obj
}

// GetOwnerAccount gets the "owner" parameter.
// The metadata pointer authority.
func (obj *Update) GetOwnerAccount() *common.AccountMeta {
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
	if obj.MetadataAddress == nil {
		return errors.New("[Update] metadataAddress param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Update] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Update] accounts.owner is not set")
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
	if err = encoder.Encode(&obj.MetadataAddress); err != nil {
		return err
	}
	return nil
}

func (obj *Update) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.MetadataAddress); err != nil {
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
						paramsBranch.Child(format.Param("MetadataAddress", *obj.MetadataAddress))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta(" mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("owner", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}