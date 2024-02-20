package service

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
	"github.com/lupguo/go-ddd-layout/app/domain/repository"
)

// IServiceImageUpload 图片上传服务接口
type IServiceImageUpload interface {
	// UploadLocalImage 本地图片上传&保存
	UploadLocalImage(ctx context.Context, uploads []*os.File) ([]*entity.UploadFile, error)

	// UploadAIImage 按域定义的prompt key生成AI图片上传&保存
	UploadAIImage(ctx context.Context, promptKey string, keywords string) ([]*entity.UploadFile, error)

	// UploadWebImage 上传了一张web url地址图片，需要下载webUrl图片，保存图片
	UploadWebImage(ctx context.Context, webImgUrl []string) ([]*entity.UploadFile, error)

	// StorageImage 将上传处理后的文件，附加标签等信息后，存储到DB中
	StorageImage(ctx context.Context, method entity.StorageMethod, imgs []*entity.UploadFile) ([]*entity.Image, error)

	// GetImages 通过图片的uuid获取指定上传的文件信息
	GetImages(uuids []uint64) ([]*entity.Image, error)
}

// ImageUploadSrv 图片上传服务
type ImageUploadSrv struct {
	dbRepos         repository.IReposStorage    // 依赖仓储存储接口，保存图片信息
	aiRepos         repository.IReposOpenAI     // 依赖ai图片仓储，生成AI图片
	httpClientRepos repository.IHttpClientRepos // 依赖http客户的，请求URL网络图片
}

func NewImageUploadSrv(
	dbRepos repository.IReposStorage,
	aiRepos repository.IReposOpenAI,
	httpClientRepos repository.IHttpClientRepos,
) *ImageUploadSrv {
	return &ImageUploadSrv{
		dbRepos:         dbRepos,
		aiRepos:         aiRepos,
		httpClientRepos: httpClientRepos,
	}
}

func (srv *ImageUploadSrv) UploadAIImage(ctx context.Context, promptKey string, keywords string) ([]*entity.UploadFile, error) {
	// TODO implement me
	panic("implement me")
}

func (srv *ImageUploadSrv) UploadWebImage(ctx context.Context, webImgUrl []string) ([]*entity.UploadFile, error) {
	// TODO implement me
	panic("implement me")
}

// UploadLocalImage UploadImage 文件上传，并生成水印，存储到指定位置中
func (srv *ImageUploadSrv) UploadLocalImage(ctx context.Context, uploads []*os.File) ([]*entity.UploadFile, error) {
	// TODO implement me
	panic("implement me")
}

func (srv *ImageUploadSrv) StorageImage(ctx context.Context, method entity.StorageMethod, uploads []*entity.UploadFile) ([]*entity.Image, error) {
	var imgs []*entity.Image

	for _, f := range uploads {
		imgs = append(imgs, &entity.Image{
			Name:            f.UploadInfo.FileName,
			UUID:            uuid.NewString(),
			Style:           f.UploadInfo.Style,
			Tag:             f.UploadInfo.Tag,
			SaveMethod:      entity.StorageToLocalFile,
			OriginImagePath: f.OriginImagePath,
			WaterImagePath:  f.WaterImagePath,
		})
	}

	// 基于仓储保存到DB中
	err := srv.dbRepos.SaveImage(imgs)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (srv *ImageUploadSrv) GetImages(uuids []uint64) ([]*entity.Image, error) {
	// TODO implement me
	panic("implement me")
}
