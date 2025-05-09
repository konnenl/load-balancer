package config

type Config struct {
	Port   string `json:"port"`
	Servers []struct{
		Url string `json: "url"`
	} `json:"servers"`
}

type ConfigLoader interface{
	Load(path string) (*Config, error)
}