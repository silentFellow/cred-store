package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func RmLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v rm <filepath>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else {
		basePath = config.Constants.EnvPath
	}

	if len(args) < 1 {
		fmt.Println("invalid usage, expected: ", usage)
		return
	}

	for _, path := range args {
		fullPath := paths.BuildPath(basePath, path)

		if !paths.CheckPathExists(fullPath) {
			fmt.Printf("%v not found\n", fullPath)
			continue
		}

		if err := os.RemoveAll(fullPath); err != nil {
			fmt.Printf("failed to remove %v: %v\n", fullPath, err)
			continue
		}

		fmt.Printf("%v deleted successfully\n", path)
	}
}
