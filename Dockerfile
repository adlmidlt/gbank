# Создаем собственный контейнер.
FROM golang:latest

# Контейнеры докера работают под линукс, используем соглашение о файловой системы линукс.
WORKDIR /usr/src/app

# Установка пакет для горячей перезагрузки браузера, что бы увидить изменения приложения.
RUN go install github.com/cosmtrek/air@latest

# Скопируем все файлы с нашего хоста.
COPY . .

# Подгрузит все бибилиотеки, которые есть в проекте.
RUN go mod tidy && go build -o gbank ./cmd/app/main.go

CMD ["./gbank"]