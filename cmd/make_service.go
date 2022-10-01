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

// makeServiceCmd represents the make:service command
var makeServiceCmd = &cobra.Command{
	Use:   "make:service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		filename := strcase.SnakeCase(name) + ".go"

		dirPath := "pkg/services"
		filePath := filepath.Join(dirPath, filename)

		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Fatal(err)
		}

		t, err := template.New("serviceTemplate").Parse(serviceTemplate)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := t.Execute(f, struct {
			ServiceName string
		}{
			ServiceName: strcase.UpperCamelCase(name),
		}); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(makeServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// make:serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// make:serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const serviceTemplate = `
package services

type {{ .ServiceName }}Interface interface {

}

type {{ .ServiceName }} struct {

}

func New{{- .ServiceName -}}() {{ .ServiceName }}Interface {
	return &{{ .ServiceName }}{}
}

`
