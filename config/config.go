package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ConfServer
	DB     ConfDB
}

type ConfServer struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG,required"`
}

type ConfDB struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=5432"`
	Username string `env:"DB_USER,default=myapp_user"`
	Password string `env:"DB_PASS,default=myapp_pass"`
	DBName   string `env:"DB_NAME,default=myapp_db"`
	Debug    bool   `env:"DB_DEBUG,default=true"`
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode conf: %s", err)
	}

	return &c
}

func NewDB() *ConfDB {
	var c ConfDB
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode confDB: %s", err)
	}

	return &c
}
