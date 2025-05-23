package config

type POSTGRES struct {
	DB       string `env-required:"false" json:"db" env:"POLLING_POSTGRES_DB"`
	User     string `env-required:"false" json:"user" env:"POLLING_POSTGRES_USER"`
	Password string `env-required:"false" json:"password" env:"POLLING_POSTGRES_PASSWORD"`
	Host     string `env-required:"false" json:"host" env:"POLLING_POSTGRES_HOST"`
	Port     uint   `env-required:"false" json:"port" env:"POLLING_POSTGRES_PORT"`
	SSLMode  string `env-required:"false" json:"sslmode" env:"POLLING_POSTGRES_SSLMODE"`
	Timezone string `env-required:"false" json:"timezone" env:"POLLING_POSTGRES_TIMEZONE"`
	Schema   string `env-required:"false" json:"schema" env:"POLLING_POSTGRES_SCHEMA"`
}

type REDIS struct {
	Host string `env-required:"false" json:"host" env:"POLLING_REDIS_HOST"`
	Port uint   `env-required:"false" json:"port" env:"POLLING_REDIS_PORT"`
}
