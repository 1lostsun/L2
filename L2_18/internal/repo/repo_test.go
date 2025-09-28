package repo

import (
	"errors"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRepo_CreateEvent(t *testing.T) {
	r := &Repo{}
	err := r.CreateEvent(entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	})

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	}
	err = r.CreateEvent(event)

	assert.NoError(t, err)
	assert.Equal(t, err, nil)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event = entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	}
	r.mapRepo[event.UserID] = []entity.Event{event}
	err = r.CreateEvent(event)

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("event already exists"))
}

func TestRepo_UpdateEvent(t *testing.T) {
	r := &Repo{}
	err := r.UpdateEvent(entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	})

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	}

	r.mapRepo[event.UserID] = []entity.Event{event}
	err = r.UpdateEvent(event)
	assert.NoError(t, err)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event = entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        time.Now(),
	}

	err = r.UpdateEvent(event)
	assert.Error(t, err)
	assert.Equal(t, err, errors.New("event does not exist"))
}

func TestRepo_DeleteEvent(t *testing.T) {
	r := &Repo{}
	now := time.Now()
	err := r.DeleteEvent(42, "Выход", now)

	assert.Error(t, err)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        now,
	}

	r.mapRepo[event.UserID] = []entity.Event{event}
	err = r.DeleteEvent(event.UserID, event.Title, now)
	assert.NoError(t, err)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event = entity.Event{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        now,
	}

	err = r.DeleteEvent(event.UserID, event.Title, event.Date)
	assert.Error(t, err)
	assert.Equal(t, err, errors.New("event not found"))
}

func TestRepo_EventsForDay(t *testing.T) {
	r := &Repo{}

	userID := uint64(42)
	date := time.Date(2025, 9, 25, 0, 0, 0, 0, time.UTC)

	events, err := r.EventsForDay(userID, date)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event1 := entity.Event{UserID: userID, Title: "Meeting", Date: date}
	event2 := entity.Event{UserID: userID, Title: "Lunch", Date: date}
	r.mapRepo[userID] = []entity.Event{event1, event2}

	events, err = r.EventsForDay(userID, date)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))
	assert.Contains(t, events, event1)
	assert.Contains(t, events, event2)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{UserID: userID, Title: "Meeting", Date: date}
	r.mapRepo[userID] = []entity.Event{event}

	targetDate := time.Date(2025, 9, 26, 0, 0, 0, 0, time.UTC)
	events, err = r.EventsForDay(userID, targetDate)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("event not found"))
}

func TestRepo_EventsForWeek(t *testing.T) {
	r := &Repo{}

	userID := uint64(42)
	dateWeek1 := time.Date(2025, 9, 25, 0, 0, 0, 0, time.UTC)
	dateWeek2 := time.Date(2025, 10, 1, 0, 0, 0, 0, time.UTC) // следующая неделя

	events, err := r.EventsForWeek(userID, dateWeek1)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event1 := entity.Event{UserID: userID, Title: "Meeting", Date: dateWeek1} // неделя 39
	event2 := entity.Event{UserID: userID, Title: "Lunch", Date: dateWeek2}   // неделя 40
	r.mapRepo[userID] = []entity.Event{event1, event2}

	events, err = r.EventsForWeek(userID, dateWeek1)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Contains(t, events, event1)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{UserID: userID, Title: "Meeting", Date: dateWeek2}
	r.mapRepo[userID] = []entity.Event{event}

	targetDate := time.Date(2025, 9, 22, 0, 0, 0, 0, time.UTC) // неделя 39
	events, err = r.EventsForWeek(userID, targetDate)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("event not found"))
}

func TestRepo_EventsForMonth(t *testing.T) {
	r := &Repo{}

	userID := uint64(42)
	dateSep := time.Date(2025, 9, 25, 0, 0, 0, 0, time.UTC)
	dateOct := time.Date(2025, 10, 5, 0, 0, 0, 0, time.UTC)

	events, err := r.EventsForMonth(userID, dateSep)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("map repo is nil"))

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event1 := entity.Event{UserID: userID, Title: "Meeting", Date: dateSep} // сентябрь
	event2 := entity.Event{UserID: userID, Title: "Lunch", Date: dateOct}   // октябрь
	r.mapRepo[userID] = []entity.Event{event1, event2}

	events, err = r.EventsForMonth(userID, dateSep)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Contains(t, events, event1)

	//-----------------------------------

	r.mapRepo = make(map[uint64][]entity.Event)
	event := entity.Event{UserID: userID, Title: "Meeting", Date: dateOct}
	r.mapRepo[userID] = []entity.Event{event}

	targetDate := time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC)
	events, err = r.EventsForMonth(userID, targetDate)
	assert.Nil(t, events)
	assert.Equal(t, err, errors.New("event not found"))
}
