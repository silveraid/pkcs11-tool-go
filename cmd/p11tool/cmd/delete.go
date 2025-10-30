package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ckaClass string
)

// rmCmd represents the generate key command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes one or more objects from the HSM",
	Long: `Removes one or more objects from the HSM

Example usage:
 p11tool rm --keyLabel "my-ec-key"
 p11tool rm --keyLabel "my-ec-key" --ckaClass CKO_PUBLIC_KEY
 p11tool rm --ckaClass ALL`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if ckaClass == "" {
			return fmt.Errorf("must provide --ckaClass")
		}

		if ckaClass != "ALL" && keyLabel == "" {
			return fmt.Errorf("must provide --keyLabel")
		}

		p11w, err := initP11Context()
		if err != nil {
			return err
		}

		defer p11w.Context.Destroy()
		defer p11w.Context.Finalize()
		defer p11w.Context.CloseSession(p11w.Session)
		defer p11w.Context.Logout(p11w.Session)
		//defer termState.cleanupPin(*slotPin, &p11Pin, *less)

		if ckaClass == "ALL" {
			// deletes all objects from a given slot
			p11w.DeleteObj("ALL", "")
		} else {
			p11w.DeleteObj(ckaClass, keyLabel)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().StringVar(&ckaClass, "ckaClass", "", "CKA CLASS (ALL, CKO_PUBLIC_KEY, CKO_PRIVATE_KEY, CKO_SECRET_KEY)")
	rmCmd.Flags().StringVar(&keyLabel, "keyLabel", "", "Key label (aka. CKA LABEL)")
}
