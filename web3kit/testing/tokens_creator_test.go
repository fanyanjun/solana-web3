package testing

import (
	"context"
	"fmt"
	"testing"

	"github.com/fanyanjun/solana-web3/irys"
	"github.com/fanyanjun/solana-web3/test"
	"github.com/fanyanjun/solana-web3/web3"
	"github.com/fanyanjun/solana-web3/web3kit"
)

func TestCreateToken(t *testing.T) {
	var args = web3kit.CreateTokenArgs{
		BasicMetadata: web3kit.BasicMetadata{
			Name:        "AA",
			Symbol:      "BB",
			Description: "A TOKEN",
			Image:       test.TestingLogo(),
			Additional: map[string]string{
				"additional": "test",
			},
		},
		Decimals:      9,
		InitialSupply: web3.Ref(uint64(9)),
	}
	var ctx = context.Background()
	var payer = test.GetYourPrivateKey()
	var owner = web3.NewComplexSigner(payer)
	var connection = web3kit.Must1(web3.NewConnection(web3.Devnet.Url(), nil))

	provider := NewIrysProvider(irys.DEV, "https://arweave.net/")
	sig, mint := web3kit.Must2(web3kit.Token.CreateToken(
		ctx, connection, payer, owner, args, provider, web3.CommitmentProcessed,
	))
	fmt.Println("Signature:", sig)
	fmt.Println("Mint:", mint)
	fmt.Printf("Explorer: https://explorer.solana.com/address/%s/metadata?cluster=devnet\n", mint.String())
}
