package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// Registry a gitlab Registry info object
type Registry struct {
	ID          int64  `json:"id"`
	Path        string `json:"path"`
	Location    string `json:"location"`
	TagsPath    string `json:"tags_path"`
	DestroyPath string `json:"destroy_path"`
}

// RegistryTag a tagged image withing the Registry
type RegistryTag struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	Revision      string `json:"revision"`
	ShortRevision string `json:"short_revision"`
	TotalSize     int64  `json:"total_size"`
	CreatedAt     string `json:"created_at"`
	DestroyPath   string `json:"destroy_path"`
}

func init() {
	rootCmd.AddCommand(cleanerCmd)
}

var cleanerCmd = &cobra.Command{
	Use:   "cleaner",
	Short: "clean the gitlab registry.",
	Long:  `This cleans a gitlab Registry. You need to set 'CI_PROJECT_URL' and 'PRIVATE_ACCESS_TOKEN'`,
	Run: func(cmd *cobra.Command, args []string) {

		registries, err := getRegistry(GetClient())
		if err != nil {
			log.Fatal(err)
		}
		for _, registry := range registries {
			fmt.Printf("\n%v*\n", registry)
			registryTags, err := getTags(GetClient(), registry)
			if err != nil {
				log.Fatal(err)
			}
			for _, registryTag := range registryTags {
				fmt.Printf("\n%v*\n", registryTag)
			}
		}
	},
}

func getTags(client *http.Client, registry Registry) ([]RegistryTag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", viper.GetString("BaseUrl"), registry.TagsPath), nil)
	SetDefaultHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var registryTags []RegistryTag
	err = json.Unmarshal(body, &registryTags)
	if err != nil {
		log.Fatal(err)
	}
	return registryTags, nil
}

func getRegistry(client *http.Client) ([]Registry, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_registry.json", viper.GetString("ProjectUrl")), nil)
	SetDefaultHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var registries []Registry
	err = json.Unmarshal(body, &registries)
	if err != nil {
		log.Fatal(err)
	}
	return registries, nil
}