package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LoadTomlFile(name string, conf interface{}) error{
	f, err := os.Open(name)
	if err != nil {
		log.Println("conf.toml, not found")
	} else {
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		if err := toml.Unmarshal(b, conf); err != nil {
			return err
		}
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
				// TODO
			case reflect.String:
				val.Field(i).SetString(v)
			case reflect.Int:
				vi, _ := strconv.Atoi(v)
				val.Field(i).SetInt(int64(vi))
			}
		}
	}
}
