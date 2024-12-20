package associated_token_account

import "github.com/fanyanjun/solana-web3/web3"

func FindAssociatedTokenAddress(
	walletAddress web3.PublicKey,
	mint web3.PublicKey,
	programId web3.PublicKey,
) (web3.PublicKey, error) {
	a, _, err := FindAssociatedTokenAddressAndBumpSeed(walletAddress, mint, programId)
	return a, err
}

func FindAssociatedTokenAddressAndBumpSeed(
	walletAddress web3.PublicKey,
	mint web3.PublicKey,
	programId web3.PublicKey,
) (web3.PublicKey, uint8, error) {
	return FindAssociatedTokenAddressAndBumpSeed2(walletAddress, mint, programId, ProgramID)
}

func FindAssociatedTokenAddressAndBumpSeed2(
	walletAddress web3.PublicKey,
	mint web3.PublicKey,
	programId web3.PublicKey,
	associatedTokenProgramId web3.PublicKey,
) (web3.PublicKey, uint8, error) {
	return web3.FindProgramAddress([][]byte{
		walletAddress[:],
		programId[:],
		mint[:],
	}, associatedTokenProgramId)
}
