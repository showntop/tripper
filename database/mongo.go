package database

import (
	// "log"
	// "os"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "strings"
	// "github.com/showntop/sun-trip-server/app/errors"
)

// Init mgo and the common DAO

// 数据连接
var Session *mgo.Session
var Mongodb *mgo.Database

// 初始化时连接数据库
func InitMongo() {
	// Log(url)

	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	url := "mongodb://showntop:1234567890a@ds023714.mlab.com:23714/tripper"

	var err error
	Session, err = mgo.Dial(url)
	// mgo.SetDebug(true)
	// mgo.SetLogger(log.New(os.Stdout, "[INFO]:", log.Lshortfile))
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	Session.SetMode(mgo.Monotonic, true)
	Mongodb = Session.DB("tripper")
}

func close() {
	Session.Close()
}

// // common DAO
// // 公用方法

// //----------------------

// func Insert(collection *mgo.Collection, result interface{}) errors.SunError {
// 	fmt.Println(result)
// 	err := collection.Insert(result)
// 	if err != nil {
// 		return errors.NewDBError(err.Error())
// 	}
// 	return nil
// }

// //----------------------

// // 适合一条记录全部更新
// func Update(collection *mgo.Collection, query interface{}, result interface{}) bool {
// 	err := collection.Update(query, result)
// 	return Err(err)
// }

// func Upsert(collection *mgo.Collection, query interface{}, result interface{}) bool {
// 	_, err := collection.Upsert(query, result)
// 	return Err(err)
// }
// func UpdateAll(collection *mgo.Collection, query interface{}, result interface{}) bool {
// 	_, err := collection.UpdateAll(query, result)
// 	return Err(err)
// }

// //
// func UpdateByQField(collection *mgo.Collection, q interface{}, field string, value interface{}) bool {
// 	_, err := collection.UpdateAll(q, bson.M{"$set": bson.M{field: value}})
// 	return Err(err)
// }
// func UpdateByQI(collection *mgo.Collection, q interface{}, v interface{}) bool {
// 	_, err := collection.UpdateAll(q, bson.M{"$set": v})
// 	return Err(err)
// }

// // 查询条件和值
// func UpdateByQMap(collection *mgo.Collection, q interface{}, v interface{}) bool {
// 	_, err := collection.UpdateAll(q, bson.M{"$set": v})
// 	return Err(err)
// }

// //------------------------

// // 删除一条
// func Delete(collection *mgo.Collection, q interface{}) bool {
// 	err := collection.Remove(q)
// 	return Err(err)
// }
// func DeleteByIdAndUserId(collection *mgo.Collection, id, userId string) bool {
// 	err := collection.Remove(GetIdAndUserIdQ(id, userId))
// 	return Err(err)
// }
// func DeleteByIdAndUserId2(collection *mgo.Collection, id, userId bson.ObjectId) bool {
// 	err := collection.Remove(GetIdAndUserIdBsonQ(id, userId))
// 	return Err(err)
// }

// // 删除所有
// func DeleteAllByIdAndUserId(collection *mgo.Collection, id, userId string) bool {
// 	_, err := collection.RemoveAll(GetIdAndUserIdQ(id, userId))
// 	return Err(err)
// }
// func DeleteAllByIdAndUserId2(collection *mgo.Collection, id, userId bson.ObjectId) bool {
// 	_, err := collection.RemoveAll(GetIdAndUserIdBsonQ(id, userId))
// 	return Err(err)
// }

// func DeleteAll(collection *mgo.Collection, q interface{}) bool {
// 	_, err := collection.RemoveAll(q)
// 	return Err(err)
// }

// //-------------------------

// func Find(collection *mgo.Collection, id string, result interface{}) {
// 	collection.FindId(bson.ObjectIdHex(id)).One(result)
// }
// func Get2(collection *mgo.Collection, id bson.ObjectId, result interface{}) {
// 	collection.FindId(id).One(result)
// }

// func FindByQ(collection *mgo.Collection, q interface{}, result interface{}) {
// 	collection.Find(q).One(result)
// }

// func ListByQ(collection *mgo.Collection, q interface{}, result interface{}) {
// 	collection.Find(q).All(result)
// }

// func ListByQLimit(collection *mgo.Collection, q interface{}, result interface{}, limit int) {
// 	collection.Find(q).Limit(limit).All(result)
// }

// // 查询某些字段, q是查询条件, fields是字段名列表
// func GetByQWithFields(collection *mgo.Collection, q bson.M, fields []string, result interface{}) {
// 	selector := make(bson.M, len(fields))
// 	for _, field := range fields {
// 		selector[field] = true
// 	}
// 	collection.Find(q).Select(selector).One(result)
// }

// // 查询某些字段, q是查询条件, fields是字段名列表
// func ListByQWithFields(collection *mgo.Collection, q bson.M, fields []string, result interface{}) {
// 	selector := make(bson.M, len(fields))
// 	for _, field := range fields {
// 		selector[field] = true
// 	}
// 	collection.Find(q).Select(selector).All(result)
// }
// func GetByIdAndUserId(collection *mgo.Collection, id, userId string, result interface{}) {
// 	collection.Find(GetIdAndUserIdQ(id, userId)).One(result)
// }
// func GetByIdAndUserId2(collection *mgo.Collection, id, userId bson.ObjectId, result interface{}) {
// 	collection.Find(GetIdAndUserIdBsonQ(id, userId)).One(result)
// }

// // 按field去重
// func Distinct(collection *mgo.Collection, q bson.M, field string, result interface{}) {
// 	collection.Find(q).Distinct(field, result)
// }

// //----------------------

// func Count(collection *mgo.Collection, q interface{}) int {
// 	cnt, err := collection.Find(q).Count()
// 	if err != nil {
// 		Err(err)
// 	}
// 	return cnt
// }

// func Has(collection *mgo.Collection, q interface{}) bool {
// 	if Count(collection, q) > 0 {
// 		return true
// 	}
// 	return false
// }

// //-----------------

// // 得到主键和userId的复合查询条件
// func GetIdAndUserIdQ(id, userId string) bson.M {
// 	return bson.M{"_id": bson.ObjectIdHex(id), "UserId": bson.ObjectIdHex(userId)}
// }
// func GetIdAndUserIdBsonQ(id, userId bson.ObjectId) bson.M {
// 	return bson.M{"_id": id, "UserId": userId}
// }

// // DB处理错误
// func Err(err error) bool {
// 	if err != nil {
// 		fmt.Println(err)
// 		// 删除时, 查找
// 		if err.Error() == "not found" {
// 			return true
// 		}
// 		return false
// 	}
// 	return true
// }

// // 检查mognodb是否lost connection
// // 每个请求之前都要检查!!
// func CheckMongoSessionLost() {
// 	// fmt.Println("检查CheckMongoSessionLostErr")
// 	err := Session.Ping()
// 	if err != nil {
// 		// Log("Lost connection to db!")
// 		Session.Refresh()
// 		err = Session.Ping()
// 		if err == nil {
// 			// Log("Reconnect to db successful.")
// 		} else {
// 			// Log("重连失败!!!! 警告")
// 		}
// 	}
// }
