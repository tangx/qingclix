package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func Test_loadConfig(t *testing.T) {
	config := loadPreinstallConfig()
	for k := range config.Instance {
		fmt.Println("k =", k)
	}

	// ChoosePreinstanllConfig(config)

	ins := config.Instance["ubuntu1804"]

	m := make(map[string]string)
	// j := InstanceItem{}
	b, _ := json.Marshal(ins)
	json.Unmarshal(b, &m)
	fmt.Println(m)

	uv := url.Values{}
	for k, v := range m {
		uv.Set(k, v)
	}
	fmt.Println(uv.Encode())
}
