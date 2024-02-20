package application

import (
	"context"

	"github.com/lupguo/go-ddd-layout/app/domain/entity"
)

// ImgSearchApp 图片搜索结构体
type ImgSearchApp struct {
}

// Search 通过关键字图片搜索
func (i *ImgSearchApp) Search(keywords string) (*entity.SearchRsp, error) {
	panic("implement me")
}

// CronIndexImage 定时索引生成图片
func (app *ImageUploadApp) CronIndexImage(ctx context.Context) error {
	return nil
}
