package main

import (
	"auth/data"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	repository := data.NewPostgresTestRepository(nil)
	testApp.Repository = repository
	os.Exit(m.Run())
}