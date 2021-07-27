package myapp

import (
	"context"

	"github.com/lupguo/ddd-layout/app/application"
	"github.com/lupguo/ddd-layout/app/interfaces"
)

func main() {

	// 依赖注入
	s := interfaces.NewSearchInf(&application.ImgSearchApp{})


	err := s.Search(context.Background(), "keywords", "")
	if err != nil {
		return
	}
}