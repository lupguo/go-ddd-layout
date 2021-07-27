package repository

import (
	"github.com/lupguo/ddd-layout/app/domain/entity"
)

type SearchImage struct {

}

type IReposSearch interface {
	Search(req *entity.SearchKeywords) []*SearchImage
	// SearchTags(req *entity.SearchTagReq) []*SearchImage
}