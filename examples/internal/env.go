package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/sei-ri/kvconf/examples/internal/config"
	"log"
)

func main()  {
	var cfg config.Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("failed to load env config %v", err)
	}

	fmt.Println(cfg.Database)
	fmt.Println(cfg.Redis)
}