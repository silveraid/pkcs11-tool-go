package cmd

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	inputFile string
)

// getSKICmd represents the generate key command
var getSKICmd = &cobra.Command{
	Use:   "getSKI",
	Short: "Calculates the SKI of private/public keys, CSRs, certificates",
	Long: `An X.509 SKI, or Subject Key Identifier, is a field within an X.509 digital certificate that provides a unique identifier for the certificate's public key. The SKI is often derived from a hash of the public key, typically using SHA-1

Example usage:
 p11tool getSKI --inputFile <YOUR PEM FILE>`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if inputFile == "" {
			return fmt.Errorf("no input file specified")
		}

		// Print out the path to the input file
		fmt.Printf("Input file: %s\n", inputFile)

		// Read the content of the pem file
		data, err := ioutil.ReadFile(inputFile)
		if err != nil {
			return err
		}

		for block, rest := pem.Decode(data); block != nil; block, rest = pem.Decode(rest) {
			fmt.Printf("Found PEM block of type: %s\n", block.Type)

			switch block.Type {
			case "CERTIFICATE":

				// Parse the certificate
				cert, err := x509.ParseCertificate(block.Bytes)
				if err != nil {
					log.Printf("Error parsing certificate: %v", err)
					continue
				}
				if cert.PublicKey == nil {
					log.Printf("Error parsing certificate: no public key found")
					continue
				}
				calculateSKIfromPubKey(cert.PublicKey)

			case "CERTIFICATE REQUEST":
				req, err := x509.ParseCertificateRequest(block.Bytes)
				if err != nil {
					log.Printf("Error parsing certificate request: %v", err)
					continue
				}
				if req.PublicKey == nil {
					log.Printf("Error parsing certificate request: no public key found")
					continue
				}
				calculateSKIfromPubKey(req.PublicKey)

			case "PRIVATE KEY": // Generic private key
				// Attempt to parse as PKCS#1 or PKCS#8
				var publicKey crypto.PublicKey
				if privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
					fmt.Println("Parsed as PKCS#1 Private Key")
					publicKey = privateKey.PublicKey
				} else if privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
					fmt.Println("Parsed as PKCS#8 Private Key")
					publicKey = privateKey.(crypto.Signer).Public()
				} else {
					fmt.Printf("Could not parse private key: %v\n", err)
				}
				if publicKey == nil {
					fmt.Printf("Error parsing private key: no public key found")
				}
				calculateSKIfromPubKey(publicKey)

			case "RSA PRIVATE KEY": // Specific RSA private key (PKCS#1)
				privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
				if err != nil {
					fmt.Printf("Error parsing RSA private key: %v\n", err)
					continue
				}
				publicKey := privateKey.Public()
				if publicKey == nil {
					fmt.Printf("Error parsing RSA private key: no public key found")
				}
				calculateSKIfromPubKey(publicKey)

			case "EC PRIVATE KEY":
				privateKey, err := x509.ParseECPrivateKey(block.Bytes)
				if err != nil {
					fmt.Printf("Error parsing RSA private key: %v\n", err)
					continue
				}
				publicKey := privateKey.Public()
				if publicKey == nil {
					fmt.Printf("Error parsing RSA private key: no public key found")
				}
				calculateSKIfromPubKey(publicKey)

			case "PUBLIC KEY":
				publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
				if err != nil {
					fmt.Printf("Error parsing public key: %v\n", err)
					continue
				}
				calculateSKIfromPubKey(publicKey)

			default:
				fmt.Printf("Unsupported or unknown block type: %s\n", block.Type)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getSKICmd)
	getSKICmd.Flags().StringVar(&inputFile, "inputFile", "", "Path of a PEM (pubKey, privKey, CSR, certificate) file")
	getSKICmd.MarkFlagRequired("inputFile")
}

// calculateSKIfromPubKey Calculates SHA-1 and SHA-256 SKI values of a public key
func calculateSKIfromPubKey(publicKey crypto.PublicKey) {
	switch pub := publicKey.(type) {
	case *rsa.PublicKey:
		fmt.Printf("Found RSA public key (%v)\n", pub.N.BitLen())
		rawBytes := pub.N.Bytes()
		fmt.Printf("SKI SHA1: %x\n", sha1.Sum(rawBytes))
		fmt.Printf("SKI SHA256: %x\n", sha256.Sum256(rawBytes))
	case *ecdsa.PublicKey:
		fmt.Printf("Found ECDSA public key (%v)\n", pub.Curve.Params().Name)
		rawBytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
		fmt.Printf("SKI SHA1: %x\n", sha1.Sum(rawBytes))
		fmt.Printf("SKI SHA256: %x\n", sha256.Sum256(rawBytes))
	default:
		log.Fatalf("Unsupported public key type: %T", pub)
	}
}
