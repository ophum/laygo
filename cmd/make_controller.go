/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

// makeControllerCmd represents the make:controller command
var makeControllerCmd = &cobra.Command{
	Use:   "make:controller",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		filename := strcase.SnakeCase(name) + ".go"

		dirPath := "pkg/controllers"
		filePath := filepath.Join(dirPath, filename)

		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Fatal(err)
		}

		t, err := template.New("controllerTemplate").Parse(controllerTemplate)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := t.Execute(f, struct {
			ControllerName string
		}{
			ControllerName: strcase.UpperCamelCase(name),
		}); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(makeControllerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// make:controllerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// make:controllerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const controllerTemplate = `
package controllers

import (
	"github.com/ophum/laygo/pkg/controllers"
)

type {{ .ControllerName }} struct {
	*controllers.BaseController
}

func New{{ .ControllerName }}() *{{ .ControllerName }} {
	router := controllers.NewRouter("")

	return &{{ .ControllerName }}{
		BaseController: controllers.NewBaseController(router),
	}
}

`
