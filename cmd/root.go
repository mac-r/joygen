package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
    "os/exec"
)

func exists(path string) (bool, error) {
    _, err := exec.Command("sh", "-c", "ls " + path).Output()
    if err == nil { return true, nil }
    if err != nil { return false, err }
    return true, err
}

var RootCmd = &cobra.Command{
    Use:   "joygen",
    Short: "joygen is a generator for all things web",
    Long: `Complete documentation is available at https://github.com/mac-r/joygen`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("joygen v0.1")
    },
}
