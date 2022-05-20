package pack

import (
	"course/course/api/internal/types"
	"course/course/rpc/types/course"
	"course/school/rpc/types/school"
)

func BuildResp(status int32, message string) *types.BaseResponse {
	return &types.BaseResponse{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

func ToCourse(m *course.Course) *types.Course {
	if m == nil {
		return nil
	}

	return &types.Course{
		Id:     m.Id,
		Name:   m.Name,
		Avatar: m.Avatar,
		Sid:    m.Sid,
		Hours:  m.Hours,
	}
}

// ToCourseList covert to Course
func ToCourseList(ms []*course.Course) []*types.Course {
	if ms == nil {
		return nil
	}

	list := make([]*types.Course, len(ms))
	for i := range ms {
		list[i] = ToCourse(ms[i])
	}

	return list
}

func SchoolIds(ms []*course.Course) []int32 {
	uIds := make([]int32, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int32]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.Sid] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}

func SchoolMap(ms []*school.School) map[int32]*school.School {
	uIdMap := make(map[int32]*school.School)
	if len(ms) == 0 {
		return uIdMap
	}
	for _, m := range ms {
		if m != nil {
			uIdMap[m.Id] = m
		}
	}
	return uIdMap
}
