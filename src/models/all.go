package models

import (
	"fmt"
	photo "models/photo"
	user "models/user"
)

var Photo photo.Photo
var User user.User

func main() {
	Photo = photo.NewPhoto()
	User = user.NewUser()
}

