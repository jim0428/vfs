package command

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	DB "vfs/internal/entity"
// )

// var cfCkCmdTests = []struct {
// 	length      int
// 	username    string
// 	foldername  string
// 	description string
// 	check       bool
// }{
// 	{-1, "", "TEST", "TEST", false},
// 	{2, "TEST", "", "TEST", false},
// 	{4, "TEST", "TEST", "", false},
// 	{3, "Luke", "Folder1", "descp1", false},
// 	{3, "Jim", "Folder2", "descp2", true},
// 	{3, "WrongUser", "Folder2", "descp2", false},
// }

// var cfECCmdTests = []struct {
// 	username    string
// 	foldername  string
// 	description string
// 	check       bool
// }{
// 	{"Luke", "Folder1", "descp1", true},
// 	{"Jim", "Folder2", "descp2", true},
// }

// func TestCreateFolder_checkCommand(t *testing.T) {

// 	db := DB.NewUserDB()
// 	//add user
// 	db.AddUser("Luke")
// 	db.AddFolder(NewCreateFolder("Luke", "Folder1", "descp1").fData)
// 	db.AddUser("Jim")

// 	for _, tt := range cfCkCmdTests {
// 		t.Run("CreateFolder_checkCommand", func(t *testing.T) {
// 			ctfolder := NewCreateFolder(tt.username, tt.foldername, tt.description)
// 			result, _ := ctfolder.Check_command(db, tt.length)
// 			assert.Equal(t, result, tt.check)
// 		})
// 	}
// }

// func TestCreateFolder_executeCommand(t *testing.T) {
// 	db := DB.NewUserDB()

// 	for _, tt := range cfECCmdTests {
// 		t.Run("CreateFolder_executeCommand", func(t *testing.T) {
// 			ctfolder := NewCreateFolder(tt.username, tt.foldername, tt.description)
// 			result := ctfolder.Execute_command(db)
// 			assert.Equal(t, result, tt.check)
// 		})
// 	}
// }
