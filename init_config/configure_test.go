package init_config

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {
	f, err := os.Open("conf_dev.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	conf := configure{}
	err = decoder.Decode(&conf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(conf.Server.Name)
}
