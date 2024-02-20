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
	intf := interfaces.NewUploadImageIntf(
		application.NewImageUploadApp(
			service.NewImageUploadSrv(
				dbInfra,
				aiInfra,
				httpClientInfra,
			),
		))

	// 上传图片Handler可以作为PB Stub接口实现、HttpHandler处理
	_, err := intf.UploadImage(context.Background())
	if err != nil {
		return
	}

	return
}
