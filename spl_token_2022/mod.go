// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package spl_token_2022

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

var ProgramID common.PublicKey = common.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")

func SetProgramID(pubkey common.PublicKey) {
	ProgramID = pubkey
	if !common.IsZero(ProgramID) {
		solanago.RegisterInstructionDecoder(common.As(ProgramID), registryDecodeInstruction)
	}
}

const ProgramName = "spl_token_2022"

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
	Instruction_InitializeMint                   uint8 = 0
	Instruction_InitializeAccount                uint8 = 1
	Instruction_InitializeMultisig               uint8 = 2
	Instruction_Transfer                         uint8 = 3
	Instruction_Approve                          uint8 = 4
	Instruction_Revoke                           uint8 = 5
	Instruction_SetAuthority                     uint8 = 6
	Instruction_MintTo                           uint8 = 7
	Instruction_Burn                             uint8 = 8
	Instruction_CloseAccount                     uint8 = 9
	Instruction_FreezeAccount                    uint8 = 10
	Instruction_ThawAccount                      uint8 = 11
	Instruction_TransferChecked                  uint8 = 12
	Instruction_ApproveChecked                   uint8 = 13
	Instruction_MintToChecked                    uint8 = 14
	Instruction_BurnChecked                      uint8 = 15
	Instruction_InitializeAccount2               uint8 = 16
	Instruction_SyncNative                       uint8 = 17
	Instruction_InitializeAccount3               uint8 = 18
	Instruction_InitializeMultisig2              uint8 = 19
	Instruction_InitializeMint2                  uint8 = 20
	Instruction_GetAccountDataSize               uint8 = 21
	Instruction_InitializeImmutableOwner         uint8 = 22
	Instruction_AmountToUiAmount                 uint8 = 23
	Instruction_UiAmountToAmount                 uint8 = 24
	Instruction_InitializeMintCloseAuthority     uint8 = 25
	Instruction_TransferFeeExtension             uint8 = 26
	Instruction_ConfidentialTransferExtension    uint8 = 27
	Instruction_DefaultAccountStateExtension     uint8 = 28
	Instruction_Reallocate                       uint8 = 29
	Instruction_MemoTransferExtension            uint8 = 30
	Instruction_CreateNativeMint                 uint8 = 31
	Instruction_InitializeNonTransferableMint    uint8 = 32
	Instruction_InterestBearingMintExtension     uint8 = 33
	Instruction_CpiGuardExtension                uint8 = 34
	Instruction_InitializePermanentDelegate      uint8 = 35
	Instruction_TransferHookExtension            uint8 = 36
	Instruction_ConfidentialTransferFeeExtension uint8 = 37
	Instruction_WithdrawExcessLamports           uint8 = 38
	Instruction_MetadataPointerExtension         uint8 = 39
	Instruction_GroupPointerExtension            uint8 = 40
	Instruction_GroupMemberPointerExtension      uint8 = 41
)

