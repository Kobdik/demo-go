package api

import (
	"demo-go/bin/config"
	"fmt"
)

func SomeRequest() error {
	conf, err := config.NewConfig()
	if err != nil {
		return err
	}
	fmt.Println(len(conf.Key))
	return nil
}
