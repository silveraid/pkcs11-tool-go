package cmd

import (
	"fmt"
	"math/rand"

	"github.com/miekg/pkcs11"
	"github.com/spf13/cobra"
)

var (
	testCase       string
	maxObjectLimit int
)

// testCmd represents the list command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Conducts various HSM related tests",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		p11w, err := initP11Context()
		if err != nil {
			return err
		}
		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)

		switch testCase {
		case "AESGCM":

			// Create a key name
			newAESKeyLabel := randomKeyName()
			fmt.Printf("Creating new AES Key: %s\n", newAESKeyLabel)

			// Create a key for testing
			_, err := p11w.CreateSymKey(newAESKeyLabel, 32, "AES")
			exitWhenError(err)

			pkcs11Attr := pkcs11.NewAttribute(pkcs11.CKA_LABEL, newAESKeyLabel)
			p11w.ListObjects(
				[]*pkcs11.Attribute{
					pkcs11Attr,
				}, maxObjectLimit,
			)

			o, _, err := p11w.FindObjects([]*pkcs11.Attribute{
				pkcs11.NewAttribute(pkcs11.CKA_LABEL, newAESKeyLabel),
			},
				1,
			)
			exitWhenError(err)

			testMsg := []byte("ThisIsATestClearTextString")
			enc, iv, err := p11w.EncAESGCM(o[0], testMsg)
			exitWhenError(err)
			fmt.Printf("successfully encrypted  message '%s' with CKM_AES_GCM and key with LABEL: %s\n CipherText %v\n IV: %v\n", testMsg, newAESKeyLabel, enc, iv)

			dec, err := p11w.DecAESGCM(o[0], enc, iv)
			exitWhenError(err)
			fmt.Printf("successfully decrypted ciptherText '%v' with CKM_AES_GCM and key with LABEL: %s\n ClearText %s\n", enc, newAESKeyLabel, dec)

			// Delete the created key
			p11w.DeleteObj("CKO_SECRET_KEY", newAESKeyLabel)

		case "HMAC384":
			pkcs11_attr := pkcs11.NewAttribute(pkcs11.CKA_LABEL, keyLabel)
			p11w.ListObjects(
				[]*pkcs11.Attribute{
					pkcs11_attr,
				}, maxObjectLimit,
			)
			o, _, err := p11w.FindObjects([]*pkcs11.Attribute{
				pkcs11.NewAttribute(pkcs11.CKA_LABEL, keyLabel),
			},
				1,
			)
			exitWhenError(err)
			testMsg := []byte("someRandomString")
			hmac, err := p11w.SignHmacSha384(o[0], testMsg)
			exitWhenError(err)
			fmt.Printf("successfully tested CKM_SHA384_HMAC on key with LABEL: %s\n HMAC %x\n", keyLabel, hmac)

		case "EC":
			return fmt.Errorf("test case not implemented: %s", testCase)
		//	case "testEc":
		//
		//		message := "Some Test Message"
		//
		//		// test SW ecdsa sign and verify
		//		ec := pw.EcdsaKey{}
		//		ec.ImportPrivKeyFromFile("contrib/testfiles/key.pem")
		//		sig, err := ec.SignMessage(message)
		//		exitWhenError(err)
		//		fmt.Println("Signature:", sig)
		//		verified := ec.VerifySignature(message, sig)
		//		fmt.Println("Verified:", verified)
		//
		//		// test PKCS11 ecdsa sign and verify
		//		// Find object
		//		id, err := hex.DecodeString("018f389d200e48536367f05b99122f355ba33572009bd2b8b521cdbbb717a5b5")
		//		exitWhenError(err)
		//
		//		o, _, err := p11w.FindObjects([]*pkcs11.Attribute{
		//			pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, "BCPRV1"),
		//			pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		//			pkcs11.NewAttribute(pkcs11.CKA_ID, id),
		//		},
		//			2,
		//		)
		//
		//		exitWhenError(err)
		//
		//		sig, err = p11w.SignMessage(message, o[0])
		//		exitWhenError(err)
		//		fmt.Println("pkcs11 Signature:", sig)
		//		verified = ec.VerifySignature(message, sig)
		//		fmt.Println("Verified:", verified)
		//
		//		// test pkcs11 verify
		//		o, _, err = p11w.FindObjects([]*pkcs11.Attribute{
		//			pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, "BCPUB1"),
		//			pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PUBLIC_KEY),
		//			pkcs11.NewAttribute(pkcs11.CKA_ID, id),
		//		},
		//			2,
		//		)
		//
		//		verified, err = p11w.VerifySignature(message, sig, o[0])
		//		exitWhenError(err)
		//		fmt.Println("pkcs11 Verified:", verified)
		//
		//		// derive test
		//		ec2 := pw.EcdsaKey{}
		//		ec2.Generate("P-256")
		//
		//		secret, err := ec.DeriveSharedSecret(ec2.PubKey)
		//		exitWhenError(err)
		//		fmt.Printf("shared secret: %x\n", secret)
		//
		//		secret, err = ec2.DeriveSharedSecret(ec.PubKey)
		//		exitWhenError(err)
		//		fmt.Printf("shared secret: %x\n", secret)

		case "RSA":
			return fmt.Errorf("test case not implemented: %s", testCase)

		//	case "testRsa":
		//		message := "Some Test Message"
		//
		//		rsa := pw.RsaKey{}
		//		//rsa.Generate(2048)
		//		err = rsa.ImportPrivKeyFromFile("contrib/testfiles/key.rsa.pem")
		//		exitWhenError(err)
		//		rsa.GenSKI()
		//
		//		err = p11w.ImportRSAKey(rsa)
		//		exitWhenError(err)
		//
		//		sig, err := rsa.SignMessage(message, 256)
		//		exitWhenError(err)
		//
		//		fmt.Println("Signature:", sig)
		//
		//		// test PKCS11 ecdsa sign and verify
		//		// Find object
		//		id, err := hex.DecodeString("0344ae0121e025d998f5923174e9e4d69b899144ac79bfdf01c065bd4d99d6cb")
		//		exitWhenError(err)
		//
		//		o, _, err := p11w.FindObjects([]*pkcs11.Attribute{
		//			pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_RSA),
		//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, "TLSPRVKEY"),
		//			pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		//			pkcs11.NewAttribute(pkcs11.CKA_ID, id),
		//		},
		//			2,
		//		)
		//		exitWhenError(err)
		//
		//		sig, err = p11w.SignMessageAdvanced([]byte(message), o[0], pkcs11.NewMechanism(pkcs11.CKM_SHA256_RSA_PKCS, nil))
		//		exitWhenError(err)
		//
		//		fmt.Println("pkcs11 Signature:", sig)

		default:
			return fmt.Errorf("unknown test case: %s", testCase)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVar(&testCase, "testcase", "", "Test case to execute (AESGCM, HMAC384, EC, RSA)")
	testCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Label for the generated key")
	testCmd.Flags().IntVar(&maxObjectLimit, "limit", 50, "Maximum number of objects to list")
	testCmd.MarkFlagRequired("testCase")
}

// randomKeyName Generates a random key name which can be used for testing.
func randomKeyName() string {
	// Generate a random integer between 10,000,000 (inclusive) and 99,999,999 (inclusive)
	// This ensures the number is always 8 digits long.
	randomNumber := rand.Intn(90000000) + 10000000

	// Return the generated random name
	return fmt.Sprintf("P11TOOL-%d", randomNumber)
}
