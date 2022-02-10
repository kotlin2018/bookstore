package controller

import (
	v1 "back/api/v1"
	"back/internal/service"
	"context"
)

type cBook struct{}

var Book = cBook{}

func (b cBook) QueryBooksByCategoryId(ctx context.Context, req *v1.BookQueryReq) (res *v1.BookQueryRes, err error) {
	err, g := service.Book().QueryByCategory(ctx, req.Category)
	if err != nil {
		return nil, err
	}

	res = &v1.BookQueryRes{
		Books: g,
	}
	return
}