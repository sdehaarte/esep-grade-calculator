package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 78, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 74, Essay)
	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"
	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("open source assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam 1", 66, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)
	actual_value := gradeCalculator.GetFinalGrade()
	
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 39, Assignment)
	gradeCalculator.AddGrade("exam 1", 37, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 33, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeType(t *testing.T) {
	expected_value := "assignment"
	if expected_value != Assignment.String() {
		t.Errorf("Expected GradeType to return '%s'; got '%s' instead", expected_value, Assignment.String())
	}

	expected_value = "exam"
	if expected_value != Exam.String() {
		t.Errorf("Expected GradeType to return '%s'; got '%s' instead", expected_value, Exam.String())
	}

	expected_value = "essay"
	if expected_value != Essay.String() {
		t.Errorf("Expected GradeType to return '%s'; got '%s' instead", expected_value, Essay.String())
	}
}
