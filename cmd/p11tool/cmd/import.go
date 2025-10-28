package cmd

//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//)
//
//// importCmd represents the import command
//var importCmd = &cobra.Command{
//	Use:   "import",
//	Short: "Import keys into the PKCS#11 token",
//	Long: `Import cryptographic keys into the PKCS#11 token.
//Supports importing EC, RSA, and symmetric keys from files or hex strings.
//
//Example usage:
//  p11tool import --keyType EC --keyFile /path/to/key.pem
//  p11tool import --keyType RSA --keyFile /path/to/key.pem
//  p11tool import --keyType AES --key "0123456789ABCDEF" --keyLabel "my-aes-key"`,
//	RunE: func(cmd *cobra.Command, args []string) error {
//		if err := initP11Context(); err != nil {
//			return err
//		}
//		defer p11w.Finalize()
//		defer p11w.CloseSession()
//
//		var err error
//		switch keyType {
//		case "RSA":
//			err = p11w.ImportRSAKeyFromFile(keyFile, keyStore)
//		case "EC":
//			err = p11w.ImportECKeyFromFile(keyFile, keyStore, keyStorepass, keyLabel)
//		case "AES", "GENERIC_SECRET", "SHA256_HMAC", "SHA384_HMAC":
//			err = p11w.ImportSymKey(keyType, key, keyStore, keyStorepass, keyLabel)
//		default:
//			return fmt.Errorf("unsupported key type: %s", keyType)
//		}
//
//		if err != nil {
//			return fmt.Errorf("failed to import key: %v", err)
//		}
//
//		fmt.Printf("Successfully imported %s key with label: %s\n", keyType, keyLabel)
//		return nil
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(importCmd)
//	importCmd.Flags().StringVar(&keyFile, "keyFile", "", "Path to the key file (for RSA/EC keys)")
//	importCmd.Flags().StringVar(&key, "key", "", "Key as hex string (for symmetric keys)")
//	importCmd.Flags().StringVar(&keyStore, "keyStore", "file", "Keystore type (file, pkcs12)")
//	importCmd.Flags().StringVar(&keyStorepass, "keyStorepass", "", "Keystore password")
//}

//	case "import":
//		if *keyType == "RSA" {
//			err = p11w.ImportRSAKeyFromFile(*keyFile, *keyStore)
//			exitWhenError(err)
//		} else if *keyType == "AES" ||
//			*keyType == "GENERIC_SECRET" ||
//			*keyType == "SHA256_HMAC" ||
//			*keyType == "SHA384_HMAC" {
//			err = p11w.ImportSymKey(*keyType, *key, *keyStore, *keyStorepass, *keyLabel)
//			exitWhenError(err)
//		} else {
//			err = p11w.ImportECKeyFromFile(*keyFile, *keyStore, *keyStorepass, *keyLabel)
//			exitWhenError(err)
//		}
//
//	case "importCert":
//		ec := pw.EcdsaKey{}
//		c := pw.GetCert(*keyFile)
//		ec.Certificate = c
//		ec.SKI.Sha256 = *keyLabel
//		Sha256Bytes, err := hex.DecodeString(ec.SKI.Sha256)
//		exitWhenError(err)
//		ec.SKI.Sha256Bytes = Sha256Bytes
//		err = p11w.ImportCertificate(ec)
//		exitWhenError(err)
//
