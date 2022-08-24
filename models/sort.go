package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ShopSort struct {
	Model

	Name *string `validate:"" gorm:"unique;comment:'商品分类名称'" query:"name" json:"name" xml:"name" form:"name"`

	ShopSubSort []ShopSubSort
}

func GetSortTotal(maps interface{}) (int64, error) {
	var count int64
	if err := Db.Model(&ShopSort{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetSorts(pageNum int, pageSize int, maps interface{}) ([]*ShopSort, error) {
	var sorts []*ShopSort
	err := Db.Preload(clause.Associations).Where(maps).Offset(pageNum).Limit(pageSize).Find(&sorts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return sorts, nil
}

func GetSort(maps interface{}) (*ShopSort, error) {
	var shop_sort ShopSort
	err := Db.Preload(clause.Associations).Where(maps).Find(&shop_sort).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &shop_sort, nil
}

func AddSort(shop_sort ShopSort) error {
	// 开始事务
	tx := Db.Begin()
	if err := Db.Create(&shop_sort).Error; err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		return err
	}
	// 否则，提交事务
	tx.Commit()
	return nil
}

func EditSort(id int, data ShopSort) error {
	// 开始事务
	tx := Db.Begin()
	if err := Db.Model(&ShopSort{}).Where("id = ?", id).Updates(data).Error; err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		return err
	}
	// 否则，提交事务
	tx.Commit()
	return nil
}

func DeleteSort(id int) error {
	if err := Db.Where("id = ?", id).Delete(ShopSort{}).Error; err != nil {
		return err
	}

	return nil
}

// 根据id判断test 对象是否存在
func ExistSortByID(id int) bool {
	var shop_sort ShopSort
	Db.Select("id").Where("id = ?", id).First(&shop_sort)

	return shop_sort.ID > 0
}
