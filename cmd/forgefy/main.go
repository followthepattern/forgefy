package main

import (
	"fmt"
	"os"
	"time"

	"github.com/followthepattern/forgefy"
	"github.com/followthepattern/forgefy/forgeio"
	"github.com/followthepattern/forgefy/plugins/gobackend"
	"github.com/followthepattern/forgefy/plugins/monorepo"
	"github.com/followthepattern/forgefy/plugins/reactfrontend"
	"github.com/spf13/cobra"
)

func main() {
	var file string
	var outputDir string
	var exclude string
	var cleanUp bool

	var root = &cobra.Command{
		Use:   "forgefy",
		Short: "forgefy generates web applications from a simple config.",
		Long:  `forgefy generates web applications using golang template framework from a forge config file, that uses yaml syntax.`,
		Run: func(cmd *cobra.Command, args []string) {
			if file == "" {
				fmt.Fprintln(os.Stdout, "file can't be empty")
				cmd.Help()
				os.Exit(1)
			}
			if outputDir == "" {
				fmt.Fprintln(os.Stdout, "output can't be empty")
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	root.Flags().StringVarP(&file, "file", "f", "", "filepath to the forge file")
	root.Flags().StringVarP(&outputDir, "output", "o", "", "output directory path")
	root.Flags().StringVarP(&exclude, "exclude", "i", "", "specify regex to exclude certain files from forging")
	root.Flags().BoolVarP(&cleanUp, "cleanup", "c", false, "set to true if you want to delete the previous forge")

	if err := root.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if cleanUp {
		if err := os.RemoveAll(outputDir); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	forgeFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	f := forgefy.New()

	f.InstallPlugins(monorepo.MonoRepo{}, gobackend.Plugin{}, reactfrontend.Plugin{})

	fw := forgeio.NewFileWriter(outputDir)

	now := time.Now()

	productName, err := f.Forge(
		string(forgeFile),
		fw,
		forgefy.WithExclude(exclude))

	if err != nil {
		fmt.Printf("during forging %s following error occured: %s", productName, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s is successfully forged to %s\n", productName, outputDir)
	fmt.Println("forged time", time.Since(now))
}
