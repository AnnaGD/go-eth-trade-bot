package keystore

import (
	"github.com/AnnaGD/go-eth-trade-bot/cmd/keystore/wallet"
	"github.com/spf13/cobra"
)

var KeystoreCmd = &cobra.Command{
	Use:   "keystore",
	Short: "Manage keystores",
	// Long:  `All software has versions. This is Hugo's`,
}

func init() {
	// Add persistent flags that will be available to all subcommands
	KeystoreCmd.PersistentFlags().StringP("keystore-dir", "d", "./keystore", "Custom keystore directory")

	// Add the wallet subcommands
	KeystoreCmd.AddCommand(wallet.CreateWalletCmd)
	// Future subcommands will be added here
	// KeystoreCmd.AddCommand(wallet.ListWalletsCmd)
	// KeystoreCmd.AddCommand(wallet.ImportWalletCmd)
}
