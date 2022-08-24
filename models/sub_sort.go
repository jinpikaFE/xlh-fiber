package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShopSubSort struct {
	Model

	Name       *string `validate:"" gorm:"unique;comment:'商品子类名称'" query:"name" json:"name" xml:"name" form:"name"`
	ShopSortID uint
}

func GetSubSorts(pageNum int, pageSize int, maps interface{}) ([]*ShopSubSort, error) {
	var subSorts []*ShopSubSort
	err := Db.Preload(clause.Associations).Where(maps).Offset(pageNum).Limit(pageSize).Find(&subSorts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return subSorts, nil
}
