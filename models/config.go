package models

type Config struct {
	Port     string
	Database Database
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}
