package main

import (
	"asyawear/server/internal/adapter"
	"asyawear/server/internal/config"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func readConf(filename string) (*config.Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &config.Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config from file %q: %w", filename, err)
	}
	return c, err
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	args := os.Args[1:]

	if len(args) == 0 {
		log.Println("you should provide config path")
		os.Exit(1)
	}

	configPath := args[0]

	if configPath == "" {
		log.Println("you should provide correct path")
		os.Exit(1)
	}

	cfg, err := readConf(configPath)
	if err != nil {
		log.Printf("cannot read config file: %v", err)
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		adapter.StartApp(ctx, cfg)
	}()

	wg.Wait()
}
