package cfg

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/creamdog/gonfig"
)

var config gonfig.Gonfig

func checkError(err *error) {
	if *err != nil {
		panic(*err)
	}
}

func environment(key string) *string {
	env := os.Getenv("ENV")

	if env == "" {
		env = "default"
	}

	envKey := strings.ToUpper(strings.Replace(key, "/", "_", -1))
	result := os.Getenv(envKey)
	return &result
}

//GetInt return key of file configuration
//if the key exists in the environment variables, it uses as predominant
func GetInt(key string, defaultValue int) *int {
	env := os.Getenv("ENV")

	if env == "" {
		env = "default"
	}

	var cfg int
	var err error

	envValue := environment(key)

	if *envValue == "" {
		cfg, err = config.GetInt(fmt.Sprintf("%v/%v", env, key), defaultValue)
		checkError(&err)
	} else {
		_int, err := strconv.Atoi(*envValue)
		checkError(&err)
		return &_int
	}

	return &cfg
}

//GetString return key of file configuration
//if the key exists in the environment variables, it uses as predominant
func GetString(key string, defaultValue string) *string {
	env := os.Getenv("ENV")

	if env == "" {
		env = "default"
	}

	var cfg string
	var err error

	envValue := environment(key)

	if *envValue == "" {
		cfg, err = config.GetString(fmt.Sprintf("%v/%v", env, key), defaultValue)
		checkError(&err)
	} else {
		return envValue
	}

	return &cfg
}

//GetBool return key of file configuration
//if the key exists in the environment variables, it uses as predominant
func GetBool(key string, defaultValue bool) *bool {
	env := os.Getenv("ENV")

	if env == "" {
		env = "default"
	}

	var cfg bool
	var err error

	envValue := environment(key)

	if *envValue == "" {
		cfg, err = config.GetBool(fmt.Sprintf("%v/%v", env, key), defaultValue)
		checkError(&err)
	} else {
		_bool, err := strconv.ParseBool(*envValue)
		checkError(&err)
		return &_bool
	}

	return &cfg
}

func init() {

	fileCFG := flag.String("c", "./config/default.json", "Configuration file")

	flag.Parse()

	fmt.Printf("Load file of configuration %v\n", *fileCFG)

	file, err := os.Open(*fileCFG)

	if err != nil {
		flag.PrintDefaults()
		log.Fatal(err)
	}

	checkError(&err)

	defer file.Close()

	config, err = gonfig.FromJson(file)

	fmt.Printf("File %v loaded !!!\n", *fileCFG)

	checkError(&err)
}
