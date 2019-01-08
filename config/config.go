package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//Conf conf
var Conf *Configs

//Configs Configs
type Configs struct {
	config map[string]string
	node   string
}

const nodeMid = "-_-"

func init() {
	Conf = new(Configs)
	Conf.config = make(map[string]string)
	Conf.LoadConfig("dev")
}

//LoadConfig getConfig in dev or pro
func (conf *Configs) LoadConfig(mode string) error {
	if mode != "dev" && mode != "pro" {
		return errors.New("No suitable env file found")
	}
	_, s, _, ok := runtime.Caller(1)
	if !ok {
		return errors.New("path is wrong")
	}
	s = filepath.Dir(s)

	env, err := os.Open(s + "/env." + mode + ".ini")
	if err != nil {
		return err
	}

	defer env.Close()

	buff := bufio.NewReader(env)

	for {
		lines, _, err := buff.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		line := strings.TrimSpace(string(lines))

		//注释
		if strings.Index(line, "#") == 0 {
			continue
		}

		//[xxx]
		n := strings.Index(line, "[")
		n1 := strings.LastIndex(line, "]")

		if n > -1 && n1 > -1 && n1 > n {
			conf.node = strings.TrimSpace(line[n+1 : n1])
			continue
		}

		if len(conf.node) == 0 || len(line) == 0 {
			continue
		}

		strArray := strings.Split(line, "=")
		originKey := strArray[0]
		value := strArray[1]
		newKey := conf.node + nodeMid + originKey
		conf.config[newKey] = value
	}
	return nil
}

//GetValue GetValue
func (conf *Configs) GetValue(node, key string) (value string, isFound bool) {
	newKey := node + nodeMid + key

	value, isFound = conf.config[newKey]

	return value, isFound
}

//GetDefaultValue GetDefaultValue
func (conf *Configs) GetDefaultValue(node, key, defaultValue string) string {
	newKey := node + nodeMid + key

	value, isFound := conf.config[newKey]

	if !isFound {
		return defaultValue
	}

	return value
}
