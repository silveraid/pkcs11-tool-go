package cmd

//	case "getSKI":
//		if *keyType == "RSA" {
//			key := pw.RsaKey{}
//			err = key.ImportPrivKeyFromFile(*keyFile)
//			exitWhenError(err)
//			key.GenSKI()
//			fmt.Printf("SKI(sha256): %s\n", key.SKI.Sha256)
//			os.Exit(0)
//		} else {
//			key := pw.EcdsaKey{}
//			err = key.ImportPrivKeyFromFile(*keyFile)
//			exitWhenError(err)
//			key.GenSKI()
//			fmt.Printf("SKI(sha256): %s\n", key.SKI.Sha256)
//			os.Exit(0)
//		}
//
//	case "getSkiFromCert":
//		key := pw.EcdsaKey{}
//		err = key.ImportPubKeyFromCertFile(*keyFile)
//		exitWhenError(err)
//		key.GenSKI()
//		fmt.Printf("SKI(sha256): %s\n", key.SKI.Sha256)
//		os.Exit(0)
//
//	case "getSkiFromB64Cert":
//		key := pw.EcdsaKey{}
//		err = key.ImportPubKeyFromBase64Cert(*keyFile)
//		exitWhenError(err)
//		key.GenSKI()
//		fmt.Printf("SKI(sha256): %s\n", key.SKI.Sha256)
//		os.Exit(0)
//
//	}
