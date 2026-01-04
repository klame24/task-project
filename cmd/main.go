package main

import (
	"context"
	"fmt"
	"task-project/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Cant load env file")
	}

	ctx := context.Background()

	// инициализиция БД
	conn, err := database.ConnectDB(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect to DB is successfully")
	fmt.Println(conn)
	// инициализация репозитория


	// инициализация сервиса, в который мы передаем репозиторий и сервис управляет какие данные мы выдаем пользователю

	// инициализация маршрутов

	// запуск сервера

}
