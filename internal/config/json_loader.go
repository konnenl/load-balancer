package config

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

// Структура загрузчика конфига из json файла
type JsonLoader struct{}

// NewJsonLoader создаёт новый экземпляр JsonLoader
func NewJsonLoader() *JsonLoader {
	return &JsonLoader{}
}

// Load загружает конфиг из указанного json файла и возвращает заполненную структуру Config
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

// IsValidUrl проводит валидацию url
func (l *JsonLoader) IsValidUrl(url string) bool {
	return regexp.MustCompile(urlPattern).MatchString(url)
}
