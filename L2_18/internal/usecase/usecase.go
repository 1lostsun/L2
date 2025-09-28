package usecase

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/entity"
	"github.com/araddon/dateparse"
	"time"
)

// Repository : Интерфейс функций репозитория
type Repository interface {
	CreateEvent(e entity.Event) error
	UpdateEvent(e entity.Event) error
	DeleteEvent(userID uint64, title string, date time.Time) error
	EventsForDay(userID uint64, date time.Time) ([]entity.Event, error)
	EventsForWeek(userID uint64, date time.Time) ([]entity.Event, error)
	EventsForMonth(userID uint64, date time.Time) ([]entity.Event, error)
}

// UseCase : Структура юзкейса
type UseCase struct {
	repo Repository
}

// New : функция конструктор, возвращающая структуру UseCase
func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

// CreateEvent : Создание события
func (uc *UseCase) CreateEvent(event entity.Request) error {
	var body entity.Event

	date, err := dateparse.ParseAny(event.Date)
	if err != nil {
		return fmt.Errorf("invalid date: %s", event.Date)
	}

	body = entity.Event{
		UserID:      event.UserID,
		Date:        date,
		Title:       event.Title,
		Description: event.Description,
	}

	return uc.repo.CreateEvent(body)
}

// UpdateEvent : Обновление события
func (uc *UseCase) UpdateEvent(event entity.Request) error {
	var body entity.Event

	date, err := dateparse.ParseAny(event.Date)
	if err != nil {
		return fmt.Errorf("invalid date: %s", event.Date)
	}

	body = entity.Event{
		UserID:      event.UserID,
		Date:        date,
		Title:       event.Title,
		Description: event.Description,
	}

	return uc.repo.UpdateEvent(body)
}

// DeleteEvent : Удаление события
func (uc *UseCase) DeleteEvent(userID uint64, title, strDate string) error {
	date, err := dateparse.ParseAny(strDate)
	if err != nil {
		return fmt.Errorf("invalid date: %s", strDate)
	}

	return uc.repo.DeleteEvent(userID, title, date)
}

// EventsForDay : Парсит дату и возвращает события календаря за день передаваемой даты или ошибку
func (uc *UseCase) EventsForDay(userID uint64, strDate string) ([]entity.Event, error) {
	date, err := dateparse.ParseAny(strDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %s", strDate)
	}

	return uc.repo.EventsForDay(userID, date)
}

// EventsForWeek : Парсит дату и возвращает события календаря за неделю передаваемой даты или ошибку
func (uc *UseCase) EventsForWeek(userID uint64, strDate string) ([]entity.Event, error) {
	date, err := dateparse.ParseAny(strDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %s", strDate)
	}

	return uc.repo.EventsForWeek(userID, date)
}

// EventsForMonth : Парсит дату и возвращает события календаря за месяц передаваемой даты или ошибку
func (uc *UseCase) EventsForMonth(userID uint64, strDate string) ([]entity.Event, error) {
	date, err := dateparse.ParseAny(strDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %s", strDate)
	}

	return uc.repo.EventsForMonth(userID, date)
}
