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

	// "path/filepath"
	"fmt"
	"html/template"
	"os"
	"os/exec"

	"github.com/PabloG6/manticore/tmpl"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// newCmd represents the new command

// var absPathPrefix, absPathErr = filepath.Abs(filepath.Dir(os.Args[0]))


var packageName string;
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a new manticore project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		appName := args[0]

		if len(args) < 1 {
			cobra.CheckErr("new command requires at least project name argument");
		}
		
		if err := os.Mkdir(appName, os.ModePerm);  err != nil {
			
			cobra.CompError(err.Error())

			os.Exit(1)
		}

		if packageName == "" {
			packageName = validateCmdName(args[0])
		}
		os.Chdir(appName)

		
		if err := os.Mkdir("controllers", os.ModePerm);  err != nil{
			
			os.Exit(1)
		}

		if err := os.Mkdir("services", os.ModePerm);  err != nil{
			cobra.CompError("An error occurred when generating this application. ")
			os.Exit(1)
		}

		if err := os.Mkdir("models", os.ModePerm);  err != nil{
			cobra.CompError("An error occurred when generating this application. ")
			os.Exit(1)
		}

		if err := os.Mkdir("middleware", os.ModePerm);  err != nil{
			cobra.CompError("An error occurred when generating this application. ")
			os.Exit(1)
		}
		

		funcMap := template.FuncMap{
			"ToLowerCase": strcase.ToLowerCamel,
		}

 
		appTemplate, err := template.New("app.go").Funcs(funcMap).Parse(string(tmpl.AppTemplate()))
		if(err != nil) {
			panic(fmt.Sprintf("Failed to generate entrypoint for application exited with error: %s", err.Error()))
			
	
		}
		
		f, err := os.Create("app.go")
		if(err != nil) {
			fmt.Printf("Failed to generate entrypoint for application exited with error: %s", err.Error())
	
		}

		if err := appTemplate.Execute(f,  ""); err != nil {
			fmt.Printf("Failed to generate entrypoint for application exited with error: %s", err.Error())
			os.Exit(1)

		 }

		 if output, err := exec.Command("go", "mod", "init", packageName).Output(); err != nil {
			 print("new error");
			 cobra.CheckErr(err.Error())
		 } else {
			 fmt.Println(output);
		 }
		
		 
		 if output, err := exec.Command("go", "get", "-u", ginPkg).Output(); err != nil {
			 print("error");
			cobra.CheckErr(err.Error())
		} else {
			fmt.Println(output);
		}
	   
		 f.Close();


	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVar(&packageName, "packageName", "", "determine the package name of ur application")

	// if absPathErr != nil {
	// 	panic(fmt.Sprintf("unable to extract get root location of application. Exited with %s", absPathErr.Error()))
	// }
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

