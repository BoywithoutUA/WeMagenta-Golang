package config

type AIMusicConfig struct {
	API              API              `mapstructure:"API"`
	SRVUser          SRVUser          `mapstructure:"SRV-User"`
	SRVCreation      SRVCreation      `mapstructure:"SRV-Creation"`
	SRVAdministrator SRVAdministrator `mapstructure:"SRV-Administrator"`
	SRVCommunity     SRVCommunity     `mapstructure:"SRV-Community"`
	MySQL            MySQLConfig      `mapstructure:"MySQL"`
	Redis            RedisConfig      `mapstructure:"Redis"`
	Consul           ConsulConfig     `mapstructure:"Consul"`
	Jaeger           JaegerConfig     `mapstructure:"Jaeger"`
	Domain           DomainConfig     `mapstructure:"Domain"`
}

type API struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SRVComposition struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SRVUser struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SRVCreation struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SRVAdministrator struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type SRVCommunity struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type DomainConfig struct {
	Host string `mapstructure:"host"`
}
