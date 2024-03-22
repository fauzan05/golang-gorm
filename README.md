Belajar gorm dengan implementasi proyek sederhana "Todo List"

Package yang harus diinstal :
- gorm.io/gorm
- gorm.io/driver/mysql
- github.com/sirupsen/logrus
- github.com/stretchr/testify
- github.com/spf13/viper
- github.com/golang-migrate/migrate
- github.com/go-playground/validator/v10

-- Migration --
Cara instal migration : go install -tags 'mysql' github.com/golang-migrate/migrate@latest

Cara membuat migration : migrate create -ext sql -dir database/migrations create_table_categories

Cara menjalankan migration : migrate -database "mysql://root@tcp(localhost:3306)/golang_restful_api" -path database/migrations up

Cara remove dirty : migrate -path database/migrations -database "mysql://root@tcp(localhost:3306)/golang_restful_api" force 20240320160949

Cara migrate ke versi tertentu (versi 1 misalnya) : migrate -database "mysql://root@tcp(localhost:3306)/golang_restful_api" -path database/migrations up 1