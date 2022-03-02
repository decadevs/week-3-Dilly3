package Courses

import "fmt"

type Course struct {
	CourseName       string
	CourseCreditLoad int32
	CourseList       []Course
}

func (c *Course) PrintCourseList() {
	for ind, val := range c.CourseList {
		fmt.Printf("%d : %s", ind, val)
	}
}

func (c *Course) CheckCourseInList(myCourse Course) bool {
	for i := 0; i < len(c.CourseList); i++ {

		if c.CourseList[i].CourseName == myCourse.CourseName {
			return true
		}
	}
	return false
}

func (c *Course) AddNewCourse(newCourse Course) {
	if isPresent := c.CheckCourseInList(newCourse); !isPresent {
		c.CourseList = append(c.CourseList, newCourse)
	}
}
