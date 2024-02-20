package repository

type IReposOpenAI interface {

	// GenerateAIImage 基于域定义的prompt，以及关键字生成图片
	GenerateAIImage(promptKey string, keywords string) (imgUrls []string, err error)
}
