package Student

import (
	"fmt"
	"school/week3/Course"
)

type Student struct {
	StudentName string
	Department  string
	PaidFees    bool
	CourseList  []Course.Course
}

func cleanUp() {
	r := recover()
	if r != nil {
		fmt.Printf("%v", r)
	}
}

func (s *Student) coursePresent(course Course.Course) bool {
	if len(s.CourseList) == 0 {
		return false
	}
	for i := 0; i < len(s.CourseList); i++ {
		if s.CourseList[i].CourseName == course.CourseName {
			return true
		}
	}
	return false
}
func (s *Student) TakeCourse(course Course.Course) []Course.Course {

	if course.CourseName == " " {
		panic("!!!!Course Name is Empty")
	}
	if !s.coursePresent(course) {
		s.CourseList = append(s.CourseList, course)
	}
	if s.CourseList[len(s.CourseList)-1].CourseName == course.CourseName {

	}
	return s.CourseList
}
