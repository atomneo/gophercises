package handler

import (
	"net/http"
	"url-shortener/models"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path]!= "" {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}



// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(mapConfiguration []models.YamlStruct, fallback http.Handler) (http.HandlerFunc) {
	var pathsToUrls = make(map[string]string)
	for _, entry := range mapConfiguration {
		pathsToUrls[entry.Path] = entry.Url
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if pathsToUrls[r.URL.Path]!= "" {
			http.Redirect(w, r, pathsToUrls[r.URL.Path], http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}