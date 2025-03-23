package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"net"
	"strconv"
)

type Host string
type Port int
type Mode int

func (h *Host) SetValue(s string) error {
	if net.ParseIP(s) != nil {
		*h = Host(s)
	} else {
		return fmt.Errorf("host is not valid")
	}

	return nil
}

func (p *Port) SetValue(s string) error {
	port, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid data type")
	}
	if port >= 0 && port <= 51820 {
		*p = Port(port)
	} else {
		return fmt.Errorf("port is not valid")
	}

	return nil
}

func (m *Mode) SetValue(s string) error {
	mode, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid data type")
	}
	if mode >= 1 && mode <= 3 {
		*m = Mode(mode)
	} else {
		return fmt.Errorf("mode is not valid")
	}

	return nil
}

type Config struct {
	Host Host `env:"HOST" env-default:"127.0.0.1"`
	Port Port `env:"PORT" env-default:"8000"`
	Mode Mode `env:"MODE" env-required:"true"`
}

var cfg Config

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Host:", cfg.Host)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("Mode:", cfg.Mode)
}
