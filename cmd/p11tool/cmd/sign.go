package cmd

//
//import (
//	"fmt"
//	"github.com/miekg/pkcs11"
//	"github.com/spf13/cobra"
//)
//
//// signCmd represents the sign command
//var signCmd = &cobra.Command{
//	Use:   "sign",
//	Short: "Sign data using a key in the PKCS#11 token",
//	Long: `Sign data using a key stored in the PKCS#11 token.
//Supports EC, RSA, and HMAC signing operations.
//
//Example usage:
//  p11tool sign --keyLabel "my-ec-key" --message "data to sign"
//  p11tool sign --keyLabel "my-hmac-key" --message "data to sign" --mechanism HMAC-SHA384`,
//	RunE: func(cmd *cobra.Command, args []string) error {
//		if err := initP11Context(); err != nil {
//			return err
//		}
//		defer p11w.Finalize()
//		defer p11w.CloseSession()
//
//		// Find the key object
//		objects, err := p11w.FindObjects([]*pkcs11.Attribute{
//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, keyLabel),
//		}, 1)
//		if err != nil {
//			return fmt.Errorf("failed to find key with label %s: %v", keyLabel, err)
//		}
//		if len(objects) == 0 {
//			return fmt.Errorf("no key found with label: %s", keyLabel)
//		}
//
//		var signature []byte
//		if mechanism == "HMAC-SHA384" {
//			signature, err = p11w.SignHmacSha384(objects[0], []byte(message))
//		} else {
//			signature, err = p11w.SignMessage(message, objects[0])
//		}
//		if err != nil {
//			return fmt.Errorf("failed to sign message: %v", err)
//		}
//
//		fmt.Printf("Successfully signed message. Signature: %x\n", signature)
//		return nil
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(signCmd)
//	signCmd.Flags().StringVar(&message, "message", "", "The message to sign")
//	signCmd.Flags().StringVar(&mechanism, "mechanism", "", "The signing mechanism (HMAC-SHA384 for HMAC keys)")
//	signCmd.MarkFlagRequired("message")
//}

//	case "SignHMAC384":
//		pkcs11_attr := pkcs11.NewAttribute(pkcs11.CKA_LABEL, *keyLabel)
//		p11w.ListObjects(
//			[]*pkcs11.Attribute{
//				pkcs11_attr,
//			}, *maxObjectsToList,
//		)
//		o, _, err := p11w.FindObjects([]*pkcs11.Attribute{
//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, *keyLabel),
//		},
//			1,
//		)
//		exitWhenError(err)
//		testMsg := []byte("someRandomString")
//		hmac, err := p11w.SignHmacSha384(o[0], testMsg)
//		exitWhenError(err)
//		fmt.Printf("successfully tested CKM_SHA384_HMAC on key with LABEL: %s\n HMAC %x\n", *keyLabel, hmac)
//
