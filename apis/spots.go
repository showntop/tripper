package apis

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
	"github.com/showntop/tripper/serializers"
	"github.com/showntop/tripper/utils"
)

func ListSpotsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var err error
	spots := make([]*models.Spot, 0)
	err = store.Spot.List(&spots)
	if err != nil {
		renderError(rw, err, http.StatusInternalServerError)
	}
	ispots := make([]interface{}, len(spots))
	for i := range spots {
		ispots[i] = spots[i]
	}
	data, err := serializers.MarshalManyPayload(ispots)
	if err != nil {
		renderError(rw, err, http.StatusInternalServerError)
	}
	render(rw, data)
}

func CreateSpotsHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var err error
	//validate params
	err = Require("title", reflect.String)
	if err != nil {
		renderError(rw, err, http.StatusBadRequest)
		return
	}
	err = Optional("description", reflect.String)
	if err != nil {
		renderError(rw, err, http.StatusBadRequest)
		return
	}
	//logic
	spot := models.NewSpot(params["title"].(string), params["description"].(string))
	err = store.Spot.Save(spot)
	if err != nil {
		renderError(rw, err, http.StatusInternalServerError)
		return
	}
	data, _ := serializers.MarshalObjectPayload(spot)
	render(rw, data)
	//render
}

func CreateOrUpdateCoversHandler(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var err error
	fileHeader := params["file"].([]*multipart.FileHeader)[0]
	file, err := fileHeader.Open()
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(rw, err, 400)
		return
	}
	newGuid := utils.NewGuid()
	filePath := "files/" + utils.GetRandomFilePath("fabnfsdlkfkl", newGuid) + "/covers"
	dir := "./public/" + filePath
	err = os.MkdirAll(dir, 0755)

	if err != nil {
		renderError(rw, err, 400)
		return
	}
	// 生成新的文件名
	filename := fileHeader.Filename
	_, ext := utils.SplitFilename(filename) // .doc
	filename = newGuid + ext
	toPath := dir + "/" + filename
	log.Println(toPath)
	err = ioutil.WriteFile(toPath, data, 0777)
	if err != nil {
		return
	}

	// add File to db
	fileType := ""
	if ext != "" {
		fileType = strings.ToLower(ext[1:])
	}
	_ = fileType
	filesize := utils.GetFilesize(toPath)
	_ = filesize
}
