package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server    *Server
		DB        *DB
		Migration *Migration
	}

	Server struct {
		Port int
	}

	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Migration struct {
		Dir    string
		Path   string
		Driver string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	})

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}

	file, err := os.Create(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the configuration to the .env file
	fmt.Fprintf(file, "GOOSE_DRIVER=%s\n", configInstance.Migration.Driver)
	fmt.Fprintf(file, "GOOSE_MIGRATION_DIR=%s\n", configInstance.Migration.Dir)
	fmt.Fprintf(file, "GOOSE_DBSTRING=%s://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s",
		configInstance.Migration.Driver,
		configInstance.DB.User,
		configInstance.DB.Password,
		configInstance.DB.Host,
		configInstance.DB.Port,
		configInstance.DB.DBName,
		configInstance.DB.SSLMode,
		configInstance.DB.TimeZone,
	)

	fmt.Println(".env file created successfully!")

	return configInstance
}

// type Config struct {
// 	DBHost     string `yaml:"db_host"`
// 	DBPort     int    `yaml:"db_port"`
// 	DBUser     string `yaml:"db_user"`
// 	DBPassword string `yaml:"db_password"`
// 	DBName     string `yaml:"db_name"`
// }
//
// func config() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath(".")
//
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("Fatal error config file: %w \n", err))
// 	}
//
// 	var config Config
// 	err = viper.Unmarshal(&config)
// 	if err != nil {
// 		panic(fmt.Errorf("Fatal error config file: %w \n", err))
// 	}
//
// 	// Open the .env file
// 	file, err := os.Create(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
//
// 	// Write the configuration to the .env file
// 	fmt.Fprintf(file, "DB_HOST=%s\n", config.DBHost)
// 	fmt.Fprintf(file, "DB_PORT=%d\n", config.DBPort)
// 	fmt.Fprintf(file, "DB_USER=%s\n", config.DBUser)
// 	fmt.Fprintf(file, "DB_PASSWORD=%s\n", config.DBPassword)
// 	fmt.Fprintf(file, "DB_NAME=%s\n", config.DBName)
//
// 	fmt.Println(".env file created successfully!")
// }
