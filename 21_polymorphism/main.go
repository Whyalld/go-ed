package main

import "fmt"

type notifier interface {
	notify(message string)
}

type baseNotifier struct {
	senderName string
}

func (b *baseNotifier) formatMessage(msg string) string {
	return fmt.Sprintf("[%s]: %s", b.senderName, msg)
}

type smsNotifier struct {
	baseNotifier
	phoneNumber string
}

func (s *smsNotifier) notify(message string) {
	formatted := s.formatMessage(message)
	fmt.Printf("Отправлено SMS на номер %s -> %s\n", s.phoneNumber, formatted)
}

type emailNotifier struct {
	baseNotifier
	emailAdress string
}

func (e *emailNotifier) notify(message string) {
	formatted := e.formatMessage(message)
	fmt.Printf("Отправлено письмо на адрес %s -> %s\n", e.emailAdress, formatted)
}

func sendAlert(n notifier, text string) {
	n.notify(text)
}

func main20() {
	sms := &smsNotifier{
		baseNotifier: baseNotifier{senderName: "Система безопасности"},
		phoneNumber: "+7-932-838-83-91",
	}
	
	email := &emailNotifier{
		baseNotifier: baseNotifier{senderName: "Маркетинг"},
		emailAdress: "user@example.com",
	}

	fmt.Println("--- Запуск уведомлений ---")
	sendAlert(sms, "Обнаружен вход в аккаунт")
	sendAlert(email, "Вам начислено 500 бонусов")
}