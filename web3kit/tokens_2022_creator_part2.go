package web3kit

import (
	"context"
	"encoding/json"

	. "github.com/fanyanjun/solana-web3/spl_token_2022"
	. "github.com/fanyanjun/solana-web3/spl_token_2022/extension"
	"github.com/fanyanjun/solana-web3/token_metadata"
	"github.com/fanyanjun/solana-web3/web3"
	"github.com/sirupsen/logrus"
)

type CreateTokenArgs2022V1 struct {
	BasicMetadataV1
	Decimals            uint8
	InitialSupply       *uint64
	EnableWhitelist     bool
	EnableBlacklist     bool
	EnableForceTransfer bool
	Fee                 *uint16 // max: 100_00(100%)
	MaximumFee          uint64
}

type BasicMetadataV1 struct {
	Image       string
	Name        string
	Symbol      string
	Uri         string
	Description string
	Additional  map[string]string
	//add fanyanjun
	Creators *[]token_metadata.Creator
}

func (t tokenKit2022) CreateTokenV1(
	ctx context.Context,
	connection *web3.Connection,
	payer web3.Signer,
	owner web3.ComplexSigner,
	args CreateTokenArgs2022V1,
	metaProvider *FileProvider,
	metaType MetadataProvider,
	commitment web3.Commitment,
) (web3.TransactionSignature, web3.PublicKey, error) {
	var mintAuthority = owner.PublicKey
	var freezeAuthority = &mintAuthority
	var permanentDelegate = owner.PublicKey
	var transferFeeConfigAuthority = &mintAuthority
	var withdrawWithheldAuthority = &mintAuthority

	/*metadataURI, err := metaProvider.MetadataURI(ctx, connection, payer, args.BasicMetadata)
	if err != nil {
		return "", web3.PublicKey{}, err
	}*/

	tokenKeypair := web3.Keypair.Generate()
	var extensions_ []ExtensionInitializationParams

	if args.EnableWhitelist {
		var state = AccountStateInitialized
		if args.EnableWhitelist {
			state = AccountStateFrozen
		}
		extensions_ = append(extensions_, NewDefaultAccountStateParams(state))
	} else if args.EnableBlacklist {
		// pass
	} else {
		freezeAuthority = nil
	}
	if args.EnableForceTransfer {
		extensions_ = append(extensions_, NewPermanentDelegateParams(permanentDelegate))
	}
	if args.Fee != nil {
		extensions_ = append(extensions_, NewTransferFeeConfigParams(
			//max(100_00 /*100%*/, *args.Fee),
			max(0 /*100%*/, *args.Fee),
			max(0, args.MaximumFee),
			transferFeeConfigAuthority,
			withdrawWithheldAuthority,
		))

	}
	e, _ := json.Marshal(extensions_)
	logrus.Info("extensions_:", string(e))

	var additionalMetadata []struct {
		Key   string
		Value string
	}
	for key, value := range args.BasicMetadataV1.Additional {
		additionalMetadata = append(additionalMetadata, struct {
			Key   string
			Value string
		}{Key: key, Value: value})
	}

	tors := make([]token_metadata.Creator, 0)
	if args.BasicMetadataV1.Creators != nil {
		for _, creator := range *args.BasicMetadataV1.Creators {
			tors = append(tors, token_metadata.Creator{
				Address:  creator.Address,
				Verified: creator.Verified,
				Share:    creator.Share,
			})
		}
	}

	tokenMetadata := token_metadata.TokenMetadata{
		UpdateAuthority:    mintAuthority,
		Mint:               tokenKeypair.PublicKey(),
		Name:               args.Name,
		Symbol:             args.Symbol,
		Uri:                args.Uri, //metadataURI
		AdditionalMetadata: additionalMetadata,
		Creators:           &tors,
	}
	sig, err := Token2022.CreateMint(
		ctx,
		connection,
		payer,
		owner,
		freezeAuthority,
		args.Decimals,
		tokenKeypair,
		extensions_,
		args.InitialSupply,
		tokenMetadata,
		metaType,
		web3.TokenProgram2022ID,
		commitment,
	)
	if err != nil {
		return "", web3.PublicKey{}, err
	}
	return sig, tokenKeypair.PublicKey(), nil
}
