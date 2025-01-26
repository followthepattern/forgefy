package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

	excludedFiles := createExcludeMap(exclude)

	if cleanUp {
		err := remove(outputDir, excludedFiles)
		if err != nil {
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
		forgefy.WithExcludedFiles(excludedFiles))

	if err != nil {
		fmt.Printf("during forging %s following error occured: %s", productName, err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s is successfully forged to %s\n", productName, outputDir)
	fmt.Println("forged time", time.Since(now))
}

func createExcludeMap(exclude string) map[string]struct{} {
	exceptions := strings.Split(exclude, ",")

	excludedFiles := make(map[string]struct{})
	for _, exception := range exceptions {
		excludedFiles[exception] = struct{}{}
	}

	return excludedFiles
}

func remove(projectPath string, excludedFiles map[string]struct{}) error {
	err := filepath.Walk(projectPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if projectPath == path {
			return nil
		}

		relPath, _ := filepath.Rel(projectPath, path)

		if info.IsDir() {
			if !pathToExclude(excludedFiles, relPath) {
				err := os.RemoveAll(path)
				if err != nil {
					return err
				}
				return filepath.SkipDir
			}
			return nil
		}

		if pathToExclude(excludedFiles, relPath) {
			return nil
		}

		return os.Remove(path)
	})

	return err
}

func pathToExclude(m map[string]struct{}, relPath string) (exclude bool) {
	if relPath == "." {
		return true
	}

	for key := range m {
		if len(key) > len(relPath) {
			exclude = strings.HasPrefix(key, relPath)
		} else {
			exclude = strings.HasPrefix(relPath, key)
		}
		if exclude {
			return
		}
	}

	return false
}
