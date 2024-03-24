package tests

import (
	"fmt"
	"golang-gorm/app/enum"
	"golang-gorm/helper"
	"golang-gorm/model/domain"
	"golang-gorm/model/web/todo"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestEnum(t *testing.T) {
	if enum.ACTIVE == "active" {
		fmt.Println(enum.ACTIVE)
	}
}

func TestExecuteSQL(t *testing.T) {
	queryInsert := "INSERT INTO todos(content,status) VALUES(?,?)"
	result := DB.Exec(queryInsert, "Makan siang", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Makan malam", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Tidur", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)

	result = DB.Exec(queryInsert, "Belajar", string(enum.ACTIVE))
	helper.SaveToLogInfo(result.RowsAffected)
	assert.Nil(t, result.Error)
}

func TestRawSQL(t *testing.T) {
	var todo todo.TodoResponse
	selectQueryWhere := "SELECT * FROM todos WHERE id = ?"
	result := DB.Raw(selectQueryWhere, "1")
	result.Scan(&todo)
	fmt.Println(todo.Created_At.Format("02 Jan 2006 15:04:05"))
	fmt.Println(todo.Updated_At.Format("02 Jan 2006 15:04:05"))
	assert.Equal(t, "1", todo.ID)
	assert.Nil(t, result.Error)
}

func TestRawSQLAll(t *testing.T) {
	var todos []todo.TodoResponse
	selectQueryAll := "SELECT * FROM todos"
	result := DB.Raw(selectQueryAll)
	result.Scan(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 16, len(todos))

}

func TestRawSQLRow(t *testing.T) {
	var todos []todo.TodoResponse
	selectQueryAll := "SELECT * FROM todos"
	result := DB.Raw(selectQueryAll)
	assert.Nil(t, result.Error)
	rows, err := result.Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		var Id int64
		var Content string
		var Status enum.Status
		var Created_At time.Time
		var Updated_At time.Time

		err := rows.Scan(&Id, &Content, &Status, &Created_At, &Updated_At)
		assert.Nil(t, err)

		if Content != "Belajar" {
			todos = append(todos, todo.TodoResponse{
				ID:         Id,
				Content:    Content,
				Status:     Status,
				Created_At: Created_At,
				Updated_At: Updated_At,
			})
		}
	}
	for _, v := range todos {
		fmt.Println(v)
	}
}

func TestCreateUser(t *testing.T) {
	user := domain.User{
		ID: 1,
		Name: domain.Name{
			FirstName:  "Fauzan",
			MiddleName: "Nur",
			LastName:   "Hidayat",
		},
		Password: "fauzan123",
	}
	// fmt.Println(user)
	// fmt.Println(&user)
	response := DB.Create(&user)
	assert.Nil(t, response.Error)
	// assert.Equal(t, 1, int(response.RowsAffected))
}

func TestBatchInsert(t *testing.T) {
	var users []domain.User
	for i := 0; i < 100; i++ {
		users = append(users, domain.User{
			Name: domain.Name{
				FirstName:  "User " + strconv.Itoa(i),
				MiddleName: "User " + strconv.Itoa(i),
				LastName:   "User " + strconv.Itoa(i),
			},
			Password: "PasswordUser" + strconv.Itoa(i),
		})
	}

	result := DB.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 100, int(result.RowsAffected))
}

func TestTransaction(t *testing.T) {
	// berhasil
	err := DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&domain.User{
			Name: domain.Name{
				FirstName:  "Fauzan",
				MiddleName: "Nur",
				LastName:   "Hidayat",
			},
			Password: "fauzan12345",
		})
		helper.HandleErrorWithPanic(result.Error)
		return nil
	})
	assert.Nil(t, err)

	// gagal
	err = DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&domain.User{
			ID: 1,
			Name: domain.Name{
				FirstName:  "Fauzan",
				MiddleName: "Nur",
				LastName:   "Hidayat",
			},
			Password: "fauzan12345",
		})
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	assert.NotNil(t, err)
}

func TestManualTransaction(t *testing.T) {
	tx := DB.Begin()
	defer tx.Rollback()

	err := tx.Create(&domain.User{
		Name: domain.Name{
			FirstName:  "Fauzan",
			MiddleName: "Nur",
			LastName:   "Hidayat",
		},
		Password: "fauzan12345",
	})

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := domain.User{}
	result := DB.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(1), user.ID)
}

func TestQuerySingleInlineCondition(t *testing.T) {
	user := domain.User{}
	result := DB.First(&user, "id = ?", "5")
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(5), user.ID)
	fmt.Println(user)

	newUser := domain.User{}
	result = DB.Where("first_name = ?", "Fauzan").First(&newUser)
	assert.Nil(t, result.Error)
	fmt.Println(user)
}

