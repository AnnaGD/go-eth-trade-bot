package keystore

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/AnnaGD/go-eth-trade-bot/cmd/keystore/wallet"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

var keystoreCmd = &cobra.Command{
	Use:   "keystore",
	Short: "Manage keystores",
	// Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		// Retrieve password from .env
		password := os.Getenv("KEYSTORE_PASSWORD")
		if password == "" {
			log.Fatal("KEYSTORE_PASSWORD is not set in .env")
		}

		//Define keystore directory
		keystoreDir := "./keystore"
		err = os.MkdirAll(keystoreDir, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating keystore directory: ", err)
		}

		// Initilize keystore manager
		ks := keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)

		// Create a new Etherum account
		account, err := ks.NewAccount(password)
		if err != nil {
			log.Fatal("Failed to create account: ", err)
		}

		fmt.Println("New account created!")
		fmt.Println("Address: ", account.Address.Hex())

		// Load the key from the keystore
		keyJson, err := os.ReadFile(account.URL.Path)
		if err != nil {
			log.Fatal("Failed to read keystore file: ", err)
		}

		// Decrypt the private key using password
		key, err := keystore.DecryptKey(keyJson, password)
		if err != nil {
			log.Fatal("Fialed to decrypt keystore: ", err)
		}

		// Extract the public key
		publicKey, ok := key.PrivateKey.Public().(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("Error extracting public key")
		}

		publicKeyHex := hexutil.Encode(crypto.FromECDSAPub(publicKey))

		// Print the public key
		fmt.Println("Public Key: ", publicKeyHex)
	},
}

func init() {
	keystoreCmd.AddCommand(wallet.CreateWalletCmd)
}
