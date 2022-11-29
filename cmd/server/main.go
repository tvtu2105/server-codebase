package main

import (
	"fmt"
	"sever-codebase/application/configs"
)

func main() {
	env, err := configs.LoadEnv()
	cfg, err := configs.LoadConfig("conf/", env)
	if err != nil {
		return
	}
	fmt.Println("Config: ", cfg)

}
