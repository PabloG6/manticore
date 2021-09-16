/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"html/template"
	"os"
	
	"unicode"

	"github.com/PabloG6/manticore/tmpl"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// controllerCmd represents the controller command
var controllerCmd = &cobra.Command{
	Use:   "controller",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Inside controller command");
	funcMap := template.FuncMap{
		"ToLowerCase": strcase.ToLowerCamel,
		"ToUpperCase": strcase.ToCamel,
	}

	cmdName := validateCmdName(args[0])
	var commandArgs = tmpl.ControllerCommandArgs{ControllerName: cmdName,}
	if len(args) < 1 {
		cobra.CheckErr("generate controller requires exactly one argument name");
	}

	controllerTemplate, err := template.New("controller.go").Funcs(funcMap).Parse(string(tmpl.ControllerTemplate()))
	if(err != nil) {
		cobra.CheckErr(err.Error())
		

	}

	wd, err := os.Getwd()
	
	if err != nil {
		cobra.CheckErr(err.Error())
	}
	f, err := os.Create(fmt.Sprintf("%s/controllers/%s.go", wd, cmdName))
	if(err != nil) {
		cobra.CheckErr(err.Error())

	}

	if err := controllerTemplate.Execute(f,  commandArgs); err != nil {
		cobra.CheckErr(err.Error())
	
	}

	
},
}

func init() {
	generateCmd.AddCommand(controllerCmd)

	
}


func validateCmdName(source string) string {
	i := 0
	l := len(source)
	// The output is initialized on demand, then first dash or underscore
	// occurs.
	var output string

	for i < l {
		if source[i] == '-' || source[i] == '_' {
			if output == "" {
				output = source[:i]
			}

			// If it's last rune and it's dash or underscore,
			// don't add it output and break the loop.
			if i == l-1 {
				break
			}

			// If next character is dash or underscore,
			// just skip the current character.
			if source[i+1] == '-' || source[i+1] == '_' {
				i++
				continue
			}

			// If the current character is dash or underscore,
			// upper next letter and add to output.
			output += string(unicode.ToUpper(rune(source[i+1])))
			// We know, what source[i] is dash or underscore and source[i+1] is
			// uppered character, so make i = i+2.
			i += 2
			continue
		}

		// If the current character isn't dash or underscore,
		// just add it.
		if output != "" {
			output += string(source[i])
		}
		i++
	}

	if output == "" {
		return source // source is initially valid name.
	}
	return output
}
