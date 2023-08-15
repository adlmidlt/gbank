Команды докер:

1. docker compose run --service-ports web bash -> команда для перехода в контейнер.

Команды golang:

1. root@bdf806b03b6b:/usr/src/app# go mod init github.com/adlmidlt/GBank -> внутри контейнера создаём и инициализируем директорию.
2. go get github.com/gofiber/fiber/v2 -> загружаем бибилиотеку для работы с rest api для проверки нашего приложения.
3. touch cmd/main.go -> создаем главный файл, и с сайта fiber.io копируем код с "Hello World".
4. 