func TestConflict(t *testing.T) {
	user := domain.User{
		ID: 21,
		Name: domain.Name{
			FirstName:  "Fauzan",
			MiddleName: "Nur",
			LastName:   "Hidayat",
		},
	}

	result := DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)
	assert.Nil(t, result.Error)
	// jika terjadi konflik, secara otomatis akan update user dengan id 21
}

func TestDelete(t *testing.T) {
	var user1 *domain.User
	result := DB.First(&user1, "id = ?", "90")
	fmt.Println(&domain.User{})
	assert.Nil(t, result.Error)

	result = DB.Delete(&domain.User{}, "id = ?", "91")
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := domain.Todo{
		Content: "Makan Siang",
		Status:  enum.ACTIVE,
	}
	result := DB.Create(&todo)
	assert.Nil(t, result.Error)

	result = DB.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.Deleted_At)

	var findDeletedTodo []domain.Todo
	result = DB.Find(&findDeletedTodo, "id = ?", todo.ID)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(findDeletedTodo))
}

func TestHardDelete(t *testing.T) {
	var todo domain.Todo
	result := DB.Unscoped().First(&todo, "id = ?", "1")
	assert.Nil(t, result.Error)

	// hard delete
	result = DB.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var findDeletedTodo []domain.Todo
	result = DB.Find(&findDeletedTodo, "id = ?", todo.ID)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(findDeletedTodo))
}

func TestLockUpdate(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var user domain.User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user).Error
		helper.HandleErrorWithPanic(err)

		user.Name.FirstName = "Fauzan"
		user.Name.LastName = "Nurhidayat"
		return tx.Save(&user).Error
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := domain.Wallet{
		UserId:  1,
		Balance: 10000,
	}
	err := DB.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user domain.User
	err := DB.Model(&domain.User{}).Preload("Wallet").First(&user, "id = ?", 1).Error
	assert.Nil(t, err)

	assert.Equal(t, uint64(1), user.ID)
	assert.Equal(t, uint64(1), user.Wallet.UserId)
	/*
		Menggunakan Preload akan menghasilkan dua query :
		SELECT * FROM `wallets` WHERE `wallets`.`user_id` = 1
		SELECT * FROM `users` WHERE id = 1 ORDER BY `users`.`id` LIMIT 1
	*/
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user domain.User
	err := DB.Model(&domain.User{}).Joins("Wallet").First(&user, "users.id = ?", 1).Error
	assert.Nil(t, err)

	assert.Equal(t, uint64(1), user.ID)
	assert.Equal(t, uint64(1), user.Wallet.UserId)

	/*
		Menggunakan Joins akan menghasilkan 1 query saja :
		SELECT `users`.`id`,`users`.`first_name`,`users`.`middle_name`,`users`.`last_name`,`users`.`password`,`users`.`created_at`,`users`.`updated_at`,`Wallet`.`id` AS `Wallet__id`,`Wallet`.`user_id` AS `Wallet__user_id`,`Wallet`.`balance` AS `Wallet__balance`,`Wallet`.`created_at` AS `Wallet__created_at`,`Wallet`.`updated_at` AS `Wallet__updated_at` FROM `users` LEFT JOIN `wallets` `Wallet` ON `users`.`id` = `Wallet`.`user_id` WHERE users.id = 1 ORDER BY `users`.`id` LIMIT 1
	*/
}

func TestAutoCreateUpdate(t *testing.T) {
	user := domain.User{
		ID: 101,
		Name: domain.Name{
			FirstName:  "Fauzan",
			MiddleName: "Nur",
			LastName:   "Hidayat",
		},
		Password: "fauzan123",
		Wallet: domain.Wallet{
			UserId:  101,
			Balance: 1000,
		},
	}
	// menggunakan method omit tidak akan mengupdate/create relasinya
	err := DB.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
	// INSERT INTO `users` (`first_name`,`middle_name`,`last_name`,`password`,`created_at`,`updated_at`,`id`) VALUES ('Fauzan','Nur','Hidayat','fauzan123','2024-03-24 07:54:21.172','2024-03-24 07:54:21.172',101
}

func TestUserAndTodos(t *testing.T) {
	// one to many
	user := domain.User{
		ID: 1,
		Name: domain.Name{
			FirstName:  "Fauzan",
			MiddleName: "Nur",
			LastName:   "Hidayat",
		},
		Password: "Fauzan123",
		Todos: []domain.Todo{
			{
				UserId:  1,
				Content: "Makan Siang",
				Status:  enum.ACTIVE,
			},
			{
				UserId:  1,
				Content: "Makan Pagi",
				Status:  enum.ACTIVE,
			},
			{
				UserId:  1,
				Content: "Makan Malam",
				Status:  enum.ACTIVE,
			},
		},
	}

	err := DB.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	// menampilkan semua data user beserta semua relasinya
	var usersResources []domain.User
	err := DB.Model(&domain.User{}).Preload("Todos").Joins("Wallet").Find(&usersResources).Error
	assert.Nil(t, err)

	for _, v := range usersResources {
		fmt.Println(v.Name.FirstName, v.Todos, v.Wallet)
	}
}

