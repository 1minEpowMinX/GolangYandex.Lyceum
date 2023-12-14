package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// Устанавливаем обработчик запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Обработка запроса

		// Логгирование информации о запросе
		log.Printf("%s %s %s %s %v\n",
			startTime.Format("2006/01/02 15:04:05"), // Форматируем время
			r.Method,                                // HTTP метод (GET, POST, и т.д.)
			r.URL.Path,                              // Путь запроса
			r.RemoteAddr,                            // IP адрес клиента
			time.Since(startTime),                   // Время, затраченное на выполнение запроса
		)
	})

	// Запуск веб-сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
