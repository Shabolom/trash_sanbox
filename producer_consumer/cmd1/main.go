package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Event struct {
	ID       uint    `json:"id"`
	CarModel string  `json:"car_model"`
	Price    float64 `json:"price"`
}

type Producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(filename string) (*Producer, error) {
	// откройте файл и создайте для него json.Encoder
	// допишите код
	// ...
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &Producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *Producer) WriteEvent(event *Event) error {
	// добавьте вызов Encode для сериализации и записи
	// допишите код
	// ...
	err := p.encoder.Encode(event)
	if err != nil {
		return err
	}
	return nil
}

type Consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(filename string) (*Consumer, error) {
	// откройте файл и создайте для него json.Decoder
	// допишите код
	// ...
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (c *Consumer) ReadEvent() (*Event, error) {
	// добавьте вызов Decode для чтения и десериализации
	// допишите код
	// ...
	event := &Event{}
	err := c.decoder.Decode(event)
	if err != nil {
		return nil, err
	}
	return event, err
}

var events = []*Event{
	{
		ID:       1,
		CarModel: "Lada",
		Price:    400000,
	},
	{
		ID:       2,
		CarModel: "Mitsubishi",
		Price:    650000,
	},
	{
		ID:       3,
		CarModel: "Toyota",
		Price:    800000,
	},
	{
		ID:       4,
		CarModel: "BMW",
		Price:    875000,
	},
	{
		ID:       5,
		CarModel: "Mercedes",
		Price:    999999,
	},
}

func main() {
	fileName := "events.log"
	defer os.Remove(fileName)

	Producer, err := NewProducer(fileName)
	if err != nil {
		log.Fatal(err)
	}

	Consumer, err := NewConsumer(fileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		if err := Producer.WriteEvent(event); err != nil {
			log.Fatal(err)
		}

		readEvent, err := Consumer.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(readEvent)
	}
}
