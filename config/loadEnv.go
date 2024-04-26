package config

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables(f ...string) error {
	var fName string

	if len(f) == 0 {
		fName = ".env"
	} else {
		fName = f[0]
	}

	err := godotenv.Load(fName)
	if err != nil {
		return err
	}

	return nil
}
