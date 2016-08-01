package cmd

import (
  "github.com/spf13/cobra"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "fmt"
  "os"
  "log"
)

// short: 'isomorphic react browserify short info'
// long: 'isomorphic react browserify long info'
// author: 'mac-r'
// docs: 'some link'
// variables:
//   - name

var ParamsStore =  make(map[string]*string)

type templateConf struct {
    Short string `yaml:"short"`
    Long string `yaml:"long"`
    Author string `yaml:"author"`
    Docs string `yaml:"docs"`
    Variables []string `yaml:"variables"`
}

func (c *templateConf) getConf(file string) *templateConf {
    yamlFile, err := ioutil.ReadFile(file)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}

func isDirectory(path string) (bool, error) {
    fileInfo, err := os.Stat(path)
    return fileInfo.IsDir(), err
}

func readTemplates() (bool, error)  {
  files, _ := ioutil.ReadDir(JoygenTemplatesPath)
  for _, f := range files {
    if f.Name() != ".git" && f.Name() != "README.md" && f.Name() != "LICENSE" {
      var c templateConf

      commandName := f.Name()
      configFile := JoygenTemplatesPath + f.Name() + "/info.yml"
      info := c.getConf(configFile)

      templateCommand := &cobra.Command{
          Use:   commandName,
          Short: info.Short,
          Long:  info.Long + "\n\n" + "Docs url: " + info.Docs + "\n" + "Author: " + info.Author,
          Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("hello from -> " + commandName)
            // fmt.Println(*ParamsStore["name"])

            // TODO:
            // validate that all variables have values
          },
      }

      RootCmd.AddCommand(templateCommand)

      for _, flagVariable := range info.Variables {
        var flagValue string
        ParamsStore[flagVariable] = &flagValue
        templateCommand.Flags().StringVarP(&flagValue, flagVariable, "", "", "All @@@@" + flagVariable + "@@@@ in a template will be replaced.")
        // TODO:
        // add a flag for the new folder name
      }
    }
  }

  return true, nil
}


func init() {
  readTemplates()
}
