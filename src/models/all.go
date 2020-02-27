package models

import (
	"fmt"
	photo "models/photo"
	user "models/user"
)

var BasePhoto *photo.Photo
var BaseUser *user.User

func main() {
	BasePhoto = photo.NewPhoto()
	BaseUser = user.NewUser()

	BasePhoto.JoinUserCond = fmt.Sprintf(
		"%v = %v.%v",
		BasePhoto.UserId.Name,
		BaseUser.TableName,
		BaseUser.Id.Name,
	)
	BaseUser.JoinPhotoCond = fmt.Sprintf(
		"%v = %v.%v",
		BaseUser.Id.Name,
		BasePhoto.TableName,
		BasePhoto.UserId.Name,
	)
}

