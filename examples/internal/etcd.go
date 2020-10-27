package main

import (
	"context"
	"fmt"
	"github.com/sei-ri/kvconf/etcd"
	"github.com/sei-ri/kvconf/examples/internal/config"
	"gopkg.in/yaml.v2"
	"log"
)

func main()  {
	ctx := context.Background()
	client, err := etcd.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var cfg config.Config
	resp, err := client.Get(ctx, "app/dev")
	if err != nil {
		log.Println(err)
	}

	if err := yaml.Unmarshal(resp, &cfg); err != nil {
		log.Println("failed to unmarshal yaml %v", err)
	}

	fmt.Println("database: ", cfg.Database)
	fmt.Println("redis: ", cfg.Redis)
}