package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Telegram struct {
	Token string `env:"TELEGRAM_BOT_TOKEN"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type SSO struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retriesCount"`
}

type Clients struct {
	SSO SSO `yaml:"sso"`
}

type Config struct {
	Env      string   `yaml:"env" env-required:"true"`
	Port     int      `yaml:"port" env-required:"true"`
	Secret   string   `yaml:"secret" env-required:"true"`
	Static   string   `yaml:"static" env-default:"./static"`
	Storage  Storage  `yaml:"storage" env-required:"true"`
	Telegram Telegram `yaml:"-"`
	Clients  Clients  `yaml:"clients" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	var telegramCfg Telegram
	if err = cleanenv.ReadEnv(&telegramCfg); err != nil {
		panic("failed to read env: " + err.Error())
	}

	cfg.Telegram = telegramCfg

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
