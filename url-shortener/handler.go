package urlshort

import (
	"net/http"
	yaml "gopkg.in/yaml.v3"
)

type YAMLPath struct {
	url string `yaml:"url"`
    path string `yaml:"path"`
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	urlsFromYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}

	mapURL := buildMap(urlsFromYaml)
	return MapHandler(mapURL, fallback), nil
}

func buildMap(urlsFromYaml []YAMLPath) (map[string]string) {
	mapURL := make(map[string]string)

	for _, url := range urlsFromYaml {
		mapURL[url.path] = url.url
	}
	return mapURL
}


func parseYAML(yml []byte) ([]YAMLPath, error){
	var urls []YAMLPath
	err := yaml.Unmarshal(yml, &urls)
	if err != nil {
		return nil, err
	}
	return urls, nil
}