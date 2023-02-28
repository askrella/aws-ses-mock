package internal

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"os"
)

type ConfigType struct {
	Port      int    `env:"PORT" envDefault:"8080"`
	OutputDir string `env:"OUTPUT_DIR"`
}

var Config ConfigType

func ReadConfigFromEnv() {
	Config = ConfigType{}
	if err := env.Parse(&Config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if Config.OutputDir == "" {
		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		Config.OutputDir = workingDir
	}

	fmt.Printf("%+v\n", Config)
}
