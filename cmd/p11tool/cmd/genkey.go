package cmd

//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//)
//
//// genKeyCmd represents the generate key command
//var genKeyCmd = &cobra.Command{
//	Use:   "genkey",
//	Short: "Generate a new key in the PKCS#11 token",
//	Long: `Generate a new cryptographic key in the PKCS#11 token.
//Supports various key types including EC, RSA, AES, and HMAC keys.
//
//Example usage:
//  p11tool genkey --keyType EC --keyLabel "my-ec-key"
//  p11tool genkey --keyType RSA --keyLen 2048 --keyLabel "my-rsa-key"
//  p11tool genkey --keyType AES --keyLen 256 --keyLabel "my-aes-key"`,
//	RunE: func(cmd *cobra.Command, args []string) error {
//		if err := initP11Context(); err != nil {
//			return err
//		}
//		defer p11w.Finalize()
//		defer p11w.CloseSession()
//
//		var err error
//		switch keyType {
//		case "EC":
//			err = p11w.GenerateECKey(keyLabel)
//		case "RSA":
//			err = p11w.GenerateRSAKey(keyLabel, keyLen)
//		case "AES":
//			err = p11w.GenerateAESKey(keyLabel, keyLen)
//		case "GENERIC_SECRET", "SHA256_HMAC", "SHA384_HMAC":
//			err = p11w.GenerateGenericKey(keyType, keyLabel, keyLen)
//		default:
//			return fmt.Errorf("unsupported key type: %s", keyType)
//		}
//
//		if err != nil {
//			return fmt.Errorf("failed to generate key: %v", err)
//		}
//
//		fmt.Printf("Successfully generated %s key with label: %s\n", keyType, keyLabel)
//		return nil
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(genKeyCmd)
//	genKeyCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Label for the generated key")
//	genKeyCmd.Flags().IntVar(&keyLen, "keyLen", 0, "Key length in bits (required for RSA, AES, and HMAC keys)")
//	genKeyCmd.MarkFlagRequired("keyLabel")
//}
