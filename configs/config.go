package configs

type Config struct {
	ServerMain struct {
		Name             string         `yaml:"name"`
		Port             int            `yaml:"port"`
		DatabasePostgres PostgresConfig `yaml:"database_postgres"`
		DatabaseRedis    RedisConfig    `yaml:"database_redis"`
	} `yaml:"server_main"`

	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`

	ServerCode struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"server_code"`
}
