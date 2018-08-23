package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler sets up our routes to redirect or go to the default page
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler converts our YAML to our map type expected by MapHandler
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var b []paths
	m := make(map[string]string)

	err := yaml.Unmarshal(yml, &b)

	if err != nil {
		return nil, err
	}

	for _, v := range b {
		m[v.Path] = v.Url
	}

	return MapHandler(m, fallback), nil
}

type paths struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
