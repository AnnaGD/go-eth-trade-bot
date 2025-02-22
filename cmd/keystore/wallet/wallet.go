package wallet

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var (
	walletPassword string
	outputDir      string
)

// Export the command variable
var CreateWalletCmd = &cobra.Command{
	Use:   "create-wallet",
	Short: "Generate a new wallaet",
	Long:  `Creates a new wallet with the provided password and saves it`,
	Run: func(cmd *cobra.Command, args []string) {

		// Validate password strength
		if err := validatePassword(walletPassword); err != nil {
			log.Fatalf("Password validation failed: %v", err)
		}

		// Ensure the keystore directory exists
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating keystore directory: %v", err)
		}

		// Initilize keystore manager
		ks := keystore.NewKeyStore(outputDir, keystore.StandardScryptN, keystore.StandardScryptP)

		// Create a new account
		account, err := ks.NewAccount(walletPassword)
		if err != nil {
			log.Fatalf("Failed to create acocunt: %v", err)
		}

		fmt.Println("🎉 New Wallet created successfully!")
		fmt.Println("📁 Keystore saved to:", account.URL.Path)
		fmt.Println("📝 Address:", account.Address.Hex())

		// Load the key from the keystore to display the public key
		keyJSON, err := os.ReadFile(account.URL.Path)
		if err != nil {
			log.Printf("Warning: Failed to read keystore file: %v", err)
			return
		}

		// Decrypt the private key using password
		key, err := keystore.DecryptKey(keyJSON, walletPassword)
		if err != nil {
			log.Printf("Warning: Failed to decrypt keystore: %v", err)
			return
		}

		// Extract and display the public key
		publicKeyBytes := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
		publicKeyHex := hexutil.Encode(publicKeyBytes)
		fmt.Println("🔑 Public Key:", publicKeyHex)

	},
}

func init() {
	// add flags to the create-wallet command
	createWalletCmd.Flags().StringVarP(&walletPassword, "password", "p", "", "Password for the new wallet (required)")
	createWalletCmd.Flags().StringVarP(&outputDir, "output", "o", "./keystore", "Directory to store the keystore file")

	// Mark password as required
	createWalletCmd.MarkFlagRequired("password")
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasDigit := strings.ContainsAny(password, "0123456789")
	hasSpecial := strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:,.<>?/")

	if !hasDigit || !hasLower || !hasSpecial || !hasUpper {
		return fmt.Errorf("password must contain uppercase, lowercase, digit, and special character")
	}

	return nil
}
