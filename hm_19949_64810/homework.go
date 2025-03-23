package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Interval int

func (m *Interval) SetValue(s string) error {
	interval, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid data type: %v", err)
	}
	if interval >= 1 && interval <= 10 {
		*m = Interval(interval)
	} else {
		return fmt.Errorf("interval is not valid, must be between 1 and 10")
	}

	return nil
}

type Config struct {
	Interval Interval `env:"INTERVAL" env-required:"true"`
}

var (
	cfg Config
	mu  sync.RWMutex
)

func init() {
	mustLoad()
}

func mustLoad() {
	if err := godotenv.Overload(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	log.Printf("Config reloaded. New interval: %d", cfg.Interval)
}

func main() {
	sigHup := make(chan os.Signal, 1)
	sigExit := make(chan os.Signal, 1)
	signal.Notify(sigHup, syscall.SIGHUP)
	signal.Notify(sigExit, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mu.RLock()
				interval := time.Duration(cfg.Interval) * time.Second
				mu.RUnlock()

				log.Println("Working...", interval)
				time.Sleep(interval)
			}
		}
	}()

	for {
		select {
		case <-sigHup:
			mustLoad()
			log.Println("Reload config.")
		case <-sigExit:
			log.Printf("Received shutdown signal")
			close(done)
			return
		}
	}

}
