package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.versioner.yaml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.SetDefault("config", "")
}

// Get us straight into the cobra command.
func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "versioner",
	Short: "versioner is a Cuttle~ tool to manage the version of your project.",
	Long:  `versioner is a Cuttle~ tool to manage the version of your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

/*
rootCmd.Run = func(cmd *cobra.Command, args []string) {
	cmd.Usage()
}
*/

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func initConfig() {
	// Add flags here
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		pwd, err := os.Getwd()
		cobra.CheckErr(err)

		homeConfig := filepath.Join(home, ".config")
		if home != pwd {
			viper.AddConfigPath(home)
		} else {
			homeConfig = filepath.Join(pwd, ".versioner")
		}

		// Search config in home directory with name ".versioner.yaml".
		viper.AddConfigPath(homeConfig)
		viper.AddConfigPath(pwd)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".versioner")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	cobra.CheckErr(err)
}
