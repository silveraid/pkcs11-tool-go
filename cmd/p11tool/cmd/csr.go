package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	pw "github.com/scottallan/p11tool-new/pkg/pkcs11wrapper"
	"github.com/spf13/cobra"
)

// csrCmd represents the generate key command
var csrCmd = &cobra.Command{
	Use:   "csr",
	Short: "Generate a new CSR using an HSM protected key (EC keys only)",
	Long: `Generate a new CSR using an HSM protected key (EC keys only)

Example usage:
 p11tool csr \
   --keyLabel <YOUR KEY LABEL> \
   --subject '{"names":[{"c":"Country","st":"State","l":"Locality","o":"Organization","ou":"Org Unit"}],"CN":"Common Name"}'`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Input validation
		if csrSubject == "" {
			return errors.New("CSR subject is required!")
		}

		// Parse the provided string
		var csrInfo pw.CSRInfo
		if err := json.Unmarshal([]byte(csrSubject), &csrInfo); err != nil {
			return err
		}

		fmt.Println(csrSubject)
		fmt.Println(csrInfo)

		//// FIXME
		//maxObjectsToList := 50
		//
		p11w, err := initP11Context()
		if err != nil {
			return err
		}

		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)
		////defer termState.cleanupPin(*slotPin, &p11Pin, *less)
		//

		//
		// Elliptic curve only
		//

		ec := pw.EcdsaKey{}
		ec.SKI.Sha256 = keyLabel
		ec.Req = &csrInfo

		fmt.Println(pw.ToJson(csrInfo))

		// Generate the CSR
		csr, _, err := p11w.GenCSR(ec)
		exitWhenError(err)

		// Save the CSR on the filesystem
		fmt.Printf("writing csr to %s\n", csrFile)
		err = ioutil.WriteFile(csrFile, csr, 0644)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(csrCmd)
	csrCmd.Flags().StringVar(&csrSubject, "subject", "", "Subject line parameters of the CSR")
	csrCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Label for the key to be used")
	csrCmd.Flags().StringVar(&csrFile, "csrFile", "/tmp/csr.pem", "Path of the PEM file containing the generated CSR")
	csrCmd.MarkFlagRequired("csrSubject")
	csrCmd.MarkFlagRequired("keyLabel")
}
