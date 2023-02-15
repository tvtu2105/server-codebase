package models

type Config struct {
	Kafka         KafkaConfig         `yaml:"kafka" mapstructure:"kafka"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch" mapstructure:"elasticsearch"`
	Redis         RedisConfig         `yaml:"redis" mapstructure:"redis"`
	MySQL         MySQLConfig         `yaml:"mysql" mapstructure:"mysql"`
	MongoDB       MongoDBConfig       `yaml:"mongodb" mapstructure:"mongodb"`
	JWT           JWT                 `yaml:"jwt" mapstructure:"jwt"`
	Server        Server              `yaml:"server" mapstructure:"server"`
}

type KafkaConfig struct {
	Addresses []string `yaml:"addresses" mapstructure:"addresses"`
	TopicA    string   `yaml:"topic_a" mapstructure:"topic_a"`
	TopicB    string   `yaml:"topic_b" mapstructure:"topic_b"`
}

type ElasticsearchConfig struct {
	Addresses   []string `yaml:"addresses" mapstructure:"addresses"`
	ClusterName string   `yaml:"cluster_name" mapstructure:"cluster_name"`
}

type RedisConfig struct {
	Addresses []string `yaml:"addresses" mapstructure:"addresses"`
	KeyA      string   `yaml:"key_a" mapstructure:"key_a"`
	KeyB      string   `yaml:"key_b" mapstructure:"key_b"`
}

type MySQLConfig struct {
	Address  string `yaml:"address" mapstructure:"address"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Auth     string `yaml:"auth" mapstructure:"auth"`
}

type MongoDBConfig struct {
	Addresses []string `yaml:"addresses" mapstructure:"addresses"`
	Username  string   `yaml:"username" mapstructure:"username"`
	Password  string   `yaml:"password" mapstructure:"password"`
}

type JWT struct {
	Secret      string `yaml:"secret" mapstructure:"secret"`
	ExpiredTime string `yaml:"expired_time" mapstructure:"expired_time"`
}

type Server struct {
	Env      string `yaml:"env" mapstructure:"env"`
	GrpcPort int    `yaml:"grpc_port" mapstructure:"grpc_port"`
	HttpPort int    `yaml:"http_port" mapstructure:"http_port"`
}
