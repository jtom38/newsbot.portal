package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	Config_API_Address = "API_ADDRESS"
)

type ConfigClient struct{}

func NewConfigClient() ConfigClient {
	c := ConfigClient{}
	c.RefreshEnv()

	return c
}

// This looks for the required key and returns it a a string.
// If its not able to be found, it will panic.
func (c ConfigClient) MustGet(key string) string {
	res, filled := os.LookupEnv(key)
	if !filled {
		msg := fmt.Sprintf("Missing the a value for '%v'.  Could generate errors.", key)
		panic(msg)
	}
	return res
}

func (cc *ConfigClient) Get(key string) string {
	res, filled := os.LookupEnv(key)
	if !filled {
		log.Printf("Missing the a value for '%v'.  Could generate errors.", key)
	}
	return res
}

func (cc *ConfigClient) GetFeature(flag string) (bool, error) {
	cc.RefreshEnv()

	res, filled := os.LookupEnv(flag)
	if !filled {
		errorMessage := fmt.Sprintf("'%v' was not found", flag)
		return false, errors.New(errorMessage)
	}

	b, err := strconv.ParseBool(res)
	if err != nil {
		return false, err
	}
	return b, nil
}

// Use this when your ConfigClient has been opened for awhile and you want to ensure you have the most recent env changes.
func (cc *ConfigClient) RefreshEnv() {
	// Check to see if we have the env file on the system
	_, err := os.Stat(".env")

	// We have the file, load it.
	if err == nil {
		_, err := os.Open(".env")
		if err == nil {
			loadEnvFile()
		}
	}
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}
