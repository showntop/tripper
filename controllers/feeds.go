package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/showntop/tripper/models"
)

type Feeds struct {
	application
}

func (c *Feeds) List(req *http.Request, ps httprouter.Params) ([]byte, *HttpError) {
	pageNo, _ := strconv.Atoi(req.URL.Query().Get("page_no"))
	pageNum, _ := strconv.Atoi(req.URL.Query().Get("page_num"))
	if pageNo == 0 {
		pageNo = 1
	}
	if pageNum == 0 {
		pageNum = 10
	}
	data, err := models.ListFeeds(pageNo, pageNum)
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}
