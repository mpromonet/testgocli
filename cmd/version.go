package cmd

import (
    _ "embed"
    "fmt"
    "github.com/spf13/cobra"
)

//go:generate sh -c "printf '%s' $(git describe --tags --always --dirty) > version.txt"
//go:embed version.txt
var version string

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of MyCLI",
    Long:  `All software has versions. This is MyCLI's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("MyCLI version: %v\n", version)
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}