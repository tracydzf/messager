package global

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"sync"
)

var (
	k   = koanf.New(".")
	cbs = make([]func(), 0)
	mtx = &sync.RWMutex{}
	p   = yaml.Parser()
)

func GetAuthConf() (confs []map[string]string, err error) {
	mtx.RLock()
	defer mtx.RUnlock()

	confs = make([]map[string]string, 0)
	err = k.Unmarshal("auths", &confs)

	return
}
