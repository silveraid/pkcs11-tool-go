package cmd

import (
	"fmt"

	"github.com/miekg/pkcs11"
	pw "github.com/scottallan/p11tool-new/pkg/pkcs11wrapper"
	"github.com/spf13/cobra"
)

// genKeyCmd represents the generate key command
var genKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate a new key in the PKCS#11 token",
	Long: `Generate a new cryptographic key in the PKCS#11 token.
Supports various key types including EC, RSA, AES, and HMAC keys.

Example usage:
 p11tool genkey --keyType EC --keyLabel "my-ec-key"
 p11tool genkey --keyType RSA --keyLen 2048 --keyLabel "my-rsa-key"
 p11tool genkey --keyType AES --keyLen 256 --keyLabel "my-aes-key"`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// FIXME
		maxObjectsToList := 50

		p11w, err := initP11Context()
		if err != nil {
			return err
		}

		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)
		//defer termState.cleanupPin(*slotPin, &p11Pin, *less)

		switch keyType {
		case "EC":
			ec := pw.EcdsaKey{}
			//TODO pass in from argument
			ec.NamedCurveAsString = "P-256"
			_, err := p11w.GenerateEC(ec)
			exitWhenError(err)

		case "RSA":
			rsa := pw.RsaKey{}
			err := p11w.GenerateRSA(rsa, keyLen, keyLabel)
			exitWhenError(err)

		case "AES":
			//Generate Key
			_, err := p11w.CreateSymKey(keyLabel, keyLen, keyType)
			exitWhenError(err)
			p11w.ListObjects(
				[]*pkcs11.Attribute{},
				maxObjectsToList,
			)

		case "GENERIC_SECRET", "SHA256_HMAC", "SHA384_HMAC":
			//Generate Key
			symKey, err := p11w.CreateSymKey(keyLabel, keyLen, keyType)
			exitWhenError(err)
			testMsg := []byte("someRandomString")
			hmac, err := p11w.SignHmacSha384(symKey, testMsg)
			exitWhenError(err)
			fmt.Printf("Successfully tested CKM_SHA384_HMAC on key with label: %s \n HMAC %x\n", keyLabel, hmac)
			p11w.ListObjects(
				[]*pkcs11.Attribute{},
				maxObjectsToList,
			)

		case "DES3":
			//Generate DES Key
			_, err := p11w.CreateSymKey(keyLabel, keyLen, keyType)
			exitWhenError(err)

		default:
			return fmt.Errorf("unsupported key type: %s", keyType)
		}

		if err != nil {
			return fmt.Errorf("failed to generate key: %v", err)
		}

		fmt.Printf("Successfully generated %s key with label: %s\n", keyType, keyLabel)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genKeyCmd)
	genKeyCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Label for the generated key")
	genKeyCmd.Flags().IntVar(&keyLen, "keyLen", 32, "Key length in bits (required for RSA, AES, and HMAC keys)")
	genKeyCmd.Flags().StringVar(&keyType, "keyType", "EC", "Type of key (EC,RSA,GENERIC_SECRET,AES,SHA256_HMAC,SHA384_HMAC,DES3)")
	genKeyCmd.MarkFlagRequired("keyLabel")
}
