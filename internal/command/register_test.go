package command

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	DB "vfs/internal/entity"
// )

// var ckcmdTests = []struct {
// 	username string
// 	length   int
// 	check    bool
// }{
// 	{"Luke", -1, false},
// 	{"Jim", 0, false},
// 	{"Ares", 1, true},
// 	{"Michael", 2, false},
// 	// //方法1，接近user的input
// 	// {"register Jim",true},
// 	// {"register Jim abc",false},
// 	// {{"Jim","Ares"},"register Jim abc",false},
// 	// //方法2，參數本身的檢查

// }

// var eccmdTests = []struct {
// 	username string
// 	check    bool
// }{
// 	{"Jim", false},
// 	{"Ares", true},
// }

// func TestRegister_checkCommand(t *testing.T) {

// 	db := DB.NewUserDB()

// 	for _, tt := range ckcmdTests {
// 		t.Run("checkCommand", func(t *testing.T) {
// 			reg := NewRegister(tt.username)
// 			result, _ := reg.Check_command(db, tt.length)
// 			assert.Equal(t, result, tt.check)
// 		})
// 	}
// }

// func TestRegister_executeCommand(t *testing.T) {
// 	db := DB.NewUserDB()
// 	db.AddUser("Jim")

// 	for _, tt := range eccmdTests {
// 		t.Run("executeCommand", func(t *testing.T) {
// 			reg := NewRegister(tt.username)
// 			result := reg.Execute_command(db)
// 			assert.Equal(t, result, tt.check)
// 		})
// 	}
// }