var InstructionImplDef = binary.NewVariantDefinition(binary.Uint8TypeIDEncoding, []binary.VariantType{
	{
		"initialize_mint", (*InitializeMint)(nil),
	},
	{
		"initialize_account", (*InitializeAccount)(nil),
	},
	{
		"initialize_multisig", (*InitializeMultisig)(nil),
	},
	{
		"transfer", (*Transfer)(nil),
	},
	{
		"approve", (*Approve)(nil),
	},
	{
		"revoke", (*Revoke)(nil),
	},
	{
		"set_authority", (*SetAuthority)(nil),
	},
	{
		"mint_to", (*MintTo)(nil),
	},
	{
		"burn", (*Burn)(nil),
	},
	{
		"close_account", (*CloseAccount)(nil),
	},
	{
		"freeze_account", (*FreezeAccount)(nil),
	},
	{
		"thaw_account", (*ThawAccount)(nil),
	},
	{
		"transfer_checked", (*TransferChecked)(nil),
	},
	{
		"approve_checked", (*ApproveChecked)(nil),
	},
	{
		"mint_to_checked", (*MintToChecked)(nil),
	},
	{
		"burn_checked", (*BurnChecked)(nil),
	},
	{
		"initialize_account2", (*InitializeAccount2)(nil),
	},
	{
		"sync_native", (*SyncNative)(nil),
	},
	{
		"initialize_account3", (*InitializeAccount3)(nil),
	},
	{
		"initialize_multisig2", (*InitializeMultisig2)(nil),
	},
	{
		"initialize_mint2", (*InitializeMint2)(nil),
	},
	{
		"get_account_data_size", (*GetAccountDataSize)(nil),
	},
	{
		"initialize_immutable_owner", (*InitializeImmutableOwner)(nil),
	},
	{
		"amount_to_ui_amount", (*AmountToUiAmount)(nil),
	},
	{
		"ui_amount_to_amount", (*UiAmountToAmount)(nil),
	},
	{
		"initialize_mint_close_authority", (*InitializeMintCloseAuthority)(nil),
	},
	{
		"transfer_fee_extension", (*TransferFeeExtension)(nil),
	},
	{
		"confidential_transfer_extension", (*ConfidentialTransferExtension)(nil),
	},
	{
		"default_account_state_extension", (*DefaultAccountStateExtension)(nil),
	},
	{
		"reallocate", (*Reallocate)(nil),
	},
	{
		"memo_transfer_extension", (*MemoTransferExtension)(nil),
	},
	{
		"create_native_mint", (*CreateNativeMint)(nil),
	},
	{
		"initialize_non_transferable_mint", (*InitializeNonTransferableMint)(nil),
	},
	{
		"interest_bearing_mint_extension", (*InterestBearingMintExtension)(nil),
	},
	{
		"cpi_guard_extension", (*CpiGuardExtension)(nil),
	},
	{
		"initialize_permanent_delegate", (*InitializePermanentDelegate)(nil),
	},
	{
		"transfer_hook_extension", (*TransferHookExtension)(nil),
	},
	{
		"confidential_transfer_fee_extension", (*ConfidentialTransferFeeExtension)(nil),
	},
	{
		"withdraw_excess_lamports", (*WithdrawExcessLamports)(nil),
	},
	{
		"metadata_pointer_extension", (*MetadataPointerExtension)(nil),
	},
	{
		"group_pointer_extension", (*GroupPointerExtension)(nil),
	},
	{
		"group_member_pointer_extension", (*GroupMemberPointerExtension)(nil),
	},
})

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id uint8) string {
	switch id {
	case Instruction_InitializeMint:
		return "InitializeMint"
	case Instruction_InitializeAccount:
		return "InitializeAccount"
	case Instruction_InitializeMultisig:
		return "InitializeMultisig"
	case Instruction_Transfer:
		return "Transfer"
	case Instruction_Approve:
		return "Approve"
	case Instruction_Revoke:
		return "Revoke"
	case Instruction_SetAuthority:
		return "SetAuthority"
	case Instruction_MintTo:
		return "MintTo"
	case Instruction_Burn:
		return "Burn"
	case Instruction_CloseAccount:
		return "CloseAccount"
	case Instruction_FreezeAccount:
		return "FreezeAccount"
	case Instruction_ThawAccount:
		return "ThawAccount"
	case Instruction_TransferChecked:
		return "TransferChecked"
	case Instruction_ApproveChecked:
		return "ApproveChecked"
	case Instruction_MintToChecked:
		return "MintToChecked"
	case Instruction_BurnChecked:
		return "BurnChecked"
	case Instruction_InitializeAccount2:
		return "InitializeAccount2"
	case Instruction_SyncNative:
		return "SyncNative"
	case Instruction_InitializeAccount3:
		return "InitializeAccount3"
	case Instruction_InitializeMultisig2:
		return "InitializeMultisig2"
	case Instruction_InitializeMint2:
		return "InitializeMint2"
	case Instruction_GetAccountDataSize:
		return "GetAccountDataSize"
	case Instruction_InitializeImmutableOwner:
		return "InitializeImmutableOwner"
	case Instruction_AmountToUiAmount:
		return "AmountToUiAmount"
	case Instruction_UiAmountToAmount:
		return "UiAmountToAmount"
	case Instruction_InitializeMintCloseAuthority:
		return "InitializeMintCloseAuthority"
	case Instruction_TransferFeeExtension:
		return "TransferFeeExtension"
	case Instruction_ConfidentialTransferExtension:
		return "ConfidentialTransferExtension"
	case Instruction_DefaultAccountStateExtension:
		return "DefaultAccountStateExtension"
	case Instruction_Reallocate:
		return "Reallocate"
	case Instruction_MemoTransferExtension:
		return "MemoTransferExtension"
	case Instruction_CreateNativeMint:
		return "CreateNativeMint"
	case Instruction_InitializeNonTransferableMint:
		return "InitializeNonTransferableMint"
	case Instruction_InterestBearingMintExtension:
		return "InterestBearingMintExtension"
	case Instruction_CpiGuardExtension:
		return "CpiGuardExtension"
	case Instruction_InitializePermanentDelegate:
		return "InitializePermanentDelegate"
	case Instruction_TransferHookExtension:
		return "TransferHookExtension"
	case Instruction_ConfidentialTransferFeeExtension:
		return "ConfidentialTransferFeeExtension"
	case Instruction_WithdrawExcessLamports:
		return "WithdrawExcessLamports"
	case Instruction_MetadataPointerExtension:
		return "MetadataPointerExtension"
	case Instruction_GroupPointerExtension:
		return "GroupPointerExtension"
	case Instruction_GroupMemberPointerExtension:
		return "GroupMemberPointerExtension"
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
