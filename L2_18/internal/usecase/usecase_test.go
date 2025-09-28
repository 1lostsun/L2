package usecase

import (
	"errors"
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateEvent(e entity.Event) error {
	return m.Called(e).Error(0)
}

func (m *MockRepo) UpdateEvent(e entity.Event) error {
	return m.Called(e).Error(0)
}

func (m *MockRepo) DeleteEvent(userID uint64, title string, date time.Time) error {
	return m.Called(userID, title, date).Error(0)
}

func (m *MockRepo) EventsForDay(userID uint64, date time.Time) ([]entity.Event, error) {
	args := m.Called(userID, date)
	return args.Get(0).([]entity.Event), args.Error(1)
}

func (m *MockRepo) EventsForWeek(userID uint64, date time.Time) ([]entity.Event, error) {
	args := m.Called(userID, date)
	return args.Get(0).([]entity.Event), args.Error(1)
}

func (m *MockRepo) EventsForMonth(userID uint64, date time.Time) ([]entity.Event, error) {
	args := m.Called(userID, date)
	return args.Get(0).([]entity.Event), args.Error(1)
}

func TestUseCase_CreateEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        "2025-09-25",
	}

	mockRepo.On("CreateEvent", mock.MatchedBy(func(e entity.Event) bool {
		return e.UserID == req.UserID &&
			e.Title == req.Title &&
			e.Description == req.Description &&
			e.Date.Format("2006-01-02") == req.Date
	})).Return(nil)

	err := uc.CreateEvent(req)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUseCase_CreateEvent_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из системы",
		Date:        "invalid date 10000",
	}

	err := uc.CreateEvent(req)
	assert.Error(t, err)
}

func TestUseCase_UpdateEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)
	req := entity.Request{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из аккаунта",
		Date:        "2025-09-25",
	}

	mockRepo.On("UpdateEvent", mock.MatchedBy(func(e entity.Event) bool {
		return e.UserID == req.UserID &&
			e.Title == req.Title &&
			e.Date.Format("2006-01-02") == req.Date
	})).Return(nil)

	err := uc.UpdateEvent(req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_UpdateEvent_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID:      42,
		Title:       "Выход",
		Description: "Пользователь вышел из аккаунта",
		Date:        "invalid date 10000",
	}

	err := uc.UpdateEvent(req)
	assert.Error(t, err)
}

func TestUseCase_UpdateEvent_NoEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID: 42,
		Title:  "invalid 10000",
		Date:   "2025-09-25",
	}

	mockRepo.On("UpdateEvent", mock.MatchedBy(func(e entity.Event) bool {
		return e.UserID == req.UserID &&
			e.Title == req.Title &&
			e.Date.Format("2006-01-02") == req.Date
	})).Return(errors.New("event not found"))

	err := uc.UpdateEvent(req)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_DeleteEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID: 42,
		Title:  "Выход",
		Date:   "2025-09-25",
	}

	mockRepo.On(
		"DeleteEvent",
		req.UserID,
		req.Title,
		mock.MatchedBy(func(d time.Time) bool {
			return d.Format("2006-01-02") == req.Date
		}),
	).Return(nil)

	err := uc.DeleteEvent(req.UserID, req.Title, req.Date)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_DeleteEvent_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID: 42,
		Title:  "Выход",
		Date:   "invalid date 10000",
	}

	mockRepo.On("DeleteEvent", req.UserID, req.Title, req.Date).Return(fmt.Errorf("invalid date: %s", req.Date))

	err := uc.DeleteEvent(req.UserID, req.Title, req.Date)
	assert.Error(t, err)
}

func TestUseCase_DeleteEvent_InvalidEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID: 42,
		Title:  "10000",
		Date:   "2025-09-25",
	}

	mockRepo.On("DeleteEvent",
		req.UserID,
		req.Title,
		mock.AnythingOfType("time.Time")).Return(errors.New("event not found"))

	err := uc.DeleteEvent(req.UserID, req.Title, req.Date)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_DeleteEvent_NoEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	req := entity.Request{
		UserID: 42,
		Title:  "10000",
		Date:   "2025-09-25",
	}

	parsedDate, _ := time.Parse("2006-01-02", req.Date)

	mockRepo.On("DeleteEvent",
		req.UserID,
		req.Title,
		parsedDate).Return(errors.New("event not found"))

	err := uc.DeleteEvent(req.UserID, req.Title, req.Date)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_EventsForDay(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)
	expectedEvents := []entity.Event{{UserID: userID, Title: "Meeting", Date: parsedDate}}

	mockRepo.On("EventsForDay", userID, parsedDate).Return(expectedEvents, nil)

	events, err := uc.EventsForDay(userID, strDate)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_EventsForDay_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "invalid date 10000"

	events, err := uc.EventsForDay(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestUseCase_EventsForDay_NoEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)

	mockRepo.On("EventsForDay", userID, parsedDate).Return([]entity.Event(nil), errors.New("event not found"))

	events, err := uc.EventsForDay(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestUseCase_EventsForWeek(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)
	expectedEvents := []entity.Event{{UserID: userID, Title: "Meeting", Date: parsedDate}}

	mockRepo.On("EventsForWeek", userID, parsedDate).Return(expectedEvents, nil)

	events, err := uc.EventsForWeek(userID, strDate)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_EventsForWeek_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "invalid date 10000"

	events, err := uc.EventsForWeek(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestUseCase_EventsForWeek_NoEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)

	mockRepo.On("EventsForWeek", userID, parsedDate).Return([]entity.Event(nil), errors.New("event not found"))

	events, err := uc.EventsForWeek(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestUseCase_EventsForMonth(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)
	expectedEvents := []entity.Event{{UserID: userID, Title: "Meeting", Date: parsedDate}}

	mockRepo.On("EventsForMonth", userID, parsedDate).Return(expectedEvents, nil)

	events, err := uc.EventsForMonth(userID, strDate)
	assert.NoError(t, err)
	assert.Equal(t, expectedEvents, events)
	mockRepo.AssertExpectations(t)
}

func TestUseCase_EventsForMonth_InvalidDate(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "invalid date 10000"

	events, err := uc.EventsForMonth(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}

func TestUseCase_EventsForMonth_NoEvent(t *testing.T) {
	mockRepo := new(MockRepo)
	uc := New(mockRepo)

	userID := uint64(42)
	strDate := "2025-09-25"
	parsedDate, _ := time.Parse("2006-01-02", strDate)

	mockRepo.On("EventsForMonth", userID, parsedDate).Return([]entity.Event(nil), errors.New("event not found"))

	events, err := uc.EventsForMonth(userID, strDate)
	assert.Error(t, err)
	assert.Nil(t, events)
}
