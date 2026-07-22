package main

import "fmt"

// Абстракция: контракт, который нужен верхнему уровню
type SMSSender interface {
	Send(phone string, msg string)
}

// Верхний уровень
type OrderService struct {
	sender SMSSender // Зависим от абстракции (интерфейса)
}

// Конструктор, куда внедряем зависимость (Dependency Injection)
func NewOrderService(s SMSSender) *OrderService {
	return &OrderService{sender: s}
}

func (s *OrderService) CreateOrder() {
	fmt.Println("Заказ создан успешно.")
	s.sender.Send("+7-999-111-22-33", "Ваш заказ готов!")
}

type TwilioSender struct{}

func (t *TwilioSender) Send(phone string, msg string) {
	fmt.Printf("[Twilio] SMS на %s: %s\n", phone, msg)
}

type SmsRuSender struct{}

func (s *SmsRuSender) Send(phone string, msg string) {
	fmt.Printf("[SMS.ru] SMS на %s: %s:\n", phone, msg)
}

type MockSender struct{}

func (m *MockSender) Send(phone string, msg string) {
	fmt.Println("[TEST] Проверка вызова отправки SMS без реальной сети.")
}

func main() {
	// Хотим работать через Twilio? Создаем его и передаем в сервис
	twilio := &TwilioSender{}
	serviceWithTwilio := NewOrderService(twilio)
	serviceWithTwilio.CreateOrder()

	// Завтра передумали и хотим SMS.ru? Меняем ОДНУ строчку в main!
	smsRu := &SmsRuSender{}
	serviceWithSmsRu := NewOrderService(smsRu)
	serviceWithSmsRu.CreateOrder() // Сам код CreateOrder мы даже не трогали!
}
