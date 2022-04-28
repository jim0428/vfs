package model

type Commandmanger interface {
	Execute_command()
	Check_command()
}

func Execute(commandmanger Commandmanger) {
	commandmanger.Check_command()
	commandmanger.Execute_command()
}
