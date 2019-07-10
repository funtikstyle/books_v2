package service

import "books_v2/domain"

type booksService struct {
	store domain.PhoneStore
}

func NewService(s domain.PhoneStore) *booksService {
	return &booksService{store: s}
}

func (bs *booksService) GetContactListService() map[string]domain.Contact {
	return bs.store.GetContactListStore()
}

func (bs *booksService) GetContactService(c string) domain.Contact {
	return bs.store.GetContactStore(c)
}

func (bs *booksService) AddContactService(c domain.Contact) {
	bs.store.AddContactStore(c)
}

func (bs *booksService) DeleteContactService(c string) {
	bs.store.DeleteContactStore(c)
}

func (bs *booksService) UpdateContactService(key string, c domain.Contact) {
	bs.store.UpdateContactStore(key, c)
}
