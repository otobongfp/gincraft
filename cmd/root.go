package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "gincraft",
	Short: "A CLI tool to scaffold Gin projects",
	Long: `GinCraft is a CLI tool that helps you quickly scaffold new Gin projects
with a standard structure and basic setup.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
