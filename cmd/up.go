package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/0xTanvir/up/cfg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// UpCmd is responsible for config loading and bootstrapping.
var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "A platform for UIU programmers",
	Long:  "up is the main command, used to build uiu programmer application.",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	AddCommands()

	if err := UpCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

var (
	// cfgName config file name
	cfgName string
	// cfgPaths config file path
	cfgPaths   string
	getVersion bool

	// Version of the application set at build time
	Version = "undefined"
	// BuildDate of the application set at build time
	BuildDate string
	// GitCommit of the application set at build time
	GitCommit string
	// GitBranch of the application set at build time
	GitBranch string
)

// AddCommands adds child commands to the root command UpCmd.
func AddCommands() {
	UpCmd.AddCommand(versionCmd)
}

// init initializes flags.
func init() {
	flags := UpCmd.PersistentFlags()
	flags.StringVar(&cfgName, "cfg-name", "development", "config file name without path and extension")
	flags.StringVar(&cfgPaths, "cfg-paths", "./etc", "paths where we search config split them by ','")
	flags.BoolVar(&getVersion, "version", false, "build information")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if getVersion {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("BuildDate: %s\n", BuildDate)
		fmt.Printf("GitCommit: %s\n", GitCommit)
		fmt.Printf("GitBranch: %s\n", GitBranch)
		os.Exit(0)
	}

	// ENV variables allow easy configuration from docker
	viper.SetEnvPrefix("UP")
	viper.AutomaticEnv()

	// We can overrule the config name with an ENV variable
	// For docker we may be used production.yaml with
	// different configuration
	if len(os.Getenv("UP_CFG_NAME")) > 0 {
		cfgName = os.Getenv("UP_CFG_NAME")
	}

	// Now we load config with viper
	if err := cfg.Load(cfgName, strings.Split(cfgPaths, ",")...); err != nil {
		panic(err.Error())
	}
	log.Infof("loaded config: %v", viper.ConfigFileUsed())
}