func TestBelongsToManyToOne(t *testing.T) {
	var todos []domain.Todo
	err := DB.Preload("User").First(&todos).Error
	assert.Nil(t, err)

	for _, v := range todos {
		fmt.Println(v)
	}
	// bisa juga menggunakan join
}

func TestBelongsToOneToOne(t *testing.T) {
	var wallets []domain.Wallet
	err := DB.Preload("User").First(&wallets).Error
	assert.Nil(t, err)

	for _, v := range wallets {
		fmt.Println(v.User)
	}
}

func TestCreateManyToMany(t *testing.T) {
	product := domain.Product{
		ID:    1,
		Name:  "Pecel",
		Price: 2000,
	}
	err := DB.Create(&product).Error
	assert.Nil(t, err)

	err = DB.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    1,
		"product_id": 1,
	}).Error
	assert.Nil(t, err)

	err = DB.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    2,
		"product_id": 1,
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToManyFromProduct(t *testing.T) {
	var products []domain.Product
	err := DB.Preload("LikedByUsers").Find(&products, "id = ?", 1).Error
	assert.Nil(t, err)
	for _, v1 := range products {
		for _, v2 := range v1.LikedByUsers {
			fmt.Println(v2)
		}
	}
}
func TestPreloadManyToManyFromUser(t *testing.T) {
	var users []domain.User
	err := DB.Preload("LikeProducts").Find(&users, "id = ?", 1).Error
	assert.Nil(t, err)
	for _, v1 := range users {
		for _, v2 := range v1.LikeProducts {
			fmt.Println(v2)
		}
	}
}

func TestManyToManyAssociation(t *testing.T) {
	var product domain.Product
	err := DB.First(&product, "id = ?", 1).Error
	assert.Nil(t, err)

	// mencari tahu siapa saja yang like produk dengan id = 1
	var users []domain.User
	err = DB.Model(&product).Where("first_name LIKE ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	for _, v := range users {
		fmt.Println(v)
	}
}

func TestAssociationAdd(t *testing.T) {
	var user domain.User
	err := DB.First(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product domain.Product
	err = DB.First(&product, "id = ?", "1").Error
	assert.Nil(t, err)

	// memasukkan data product dan user ke tabel user_like_product secara otomatis
	err = DB.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var user domain.User
		err := tx.First(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := domain.Wallet{
			ID:      2,
			UserId:  user.ID,
			Balance: 20000,
		}
		// jika walletnya tidak ada, maka akan dibuat. jika ada maka akan dhapus lalu dibuat ulang dengan user_id sesuai dengan &user
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		// UPDATE `wallets` SET `user_id`=NULL WHERE `wallets`.`id` <> 2 AND `wallets`.`user_id` = 1
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	var user domain.User
	err := DB.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	var product domain.Product
	err = DB.Take(&product, "id = ?", "1").Error
	assert.Nil(t, err)

	// menghapus data relasi hanya yang terkait dengan produk, misalnya user dengan id 1
	err = DB.Model(&product).Association("LikedByUsers").Delete(&user)
	// DELETE FROM `user_like_product` WHERE `user_like_product`.`product_id` = 1 AND `user_like_product`.`user_id` = 1
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product domain.Product
	err := DB.First(&product, "id = ?", "1").Error
	assert.Nil(t, err)

	// akan menghapus semua data relasi ke product, user apapun itu
	err = DB.Model(&product).Association("LikedByUsers").Clear()
	// DELETE FROM `user_like_product` WHERE `user_like_product`.`product_id` = 1
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user domain.User
	err := DB.Preload("Wallet.User", "first_name LIKE ?", "User%").Preload("Todos").First(&user, "id = ?", 1).Error
	assert.Nil(t, err)
}

func TestNestedPreloading(t *testing.T) {
	var wallet domain.Wallet
	err := DB.Preload("User.Todos").Find(&wallet, "id = ?", 1).Error
	assert.Nil(t, err)

	fmt.Println(wallet.User)
}

func TestJoinQuery(t *testing.T) {
	var users []domain.User
	err := DB.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))

	// atau
	users = []domain.User{}
	err = DB.Joins("Wallet").Find(&users).Error // left join defaultnya
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestGroupByHaving(t *testing.T) {
	var result []AggregationResult
	// mencari semua user yang memiliki wallet dengan saldo diatas 10000
	err := DB.Model(&domain.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 10000).
		Find(&result).Error
	assert.Nil(t, err)
	fmt.Println(result)
}
