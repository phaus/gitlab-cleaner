package cmd

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Verbose - do verbose output.
var Verbose bool

// DryRun - do a dry run - don't change anything.
var DryRun bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A brief description of your application",
	Long:  `This Application cleans the Gitlab Registry based on some simple rules.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry", "d", false, "dry run")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	_, hasToken := os.LookupEnv("PRIVATE_ACCESS_TOKEN")
	_, hasURL := os.LookupEnv("CI_PROJECT_URL")
	if !hasToken || !hasURL {
		log.Fatal("You need to set 'CI_PROJECT_URL' and 'PRIVATE_ACCESS_TOKEN'")
	}

	viper.Set("Token", os.Getenv("PRIVATE_ACCESS_TOKEN"))
	viper.Set("ProjectUrl", os.Getenv("CI_PROJECT_URL"))

	u, err := url.Parse(viper.GetString("ProjectUrl"))
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("BaseUrl", fmt.Sprintf("%s://%s", u.Scheme, u.Host))
	viper.Set("RegistryUrl", fmt.Sprintf("%s/container_registry.json", viper.GetString("ProjectUrl")))

}
