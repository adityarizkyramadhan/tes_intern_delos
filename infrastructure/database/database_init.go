package database

import (
	"github.com/joho/godotenv"
)

type DriverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func readEnvSupabase() (DriverSupabase, error) {
	envSupabase, err := godotenv.Read()
	if err != nil {
		return DriverSupabase{}, err
	}
	return DriverSupabase{
		User:     envSupabase["SUPABASE_USER"],
		Password: envSupabase["SUPABASE_PASSWORD"],
		Host:     envSupabase["SUPABASE_HOST"],
		Port:     envSupabase["SUPABASE_PORT"],
		DbName:   envSupabase["SUPABASE_DB_NAME"],
	}, nil
}

func NewDriverSupabase() (DriverSupabase, error) {
	env, err := readEnvSupabase()
	if err != nil {
		return DriverSupabase{}, err
	}
	return env, nil
}

type DriverMysql struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func readEnvMysql() (DriverMysql, error) {
	envMysql, err := godotenv.Read()
	if err != nil {
		return DriverMysql{}, err
	}
	return DriverMysql{
		User:     envMysql["DB_USER"],
		Password: envMysql["DB_PASSWORD"],
		Host:     envMysql["DB_HOST"],
		Port:     envMysql["DB_PORT"],
		DbName:   envMysql["DB_NAME"],
	}, nil
}

func NewDriverMysql() (DriverMysql, error) {
	env, err := readEnvMysql()
	if err != nil {
		return DriverMysql{}, err
	}
	return env, nil
}
