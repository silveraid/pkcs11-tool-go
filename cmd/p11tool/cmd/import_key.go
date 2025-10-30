package cmd

import (
	"fmt"

	pw "github.com/scottallan/p11tool-new/pkg/pkcs11wrapper"
	"github.com/spf13/cobra"
)

// importKeyCmd represents the import command
var importKeyCmd = &cobra.Command{
	Use:   "importKey",
	Short: "Import keys into the PKCS#11 token",
	Long: `Import cryptographic keys into the PKCS#11 token.
Supports importing EC, RSA, and symmetric keys from files or hex strings.

Note that if neither the --keyType nor the --key parameters are defined, the
function will generate new key material and will attempt to import it into
the PKCS#11 token.

Example usage:
 p11tool importKey --keyType EC --keyFile /path/to/key.pem
 p11tool importKey --keyType RSA --keyFile /path/to/key.pem
 p11tool importKey --keyType AES --key "0123456789ABCDEF" --keyLabel "my-aes-key"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		p11w, err := initP11Context()
		if err != nil {
			return err
		}
		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)
		//defer termState.cleanupPin(*slotPin, &p11Pin, *less)

		// If keyType is defined, but keyFile and key does not, we will
		// generate and import a new key into the HSM.
		if keyType != "" && keyFile == "" && key == "" {

			switch keyType {
			// FIXME: we should probably respect the parameters
			case "RSA":
				rsa := pw.RsaKey{}
				rsa.Generate(2048)
				p11w.ImportRSAKey(rsa)
			case "EC":
				ec := pw.EcdsaKey{}
				// TODO: fix non working curves (P-521)
				ec.Generate("P-256")
				p11w.ImportECKey(ec)
			default:
				return fmt.Errorf("unsupported key type: %s", keyType)
			}
		}

		switch keyType {
		case "RSA":
			err = p11w.ImportRSAKeyFromFile(keyFile, keyStore)
		case "EC":
			err = p11w.ImportECKeyFromFile(keyFile, keyStore, keyStorepass, keyLabel)
		case "AES", "GENERIC_SECRET", "SHA256_HMAC", "SHA384_HMAC":
			err = p11w.ImportSymKey(keyType, key, keyStore, keyStorepass, keyLabel)
		default:
			return fmt.Errorf("unsupported key type: %s", keyType)
		}

		exitWhenError(err)

		fmt.Printf("Successfully imported %s key with label: %s\n", keyType, keyLabel)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(importKeyCmd)
	importKeyCmd.Flags().StringVar(&keyType, "keyType", "EC", "Type of key (EC,RSA,GENERIC_SECRET,AES,SHA256_HMAC,SHA384_HMAC,DES3)")
	importKeyCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Label for the imported key")
	importKeyCmd.Flags().StringVar(&keyFile, "keyFile", "", "Path to the key file (for RSA/EC keys)")
	importKeyCmd.Flags().StringVar(&key, "key", "", "Key as hex string (for symmetric keys)")
	importKeyCmd.Flags().StringVar(&keyStore, "keyStore", "file", "Keystore type (file, pkcs12)")
	importKeyCmd.Flags().StringVar(&keyStorepass, "keyStorepass", "", "Keystore password")
	importKeyCmd.MarkFlagRequired("keyLabel")
}
