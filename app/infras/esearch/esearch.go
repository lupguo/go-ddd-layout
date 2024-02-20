package esearch

import (
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
)

type ESearch struct {
}

func (e *ESearch) Search(req *entity.SearchKeywords) []*entity.SearchImage {
	panic("implement me")
}
