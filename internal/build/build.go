package build

import (
	"io/ioutil"
	"log"
	"os"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
)

func init() {
	// specify build params
	_ = os.Setenv("GOPROXY", "https://goproxy.io")
	_ = os.Setenv("GOSUMDB", "off")
	_ = os.Setenv("GOPRIVATE", "gitlab.papegames.com/*")
}

// BuildBinary build a binary in tmp
func BuildBinary() (string, error) {
	f, err := ioutil.TempFile("", "*") // linux is /tmp/xxxx
	if err != nil {
		return "", err
	}
	f.Close()
	if err, o, e := command.DoCommand("go", "build", "-o", f.Name(), "./"); err != nil {
		log.Printf("build to %v err: %v, err: %v, out: %v\n", f.Name(), err, e, o)
		return "", err
	}
	return f.Name(), nil
}
