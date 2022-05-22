package command

import (
	"testing"

	"github.com/stretchr/testify/assert"

	DB "vfs/internal/entity"
)

var ckcmdTests = []struct {
	username string
	length   int
	check    bool
}{
	{"Luke", -1, false},
	{"Jim", 0, false},
	{"Ares", 1, true},
	{"Michael", 2, false},
}

var eccmdTests = []struct {
	username string
	check    bool
}{
	{"Jim", false},
	{"Ares", true},
}

func TestRegister_checkCommand(t *testing.T) {

	db := DB.NewUserDB()

	for _, tt := range ckcmdTests {
		t.Run("checkCommand", func(t *testing.T) {
			reg := NewRegister(tt.username)
			result, _ := reg.Check_command(db, tt.length)
			assert.Equal(t, result, tt.check)
		})
	}
}

func TestRegister_executeCommand(t *testing.T) {
	db := DB.NewUserDB()
	db.AddUser("Jim")

	for _, tt := range eccmdTests {
		t.Run("executeCommand", func(t *testing.T) {
			reg := NewRegister(tt.username)
			result := reg.Execute_command(db)
			assert.Equal(t, result, tt.check)
		})
	}
}
