package config

import (	
	"fmt"
	user "os/user"
	ioutil "io/ioutil"
	yaml "gopkg.in/yaml.v2"
	proto "github.com/golang/protobuf/proto"
)

var defaultServerConfigPath = "/etc/leveler/config.yml"

func Read(path string, component string, config interface{}) error {
	var contents []byte
	var err error

	if len(path) == 0 {
		// read the default config
		if component == "client" {
			path, err = getDefaultClientConfigPath()
		} else {
			path = defaultServerConfigPath
		}
		
		if err != nil {
			return err
		}
	} 

	contents, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(contents, config.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func getDefaultClientConfigPath() (string, error) {
	u, err := user.Current() 
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/.leveler/client.yml", u.HomeDir), nil
}