package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"reflect"
)

type FilterConfigWrapper struct {
	FilterConfig
}

type FilterConfig interface {
}

type RegexpFilterConfig struct {
	Regexp string `json:"regexp"`
}

type HashFilterConfig struct {
	Salt    int64   `json:"salt,omitempty"`
	Percent float64 `json:"percent"`
}

type Config struct {
	Filter FilterConfigWrapper `json:"filter"`
}

func (w *FilterConfigWrapper) UnmarshalJSON(data []byte) error {
	var wrapper struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return err
	}
	switch wrapper.Type {
	case "regexp":
		var filter RegexpFilterConfig
		if err := json.Unmarshal(wrapper.Value, &filter); err != nil {
			return err
		}
		w.FilterConfig = filter
		return nil
	case "hash":
		var filter HashFilterConfig
		if err := json.Unmarshal(wrapper.Value, &filter); err != nil {
			return err
		}
		w.FilterConfig = filter
		return nil
	default:
		return errors.New("Unknown polymorph type: " + wrapper.Type)
	}
}

func (w FilterConfigWrapper) MarshalJSON() ([]byte, error) {
	var wrapper struct {
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	}
	switch w.FilterConfig.(type) {
	case RegexpFilterConfig:
		wrapper.Type = "regexp"
	case HashFilterConfig:
		wrapper.Type = "hash"
	default:
		return nil, errors.New("Unknown polymorph type: " + reflect.TypeOf(w.FilterConfig).String())
	}
	wrapper.Value = w.FilterConfig
	return json.Marshal(wrapper)
}

func main() {
	var config Config
	//language=json
	if err := json.Unmarshal([]byte(`{
	"filter": {
		"type": "regexp",
		"value": {
			"regexp": "^[a-z]+$"
		}
	}
}`), &config); err != nil {
		panic(err)
	}
	spew.Dump(config)
	data, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
