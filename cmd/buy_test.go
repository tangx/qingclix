package cmd

import (
	"fmt"
	"testing"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
)

func Test_LoadPresetCofnig(t *testing.T) {
	preset := LoadPresetConfig()
	// config := ChooseConfig(preset)
	config := preset.Configs["dev--ubuntu1604-2c4g-200g"]
	fmt.Println(config)
	values, err := query.Values(config)
	if err != nil {
		logrus.Fatal("query.Values=", err)
	}
	// fmt.Println(values)
	fmt.Println(values.Encode())

}
