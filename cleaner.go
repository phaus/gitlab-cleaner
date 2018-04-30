package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// ProjectURL url of the gitlab project
var ProjectURL string

// BaseURL the base url of the gitlab installation
var BaseURL string

// PrivateToken the private access token for the gitlab api
var PrivateToken string

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

func main() {
	ProjectURL = os.Getenv("CI_PROJECT_URL")
	PrivateToken = os.Getenv("PRIVATE_ACCESS_TOKEN")

	client := &http.Client{}

	u, err := url.Parse(ProjectURL)
	if err != nil {
		log.Fatal(err)
	}

	BaseURL = fmt.Sprintf("%s://%s", u.Scheme, u.Host)

	fmt.Printf("calling %s on BaseURL %s", ProjectURL, BaseURL)

	registries, err := getRegistry(client)
	if err != nil {
		log.Fatal(err)
	}
	for _, registry := range registries {
		fmt.Printf("\n%v*\n", registry)
		registryTags, err := getTags(client, registry)
		if err != nil {
			log.Fatal(err)
		}
		for _, registryTag := range registryTags {
			fmt.Printf("\n%v*\n", registryTag)
		}
	}
}

func getTags(client *http.Client, registry Registry) ([]RegistryTag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BaseURL, registry.TagsPath), nil)
	setHeaders(req)
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
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/container_registry.json", ProjectURL), nil)
	setHeaders(req)
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

func setHeaders(req *http.Request) {
	req.Header.Add("Private-Token", PrivateToken)
	req.Header.Add("accept", "application/json, text/plain, */*")
}
