package esearch

import (
	"github.com/lupguo/ddd-layout/app/domain/entity"
	"github.com/lupguo/ddd-layout/app/domain/repository"
)

type ESearch struct {

}

func (e *ESearch) Search(req *entity.SearchKeywords) []*repository.SearchImage {
	panic("implement me")
}
