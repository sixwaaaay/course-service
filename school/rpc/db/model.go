package db

// import gorm

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// use gorm as the ORM layer
// the table is "s_school"
// the primary key is "id"
// it's struct is below:

type School struct {
	ID   int32  `gorm:"primary_key"`
	Name string `gorm:"type:varchar(40)"`
}

type SchoolModel struct {
	db *gorm.DB
}

// NewSchoolModel create a new SchoolModel
func NewSchoolModel(datasource string) *SchoolModel {
	// connect to the database
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &SchoolModel{db: db}
}

// TableName  "s_school"
func (School) TableName() string {
	return "s_school"
}

// MGetSchool get all schools whose id in given id array
func (sm *SchoolModel) MGetSchool(ctx context.Context, ids []int32) ([]*School, error) {
	var schools []*School
	if len(ids) == 0 {
		return schools, nil
	}
	if err := sm.db.WithContext(ctx).Where("id in (?)", ids).Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}

// GetIdByName get school id by name
func (sm *SchoolModel) GetIdByName(ctx context.Context, name string) (int32, error) {
	var school School
	if err := sm.db.WithContext(ctx).Where("name = ?", name).First(&school).Error; err != nil {
		return 0, err
	}
	return school.ID, nil
}
