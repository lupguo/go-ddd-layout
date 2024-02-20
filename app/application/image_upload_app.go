package application

import (
	"context"
	"os"

	"github.com/pkg/errors"

	"github.com/lupguo/go-ddd-layout/app/domain/entity"
	"github.com/lupguo/go-ddd-layout/app/domain/service"
)

type ImageUploadApp struct {
	srv service.IServiceImageUpload
}

func NewImageUploadApp(srv service.IServiceImageUpload) *ImageUploadApp {
	return &ImageUploadApp{srv: srv}
}

// UploadImage 图片上传
func (app *ImageUploadApp) UploadImage(ctx context.Context, updInfo *entity.UploadInfo) ([]*entity.Image, error) {
	// 1. 图片解析&上传&水印
	var uploadImgFiles []*entity.UploadFile
	var err error
	from := updInfo.UploadFrom
	switch from {
	case entity.UploadFromAI:
		// 获取prompt 配置
		uploadImgFiles, err = app.srv.UploadAIImage(ctx, "key", "macbook in the sky")
		if err != nil {
			return nil, errors.Wrap(err, "upload ai image got err")
		}
	case entity.UploadFromLocalFile:
		// 获取上传图片
		var uploads []*os.File
		uploadImgFiles, err = app.srv.UploadLocalImage(ctx, uploads)
		if err != nil {
			return nil, errors.Wrap(err, "upload local image got err")
		}
	case entity.UploadFromWebUrl:
		// 从data 获取weburl
		var webUrls []string
		uploadImgFiles, err = app.srv.UploadWebImage(ctx, webUrls)
		if err != nil {
			return nil, errors.Wrap(err, "upload web url image got err")
		}
	default:
		return nil, errors.New("error upload image type")
	}

	// 2. 存储到DB
	imgs, err := app.srv.StorageImage(ctx, 0, uploadImgFiles)
	if err != nil {
		return nil, errors.Wrap(err, "storage image got err")
	}

	return imgs, nil
}
