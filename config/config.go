package config

import "fmt"

type APIConfig struct {
	Host string
	Port int
}

var (
	cookiePassword = "abc1234&&&&"
	studentPassword = "dshs1234@"

	serverAddress = "127.0.0.1"
	serverPort = 80
)

func BuildAPIConfig() *APIConfig {
	apiConfig := APIConfig{
		Host: "127.0.0.1",
		Port: 7557,
	}
	return &apiConfig
}

func GetServerAddress() string {
	return fmt.Sprintf("%s:%d", serverAddress, serverPort)
}

func GetSturentPassword() string {
	return studentPassword
}

func GetCookiePassword() string {
	return cookiePassword
}
