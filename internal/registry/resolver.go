package registry

import (
	"fmt"
)

type Package struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var (
	defaultRegistry = map[string]string{
		"mux":  "github.com/gorilla/mux",
		"echo": "github.com/labstack/echo/v4",
	}
)

func Resolve(name string) (string, error) {
	if url, ok := defaultRegistry[name]; ok {
		return url, nil
	}

	return "", fmt.Errorf("package not found: %s", name)
}
