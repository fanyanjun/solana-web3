// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package group_member_pointer

import (
	"bytes"
	"fmt"
	spew "github.com/davecgh/go-spew/spew"
	common "github.com/fanyanjun/solana-web3/common"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
	text "github.com/gagliardetto/solana-go/text"
	treeout "github.com/gagliardetto/treeout"
)

var ProgramID common.PublicKey = common.MustPublicKeyFromBase58("11111111111111111111111111111111")

func SetProgramID(pubkey common.PublicKey) {
	ProgramID = pubkey
	if !common.IsZero(ProgramID) {
		solanago.RegisterInstructionDecoder(common.As(ProgramID), registryDecodeInstruction)
	}
}

const ProgramName = "group_member_pointer"

func init() {
	if !common.IsZero(ProgramID) {
		solanago.RegisterInstructionDecoder(common.As(ProgramID), registryDecodeInstruction)
	}
}

func btou32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

var (
	Instruction_Initialize uint8 = 0
	Instruction_Update     uint8 = 1
)

var InstructionImplDef = binary.NewVariantDefinition(binary.Uint8TypeIDEncoding, []binary.VariantType{
	{
		"initialize", (*Initialize)(nil),
	},
	{
		"update", (*Update)(nil),
	},
})

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id uint8) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_Update:
		return "Update"
	default:
		return ""
	}
}

func registryDecodeInstruction(accounts []*solanago.AccountMeta, data []byte) (interface{}, error) {
	obj, err := DecodeInstruction(common.ConvertMeta(accounts), data)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func DecodeInstruction(accounts []*common.AccountMeta, data []byte) (*Instruction, error) {
	obj := new(Instruction)
	if err := binary.NewBorshDecoder(data).Decode(obj); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := obj.Impl.(common.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return obj, nil
}

type Instruction struct {
	binary.BaseVariant
	programId *common.PublicKey
	typeIdLen uint8
}

func (obj *Instruction) EncodeToTree(parent treeout.Branches) {
	if enToTree, ok := obj.Impl.(text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(spew.Sdump(obj))
	}
}

func (obj *Instruction) ProgramID() common.PublicKey {
	if obj.programId != nil {
		return *obj.programId
	}
	return ProgramID
}

func (obj *Instruction) Accounts() (out []*common.AccountMeta) {
	return obj.Impl.(common.AccountsGettable).GetAccounts()
}

func (obj *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.NewBorshEncoder(buf).Encode(obj); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (obj *Instruction) TextEncode(encoder *text.Encoder, option *text.Option) error {
	return encoder.Encode(obj.Impl, option)
}

func (obj *Instruction) UnmarshalWithDecoder(decoder *binary.Decoder) error {
	return obj.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (obj *Instruction) MarshalWithEncoder(encoder *binary.Encoder) error {
	err := encoder.WriteBytes(obj.TypeID.Bytes()[:obj.typeIdLen], false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(obj.Impl)
}
