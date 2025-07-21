package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite" // SQLite без CGO
)

var DB *sql.DB

func InitDB() {
	var err error

	// Подключаемся к SQLite
	DB, err = sql.Open("sqlite", "file.db?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}

	// Создаем папку для миграций, если её нет
	if _, err := os.Stat("migrations"); os.IsNotExist(err) {
		err = os.Mkdir("migrations", 0755)
		if err != nil {
			log.Fatal("Не удалось создать папку migrations:", err)
		}
	}

	// Запускаем Goose миграции
	if err := goose.Up(DB, "migrations"); err != nil {
		log.Fatal("Ошибка выполнения миграций:", err)
	}
}
