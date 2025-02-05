package git

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
)

func AutoGit(cmd *cobra.Command) error {
	cmdName := cmd.Name()

	if !CheckGitExists() {
		cmd.SilenceUsage = true
		return fmt.Errorf("git is not installed")
	}

	if !IsValidGitPath(config.Constants.StorePath) {
		var choice string
		fmt.Print(
			"github repository not found in the store path. Do you want to initialize a new repository? (y/n): ",
		)
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			cmd.SilenceUsage = true
			return fmt.Errorf("git repository not found")
		}

		if err := InitRepo(config.Constants.StorePath); err != nil {
			cmd.SilenceUsage = true
			return err
		}
	}

	if !HaveDiff(config.Constants.StorePath) {
		return nil
	}

	if err := AddFiles(config.Constants.StorePath); err != nil {
		cmd.SilenceUsage = true
		return err
	}

	commitMessage := fmt.Sprintf("Auto commit by %v command", cmdName)
	if err := CommitFiles(config.Constants.StorePath, commitMessage); err != nil {
		cmd.SilenceUsage = true
		return err
	}

	if err := PushRepo(config.Constants.StorePath); err != nil {
		cmd.SilenceUsage = true
		return err
	}

	return nil
}
