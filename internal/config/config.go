package config

// Регулярное выражение для проверки корректности url, указанных в конфиге
const (
	urlPattern = `^(http)://([a-zA-Z0-9\-]+(\.[a-zA-Z0-9\-]+)*|(\d{1,3}\.){3}\d{1,3}):\d+$`
)

// Структура конфига (используется в main)
type Config struct {
	Algorithm string `json:"algorithm"`
	Port      string `json:"port"`
	Servers   []struct {
		Url string `json: "url"`
	} `json:"servers"`
}

// Интерфейс для загрузчиков конфигурации
type ConfigLoader interface {
	Load(path string) (*Config, error)
}

// Функция, возвращающая загрузчик для указанного формата
func NewLoader(format string) ConfigLoader {
	switch format {
	case "json":
		return NewJsonLoader()
	default:
		return NewJsonLoader()
	}
}
