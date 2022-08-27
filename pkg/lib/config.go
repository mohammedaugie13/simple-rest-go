package lib

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type applicationEnv struct {
	ServerPort  int    `mapstructure:"SERVER_PORT"`
	FrontEndURL string `mapstructure:"FRONTEND_URL"`
	Env         string
	Migration   string

	// Postgres Database
	PgHost     string `mapstructure:"PG_HOST"`
	PgPort     int    `mapstructure:"PG_PORT"`
	PgUser     string `mapstructure:"PG_USER"`
	PgPassword string `mapstructure:"PG_PASSWORD"`
	PgDB       string `mapstructure:"PG_DB"`
}

type signKey struct {
	privateKey []byte
	PublicKey  []byte
}

type appConfig struct {
	App  applicationEnv
	Cert signKey
}

var AppConfig = appConfig{}

func InitAppConfig() *appConfig {
	var configStruct = &AppConfig

	configStruct.Cert = *GetCertKey()
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read configuration")
	}

	err = viper.Unmarshal(&configStruct.App)
	if err != nil {
		log.Fatal("AppConfig file can't be loaded", err)
	}

	configStruct.App.Env = os.Getenv("env")
	configStruct.App.Migration = os.Getenv("migration")

	return configStruct
}

func GetCertKey() *signKey {
	rootPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatalln(err)
	}

	privKey, err := ioutil.ReadFile(rootPath + "/cert/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	pubKey, err := ioutil.ReadFile(rootPath + "/cert/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	return &signKey{
		privateKey: privKey,
		PublicKey:  pubKey,
	}
}
