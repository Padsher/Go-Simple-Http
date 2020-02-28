package photo

import (
	"database/sql"
	"math/big"
)

type Photo struct {
	TableName string

	Id struct {
		Name string
		Value big.Int
	}

	FullPath struct {
		Name string
		Value string
	}

	PreviewPath struct {
		Name string
		Value string
	}

	UserId struct {
		Name string
		Value big.Int
	}
}

func NewPhoto() Photo {
	photo := Photo{}
	photo.TableName = "photos"
	photo.Id.Name = "id"
	photo.FullPath.Name = "full_path"
	photo.PreviewPath.Name = "preview_path"
	photo.UserId.Name = "user_id"
	return photo
}

func (this *Photo) Migrate(db *sql.DB) {

}