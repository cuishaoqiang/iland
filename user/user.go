package user

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
	"time"
)

type UserOrm struct {
	ID        int64  // xorm默认自动递增
	Name      string    `xorm:"name"`
	Password  string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

var ormEngine *xorm.Engine = nil
const ADMIN_ID int64 = 10000

func init(){
	orm, err := xorm.NewEngine("sqlite3", "./dbs/user.db")
	if err != nil {
		fmt.Println("orm failed to initialized: %v", err)
	}
	ormEngine = orm

	ok, _ := ormEngine.IsTableExist(UserOrm{
		ID:        10000,
	})

	if !ok {
		err = ormEngine.Sync2(new(UserOrm))
		if err != nil {
			fmt.Printf("orm failed to initialized User table: %v", err)
		}

		user := &UserOrm{
			ID: ADMIN_ID,
			Name: "admin",
			Password: "123456",
		}
		_, err = Insert(user)
		if err != nil {
			os.Exit(0)
		}
	}
}

func Insert(user *UserOrm) (int64, error) {
	return ormEngine.Insert(user)
}

func Get(user *UserOrm) (bool, error) {
	return ormEngine.Get(user)
}