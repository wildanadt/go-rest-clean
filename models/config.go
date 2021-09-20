package models

type ServerConfig struct {
	Name           string `env:"SERVER_NAME"`
	Port           string `env:"PORT_SERVER,required"`
	Host           string `env:"HOST_SERVER,required"`
	PostgresConfig PostgresConfig
	JWTConfig      JWTConfig
}

type PostgresConfig struct {
	Host     string `env:"HOST_POSTGRES,required"`
	Port     string `env:"PORT_POSTGRES,required"`
	Name     string `env:"NAME_POSTGRES,required"`
	User     string `env:"USER_POSTGRES,required"`
	Password string `env:"PASS_POSTGRES,required"`
}

type JWTConfig struct {
	APISecret     string `env:"API_SECRET,required"`
	TokenLifespan string `env:"TOKEN_HOUR_LIFESPAN,required"`
}
