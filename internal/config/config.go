package config

func NewLoader(format string) ConfigLoader {
	switch format {
	default:
		return &JsonLoader{}
	}
}