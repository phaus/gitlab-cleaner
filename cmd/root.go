package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cleaner",
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
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	_, hasToken := os.LookupEnv("PRIVATE_ACCESS_TOKEN")
	_, hasURL := os.LookupEnv("CI_PROJECT_URL")
	if !hasToken || !hasURL {
		log.Println("You need to set 'CI_PROJECT_URL' and 'PRIVATE_ACCESS_TOKEN'")
		os.Exit(0)
	}

	viper.Set("Token", os.Getenv("PRIVATE_ACCESS_TOKEN"))
	viper.Set("ProjectUrl", os.Getenv("CI_PROJECT_URL"))

	u, err := url.Parse(viper.GetString("ProjectUrl"))
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("BaseUrl", fmt.Sprintf("%s://%s", u.Scheme, u.Host))
	viper.Set("RegistryUrl", fmt.Sprintf("%s/container_registry.json", viper.GetString("ProjectUrl")))

	fmt.Printf("BaseUrl %s\n", viper.GetString("BaseUrl"))
	fmt.Printf("ProjectUrl %s\n", viper.GetString("ProjectUrl"))
	fmt.Printf("RegistryUrl %s\n", viper.GetString("RegistryUrl"))
}

// GetClient - returns the default HTTP Client
func GetClient() *http.Client {
	client := &http.Client{}
	return client
}

// SetDefaultHeaders - adds the default HTTP Headers to the Request
func SetDefaultHeaders(req *http.Request) {
	req.Header.Add("Private-Token", viper.GetString("Token"))
	req.Header.Add("accept", "application/json, text/plain, */*")
}
