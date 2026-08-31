package modelos

import (
	"time"
	"uuid"
)

type User struct {
	Id        int
	IdPublic  uuid.UUID
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (this *User) AddUser(id int, idPublic uuid.UUID, name string, createdAt time.Time, status bool) {
	this.Id = id
	this.IdPublic = idPublic
	this.Name = name
	this.CreatedAt = createdAt
	this.Status = status
}
