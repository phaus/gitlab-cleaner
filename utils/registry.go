package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/spf13/viper"
)

// Registry a gitlab Registry info object
type Registry struct {
	ID          uint64 `json:"id"`
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
	TotalSize     uint64 `json:"total_size"`
	CreatedAt     string `json:"created_at"`
	DestroyPath   string `json:"destroy_path"`
}

// GetTags - GETs the Tags of a specific (Docker) Registry.
func GetTags(client *http.Client, registry Registry) (map[string]RegistryTag, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", viper.GetString("BaseUrl"), registry.TagsPath), nil)
	SetDefaultHeaders(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	totalPages, err := strconv.Atoi(resp.Header.Get("x-total-pages"))
	if err != nil {
		return nil, err
	}
	totalCount, err := strconv.Atoi(resp.Header.Get("x-total"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Printf("\n%d %s on %d %s. Loading Page: ",
		totalCount,
		ImageLabel(totalCount),
		totalPages,
		PageLabel(totalPages))

	var registryTags = make(map[string]RegistryTag)
	var count = 1
	for page := 1; page <= totalPages; page++ {
		fmt.Printf(" %d", count)
		req, err := http.NewRequest("GET", fmt.Sprintf("%s%s&page=%d", viper.GetString("BaseUrl"), registry.TagsPath, page), nil)
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
		err = parseTags(body, registryTags)
		if err != nil {
			return nil, err
		}
		count++
	}
	fmt.Printf("\n\n")
	return registryTags, nil
}

// GetTagsAndSortedKeys - returns the tag and sorted keys of a registry
func GetTagsAndSortedKeys(client *http.Client, registry Registry) (map[string]RegistryTag, []string) {
	registryTags, err := GetTags(GetClient(), registry)
	if err != nil {
		log.Fatal(err)
	}
	keys := SortedKeys(registryTags)
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return registryTags, keys
}

// PageLabel returns singular/plural according to count.
func PageLabel(count int) string {
	if count > 1 {
		return "pages"
	}
	return "page"
}

// ImageLabel returns singular/plural according to count.
func ImageLabel(count int) string {
	if count > 1 {
		return "images"
	}
	return "image"
}

// GetRegistry - GETs the (Docker) Registry for a specific project url.
func GetRegistry(client *http.Client) ([]Registry, error) {
	req, err := http.NewRequest("GET", viper.GetString("RegistryUrl"), nil)
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

// SortedKeys - returns the keys of a RegistryTags Map in sorted order.
func SortedKeys(registryTags map[string]RegistryTag) []string {
	var keys []string
	for _, v := range registryTags {
		keys = append(keys, v.CreatedAt)
	}
	sort.Strings(keys)
	return keys
}

func parseTags(body []byte, registryTags map[string]RegistryTag) error {
	var innerTags []RegistryTag
	var err = json.Unmarshal(body, &innerTags)
	if err != nil {
		log.Fatal(err)
	}
	for _, registryTag := range innerTags {
		registryTags[registryTag.CreatedAt] = registryTag
	}
	return nil
}
