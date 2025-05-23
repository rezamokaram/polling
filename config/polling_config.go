package config

type PollingConfig struct {
	Polling  POLLING  `env-required:"false" json:"polling"`
	Redis    REDIS    `env-required:"false" json:"redis"`
	Postgres POSTGRES `env-required:"false" json:"postgres"`
}

func (PollingConfig) configSignature() {}

type POLLING struct {
	Name    string `env-required:"false" json:"name" env:"POLLING_APP_NAME"`
	Version string `env-required:"false" json:"version" env:"POLLING_APP_VERSION"`
	Host    string `env-required:"false" json:"host" env:"POLLING_APP_HOST"`
	Port    uint   `env-required:"false" json:"port" env:"POLLING_APP_PORT"`
}
