package login

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/kataras/iris/v12"
)

var boltDB *bolt.DB
const BUCKET_NAME string = "./dbs/tokens"
func init(){
	db, err := bolt.Open(BUCKET_NAME, 0600, nil)
	if err != nil {
		fmt.Println("## tokens.db open error, ", err)
	}
	boltDB = db

}

func CloseBoltDB(){
	boltDB.Close()
}

func appendTokens(token string, username string) error {
	if err := boltDB.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME)); err != nil {
			fmt.Println("## tokens bucket create failed", err)
			return err
		}

		b := tx.Bucket([]byte(BUCKET_NAME))
		err := b.Put([]byte(token), []byte(username))
		return err
	}); err != nil {
		fmt.Println("## token append failed, ", err)
	}
	return nil
}

func CheckInTokens(token string) bool {
	if err := boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BUCKET_NAME))
		v := b.Get([]byte(token))
		fmt.Printf("## the value is :%s\n", v)

		return nil
	}); err != nil {
		fmt.Println("## view error, ", err)
		return false
	}
	return true
}

// TODO
func Logout(ctx iris.Context){
	// 获取token信息
	// 将token 加入到黑名单
	// 返回正确响应
}

