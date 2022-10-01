/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
	"golang.org/x/mod/modfile"
)

// makeGatewayCmd represents the make:gateway command
var makeGatewayCmd = &cobra.Command{
	Use:   "make:gateway",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goModBytes, err := os.ReadFile("go.mod")
		if err != nil {
			log.Fatal(err)
		}

		modulePath := modfile.ModulePath(goModBytes)

		p := pluralize.NewClient()
		gateway := args[0]
		gateway = p.Plural(gateway)

		name := args[1]
		filename := strcase.SnakeCase(name) + ".go"

		interfaceDirPath := "pkg/services/" + gateway
		interfaceFilePath := filepath.Join(interfaceDirPath, filename)

		if err := os.MkdirAll(interfaceDirPath, 0755); err != nil {
			log.Fatal(err)
		}

		interfaceTemplate, err := template.New("interfaceTemplate").Parse(interfaceTemplate)
		if err != nil {
			log.Fatal(err)
		}

		interfaceFile, err := os.Create(interfaceFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer interfaceFile.Close()

		if err := interfaceTemplate.Execute(interfaceFile, struct {
			Gateway     string
			GatewayName string
		}{
			Gateway:     gateway,
			GatewayName: strcase.UpperCamelCase(name),
		}); err != nil {
			log.Fatal(err)
		}

		implementsDirPath := "pkg/gateways/" + gateway
		implementsFilePath := filepath.Join(implementsDirPath, filename)

		if err := os.MkdirAll(implementsDirPath, 0755); err != nil {
			log.Fatal(err)
		}

		implementsFile, err := os.Create(implementsFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer implementsFile.Close()

		implementsTemplate, err := template.New("implementsTemplate").Parse(implementsTemplate)
		if err != nil {
			log.Fatal(err)
		}

		if err := implementsTemplate.Execute(implementsFile, struct {
			Gateway     string
			GatewayName string
			ModulePath  string
		}{
			Gateway:     gateway,
			GatewayName: strcase.UpperCamelCase(name),
			ModulePath:  modulePath,
		}); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(makeGatewayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// make:repositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// make:repositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const interfaceTemplate = `
package {{ .Gateway }}

type {{ .GatewayName }}Interface interface {

}

`

const implementsTemplate = `
package {{ .Gateway }}

import "{{ .ModulePath }}/pkg/services/{{ .Gateway }}"

type {{ .GatewayName }} struct {

}

func New{{ .GatewayName }}() {{ .Gateway }}.{{ .GatewayName }}Interface {
	return &{{ .GatewayName }}{}
}

`
