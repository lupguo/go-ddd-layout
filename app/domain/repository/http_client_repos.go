package repository

import (
	"os"
)

type IHttpClientRepos interface {

	// DownloadUrl 批量下载网络图片
	DownloadUrl(webUrl []string) (map[string]*os.File, error)
}
