package main

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
		App       *App
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

	App struct {
		Name string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func main() {
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
	fmt.Fprintf(file, "BINARY_NAME=%s\n", configInstance.App.Name)
	fmt.Fprintf(file, "GOOSE_DRIVER=%s\n", configInstance.Migration.Driver)
	fmt.Fprintf(file, "GOOSE_MIGRATION_DIR=%s\n", configInstance.Migration.Dir)
	fmt.Fprintf(file, "GOOSE_DBSTRING='%s://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=%s'",
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
}
