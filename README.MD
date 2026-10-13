# ESEP Grade Calculator

A Go library for calculating a final letter grade from assignment, exam, and essay scores.

## Requirements

- Go 1.24.4 or later
- Git to clone the repository

The project uses Go modules and the standard-library `testing` package. No third-party dependencies are required.

## Grading Rules

The intended final numerical grade is calculated by averaging the scores within each category, then applying these weights:

| Category | Weight |
|----------|--------|
| Assignments | 50% |
| Exams | 35% |
| Essays | 15% |

The calculator uses the following letter-grade thresholds:

| Letter Grade | Minimum Numerical Grade |
|--------------|-------------------------|
| A | 90 |
| B | 80 |
| C | 70 |
| D | 60 |
| F | Below 60 |

The current implementation uses integer averages and truncates the final weighted score to an integer.

## Project Structure

- `README.MD`: Project overview, usage, and testing instructions.
- `grade-calculator/go.mod`: Go module name and required Go version.
- `grade-calculator/grade_calculator.go`: Calculator implementation.
- `grade-calculator/grade_calculator_test.go`: Unit tests.

## Usage

The following example can be used within the `esepunittests` package:

```go
calculator := NewGradeCalculator()

calculator.AddGrade("Assignment 1", 90, Assignment)
calculator.AddGrade("Exam 1", 80, Exam)
calculator.AddGrade("Essay 1", 100, Essay)

finalGrade := calculator.GetFinalGrade()
```

According to the grading rules, these scores produce a weighted result of 88 and a letter grade of B.

## Run Unit Tests

From the repository root:

```powershell
cd grade-calculator
go test -v ./...
```

The `-v` option displays individual test results. The `./...` pattern includes the current package and its subpackages.

## Measure Code Coverage

From the `grade-calculator` directory:

```powershell
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
```

To open a visual coverage report in your browser:

```powershell
go tool cover -html=coverage.txt
```

Coverage reports which source statements executed during testing. It does not guarantee that every input or behavior has been tested.

## Current Limitations

The initial implementation contains known calculation bugs and an incorrectly formed failing-grade test. The usage example describes the intended result; run the tests to evaluate the current implementation.

Each category must contain at least one score to avoid division by zero. Score validation and a policy for missing categories are not implemented.