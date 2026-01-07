package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	path, err1 := os.UserHomeDir()
	if err1 != nil {
		return Config{}, err1
	}

	data, err := os.Open(path + configFileName)

	if err != nil {
		return Config{}, err
	}
	defer data.Close()
	var res Config //idhar create hua

	decoder := json.NewDecoder(data)
	erro := decoder.Decode(&res)
	if erro != nil {
		return Config{}, erro
	}

	return res, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(cfg)
}

func write(cfg *Config) error {
	path, err1 := os.UserHomeDir()
	if err1 != nil {
		return err1
	}

	file, err1 := os.OpenFile(
		path+configFileName,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err1 != nil {
		return err1
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err := encoder.Encode(cfg)

	return err
}
