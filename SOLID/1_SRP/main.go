package main

import "fmt"

// Чистая структура данных. Она просто хранит информацию
type Employee struct {
	ID int
	Name string
}

// Отдельная структура для Бухгалтерии (Актор: Финансы)
type FiancialReportal struct{}

func (fr *FiancialReportal) CalculatePay(e *Employee) int {
	// Здесь пишется реальная формула расчета зарплаты
	baseSalary := 50_000
	fmt.Printf("[Бухгалтерия]: Считаем зарплату для %s\n", e.Name)
	return baseSalary
}

// Отдельная структура для Отдела кадров (Актор: HR)
type HourReporter struct{}

func (hr *HourReporter) ReportHours(e *Employee) int {
	// Здесь пишется реальная логика подсчета отработанных часов
	workedHours := 160
	fmt.Printf("[Отдел кадров]: Считаем рабочие часы для %s\n", e.Name)
	return workedHours
}

// Отдельная структура для Базы данных (Актор: Сисадмины/DBA)
type EmployeeRepository struct{}

func (repo *EmployeeRepository) Save(e *Employee) {
	// Здесь пишется код сохранения сотрудника в базу данных
	fmt.Printf("[База данных]: Сотрудник %s (ID: %d) успешно схоранен!\n", e.Name, e.ID)
}

func main() {
	// Создаем объект сотрудника
	emp := &Employee{ID: 1, Name: "Саид"}

	// Создаем инструменты для разных отделов
	finance := FiancialReportal{}
	hr := HourReporter{}
	db := EmployeeRepository{}

	// Каждый отдел делает свою работу независимую друг друга
	salary := finance.CalculatePay(emp)
	hours := hr.ReportHours(emp)
	db.Save(emp)

	fmt.Printf("\nИтог: %s заработал %d руб. за %d часов.\n", emp.Name, salary, hours)
}
