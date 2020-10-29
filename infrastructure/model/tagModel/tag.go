package tagModel

import (
	"database/sql"
	"log"

	"ToDoList/infrastructure/db"
)

type Tag struct {
	ID      int64
	TagData string
}

// SelectTag ユーザを取得
func SelectTag(tagID int) (error, *Tag) {
	row := db.Conn.QueryRow("SELECT id, tag_data FROM tag WHERE id = ?", tagID)

	tag := Tag{}
	err := row.Scan(&tag.ID, &tag.TagData)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return err, nil
	}
	return nil, &tag
}
