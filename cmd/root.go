/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	env       string
	migration string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:       "Bitwyre",
	Short:     "Bitwyre CLI Utility System",
	Long:      `Bitwyre CLI Utility System`,
	ValidArgs: []string{"local", "dev", "prod"},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		startServer(cmd)
	},
}

var serveCmd = &cobra.Command{
	Use:       "serve",
	Short:     "Run Application Server via :3000",
	ValidArgs: []string{"local", "dev", "prod"},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		startServer(cmd)
	},
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Execute Database Action",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&env, "env", "e", "local", "Environment: local | dev | prod")
	serveCmd.Flags().StringVarP(&migration, "migration", "m", "n", "Auto Database Migration: y | n")
	rootCmd.AddCommand(databaseCmd)

}

func startServer(cmd *cobra.Command) bool {
	// Check Database Migration
	if migration != "" {
		os.Setenv("migration", migration)
	}

	for _, v := range cmd.ValidArgs {
		if v == env {
			os.Setenv("env", env)
			Server()
			return true
		}
	}

	log.Fatal("Env value only valid with: ", cmd.ValidArgs)
	return false
}


