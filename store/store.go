package store

import (
	"books_v2/domain"
)

type PhoneBook struct {
	list map[string]domain.Contact
}

func NewBooks() *PhoneBook {
	return &PhoneBook{list: map[string]domain.Contact{}}
}

func (pb *PhoneBook) GetContactListStore() map[string]domain.Contact {

	return pb.list
}

func (pb *PhoneBook) GetContactStore(c string) domain.Contact {
	return pb.list[c]
}

func (pb *PhoneBook) AddContactStore(c domain.Contact) {
	pb.list[c.Name] = c
}

func (pb *PhoneBook) DeleteContactStore(c string) {
	delete(pb.list, c)

}

func (pb *PhoneBook) UpdateContactStore(key string, c domain.Contact) {
	pb.list[key] = c
}
