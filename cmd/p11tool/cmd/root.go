package cmd

import (
	"fmt"
	"math"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"

	pw "github.com/scottallan/p11tool-new/pkg/pkcs11wrapper"
	"github.com/spf13/cobra"
)

const defaultLibPaths = `
/usr/lib/softhsm/libsofthsm2.so,
/usr/lib/x86_64-linux-gnu/softhsm/libsofthsm2.so,
/usr/lib/s390x-linux-gnu/softhsm/libsofthsm2.so,
/usr/lib/powerpc64le-linux-gnu/softhsm/libsofthsm2.so,
/opt/homebrew/Cellar/softhsm/2.6.1/lib/softhsm/libsofthsm2.so,
/usr/local/lib/softhsm/libsofthsm2.so`

var (
	// flags
	pkcs11Library string
	slotLabel     string
	slotPin       string
	keyFile       string
	keyType       string
	keyLen        int
	keyLabel      string
	keyStore      string
	keyStorepass  string
	key           string
	csrInfo       string
	wrapKey       string
	objClass      string
	outF          string
	noDec         bool
	less          bool
	byCKAID       bool
	mechOver      string
	maxObjects    int
	message       string
	mechanism     string
)

type termInfo struct {
	termState *terminal.State
	curState  *terminal.State
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "p11tool",
	Short: "A PKCS#11 tool for key management and operations",
	Long: ` p11tool is a command-line utility for managing PKCS#11 tokens and performing
cryptographic operations. It supports various key types including RSA, EC, and AES,
and provides functionality for key generation, import/export, signing, and more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVar(&pkcs11Library, "lib", "", "Location of pkcs11 library")
	rootCmd.PersistentFlags().StringVar(&slotLabel, "slot", "p11tool", "Slot Label")
	rootCmd.PersistentFlags().StringVar(&slotPin, "pin", "", "Slot PIN")
	rootCmd.PersistentFlags().StringVar(&keyType, "keyType", "EC", "Type of key (EC,RSA,GENERIC_SECRET,AES,SHA256_HMAC,SHA384_HMAC,DES3)")
	rootCmd.PersistentFlags().BoolVar(&less, "less", true, "Don't show password preamble")
}

// initP11Context initializes the PKCS#11 context and session
func initP11Context() (*pw.Pkcs11Wrapper, error) {
	var err error
	var p11Lib string
	var p11Pin string

	// override command line parameter with environment variable
	if x, y := os.LookupEnv("P11TOOL_PKCS11_LIBRARY"); y {
		pkcs11Library = x
	}

	if pkcs11Library == "" {
		// The PKCS11 library has no value, let's try to find a library by
		// testing all the possible locations.
		p11Lib, err = searchForLib(defaultLibPaths)
		if err != nil {
			return nil, err
		}
	} else {
		// The PKCS11 library had value, let's try to open it.
		p11Lib, err = searchForLib(pkcs11Library)
		if err != nil {
			return nil, err
		}
	}

	// override command line parameter with environment variable
	if x, y := os.LookupEnv("P11TOOL_SLOT_LABEL"); y {
		slotLabel = x
	}

	fmt.Printf("Using PKCS#11 library: %s\n", p11Lib)
	fmt.Printf("Using Slot: %s\n", slotLabel)

	// override command line parameter with environment variable
	if x, y := os.LookupEnv("P11TOOL_SLOT_PIN"); y {
		slotPin = x
	}

	if slotPin == "" {
		termState := termInfo{}
		tState, err := terminal.GetState(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}
		termState.termState = tState
		p11Pin, err = termState.askForPin(less)
		if err != nil {
			return nil, err
		}
	} else {
		p11Pin = slotPin
	}

	p11w := pw.Pkcs11Wrapper{
		Library: pw.Pkcs11Library{
			Path: p11Lib,
		},
		SlotLabel: slotLabel,
		SlotPin:   p11Pin,
	}

	if err = p11w.InitContext(); err != nil {
		return nil, err
	}

	if err = p11w.InitSession(); err != nil {
		return nil, err
	}

	if err = p11w.Login(); err != nil {
		return nil, err
	}

	return &p11w, nil
}

// askForPin
func (t *termInfo) askForPin(less bool) (slotPin string, err error) {
	if !less {
		fmt.Println("*** High Security Password Mode Detected ***")
		fmt.Println("*** Preparing SecureRandom Encrypted Memory Space ***")

		for i := 1; i <= 10; i++ {
			if math.Mod(float64(i), 2) == 1 {
				fmt.Printf(". %d%%", i*10)
			} else {
				fmt.Print("...")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("\nEnter Token Password (Pin):")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", fmt.Errorf("Error Getting PIN from Terminal: %v", err)
	}
	slotPin = string(bytePassword)
	bytePassword = []byte{}
	fmt.Println()
	return
}

// searchForLib search comma-separated list of paths for pkcs11 lib
func searchForLib(paths string) (string, error) {
	libPaths := strings.Split(paths, ",")
	for _, path := range libPaths {
		if _, err := os.Stat(strings.TrimSpace(path)); !os.IsNotExist(err) {
			return strings.TrimSpace(path), nil
		}
	}
	return "", fmt.Errorf("no suitable paths for pkcs11 library found: %s", paths)
}
