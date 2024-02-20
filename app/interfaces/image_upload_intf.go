package interfaces

import (
	"context"
	"encoding/json"

	"github.com/lupguo/go-ddd-layout/app/application"
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
)

type UploadImageIntf struct {
	app *application.ImageUploadApp
}

func NewUploadImageIntf(app *application.ImageUploadApp) *UploadImageIntf {
	return &UploadImageIntf{app: app}
}

func (i *UploadImageIntf) UploadImage(req interface{}) (rsp interface{}, err error) {
	ctx := context.Background()

	// fill data from info
	images, err := i.app.UploadImage(ctx, &entity.UploadInfo{
		FileName:   "",
		WebUrl:     "",
		UploadFrom: 0,
		Method:     0,
		Style:      nil,
		Tag:        nil,
	})
	if err != nil {
		return nil, err
	}

	// json rsp
	data, _ := json.Marshal(images)
	_ = data

	// rsp.data = data
	return rsp, nil
}
