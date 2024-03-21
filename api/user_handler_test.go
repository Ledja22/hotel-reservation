package api

import (
	"testing"

	"github.com/Ledja22/hotel-reservation/db"
)

type testdb struct {
	db.UserStore
}

func setup(t *testing.T) *testdb {
	return nil
}

func TestPostUser(t *testing.T) {
	t.Fail()
}
