package api

import (
	"demo-go/bin/config"
	"fmt"
)

func SomeRequest(conf *config.Config) error {
	fmt.Println(len(conf.Key))
	return nil
}
