package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type PRTGConfig struct {
	IpBIB        string
	IpKIM        string
	IpMAL        string
	IpBSL        string
	IpSML        string
	IpMSIG       string
	IpBCHO       string
	User         string
	Password     string
	UserBCHO     string
	PasswordBCHO string
	PasswordSML  string
}

type Config struct {
	ApiConfig
	PRTGConfig
	DbConfig
}

func LoadEnv(key string) string {
	// typeEnv := os.Args[1]
	// if typeEnv != "dev" && typeEnv != "prod" && typeEnv != "test" {
	// 	typeEnv = "dev"
	// }
	// env := fmt.Sprintf(".%s.env", typeEnv)
	viper.SetConfigFile("config/config.env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file%s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion %s", key)
	}

	return value
}

func (c *Config) readConfig() {

	api := LoadEnv("API_URL")
	ipBIB := LoadEnv("IP_BIB")
	ipKIM := LoadEnv("IP_KIM")
	ipMAL := LoadEnv("IP_MAL")
	ipBSL := LoadEnv("IP_BSL")
	ipSML := LoadEnv("IP_SML")
	ipMSIG := LoadEnv("IP_MSIG")
	ipBCHO := LoadEnv("IP_BCHO")
	user := LoadEnv("USER")
	password := LoadEnv("PASSWORD")
	userBCHO := LoadEnv("USER_BCHO")
	passwordBCHO := LoadEnv("PASSWORD_BCHO")
	passwordSML := LoadEnv("PASSWORD_SML")

	dbHost := LoadEnv("DB_HOST")
	dbPort := LoadEnv("DB_PORT")
	dbUser := LoadEnv("DB_USER")
	dbPassword := LoadEnv("DB_PASSWORD")
	dbName := LoadEnv("DB_NAME")
	// log.Println("==============")
	// log.Println(ipBIB)
	// log.Println("==============")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	c.ApiConfig = ApiConfig{Url: api}
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.PRTGConfig = PRTGConfig{
		IpBIB:        ipBIB,
		IpKIM:        ipKIM,
		IpMAL:        ipMAL,
		IpBSL:        ipBSL,
		IpSML:        ipSML,
		IpMSIG:       ipMSIG,
		IpBCHO:       ipBCHO,
		User:         user,
		Password:     password,
		UserBCHO:     userBCHO,
		PasswordBCHO: passwordBCHO,
		PasswordSML:  passwordSML,
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
