package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of MyCLI",
    Long:  `All software has versions. This is MyCLI's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("MyCLI version: %s\n", version)
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}