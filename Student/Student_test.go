package Student

import (
	"school/week3/Course"
	"testing"
)

func TestStudent_TakeCourse(t *testing.T) {
	//defer cleanUp()
	s := Student{"Emma", "Geology", true, []Course.Course{}}

	table := []struct {
		input    Course.Course
		expected string
	}{
		{Course.Course{"Biology", 5}, "Biology"},
		{Course.Course{"Chemistry", 4}, "Chemistry"},
		{Course.Course{"Chemistry", 4}, "Chemistry"},
	}
	for _, val := range table {
		list := s.TakeCourse(val.input)
		if list[len(s.CourseList)-1].CourseName != val.expected {
			t.Errorf("TakeCourse: Expected %v Got %v", val.expected, list[len(s.CourseList)-1].CourseName)
		}
	}
}
