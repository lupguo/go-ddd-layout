package application

import (
	"github.com/lupguo/ddd-layout/app/domain/entity"
)

// ISearchApp 搜索接口
type ISearchApp interface {
	Search(keywords string) (*entity.SearchRsp, error)
}

// ImgSearchApp 图片搜索结构体
type ImgSearchApp struct {
	
}

// Search 接口实现
func (i *ImgSearchApp) Search(keywords string) (*entity.SearchRsp, error) {
	// todo 业务聚合(rpc或者domain service聚合)
	panic("implement me")
}
