package config

import (
	"github.com/caarlos0/env/v8"
)

func Read() (Variables, error) {
	var v Variables
	if err := env.Parse(&v); err != nil {
		var zero Variables
		return zero, err
	}
	return v, nil
}
