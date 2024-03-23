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

Cara membuat migration : 
- Prod
migrate create -ext sql -dir database/prod/migrations create_table_users
migrate create -ext sql -dir database/prod/migrations create_table_todos
- Test
migrate create -ext sql -dir database/test/migrations create_table_users
migrate create -ext sql -dir database/test/migrations create_table_todos

Cara menjalankan migration :
- Prod
migrate -database "mysql://root@tcp(localhost:3306)/golang_gorm" -path database/prod/migrations up 
- Test
 migrate -database "mysql://root@tcp(localhost:3306)/golang_gorm_test" -path database/test/migrations up

Cara remove dirty : 
- Prod
migrate -path database/prod/migrations -database "mysql://root@tcp(localhost:3306)/golang_gorm" force V
-Test
migrate -path database/test/migrations -database "mysql://root@tcp(localhost:3306)/golang_gorm_test" force V

Cara migrate ke versi tertentu (versi 1 misalnya) : 
- Prod
migrate -database "mysql://root@tcp(localhost:3306)/golang_gorm" -path database/prod/migrations up 1
- Test
migrate -database "mysql://root@tcp(localhost:3306)/golang_gorm_test" -path database/test/migrations up 1

-- GORM --
Field Permission
<-:create untuk create only
<-:update untuk update only
<- untuk create dan update
- ignore/tidak bisa di read/write