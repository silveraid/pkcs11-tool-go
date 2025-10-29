package cmd

import (
	"encoding/hex"
	"fmt"

	pw "github.com/scottallan/p11tool-new/pkg/pkcs11wrapper"
	"github.com/spf13/cobra"
)

var (
	certLabel string
	certFile  string
)

// importCertCmd handles the import of certificates
var importCertCmd = &cobra.Command{
	Use:   "importCert",
	Short: "Import certificates into the PKCS#11 token",
	Long: `Import certificates into the PKCS#11 token

Example usage:
  p11tool importCert --certLabel <LABEL> --certFile /path/to/cert.pem`,
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

		ec := pw.EcdsaKey{}
		c := pw.GetCert(certFile)
		ec.Certificate = c
		ec.SKI.Sha256 = certLabel
		Sha256Bytes, err := hex.DecodeString(ec.SKI.Sha256)
		exitWhenError(err)
		ec.SKI.Sha256Bytes = Sha256Bytes
		err = p11w.ImportCertificate(ec)
		exitWhenError(err)

		fmt.Printf("Successfully imported certificate with label: %s\n", certLabel)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCertCmd)
	importCertCmd.Flags().StringVar(&certLabel, "certLabel", "", "Label for the imported certificate")
	importCertCmd.Flags().StringVar(&certFile, "certFile", "", "Path to the certificate file")
	importCertCmd.MarkFlagRequired("certLabel")
	importCertCmd.MarkFlagRequired("certFile")
}
