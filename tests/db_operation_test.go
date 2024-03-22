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
				ID: Id,
				Content: Content,
				Status: Status,
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
			FirstName: "Fauzan",
			MiddleName: "Nur",
			LastName: "Hidayat",
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
				FirstName: "User " + strconv.Itoa(i),
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
				FirstName: "Fauzan",
				MiddleName: "Nur",
				LastName: "Hidayat",
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
				FirstName: "Fauzan",
				MiddleName: "Nur",
				LastName: "Hidayat",
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
			FirstName: "Fauzan",
			MiddleName: "Nur",
			LastName: "Hidayat",
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