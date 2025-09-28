package repo

import (
	"errors"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/entity"
	"sync"
	"time"
)

// Repo : Структура локального репозитория
type Repo struct {
	mapRepo map[uint64][]entity.Event
	mx      sync.RWMutex
}

// New : функция конструктор, возвращающая структуру Repo
func New() *Repo {
	return &Repo{
		mapRepo: make(map[uint64][]entity.Event),
		mx:      sync.RWMutex{},
	}
}

// CreateEvent : Создание события
func (r *Repo) CreateEvent(e entity.Event) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return errors.New("map repo is nil")
	}

	for _, v := range r.mapRepo[e.UserID] {
		if v == e {
			return errors.New("event already exists")
		}
	}

	r.mapRepo[e.UserID] = append(r.mapRepo[e.UserID], e)
	return nil
}

// UpdateEvent : Обновление события
func (r *Repo) UpdateEvent(e entity.Event) error {
	flag := false
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return errors.New("map repo is nil")
	}

	for i, v := range r.mapRepo[e.UserID] {
		if v.Date.Equal(e.Date) && v.Title == e.Title {
			r.mapRepo[e.UserID][i] = e
			flag = true
			break
		}
	}

	if !flag {
		return errors.New("event does not exist")
	}

	return nil
}

// DeleteEvent : Удаление событие
func (r *Repo) DeleteEvent(userID uint64, title string, date time.Time) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return errors.New("map repo is nil")
	}

	events := make([]entity.Event, 0)

	for _, v := range r.mapRepo[userID] {
		if v.Date.Equal(date) && v.Title == title {
			continue
		}

		events = append(events, v)
	}

	if len(events) == len(r.mapRepo[userID]) {
		return errors.New("event not found")
	}

	r.mapRepo[userID] = events
	return nil
}

// EventsForDay : Возвращает события календаря за день передаваемой даты или ошибку
func (r *Repo) EventsForDay(userID uint64, date time.Time) ([]entity.Event, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return nil, errors.New("map repo is nil")
	}

	events := make([]entity.Event, 0)
	for _, v := range r.mapRepo[userID] {
		if v.Date.Day() == date.Day() && v.Date.Month() == date.Month() && v.Date.Year() == date.Year() {
			events = append(events, v)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("event not found")
	}

	return events, nil
}

// EventsForWeek : Возвращает события календаря за неделю передаваемой даты или ошибку
func (r *Repo) EventsForWeek(userID uint64, date time.Time) ([]entity.Event, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return nil, errors.New("map repo is nil")
	}

	events := make([]entity.Event, 0)
	for _, v := range r.mapRepo[userID] {
		repoYear, repoWeek := v.Date.ISOWeek()
		requestYear, requestWeek := date.ISOWeek()
		if repoWeek == requestWeek && v.Date.Month() == date.Month() && repoYear == requestYear {
			events = append(events, v)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("event not found")
	}

	return events, nil
}

// EventsForMonth : Возвращает события календаря за месяц передаваемой даты или ошибку
func (r *Repo) EventsForMonth(userID uint64, date time.Time) ([]entity.Event, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	if r.mapRepo == nil {
		return nil, errors.New("map repo is nil")
	}

	events := make([]entity.Event, 0)
	for _, v := range r.mapRepo[userID] {
		if v.Date.Month() == date.Month() && v.Date.Year() == date.Year() {
			events = append(events, v)
		}
	}

	if len(events) == 0 {
		return nil, errors.New("event not found")
	}

	return events, nil
}
