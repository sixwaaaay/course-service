package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Course struct {
	ID     int32  `gorm:"primary_key"`
	Name   string `gorm:"type:varchar(100);unique_index"`
	Avatar string `gorm:"type:varchar(100)"`
	Hours  int32  `gorm:"type:int(32)"`
	Sid    int32  `gorm:"type:int(32)"`
}

// TableName c_course
func (Course) TableName() string {
	return "c_course"
}

type CourseModel struct {
	db *gorm.DB
}

// NewCourseModel returns a new CourseModel
func NewCourseModel(datasource string) *CourseModel {
	// connect to the database
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &CourseModel{db}
}

/*// MGetNotes multiple get list of note info
  func MGetNotes(ctx context.Context, noteIDs []int64) ([]*Note, error) {
  	var res []*Note
  	if len(noteIDs) == 0 {
  		return res, nil
  	}

  	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&res).Error; err != nil {
  		return res, err
  	}
  	return res, nil
  }

  // UpdateNote update note info
  func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
  	params := map[string]interface{}{}
  	if title != nil {
  		params["title"] = *title
  	}
  	if content != nil {
  		params["content"] = *content
  	}
  	return DB.WithContext(ctx).Model(&Note{}).Where("id = ? and user_id = ?", noteID, userID).
  		Updates(params).Error
  }

  // DeleteNote delete note info
  func DeleteNote(ctx context.Context, noteID, userID int64) error {
  	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&Note{}).Error
  }

  // QueryNote query list of note info
  func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Note, int64, error) {
  	var total int64
  	var res []*Note
  	conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id = ?", userID)

  	if searchKey != nil {
  		conn = conn.Where("title like ?", "%"+*searchKey+"%")
  	}

  	if err := conn.Count(&total).Error; err != nil {
  		return res, total, err
  	}

  	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
  		return res, total, err
  	}

  	return res, total, nil
  }*/

// CreateCourse

func (m *CourseModel) CreateCourse(ctx context.Context, course *Course) error {
	return m.db.WithContext(ctx).Create(course).Error
}

// MGetCourses multiple get list of course info
func (m *CourseModel) MGetCourses(ctx context.Context, courseIDs []int32) ([]*Course, error) {
	var res []*Course
	if len(courseIDs) == 0 {
		return res, nil
	}
	if err := m.db.WithContext(ctx).Where("id in ?", courseIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateCourse update course info
func (m *CourseModel) UpdateCourse(ctx context.Context, courseID int32, name *string, avatar *string, sid *int32, hours *int32) error {
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if avatar != nil {
		params["avatar"] = *avatar
	}
	if sid != nil {
		params["sid"] = *sid
	}
	if hours != nil {
		params["hours"] = *hours
	}
	return m.db.WithContext(ctx).Model(&Course{}).Where("id = ?", courseID).Updates(params).Error
}

// DeleteCourse delete course info
func (m *CourseModel) DeleteCourse(ctx context.Context, courseID int32) error {
	return m.db.WithContext(ctx).Where("id = ?", courseID).Delete(&Course{}).Error
}

// QueryCourseBySid query list of course info
func (m *CourseModel) QueryCourseBySid(ctx context.Context, sid *int32, limit, offset int) ([]*Course, int64, error) {
	var total int64
	var res []*Course
	var conn *gorm.DB
	if sid != nil {
		conn = m.db.WithContext(ctx).Model(&Course{}).Where("sid = ?", *sid)
	} else {
		conn = m.db.WithContext(ctx).Model(&Course{})
	}
	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}
	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

// QueryDuplicateName query the course name is duplicate
func (m *CourseModel) QueryDuplicateName(ctx context.Context, name string) (bool, error) {
	var res []*Course
	if err := m.db.WithContext(ctx).Where("name = ?", name).Find(&res).Error; err != nil {
		return false, err
	}
	if len(res) > 0 {
		return true, nil
	}
	return false, nil
}
