package service

import (
	"secondlevel/develop/dev11/internal/model"
	"secondlevel/develop/dev11/internal/repository"
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

type Service struct {
	IEvent
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		IEvent: NewUserService(repos.IEvent),
	}
}

///////////////////////////////////////////////////

type UserService struct {
	repo repository.IEvent
}

func NewUserService(repo repository.IEvent) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateEvent(event *model.Event) error {
	return s.repo.CreateEvent(event)
}

func (s *UserService) UpdateEvent(event *model.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *UserService) DeleteEvent(event *model.Event) error {
	return s.repo.DeleteEvent(event)
}

func (s *UserService) EventsForDay(userId int, date time.Time) ([]model.Event, error) {
	return s.repo.EventsForDay(userId, date)
}

func (s *UserService) EventsForWeek(userId int, date time.Time) ([]model.Event, error) {
	return s.repo.EventsForWeek(userId, date)
}

func (s *UserService) EventsForMonth(userId int, date time.Time) ([]model.Event, error) {
	return s.repo.EventsForMonth(userId, date)
}
