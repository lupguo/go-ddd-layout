package main

import (
	"context"

	"github.com/lupguo/go-ddd-layout/app/application"
	"github.com/lupguo/go-ddd-layout/app/domain/service"
	"github.com/lupguo/go-ddd-layout/app/infras/dbs"
	"github.com/lupguo/go-ddd-layout/app/infras/httpclient"
	"github.com/lupguo/go-ddd-layout/app/infras/openaix"
	"github.com/lupguo/go-ddd-layout/app/interfaces"
)

func main() {
	// config parse
	mysqlDSN := `gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local`

	// 基础设施
	dbInfra, _ := dbs.NewMysqlInfra(mysqlDSN)
	aiInfra := openaix.NewOpenAIInfra("socks5h://127.0.0.1:10080", "auth-token")
	httpClientInfra := httpclient.NewHttpClientInfra("socks5h://127.0.0.1:10080")

	// 服务接口实现
	uploadServiceIntf := interfaces.NewUploadImageIntf(
		application.NewImageUploadApp(
			service.NewImageUploadSrv(
				dbInfra,
				aiInfra,
				httpClientInfra,
			),
		))

	// todo :
	//  - 如果是RPC类服务，可以将 uploadServiceIntf 注入到服务接口
	//  - 如供是HTTP类型服务，可以将 uploadServiceIntf 包装注入到 Web Router中，提供HttpHandle处理能力
	_, err := uploadServiceIntf.UploadImage(context.Background())
	if err != nil {
		return
	}

	return
}
