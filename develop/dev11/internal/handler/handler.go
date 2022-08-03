package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"secondlevel/develop/dev11/internal/model"
	"secondlevel/develop/dev11/internal/service"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)

	handler := Logging(mux)
	return handler
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respErr400(w)
		return
	}
	event, err := DecodeJSON(r)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}
	err = h.services.CreateEvent(event)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respPost(w)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respErr400(w)
		return
	}

	event, err := DecodeJSON(r)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}

	err = h.services.UpdateEvent(event)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respPost(w)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respErr400(w)
		return
	}

	event, err := DecodeJSON(r)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}

	err = h.services.DeleteEvent(event)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respPost(w)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respErr400(w)
		return
	}

	userId, date, err := getParams(r.URL)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}

	events, err := h.services.EventsForDay(userId, date)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respGet(w, events)

}
func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respErr400(w)
		return
	}

	userId, date, err := getParams(r.URL)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}

	events, err := h.services.EventsForWeek(userId, date)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respGet(w, events)
}
func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respErr400(w)
		return
	}

	userId, date, err := getParams(r.URL)
	if err != nil {
		logrus.Println(err)
		respErr400(w)
		return
	}

	events, err := h.services.EventsForMonth(userId, date)
	if err != nil {
		logrus.Println(err)
		respErr503(w)
		return
	}

	respGet(w, events)
}

func DecodeJSON(r *http.Request) (*model.Event, error) {
	event := &model.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		return nil, err
	}
	if event.UserID < 1 && event.EventID < 1 {
		return nil, errors.New("incorrect Id value")
	}

	return event, nil
}

func chekDate(date string) (time.Time, error) {
	//myTime, err := time.Parse("2006-01-02", date)
	myTime, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return time.Time{}, errors.New("incorrect date, layout = 2022-07-27T15:04:05Z")
	}

	return myTime, nil
}

func getParams(url *url.URL) (int, time.Time, error) {
	userId := url.Query().Get("user_id")
	date := url.Query().Get("date")
	//date := url.Query().Get("date") + "T00:00:00Z"

	id, err := strconv.Atoi(userId)
	if err != nil || id < 1 {
		if id < 1 {
			err = errors.New("incorrect id")
		}
		return 0, time.Time{}, err
	}

	myTime, err := chekDate(date)
	if err != nil {
		logrus.Println(err)
		return 0, time.Time{}, err
	}
	return id, myTime, nil
}

func respErr400(w http.ResponseWriter){
	resp, _ := json.Marshal(map[string]int{
			"error": http.StatusBadRequest,
		})
		w.Write(resp)
}

func respErr503(w http.ResponseWriter){
	resp, _ := json.Marshal(map[string]int{
			"error": http.StatusServiceUnavailable,
		})
		w.Write(resp)
}

func respGet(w http.ResponseWriter, events []model.Event){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(map[string][]model.Event{
			"result": events,
	})
	w.Write(result)
}

func respPost(w http.ResponseWriter){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, _ := json.Marshal(map[string]string{
			"result": "successfully",
	})
	w.Write(result)
}
