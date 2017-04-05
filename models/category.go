package models

type Category struct {
	Id        int `bson:"_id" json:"id"`
	Name string `json:"name"`
}


func GetCategories() ([]*Category, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Category = make([]*Category, 0)
	err := sess.DB(DBNAME).C(C_CATEGORY_NAME).Find(nil).All(&result)

	return result, err
}
