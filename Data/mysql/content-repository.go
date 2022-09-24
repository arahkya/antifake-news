package mysql

import (
	"database/sql"
	"fmt"

	"github.com/arahkya/antifake-news/models"
	_ "github.com/go-sql-driver/mysql"
)

var connectionString = "root:vkiydKN123456@tcp(antifake-news-db:3306)/AntifakeNews"

func CreateContent(content *models.ContentModel) {
	db, err := sql.Open("mysql", connectionString)
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

func ListContent() []models.ContentModel {
	contents := []models.ContentModel{}

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("Database open error.")
	}
	defer db.Close()

	statement, statementErr := db.Prepare(`select Id, TitleEn, TitleTh, Detail, Author, Organize from Contents`)
	if statementErr != nil {
		panic(statementErr.Error())
	}
	defer statement.Close()

	rows, queryErr := statement.Query()
	if queryErr != nil {
		panic(queryErr.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var content models.ContentModel

		if scanError := rows.Scan(&content.Id, &content.TitleEn, &content.TitleTh, &content.Detail, &content.Author, &content.Organize); scanError != nil {
			panic(scanError.Error())
		}

		contents = append(contents, content)
	}

	return contents
}
