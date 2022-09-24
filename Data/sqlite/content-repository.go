package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/arahkya/antifake-news/models"

	_ "github.com/mattn/go-sqlite3"
)

var connectionString = "/usr/src/app/Data/sqlite/antifake-news.db"

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		panic("Database open error.")
	}
	defer db.Close()

	return db
}

func DeleteContent(db *sql.DB, contentId *int) {
	statement, statementErr := db.Prepare(`delete from contents where Id = ?`)
	if statementErr != nil {
		panic(statementErr.Error())
	}
	defer statement.Close()

	_, execErr := statement.Exec(contentId)

	if execErr != nil {
		panic("Exec statement error")
	}

	fmt.Println("Delete Content")
}

func CreateContent(db *sql.DB, content *models.ContentModel) {
	statement, statementErr := db.Prepare(`insert into contents(TitleTh, TitleEn, Detail, Author, Organize) values(?,?,?,?,?)`)
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

func UpdateContent(db *sql.DB, content *models.ContentModel) {
	statement, statementErr := db.Prepare(`update contents set TitleTh = ?, TitleEn = ?, Detail = ?, Author = ?, Organize = ? where Id = ?`)
	if statementErr != nil {
		panic(statementErr.Error())
	}
	defer statement.Close()

	_, execErr := statement.Exec(content.TitleTh, content.TitleEn, content.Detail, content.Author, content.Organize, content.Id)

	if execErr != nil {
		panic("Exec statement error")
	}

	fmt.Println("Update Content")
}

func ListContent(db *sql.DB) []models.ContentModel {
	contents := []models.ContentModel{}

	statement, statementErr := db.Prepare(`select Id, TitleEn, TitleTh, Detail, Author, Organize from contents`)
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

func GetContent(db *sql.DB, contentId *int) models.ContentModel {
	contents := []models.ContentModel{}

	statement, statementErr := db.Prepare(`select Id, TitleEn, TitleTh, Detail, Author, Organize from contents where Id = ?`)
	if statementErr != nil {
		panic(statementErr.Error())
	}
	defer statement.Close()

	rows, queryErr := statement.Query(&contentId)
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

	return contents[0]
}
