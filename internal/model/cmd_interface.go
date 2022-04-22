package model

type Commandmanger interface {
	Execute_command()
}

func Execute(commandmanger Commandmanger) {
	commandmanger.Execute_command()
}
