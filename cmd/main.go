package main

import (
	"context"
	"fmt"
	"net/http"
	"task-project/internal/database"
	"task-project/internal/handlers"
	"task-project/internal/repositories"
	"task-project/internal/routes"
	"task-project/internal/services"

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
	taskR := repositories.NewTaskRepository(conn)

	// инициализация сервиса, в который мы передаем репозиторий и сервис управляет какие данные мы выдаем пользователю
	taskS := services.NewTaskService(taskR)
	title := "test 2"
	descr := "test 2 descr"

	_, err = taskS.CreateTask(ctx, title, descr)
	if err != nil {
		fmt.Println("Не получилось создать таску из сервиса")
	}
	fmt.Println("Получилось создать таску из сервиса!")
	// инициализация маршрутов
	taskH := handlers.NewTaskHandler(taskS)
	// запуск сервера
	router := routes.SetupRoutes(taskH)

	err = http.ListenAndServe(":5050", router)
	if err != nil {
		panic(err)
	}
}
