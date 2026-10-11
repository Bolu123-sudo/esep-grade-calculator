# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

## Grade Categories
- Assignments: 50%
- Exams: 35%
- Essays: 15%

## Grade Scale
- A: 90–100
- B: 80–89
- C: 70–79
- D: 60–69
- F: Below 60

A grade of C or higher is considered passing.

## Project Structure
- `grade-calculator/grade_calculator.go`: Grade calculation implementation
- `grade-calculator/grade_calculator_test.go`: Unit tests
- `grade-calculator/go.mod`: Go module configuration

## Running Tests

From the repository root:

```bash
cd grade-calculator
go test ./... -v
```

## Measuring Code Coverage

From the `grade-calculator` directory:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

## How It Works
1. Create a calculator using `NewGradeCalculator()`.
2. Add grades using `AddGrade()`.
3. Calculate the final letter grade using `GetFinalGrade()`.
