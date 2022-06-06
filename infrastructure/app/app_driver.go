package app

import "github.com/joho/godotenv"

type DriverApp struct {
	Port      string
	SecretKey string
}

func readEnvApp() (DriverApp, error) {
	envApp, err := godotenv.Read()
	if err != nil {
		return DriverApp{}, err
	}
	return DriverApp{
		Port:      envApp["PORT"],
		SecretKey: envApp["SECRET_KEY"],
	}, nil
}

func NewDriverApp() (DriverApp, error) {
	env, err := readEnvApp()
	if err != nil {
		return DriverApp{}, err
	}
	return env, nil

}
