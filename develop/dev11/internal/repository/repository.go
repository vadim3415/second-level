package repository

import (
	"errors"
	"fmt"
	"secondlevel/develop/dev11/internal/model"
	"sync"
	"time"
)

type IEvent interface {
	CreateEvent(event *model.Event) error
	UpdateEvent(event *model.Event) error
	DeleteEvent(event *model.Event) error

	EventsForDay(userId int, date time.Time) ([]model.Event, error)
	EventsForWeek(userId int, date time.Time) ([]model.Event, error)
	EventsForMonth(userId int, date time.Time) ([]model.Event, error)
}

type Repository struct {
	IEvent
}

func NewRepository() *Repository {
	return &Repository{
		IEvent: NewStorage(),
	}
}

//////////////////////////////////////////////////

type Storage struct {
	data map[string]*model.Event
	mtx  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]*model.Event),
		mtx:  sync.RWMutex{},
	}
}

func (d *Storage) CreateEvent(event *model.Event) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)

	if _, ok := d.data[id]; ok {
		return errors.New("data exist")
	}
	d.data[id] = event
	return nil
}

func (d *Storage) UpdateEvent(event *model.Event) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)

	if _, ok := d.data[id]; !ok {
		return errors.New("data not exist")

	}

	d.data[id] = event

	return nil
}

func (d *Storage) DeleteEvent(event *model.Event) error {
	d.mtx.Lock()
	defer d.mtx.Unlock()

	id := fmt.Sprintf("%d%d", event.UserID, event.EventID)

	if _, ok := d.data[id]; !ok {
		return errors.New("data not exist")
	}

	delete(d.data, id)

	return nil
}

func (d *Storage) EventsForDay(userId int, date time.Time) ([]model.Event, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var events []model.Event

	for _, event := range d.data {
		subDate := event.Date.Sub(date)
		if subDate == 0 && event.UserID == userId {
			events = append(events, *event)
		}
	}
	if len(events) < 1 {
		return nil, errors.New("not date event")
	}

	return events, nil
}

func (d *Storage) EventsForWeek(userId int, date time.Time) ([]model.Event, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var events []model.Event
	dateWeek := date.AddDate(0, 0, 7)

	for _, event := range d.data {
		subDate := event.Date.Sub(date)
		beforeDate := event.Date.Before(dateWeek)
		afterDate := event.Date.After(date)

		if (beforeDate && afterDate && event.UserID == userId) || (subDate == 0 && event.UserID == userId) {
			events = append(events, *event)
		}
	}

	if len(events) < 1 {
		return nil, errors.New("not date event")
	}

	return events, nil
}

func (d *Storage) EventsForMonth(userId int, date time.Time) ([]model.Event, error) {
	d.mtx.RLock()
	defer d.mtx.RUnlock()

	var events []model.Event
	dateMonth := date.AddDate(0, 1, 0)

	for _, event := range d.data {
		subDate := event.Date.Sub(date)
		beforeDate := event.Date.Before(dateMonth)
		afterDate := event.Date.After(date)

		if (beforeDate && afterDate && event.UserID == userId) || (subDate == 0 && event.UserID == userId) {
			events = append(events, *event)
		}
	}

	if len(events) < 1 {
		return nil, errors.New("not date event")
	}

	return events, nil
}
