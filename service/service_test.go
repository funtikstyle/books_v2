package service

import (
	"books_v2/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testUser struct {
	u   domain.Contact
	res error
}

var testData = []testUser{
	{u: domain.Contact{Name: ""}, res: ErrNoName},
	{u: domain.Contact{Name: "    "}, res: ErrNoName},
	{u: domain.Contact{Name: "                 "}, res: ErrNoName},
	{u: domain.Contact{Name: "name           "}, res: nil},
	{u: domain.Contact{Name: "name"}, res: ErrIsExist},
}
