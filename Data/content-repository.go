package data

import (
	"database/sql"
	"fmt"

	"github.com/arahkya/antifake-news/models"
	_ "github.com/go-sql-driver/mysql"
)

func CreateContent(content *models.ContentModel) {
	db, err := sql.Open("mysql", "root:vkiydKN8986@tcp(127.0.0.1:3306)/AntifakeNews")
	if err != nil {
		panic("Database open error.")
	}
	defer db.Close()

	statement, statementErr := db.Prepare(`insert into Contents(TitleTh, TitleEn, Detail, Author, Organize) values(?,?,?,?,?)`)
	if statementErr != nil {
		panic(statementErr.Error())
	}

	defer statement.Close()

	_, execErr := statement.Exec(content.TitleTh, content.TitleEn, content.Detail, content.Author, content.Organize)

	if execErr != nil {
		panic("Exec statement error")
	}

	fmt.Println("Save Content")
}
