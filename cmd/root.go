package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tix",
	Short: "tix is a to-do list app.",
	Long:  `tix is a to-do list app.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile  string
	dataFile string
)

func initConfig() {
	viper.SetConfigName(".tix")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("tix")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config: ", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	home, err := user.Current()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home.HomeDir+string(os.PathSeparator)+".todo.json", "File to store todo items")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (defaults to $HOME/.tix.yaml")
}
