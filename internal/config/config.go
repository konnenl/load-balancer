package config

const(
	urlPattern = `^(http)://([a-zA-Z0-9\-]+(\.[a-zA-Z0-9\-]+)*|(\d{1,3}\.){3}\d{1,3}):\d+$`
)

type Config struct {
	Port   string `json:"port"`
	Servers []struct{
		Url string `json: "url"`
	} `json:"servers"`
}

type ConfigLoader interface{
	Load(path string) (*Config, error)
}

func NewLoader(format string) ConfigLoader {
	switch format {
	case "json":
		return NewJsonLoader()
	default:
		return NewJsonLoader()
	}
}