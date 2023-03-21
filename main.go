package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type conversion struct {
	from     string
	ingester func([]byte, any) error
	to       string
	spewer   func(any) ([]byte, error)
}

var converters = map[string]conversion{
	"yj": {"yaml", yaml.Unmarshal, "json", json.Marshal},
	"jy": {"json", json.Unmarshal, "yaml", yamlMarshal},
}

// yamlMarshal wraps creating an encoder without horrible 4-space indents
func yamlMarshal(object any) ([]byte, error) {
	buf := &bytes.Buffer{}
	yamlencoder := yaml.NewEncoder(buf)
	yamlencoder.SetIndent(2)
	err := yamlencoder.Encode(object)
	return buf.Bytes(), err
}

func convert(in []byte, converter string) ([]byte, error) {
	var object any
	if _, ok := converters[converter]; !ok {
		return []byte{}, fmt.Errorf("conversion \"%s\" is not supported", converter)
	}
	c := converters[converter]
	err := c.ingester(in, &object)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to unmarshal %s: %w", c.from, err)
	}
	out, err := c.spewer(object)
	if err != nil {
		return out, fmt.Errorf("failed to marshal %s: %w", c.to, err)
	}
	return out, nil
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read: %s\n", err)
		os.Exit(1)
	}
	l := len(os.Args[0])
	out, err := convert(in, os.Args[0][l-2:l])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert: %s\n", err)
		os.Exit(1)
	}
	fmt.Print(string(out))
}
