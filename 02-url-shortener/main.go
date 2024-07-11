package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"url-shortener/handler"
	"url-shortener/models"

	"gopkg.in/yaml.v2"
)

func main() {
	yamlFile := flag.String("y", "config.yml", "Yaml config file with mappings")
	jsonFile := flag.String("j", "config.json", "Json config file with mappings")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/cats":           "https://google.com/search?q=cats",
	 	"/dogs":           "https://google.com/search?q=dogs",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Bonus task, add json handler
	json, err := loadFile(*jsonFile)
	if err!= nil {
		panic(err)
	}

	parsedJson, err := parseJson([]byte(json))
	if err!= nil {
		panic(err)
	}

	// Use existing YAMLHandler, because data format is the same
	jsonHandler := handler.YAMLHandler(parsedJson, mapHandler)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := loadFile(*yamlFile)
	if err!= nil {
		panic(err)
	}

	parsedYaml, err := parseYaml([]byte(yaml))
	if err!= nil {
		panic(err)
	}

	yamlHandler := handler.YAMLHandler(parsedYaml, jsonHandler)

	fmt.Println("Starting the server on :8080")
 	http.ListenAndServe(":8080", yamlHandler)
}

func loadFile(filename string) (string, error){
	fmt.Printf("Loading file %s\n", filename)
	fileContent, loadFileError := os.ReadFile(filename)
	return string(fileContent), loadFileError
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseYaml(yamlContent []byte) ([]models.YamlStruct, error) {
	var data []models.YamlStruct
	err := yaml.Unmarshal(yamlContent, &data)
	return data, err
}

func parseJson(jsonContent []byte) ([]models.YamlStruct, error) {
	var data []models.YamlStruct
	err := json.Unmarshal(jsonContent, &data)
	return data, err
}