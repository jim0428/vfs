package model

type Commandmanger interface {
	Execute_command()
	Check_command(int) string
}

//func Check_command(length int) string { return "" }
