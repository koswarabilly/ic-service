package setup

type Configuration struct {
	Env                     string   `envconfig:"go_env" default:"development"`
	LogLevel                LogLevel `envconfig:"log_level"`
	ServiceName             string   `envconfig:"service_name" required:"true"`
	ServicePort             int64    `envconfig:"service_port" default:"30002"`
	RedisHost               string   `envconfig:"redis_host" default:"localhost"`
	RedisPort               int      `envconfig:"redis_port" default:"6379"`
	RedisPassword           string   `envconfig:"redis_password"`
	RedisDatabaseNumber     int      `envconfig:"redis_db_number" default:"0"`
	RedisHostTest           string   `envconfig:"redis_host_test" default:"localhost"`
	RedisPortTest           int      `envconfig:"redis_port_test" default:"6379"`
	RedisPasswordTest       string   `envconfig:"redis_password_test"`
	RedisDatabaseNumberTest int      `envconfig:"redis_db_number_test" default:"15"`
}
