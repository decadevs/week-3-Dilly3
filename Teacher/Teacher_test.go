package Teacher

import (
	"fmt"
	"testing"
)
import "school/week3/Course"

func TestTeachers_GetDistinctionStudents(t *testing.T) {
	teac := Teacher{"Okoli", 15, 43,
		Course.Course{"Biology", 3},
		map[name]float64{"mike": 79, "Tim": 56, "George": 65, "diaz": 32}, map[name]int32{}}

	table := []struct {
		cutOffMark float64
	}{
		{70}, {60}, {50}, {40},
	}

	for _, v := range table {
		m := teac.GetDistinctionStudents(v.cutOffMark)
		for _, val := range m {
			if val < v.cutOffMark {
				t.Errorf("grade cant be less than cut off")
			}
		}

	}
}

func TestTeachers_GetFailingtudents(t *testing.T) {
	teac := Teacher{"Okoli", 15, 43,
		Course.Course{"Biology", 3},
		map[name]float64{"mike": 79, "Tim": 56, "George": 65, "diaz": 32}, map[name]int32{}}

	table := []struct {
		cutOffMark float64
	}{
		{70}, {60}, {50}, {40},
	}

	for _, v := range table {
		m := teac.GetFailingStudents(v.cutOffMark)
		for _, val := range m {
			if val >= v.cutOffMark {
				t.Errorf("grade cant be more  than cut off")
			}
		}

	}
}

func TestTeacher_GetAttendancePercentage(t *testing.T) {
	r := recover()
	if r != nil {
		fmt.Printf("%v", r)
	}

	teach := Teacher{"Okoli", 15, 43,
		Course.Course{"Biology", 3},
		map[name]float64{"mike": 79, "Tim": 56, "George": 65, "diaz": 32},
		map[name]int32{"mike": 4, "Tim": 6, "George": 9, "diaz": 3}}

	table := []struct {
		name     string
		expected int32
	}{
		{"mike", 4}, {"Tim", 6}, {"George", 9},
	}
	for _, val := range table {

		m := teach.GetAttendancePercentage(val.name)

		if m != val.expected {
			t.Errorf("GetAttendancePercentage: Expected %v , Got %v", val.expected, m)
		}
	}
}
