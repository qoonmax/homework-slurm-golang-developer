package main

//go:generate mockgen -source $GOFILE -destination ./mock.go -package ${GOPACKAGE}

import (
	"errors"
	"fmt"
)

var (
	ErrDataNotFound = errors.New("data not found")
	ErrInvalidInput = errors.New("invalid input")
)

type MemoryStorage struct {
	data map[string]interface{}
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]interface{}),
	}
}

func (s *MemoryStorage) Get(key string) (interface{}, error) {
	if value, ok := s.data[key]; ok {
		return value, nil
	}
	return nil, ErrDataNotFound
}

func (s *MemoryStorage) Set(key string, value interface{}) {
	s.data[key] = value
}

type OutputNotifier struct{}

func NewOutputNotifier() *OutputNotifier {
	return &OutputNotifier{}
}

func (o *OutputNotifier) Send(key string) {
	fmt.Println(fmt.Sprintf("data saved for key: %s", key))
}

type Storage interface {
	Get(string) (interface{}, error)
	Set(string, interface{})
}

type Notifier interface {
	Send(message string)
}

type Processor struct {
	storage  Storage
	notifier Notifier
}

func NewProcessor(s Storage, n Notifier) *Processor {
	return &Processor{
		storage:  s,
		notifier: n,
	}
}

func (p *Processor) Process(key string, value interface{}) (interface{}, error) {
	if key == "" {
		return nil, ErrInvalidInput
	}

	data, err := p.storage.Get(key)
	if err != nil && !errors.Is(err, ErrDataNotFound) {
		return nil, err
	}

	if data != nil {
		return data, nil
	}

	p.storage.Set(key, value)

	if value != nil {
		p.notifier.Send(key)
	}

	return value, nil
}

func main() {
	memoryStorage := NewMemoryStorage()
	outputNotifier := NewOutputNotifier()

	processor := NewProcessor(
		memoryStorage,
		outputNotifier,
	)

	_, err := processor.Process("key1", 1)
	if err != nil {
		fmt.Println(fmt.Errorf("error for `key1`: %w", err))
	}

	_, err = processor.Process("key2", 2)
	if err != nil {
		fmt.Println(fmt.Errorf("error for `key2`: %w", err))
	}

	_, err = processor.Process("", 1)
	if err != nil {
		fmt.Println(fmt.Errorf("error for ``: %w", err))
	}

	_, err = processor.Process("key4", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("error for `key4`: %w", err))
	}
}
