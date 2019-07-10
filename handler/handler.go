package handler

import (
	"books_v2/domain"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type PhoneService struct {
	srv domain.PhoneService
}

func NewHandlers(s domain.PhoneService) *PhoneService {
	return &PhoneService{srv: s}
}

func (h *PhoneService) GetContactList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//if h.srv.CountContacts() != 0 {

	fmt.Fprintln(w, "Список контактов:")
	//fmt.Fprintln(w, "Количество контактов: ", h.srv.CountContacts())
	for _, v := range h.srv.GetContactListService() {
		fmt.Fprintf(w, "Телефон: - %s Имя: - %s\n", v.Phone, v.Name)
	}
	return
	//}
	//fmt.Fprintf(w, "Список контактов пуст.")
}

func (h *PhoneService) GetContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	h.srv.GetContactListService()
}

func (h *PhoneService) AddContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contact := domain.Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Println(err.Error())
		return
	}
	h.srv.AddContactService(contact)
}

func (h *PhoneService) DeleteContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := ps.ByName("id")
	h.srv.DeleteContactService(c)
}

func (h *PhoneService) UpdateContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contact := domain.Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := ps.ByName("id")
	h.srv.UpdateContactService(c, contact)
}
