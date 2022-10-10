package utils

import (
	"encoding/json"
	"errors"
	"os"
)

type config struct {
	config map[string]interface{}
}

func NewConfig() *config {
	t, err := os.ReadFile("config.json")
	if err != nil {
		panic("Configuration file error. config.json file shoulb be on exec path")
	}
	c := config{}
	if err := json.Unmarshal(t, &c.config); err != nil {
		panic("Configuration file invalid json")
	}
	return &c
}

func (c *config) GetGroup(group string) (map[string]interface{}, error) {
	g, ok := c.config[group]
	if !ok {
		return nil, errors.New("valid group")
	}
	return g.(map[string]interface{}), nil
}

func (c *config) GetField(group string, field string) (string, error) {
	g, err := c.GetGroup(group)
	if err != nil {
		return "", err
	}
	g1, ok := g[field]
	if !ok {
		return "", errors.New("valid field")
	}
	return g1.(string), nil
}
