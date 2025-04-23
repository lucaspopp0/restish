package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

const (
	shellBash       = "bash"
	shellZsh        = "zsh"
	shellFish       = "fish"
	shellPowershell = "powershell"
)

var (
	completableShells = []string{
		shellBash,
		shellZsh,
		shellFish,
		shellPowershell,
	}
)

var completionCmdLong = strings.Join([]string{
	"Generate a completion script for the specified shell, and write to stdout.",
	"",
	"This script can be used to enable tab completion of your current restish configuration.",
	"",
	"For details on how to use the output of this command, review your shell's documentation regarding command completion.",
}, "\n")

func completionCmdUsage() string {
	shellsList := []string{}
	for _, shell := range completableShells {
		shellsList = append(shellsList, fmt.Sprintf("  %s", shell))
	}

	return strings.Join([]string{
		"completion <shell>",
		"",
		"Supported Shells:",
		strings.Join(shellsList, "\n"),
	}, "\n")
}

func completion(
	shell string,
	cmd *cobra.Command,
) {
	switch shell {
	case shellBash:
		cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	case shellZsh:
		cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case shellFish:
		cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
	case "powershell":
		cmd.Root().GenPowerShellCompletionWithDesc(cmd.OutOrStdout())
	}
}
