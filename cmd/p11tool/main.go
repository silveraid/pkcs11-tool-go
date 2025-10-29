package main

import (
	"fmt"

	"github.com/scottallan/p11tool-new/cmd/p11tool/cmd"
)

func main() {

	// Print out header
	fmt.Println("p11tool v0.0.7 (2025-10-29)")
	fmt.Println("by George Bolo, Scott Alan, and Frank Felhoffer")
	fmt.Println("-----------------------------------------------")
	fmt.Println()

	// Cobra to do it's magic
	cmd.Execute()
}

//
//func main() {
//
//	// get flags
//	pkcs11Library := flag.String("lib", "", "Location of pkcs11 library")
//	slotLabel := flag.String("slot", "ForFabric", "Slot Label")
//	slotPin := flag.String("pin", "", "Slot PIN")
//	action := flag.String("action", "list", "list,import,generate,generateAndImport,generateSecret,generateAES,generateDES,wrapKeyWithDES3,unwrapASYMWithDES3,getSKI,getSkiFromCert,getSkiFromB64Cert,SignHMAC384,TestAESGCM,generateCSR,importCert,deleteObj")
//	keyFile := flag.String("keyFile", "/some/dir/key.pem)", "path to key you want to import or getSKI")
//	keyType := flag.String("keyType", "EC", "Type of key (EC,RSA,GENERIC_SECRET,AES,SHA256_HMAC,SHA384_HMAC,DES3)")
//	keyLen := flag.Int("keyLen", 32, "Key Length for CKK_GENERIC_SECRET (32,48,...)")
//	keyLabel := flag.String("keyLabel", "tmpkey", "Label of CKK_GENERIC_SECRET")
//	keyStore := flag.String("keyStore", "file", "Keystore Type (file,pkcs12)")
//	keyStorepass := flag.String("keyStorepass", "securekey", "Keystore Storepass")
//	key := flag.String("key", "", "Key as HEX String")
//	csrInfo := flag.String("csrInfo", "", "json file with values for CSR Creation")
//	wrapKey := flag.String("wrapKey", "wrapKey", "DES3 Wrapping Key for unwrapping key material onto Gemalto")
//	objClass := flag.String("objClass", "", "CKA_CLASS for Deletion of Objects")
//	outF := flag.String("outFile", "out.pem", "output file for CSR Generation")
//	noDec := flag.Bool("noDec", false, "when set wrapped material will remain encrypted")
//	less := flag.Bool("less", true, "Dont show password preamble")
//
//	byCKAID := flag.Bool("byCKAID", false, "when set we will assume keyLabel is a CKA_ID represented as a string")
//
//	mechOver := flag.String("mechanismOverride", "", "Allow override of mechanism - only supported on certain operations [wrapKeyWithDES3, wrapKeyWithAES]")
//
//	maxObjectsToList := flag.Int("maxObjectsToList", 50, "Paramter to be used with -action list to specify how many objects to print")
//
//	var gracefulStop = make(chan os.Signal)
//	signal.Notify(gracefulStop, syscall.SIGTERM)
//	signal.Notify(gracefulStop, syscall.SIGINT)
//
//	flag.Parse()
//
//	var err error
//	//Neet to Get State of the Existing Terminal
//	termState := termInfo{}
//	tState, err := terminal.GetState(int(syscall.Stdin))
//	termState.termState = tState
//	go func() {
//		sig := <-gracefulStop
//		var err error
//		fmt.Printf("\n**********caught signal: %+v  EXITING\n", sig)
//		cState, err := terminal.GetState(int(syscall.Stdin))
//		if err != nil {
//			panic(err)
//		}
//		termState.curState = cState
//		if termState.curState == termState.termState {
//			fmt.Println("Terminal State OK!  Exiting Normally")
//			panic(err)
//		} else {
//			fmt.Printf("Terminal State Changed!\n[Current State: %v]\n[Original State :%v] Reverting before Exiting\n", *termState.curState, *termState.termState)
//			err = terminal.Restore(int(syscall.Stdin), termState.termState)
//			panic(err)
//		}
//	}()
//
//	// complete actions which do not require HSM
//	switch *action {
//
//
//	// complete actions which require HSM
//
//	// initialize pkcs11
//	var p11Lib string
//	var p11Pin string
//
//	if *pkcs11Library == "" {
//		p11Lib, err = searchForLib(defaultLibPaths)
//		exitWhenError(err)
//	} else {
//		p11Lib, err = searchForLib(*pkcs11Library)
//		exitWhenError(err)
//	}
//	if *slotPin == "" {
//		p11Pin, err = termState.askForPin(*less)
//		if err != nil {
//			exitWhenError(err)
//		}
//	} else {
//		p11Pin = *slotPin
//	}
//
//	p11w = pw.Pkcs11Wrapper{
//		Library: pw.Pkcs11Library{
//			Path: p11Lib,
//		},
//		SlotLabel: *slotLabel,
//		SlotPin:   p11Pin,
//	}
//
//	err = p11w.InitContext()
//	exitWhenError(err)
//
//	err = p11w.InitSession()
//	exitWhenError(err)
//
//	err = p11w.Login()
//	exitWhenError(err)
//
//	// defer cleanup
//	defer p11w.Context.Destroy()
//	defer p11w.Context.Finalize()
//	defer p11w.Context.CloseSession(p11w.Session)
//	defer p11w.Context.Logout(p11w.Session)
//	defer termState.cleanupPin(*slotPin, &p11Pin, *less)
//
//	switch *action {
//

//
//	case "generateAndImport":
//		if *keyType == "RSA" {
//			rsa := pw.RsaKey{}
//			rsa.Generate(2048)
//			p11w.ImportRSAKey(rsa)
//		} else {
//			ec := pw.EcdsaKey{}
//			// TODO: fix non working curves (P-521)
//			ec.Generate("P-256")
//			p11w.ImportECKey(ec)
//		}
//
