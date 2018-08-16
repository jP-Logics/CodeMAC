// Package configuration is used to keep configuration in a map that can easily accessable
package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Configuration is used to maintain configuration
type Configuration map[string]interface{}

// New loads information from the file
func New(fn string) (Configuration, error) {
	bArr, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	var configData Configuration
	err = json.Unmarshal(bArr, &configData)
	if err != nil {
		return nil, err
	}
	return configData, nil
}

// Key fetches all keywords
func (c Configuration) Key(key string) (Configuration, error) {
	if val, ok := c[key]; ok {
		return Configuration((val).(map[string]interface{})), nil
	}
	return nil, errors.New("no data")
}

// Get fetches all keywords
func (c Configuration) Get(key string) Configuration {
	if val, ok := c[key]; ok {
		return Configuration((val).(map[string]interface{}))
	}
	return nil
}

// Value gives value for a key . Usually Value is used for the leaf value
func (c Configuration) Value(key string) interface{} {
	if val, ok := c[key]; ok {
		return val
	}
	return nil
}
