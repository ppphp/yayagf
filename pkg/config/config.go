package config

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

var ErrorNoToml = errors.New("toml not found")

func LoadTomlFile(name string, conf interface{}) error {

	b, err := ioutil.ReadFile(name)
	if err != nil {
		return ErrorNoToml
	}

	if err := toml.Unmarshal(b, conf); err != nil {
		return err
	}

	return nil
}

func LoadEnv(conf interface{}) {
	val := reflect.ValueOf(conf).Elem()
	for i := 0; i < val.NumField(); i++ {
		name := val.Type().Field(i).Name
		v, ok := os.LookupEnv(strings.ToUpper(name))
		if ok && v != "" {
			switch val.Type().Field(i).Type.Kind() {
			case reflect.Struct:
				panic("shall not")
			case reflect.String:
				val.Field(i).SetString(v)
			case reflect.Int:
				vi, _ := strconv.Atoi(v)
				val.Field(i).SetInt(int64(vi))
			}
		}
	}
}

// only support ini like config
func LoadConfig(conf interface{}) error {
	if err := LoadTomlFile("conf.toml", conf); err != nil {
		return err
	}

	LoadEnv(conf)

	return nil
}
