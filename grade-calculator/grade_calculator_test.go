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

func TestGetGradesHighScores(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("Assignment 1", 75, Assignment)
    gc.AddGrade("Exam 1", 75, Exam)
    gc.AddGrade("Essay 1", 75, Essay)

    if actual := gc.GetFinalGrade(); actual != "C" {
        t.Errorf("Expected C, got %s", actual)
    }
}

func TestGetGradeD(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("Assignment 1", 65, Assignment)
    gc.AddGrade("Exam 1", 65, Exam)
    gc.AddGrade("Essay 1", 65, Essay)

    if actual := gc.GetFinalGrade(); actual != "D" {
        t.Errorf("Expected D, got %s", actual)
    }
}

func TestGetGradeFWithLowScores(t *testing.T) {
    gc := NewGradeCalculator()
    gc.AddGrade("Assignment 1", 40, Assignment)
    gc.AddGrade("Exam 1", 40, Exam)
    gc.AddGrade("Essay 1", 40, Essay)

    if actual := gc.GetFinalGrade(); actual != "F" {
        t.Errorf("Expected F, got %s", actual)
    }
}

func TestGradeTypeString(t *testing.T) {
    tests := []struct {
        gradeType GradeType
        expected  string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }

    for _, test := range tests {
        actual := test.gradeType.String()
        if actual != test.expected {
            t.Errorf("Expected %s, got %s", test.expected, actual)
        }
    }
}

func TestComputeAverageEmpty(t *testing.T) {
    actual := computeAverage([]Grade{})

    if actual != 0 {
        t.Errorf("Expected 0 for empty grades, got %d", actual)
    }
}

