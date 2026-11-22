package main

import (
	"fmt"
	"sync"
)

type SensorData struct {
	ID    string
	Value int
}

type Subscriber interface {
	Notify(data string)
}

type TgSubscriber struct{}
func (t *TgSubscriber) Notify(d string) { fmt.Printf("[Tg] %s\n", d) }

type EmailSubscriber struct{}
func (e *EmailSubscriber) Notify(d string) { fmt.Printf("[Email] %s\n", d) }

type Handler interface {
	SetNext(Handler) Handler
	Execute(*SensorData) bool
}

type BaseHandler struct {
	next Handler
}
func (b *BaseHandler) SetNext(next Handler) Handler {
	b.next = next
	return next
}
func (b *BaseHandler) ExecuteNext(data *SensorData) bool {
	if b.next != nil {
		return b.next.Execute(data)
	}
	return true
}

type LoggerHandler struct{ BaseHandler }
func (l *LoggerHandler) Execute(data *SensorData) bool {
	fmt.Printf("[Log] ID: %s, Val: %d\n", data.ID, data.Value)
	return l.ExecuteNext(data)
}

type ValidateHandler struct { BaseHandler }
func (v *ValidateHandler) Execute(data *SensorData) bool {
	if data.Value < 0 {
		return false
	}
	return v.ExecuteNext(data)
}

type AccessHandler struct {
	BaseHandler
	whiteList map[string]bool
}
func (a *AccessHandler) Execute(data *SensorData) bool {
	if !a.whiteList[data.ID] {
		fmt.Printf("[Access Denied] %s\n", data.ID)
		return false
	}
	return a.ExecuteNext(data)
}

type ReporterHandler struct {
	BaseHandler
	subs []Subscriber
}
func (r *ReporterHandler) Execute(data *SensorData) bool {
	if data.Value > 10 {
		msg := fmt.Sprintf("Sensor %s triggered: %d", data.ID, data.Value)
		for _, sub := range r.subs {
			sub.Notify(msg)
		}
		return r.ExecuteNext(data)
	}
	return false
}

type SensorDataProcessor struct {
	chain Handler
}

var (
	processorInstance *SensorDataProcessor
	processorOnce     sync.Once
)

func GetProcessorInstance(allowedIDs []string) *SensorDataProcessor {
	processorOnce.Do(func() {
		wMap := make(map[string]bool)
		for _, id := range allowedIDs {
			wMap[id] = true
		}

		logger := &LoggerHandler{}
		validator := &ValidateHandler{}
		access := &AccessHandler{whiteList: wMap}
		reporter := &ReporterHandler{
			subs: []Subscriber{&TgSubscriber{}, &EmailSubscriber{}},
		}

		logger.SetNext(validator).SetNext(access).SetNext(reporter)

		processorInstance = &SensorDataProcessor{chain: logger}
	})
	return processorInstance
}

func (p *SensorDataProcessor) Process(data *SensorData) {
	p.chain.Execute(data)
}
