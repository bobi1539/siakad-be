package search

import (
	"net/http"
	"siakad/constant"
	"siakad/helper"
)

type PageSize struct {
	Page int
	Size int
}

func BuildPageSize(page int, size int) PageSize {
	return PageSize{
		Page: page,
		Size: size,
	}
}

func BuildPageSizeFromRequest(httpRequest *http.Request) PageSize {
	return PageSize{
		Page: helper.GetPageOrSize(httpRequest, constant.PAGE),
		Size: helper.GetPageOrSize(httpRequest, constant.SIZE),
	}
}
