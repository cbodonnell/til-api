package workers

import (
	"log"
	"time"

	"github.com/cbodonnell/til-api/config"
)

type ExampleWorker struct {
	conf    config.Configuration
	channel chan interface{}
}

func NewExampleWorker(_conf config.Configuration) Worker {
	return &ExampleWorker{
		conf:    _conf,
		channel: make(chan interface{}),
	}
}

func (f *ExampleWorker) Start() {
	go func() {
		for {
			log.Println("Doing some work!")
			time.Sleep(5 * time.Second)
		}
	}()
}

func (f *ExampleWorker) GetChannel() chan interface{} {
	return f.channel
}
