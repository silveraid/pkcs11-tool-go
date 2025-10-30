package cmd

//	case "unwrapASYMWithDES3":
//		w, _, err := p11w.FindObjects([]*pkcs11.Attribute{
//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, *wrapKey),
//		},
//			1,
//		)
//		exitWhenError(err)
//
//		switch *keyType {
//		case "EC":
//			err := p11w.UnWrapECKeyFromFile(*keyFile, *keyStore, *keyStorepass, *keyLabel, w[0])
//			exitWhenError(err)
//		case "RSA":
//			err := p11w.UnWrapRSAKeyFromFile(*keyFile, *keyStore, *keyStorepass, *keyLabel, w[0])
//			exitWhenError(err)
//		case "AES":
//		   wrappedKey, err := p11w.WrapSymKey("AES", *key, *keyLen, w[0])
//		   if err != nil {
//			   fmt.Printf("Unable to Wrap key: %v\n", *key)
//		   }
//		   fmt.Printf("Wrapped Key to Value: %v\n", wrappedKey)
//		   //Unwrap key onto HSM
//		   err = p11w.UnwrapSymKey("AES", wrappedKey, w[0], *keyLabel)
//		   exitWhenError(err)
//		}
//
//	case "wrapKeyWithDES3":
//		w, _, err := p11w.FindObjects([]*pkcs11.Attribute{
//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, *wrapKey),
//		},
//			1,
//		)
//		exitWhenError(err)
//
//		var wrappedKey []byte
//		switch *keyType {
//		case "RSA":
//			wrappedKey, err = p11w.WrapP11Key("DES3", *objClass, *keyLabel, w[0], *byCKAID, *mechOver)
//			exitWhenError(err)
//			decryptedKey, err := p11w.DecryptP11Key("DES3", wrappedKey, w[0], *mechOver)
//			outFile, err := os.Create(*outF)
//			if err != nil {
//				fmt.Printf("Unable to write key %s", err.Error())
//				return
//			}
//			defer outFile.Close()
//
//			fmt.Printf("writing key to %s\n", *outF)
//			if *noDec {
//				fmt.Printf("writing encrypted\n?")
//				err = ioutil.WriteFile(*outF, wrappedKey, 0644)
//			} else {
//				fmt.Printf("writing decrypted\n")
//				err = ioutil.WriteFile(*outF, decryptedKey, 0644)
//			}
//
//			if err != nil {
//				return
//			}
//		case "EC":
//			fmt.Printf("Need to Implement EC Key Wrapping")
//
//		}
//
//	case "wrapKeyWithAES":
//		w, _, err := p11w.FindObjects([]*pkcs11.Attribute{
//			pkcs11.NewAttribute(pkcs11.CKA_LABEL, *wrapKey),
//		},
//			1,
//		)
//		exitWhenError(err)
//
//		var wrappedKey []byte
//		switch *keyType {
//		case "RSA":
//			wrappedKey, err = p11w.WrapP11Key("AES", *objClass, *keyLabel, w[0], *byCKAID, *mechOver)
//			exitWhenError(err)
//			decryptedKey, err := p11w.DecryptP11Key("AES", wrappedKey, w[0], *mechOver)
//			outFile, err := os.Create(*outF)
//			if err != nil {
//				fmt.Printf("Unable to write key %s", err.Error())
//				return
//			}
//			defer outFile.Close()
//
//			fmt.Printf("writing key to %s\n", *outF)
//			if *noDec {
//				fmt.Printf("writing encrypted\n?")
//				err = ioutil.WriteFile(*outF, wrappedKey, 0644)
//			} else {
//				fmt.Printf("writing decrypted\n")
//				err = ioutil.WriteFile(*outF, decryptedKey, 0644)
//			}
//
//			if err != nil {
//				return
//			}
//		case "EC":
//			fmt.Printf("Need to Implement EC Key Wrapping")
//
//		}
//
