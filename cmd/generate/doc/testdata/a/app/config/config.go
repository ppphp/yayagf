package config

import (
	"sync"

	"gitlab.papegames.com/fengche/yayagf/pkg/config"
	"gitlab.papegames.com/fengche/yayagf/pkg/log"
)

var lock sync.RWMutex

type Config struct {
	DB   string
	Port int
	Log  log.Config
}

var conf = new(Config)

// only support ini like config
func LoadConfig() error {
	lock.Lock()
	defer lock.Unlock()
	if err := config.LoadConfig("conf.toml", conf); err != nil {
		return err
	}

	log.Infof("%v", conf)

	return nil
}

func GetConfig() Config {
	lock.RLock()
	defer lock.RUnlock()
	return *conf
}
