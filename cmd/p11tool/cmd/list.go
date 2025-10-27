package cmd

import (
	"github.com/miekg/pkcs11"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List objects in the PKCS#11 token",
	Long: `List all objects or specific types of objects in the PKCS#11 token.
For example:
  p11tool list --keyType EC    # List only EC keys
  p11tool list                 # List all objects`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// FIXME
		maxObjectsToList := 100

		p11w, err := initP11Context()
		if err != nil {
			return err
		}
		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)
		//defer termState.cleanupPin(*slotPin, &p11Pin, *less)

		p11w.ListObjects(
			[]*pkcs11.Attribute{},
			maxObjectsToList,
		)

		//defer p11w.Finalize()
		//defer p11w.CloseSession()
		//
		//objects, err := p11w.ListObjects(keyType, objClass, maxObjects)
		//if err != nil {
		//	return fmt.Errorf("failed to list objects: %v", err)
		//}
		//
		//if len(objects) == 0 {
		//	fmt.Println("No objects found")
		//	return nil
		//}
		//
		//for _, obj := range objects {
		//	fmt.Printf("CKA_LABEL: %s\n", obj.Label)
		//	fmt.Printf("CKA_CLASS: %s\n", obj.Class)
		//	fmt.Printf("CKA_ID: %x\n", obj.CkaID)
		//	if obj.KeyType != "" {
		//		fmt.Printf("CKA_KEY_TYPE: %s\n", obj.KeyType)
		//	}
		//	fmt.Println("----------------------------------------")
		//}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&objClass, "class", "", "Object class (public,private,secret,cert,data)")
	listCmd.Flags().IntVar(&maxObjects, "max", 100, "Maximum number of objects to list")
}
