package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type JsonLoader struct{}

func NewJsonLoader() *JsonLoader {
	return &JsonLoader{}
}

func (l *JsonLoader) Load(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err = json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	for i, s := range cfg.Servers {
		if !l.IsValidUrl(s.Url) {
			return nil, fmt.Errorf("invalid url index: %d", i)
		}
	}

	return &cfg, nil
}

func (l *JsonLoader) IsValidUrl(url string) bool {
	return regexp.MustCompile(urlPattern).MatchString(url)
}
