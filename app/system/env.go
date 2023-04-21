package system

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var (
	_RELATIVE_ENV_PATH = []string{"..", "..", ".env"}
	ENV                Enviornment
)

type Enviornment struct {
	AWS struct {
		ExeBucket string `env:"AWS_S3_EXE_BUCKET"`
		DecBucket string `env:"AWS_S3_DEC_BUCKET"`
	}
	Rabbit struct {
		User string `env:"AMQP_USER"`
		Pass string `env:"AMQP_PASS"`
		Host string `env:"AMQP_HOST"`
		Port string `env:"AMQP_PORT"`
		Chan string `env:"AMQP_Chan"`
		Exch string `env:"AMQP_EXCH"`
	}
	Extras env.EnvSet
}

func loadDotEnv() {
	_, file, _, _ := runtime.Caller(0)

	envPath := filepath.Join(_RELATIVE_ENV_PATH...)
	envPath = filepath.Join(file, envPath)
	fmt.Printf("Path to env %s", envPath)
	godotenv.Load(envPath)
}

func loadEnv() {
	es, err := env.UnmarshalFromEnviron(ENV)
	if err != nil {
		log.Fatalf("env error: %v", err)
	}

	ENV.Extras = es
}

// Unwrap function     Panic if key is empty
func Unwrap(key string) string {
	if key == "" {
		log.Fatalf("MISSING KEY %s", key)
	}
	return key
}

func init() {
	loadDotEnv()
	loadEnv()

}
