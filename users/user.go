package users

import (
	"fmt"
	"time"
	"uuid"

	"github.com/nicos92/gouno/modelos"
)

func AltaUsuario() {
	user := new(modelos.User)

	user.AddUser(1, uuid.New(), "Nicolás", time.Now(), true)

	fmt.Printf("\n%v", user)
}
