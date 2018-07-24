package user

import (
	"fmt"
	"net/url"

	"github.com/alyssonggs/golang-fast-startup/role"
	"github.com/alyssonggs/golang-fast-startup/settings"
)

type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Username string      `json:"username"`
	Lastname string      `json:"lastName"`
	Active   int         `json:"active"`
	Roles    []role.Role `json:"roles"`
}

func connection() string {
	return fmt.Sprintf(
		"mysql://%s:%s@%s:%s/%s?sslmode=disable",
		settings.Settings.Database.UserName,
		url.QueryEscape(settings.Settings.Database.Password),
		settings.Settings.Database.Host,
		settings.Settings.Database.Port,
		settings.Settings.Database.DatabaseName,
	)
}
