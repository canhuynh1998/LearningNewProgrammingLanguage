package urlshort

import (
	"encoding/json"
	"net/http"
	"database/sql"
	pg"github.com/lib/pq"
	yaml "gopkg.in/yaml.v3"
)

type Path interface {
	GetPath() string
	GetUrl() string
}

type YAMLPath struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func (yml YAMLPath) GetPath() string {
	return yml.Path
}

func (yml YAMLPath) GetUrl() string {
	return yml.Url
}

type JSONPath struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func (yml JSONPath) GetPath() string {
	return yml.Path
}
func (yml JSONPath) GetUrl() string {
	return yml.Url
}


// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		currentPath := r.URL.Path
		if path, ok := pathsToUrls[currentPath]; ok {
			http.Redirect(w, r, path, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func buildPathMap(paths []Path) map[string]string {
	pathMap := make(map[string]string)
	for _, p := range paths {
		pathMap[p.GetPath()] = p.GetUrl()
	}
	return pathMap
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this..
	urlsFromYaml, err := parseYAMLtoPath(yml)
	if err != nil {
		return nil, err
	}
	mapURL := buildPathMap(urlsFromYaml)
	return MapHandler(mapURL, fallback), nil
}

func parseYAMLtoPath(yml []byte) ([]Path, error) {
	var urls []YAMLPath
	err := yaml.Unmarshal(yml, &urls)
	if err != nil {
		return nil, err
	}

	return yamlToPath(urls), nil
}

func yamlToPath(yaml []YAMLPath) ([]Path){
	paths := make([]Path, len(yaml))
	for i, url := range yaml {
		paths[i] = url
	}
	return paths
}

func JSONHandler(js []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlsFromJson, err := parseJSONtoPath(js)
	if err != nil {
		return nil, err
	}
	mapURL := buildPathMap(urlsFromJson)
	return MapHandler(mapURL, fallback), nil
}

func parseJSONtoPath(js []byte) ([]Path, error) {
	var urls []JSONPath
	err := json.Unmarshal(js, &urls)
	if err != nil {
		return nil, err
	}
	return jsonToPath(urls), nil
}

func jsonToPath(yaml []JSONPath) ([]Path){
	paths := make([]Path, len(yaml))
	for i, url := range yaml {
		paths[i] = url
	}
	return paths
}