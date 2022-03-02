package Teacher

import (
	"fmt"
	"school/week3/Course"
)

type name = string
type age = int32
type yearOfService = int32

func cleanup() {
	r := recover()
	if r != nil {
		fmt.Printf("%v", r)
	}
}

type Teacher struct {
	TeacherName       string
	YearsOfService    int32
	TeacherAge        int32
	Course            Course.Course
	StudentScores     map[string]float64
	StudentAttendance map[string]int32
}

func (t *Teacher) GetDistinctionStudents(cutoffMark float64) map[string]float64 {
	distinctionStudent := make(map[string]float64)

	for k, v := range t.StudentScores {
		if v >= cutoffMark {
			distinctionStudent[k] = v
		}
	}
	return distinctionStudent

}

func (t *Teacher) GetFailingStudents(cutoffMark float64) map[string]float64 {
	failingStudents := make(map[string]float64)

	for k, v := range t.StudentScores {
		if v <= cutoffMark {
			failingStudents[k] = v
		}
	}
	return failingStudents

}

func (t *Teacher) GetAttendancePercentage(s name) int32 {
	defer cleanup()
	if s == "" {
		panic("!!!!!Empty String as input!!!!!")
	}
	if val, ok := t.StudentAttendance[s]; ok {
		return val
	}
	return 0
}
