package dbs

import (
	"github.com/lupguo/go-ddd-layout/app/domain/entity"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInfra struct {
	db *gorm.DB
}

func NewMysqlInfra(dsn string) (*MysqlInfra, error) {
	// https://github.com/go-sql-driver/mysql
	// dsn := "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrapf(err, "db.Open(%s) got err", dsn)
	}
	return &MysqlInfra{db: db}, nil
}

func (infra *MysqlInfra) SaveImage(imgs []*entity.Image) error {
	err := infra.db.Debug().CreateInBatches(imgs, 100).Error
	if err != nil {
		return errors.Wrap(err, "db sql[SaveImage] got err")
	}

	return nil
}

func (infra *MysqlInfra) FindImages(ids []uint64) ([]*entity.Image, error) {
	var records []*entity.Image
	err := infra.db.Debug().
		Find(&records, "uuid in (?)", ids).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "db sql[FindImages] got err")
	}

	return records, nil
}
