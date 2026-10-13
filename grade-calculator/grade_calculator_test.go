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

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeCAndD(t *testing.T) {
	cases := []struct {
		name     string
		score    int
		expected string
	}{
		{name: "C boundary", score: 70, expected: "C"},
		{name: "D boundary", score: 60, expected: "D"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calculator := NewGradeCalculator()
			calculator.AddGrade("assignment", tc.score, Assignment)
			calculator.AddGrade("exam", tc.score, Exam)
			calculator.AddGrade("essay", tc.score, Essay)

			if actual := calculator.GetFinalGrade(); actual != tc.expected {
				t.Errorf("GetFinalGrade() = %q; want %q", actual, tc.expected)
			}
		})
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := []struct {
		gradeType GradeType
		expected  string
	}{
		{gradeType: Assignment, expected: "assignment"},
		{gradeType: Exam, expected: "exam"},
		{gradeType: Essay, expected: "essay"},
	}

	for _, tc := range cases {
		if actual := tc.gradeType.String(); actual != tc.expected {
			t.Errorf("GradeType(%d).String() = %q; want %q", tc.gradeType, actual, tc.expected)
		}
	}
}

func TestComputeAverageEmpty(t *testing.T) {
	if actual := computeAverage(nil); actual != 0 {
		t.Errorf("computeAverage(nil) = %d; want 0", actual)
	}
}

func TestGetFinalGradeMultipleScores(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("assignment 1", 80, Assignment)
	calculator.AddGrade("assignment 2", 100, Assignment)
	calculator.AddGrade("exam 1", 70, Exam)
	calculator.AddGrade("exam 2", 90, Exam)
	calculator.AddGrade("essay 1", 60, Essay)
	calculator.AddGrade("essay 2", 80, Essay)

	// Category averages are 90, 80, and 70; the weighted grade is 83.5.
	if actual := calculator.GetFinalGrade(); actual != "B" {
		t.Errorf("GetFinalGrade() = %q; want B", actual)
	}
}

func TestGetFinalGradeDistinctCategoryScores(t *testing.T) {
	calculator := NewGradeCalculator()
	calculator.AddGrade("assignment", 100, Assignment)
	calculator.AddGrade("exam", 80, Exam)
	calculator.AddGrade("essay", 0, Essay)

	// 100*0.50 + 80*0.35 + 0*0.15 = 78, which is a C.
	if actual := calculator.GetFinalGrade(); actual != "C" {
		t.Errorf("GetFinalGrade() = %q; want C", actual)
	}
}
