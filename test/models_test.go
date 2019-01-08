package models

import (
	"GoModDemo/models"
	"testing"
)

func TestInitDB(t *testing.T) {
	db := models.DB
	if db == nil {
		t.Errorf("db can not be null")
	}
}
