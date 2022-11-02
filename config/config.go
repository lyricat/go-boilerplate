package config

type (
	Config struct {
		DB DB `yaml:"db"`
	}

	DB struct {
		Driver     string `json:"driver"`
		Datasource string `json:"datasource"`
	}
)
