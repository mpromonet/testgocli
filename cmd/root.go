package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "mycli",
    Short: "MyCLI is a simple CLI application",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Do Stuff Here
        fmt.Println("Hello from MyCLI!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
}
