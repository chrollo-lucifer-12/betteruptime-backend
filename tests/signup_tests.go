package tests

import (
	"testing"

	"github.com/chrollo-lucifer-12/betteruptime/db"
	"github.com/chrollo-lucifer-12/betteruptime/server"
	"github.com/go-playground/assert/v2"
)

func TestSignup(t *testing.T) {
	testDB, _ := db.NewTestDB()

	s := server.NewServer(server.ServerOpts{
		DB: testDB,
	})

	reqBody := `{"username":"test","password":"pass123"}`
	req, w := JSONRequest("POST", "/signup", reqBody)

	s.ServerHttp(w, req)

	assert.Equal(t, 201, w.Code)
}
