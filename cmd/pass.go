package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/cmd/pass"
	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// passCmd represents the pass command
var passCmd = &cobra.Command{
	Use:   "pass",
	Short: "A command to manage passwords",
	Long: `The pass command allows you to manage your passwords efficiently.
It provides functionalities to create, update, and delete passwords.

Examples:
- Create a new password: pass {insert/generate}
- Update an existing password: pass edit
- Delete a password: pass rm`,
	Run: func(cmd *cobra.Command, args []string) {
		passPath := config.Constants.PassPath

		if utils.CheckPathExists(passPath) {
			err := utils.PrintTree(passPath, "", true)
			if err != nil {
				fmt.Printf("Failed to parse password store: %v\n", err)
			}
		}
	},
}

func init() {
	passCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !utils.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("GPG key not found, try [cred init <gpg-key-id>]")
		}

		if !utils.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("Invalid GPG key, try [cred init <gpg-key-id>]")
		}

		return nil
	}

	passCmd.AddCommand(pass.GenerateCmd)
	passCmd.AddCommand(pass.InsertCmd)
	passCmd.AddCommand(pass.ShowCmd)
	passCmd.AddCommand(pass.CopyCmd)
	passCmd.AddCommand(pass.EditCmd)
	passCmd.AddCommand(pass.LsCmd)
	rootCmd.AddCommand(passCmd)
}
