package interfaces

import (
	"context"

	"github.com/lupguo/ddd-layout/app/application"
)

// SearchInf 搜索接口
type SearchInf struct {
	app *application.ImgSearchApp
}

func NewSearchInf(app *application.ImgSearchApp) *SearchInf {
	return &SearchInf{app: app}
}

// Search 搜索
func (i *SearchInf) Search(ctx context.Context, req interface{}, rsp interface{}) error  {
	// 1. 入参校验

	// 2. 业务转发
	keywords := ""
	search, err := i.app.Search(keywords)
	if err != nil {
		return err
	}
	rsp = search
	return nil
}