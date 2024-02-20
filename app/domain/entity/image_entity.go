package entity

// Style 图片风格
type Style struct {
	ID        string
	StyleName string // 风格名称，国风、古装..
	StyleDesc string // 风格描述
}

// Tag 图片标签
type Tag struct {
	TagID   uint64
	TagName string
}

// Image 图片文件
type Image struct {
	ID        uint64 `gorm:"id" json:"id,omitempty"`
	CreatedAt string `gorm:"created_at"`
	UpdatedAt string `gorm:"updated_at"`
	DeletedAt string `gorm:"deleted_at"`

	Name  string `gorm:"name" json:"name,omitempty"`   // 图片名称
	UUID  string `gorm:"uuid" json:"uuid,omitempty"`   // 图片UUID
	Style *Style `gorm:"style" json:"style,omitempty"` // 图片风格
	Tag   *Tag   `gorm:"tag" json:"tag,omitempty"`     // 图片标签

	SaveMethod      StorageMethod `gorm:"save_method" json:"save_method,omitempty"`             // 原图存储方式
	OriginImagePath string        `gorm:"origin_image_path" json:"origin_image_path,omitempty"` // 原图路径
	WaterImagePath  string        `gorm:"water_image_path" json:"water_image_path,omitempty"`   // 水印图路径
}
