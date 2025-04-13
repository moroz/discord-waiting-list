package services

import (
	"context"
	"database/sql"
	"net/mail"
	"net/url"
	"strings"

	"github.com/moroz/discord-waiting-list/db/queries"
	"modernc.org/sqlite"
)

type SignUpParams struct {
	Email  string
	Name   string
	Bio    *string
	Region *string
}

type ValidationErrors []ValidationError

type ValidationError struct {
	Key     string
	Message string
}

func (v ValidationErrors) Error() string {
	if v == nil {
		return ""
	}
	return v[0].Message
}

func (v ValidationErrors) Get(key string) string {
	for _, err := range v {
		if err.Key == key {
			return err.Message
		}
	}
	return ""
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (p *SignUpParams) Parse(f url.Values) {
	p.Email = strings.ToLower(f.Get("email"))
	p.Name = f.Get("name")
	if b := f.Get("bio"); b != "" {
		p.Bio = &b
	}
	if r := f.Get("region"); r != "" {
		p.Region = &r
	}
}

func (p *SignUpParams) Validate() (valid bool, errors ValidationErrors) {
	valid = true
	if p.Email == "" {
		valid = false
		errors = append(errors, ValidationError{
			Key:     "email",
			Message: "can't be blank",
		})
	}
	if !validateEmail(p.Email) {
		valid = false
		errors = append(errors, ValidationError{
			Key:     "email",
			Message: "is not a valid email",
		})
	}
	if p.Name == "" {
		valid = false
		errors = append(errors, ValidationError{
			Key:     "name",
			Message: "can't be blank",
		})
	}
	if p.Region != nil && len(*p.Region) != 2 {
		valid = false
		errors = append(errors, ValidationError{
			Key:     "region",
			Message: "must be a 2-letter ISO country code",
		})
	}
	return
}

type waiterService struct {
	db *sql.DB
}

func WaiterService(db *sql.DB) *waiterService {
	return &waiterService{db}
}

func (s *waiterService) CreateWaiter(ctx context.Context, params *SignUpParams) (*queries.WaitingList, error) {
	if ok, errors := params.Validate(); !ok {
		return nil, errors
	}

	waiter, err := queries.New(s.db).InsertWaiter(ctx, queries.InsertWaiterParams(*params))

	if err, ok := err.(*sqlite.Error); ok && err.Code() == 2067 {
		return nil, ValidationErrors{{Key: "email", Message: "This e-mail is already on the list."}}
	}

	return waiter, err
}
