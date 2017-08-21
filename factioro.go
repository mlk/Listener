package main

type Factorio struct {
	CurrentPlayers map[string]bool
	NewUserCallback func(user string)
	DeletedUserCallback func(user string)
}

func CreateFactorio(newUserCallback func(user string), deletedUserCallback func(user string)) Factorio {
	return Factorio{
		CurrentPlayers: make(map[string]bool),
		NewUserCallback: newUserCallback,
		DeletedUserCallback: deletedUserCallback,
	}
}

func (this *Factorio) AddUser(user string) {
	if !this.CurrentPlayers[user] {
		this.CurrentPlayers[user] = true
		this.NewUserCallback(user)
	}
}

func (this *Factorio) DeleteUser(user string) {
	if this.CurrentPlayers[user] {
		this.CurrentPlayers[user] = false
		this.DeletedUserCallback(user)
	}
}