package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var buildDocs bool

var rootCmd = &cobra.Command{
	Use:   "cred",
	Short: "A password and environment variables manager",
	Long: `Cred is a powerful CLI tool built in Go for managing passwords and environment variables.
It uses GPG encryption to securely store and manage sensitive information.

Examples and usage:
- Initialize with a GPG key: cred init <gpg-key-id>
- Store a new credentials: cred {pass/env} insert <file-name>
- Retrieve a credentials: cred {pass/env} show <file-name>
- Retrieve a credentials: cred {pass/env} copy <file-name>
- List all stored credentials: cred {pass/env} list`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	rootCmd.DisableAutoGenTag = true
	if buildDocs {
		if err := doc.GenMarkdownTree(rootCmd, "./docs/src"); err != nil {
			log.Fatalln("failed to generate docs: ", err)
		}
		log.Println("Documentation generated at ./docs/src")
	}
}

func init() {
	rootCmd.Flags().
		BoolVar(&buildDocs, "generate-docs", false, "Creates markdown documentation for the CLI")
}
