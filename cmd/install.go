package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "bufio"
    "os/exec"
)

func init() {
    RootCmd.AddCommand(installCmd)
    installCmd.Flags().StringVarP(&GithubUser, "githubUser", "u", "mac-r", "Github user to fetch templates from")
}

var GithubUser string
var RewriteFolder bool = false

func checkGit() (bool, error) {
  _, err := exec.Command("sh", "-c", "git --version").Output()
  if err == nil { return true, nil }
  if err != nil {
    fmt.Printf("Git is not installed!\n")
    return false, err
  }
  return true, err
}

var installCmd = &cobra.Command{
    Use:   "install",
    Short: "Fetch templates from a github user",
    Long:  `Details:
  Initial set of templates is introduced at https://github.com/mac-r/joygen-templates.
  You can easily create your own by forking it and providing your github user name.
  Templates will be installed at ~/.joygen/joygen-templates.
    `,
    Run: func(cmd *cobra.Command, args []string) {
        if RewriteFolder == true {
          coreDirStatus, _ := exists("~/.joygen/")
          if coreDirStatus == false { exec.Command("sh", "-c", "mkdir ~/.joygen").Output() }
          if coreDirStatus == true {
            templatesDirStatus, _ := exists("~/.joygen/joygen-templates")
            if templatesDirStatus == true {
              exec.Command("sh", "-c", "rm -rf ~/.joygen/joygen-templates").Output()
            }
          }
          exec.Command("sh", "-c", "mkdir ~/.joygen/joygen-templates").Output()
          fmt.Println("Cleaned up templates folder at ~/.joygen/joygen-templates. Fetching templates from "+ GithubUser +"/joygen-templates...")
          _, err := exec.Command("sh", "-c", "git clone git@github.com:" + GithubUser + "/joygen-templates.git ~/.joygen/joygen-templates").Output()
          if err == nil {
            fmt.Println("Your local templates are updated. Yay! :)")
          } else {
            fmt.Println("Some problems during update. Does " + GithubUser + " have this repo? Also you might have internet connection issues.")
          }
        }
    },
    PreRun: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Updating local joygen templates...\n")
        fmt.Printf("Using " + GithubUser + "/joygen-templates as a source. \n")
        fmt.Printf("Would you like to rewrite your local templates if they exist? y/N\n")

        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
          if scanner.Text() == "n" || scanner.Text() == "N" || scanner.Text() == "" {
            fmt.Printf("Installation was cancelled.\n")
            os.Exit(-1)
          } else if scanner.Text() == "y" || scanner.Text() == "Y" {
            gitStatus, _ := checkGit()
            if gitStatus == false { os.Exit(-1) }
            RewriteFolder = true
            break
          }
        }
    },

}
