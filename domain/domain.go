package domain

type Contact struct {
	Name  string
	Phone string
}
type PhoneService interface {
	GetContactListService() map[string]Contact
	AddContactService(c Contact)
	DeleteContactService(c string)
	GetContactService(c string) Contact
	UpdateContactService(key string, c Contact)
}

type PhoneStore interface {
	GetContactListStore() map[string]Contact
	AddContactStore(c Contact)
	DeleteContactStore(c string)
	GetContactStore(c string) Contact
	UpdateContactStore(key string, c Contact)
}
