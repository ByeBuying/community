package conf

import (
	"github.com/naoina/toml"
	"os"
	"path"

	"go-common/klay/util"
)

type Config struct {
	Datadir struct {
		Root     string
		Keystore string
		Journal  string
		Log      string
	}

	Port struct {
		Server string
		Http   int
	}

	Common struct {
		SwgHost        string
		ServerFullHost string
		ServiceId      string
		Priv           string
	}

	Log struct {
		Terminal struct {
			Use       bool
			Verbosity int
		}
		File struct {
			Use       bool
			Verbosity int
			FileName  string
		}
	}

	Repositories map[string]map[string]interface{}
}

func NewConfig(file string) *Config {
	c := new(Config)

	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			c.sanitize()
			return c
		}
	}
}

func (p *Config) sanitize() {
	if p.Datadir.Root[0] == byte('~') {
		p.Datadir.Root = path.Join(util.HomeDir(), p.Datadir.Root[1:])
	}
}
