package author

import (
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	AuthorName string `gorm:"type:varchar(100);column:AuthorName"`
}

func (Author) TableName() string {
	return "Author"
}

func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, CreatedAt : %s", a.ID, a.AuthorName, a.CreatedAt.Format("2006-01-02 15:04:05"))
}
