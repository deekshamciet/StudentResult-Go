# Student Result Processing System Using Go

## 1. Project Overview
The Student Result Processing System is a command-line application developed using the Go programming language. It accepts student details and subject marks, calculates total marks and average, assigns grades and displays the final result.

## 2. Objectives
- Understand Go programming fundamentals.
- Accept student information and subject marks.
- Calculate total and average marks.
- Assign grades automatically.
- Check student pass or fail status.
- Implement modular programming using packages.
- Maintain source code using Git and GitHub.

## 3. Technologies Used
- Programming Language: Go
- Code Editor: Visual Studio Code
- Version Control: Git
- Repository: GitHub

## 4. Features
- Student name and roll number input.
- Five subject marks input.
- Automatic total calculation.
- Average percentage calculation.
- Automatic grade assignment.
- Pass or fail checking.
- Invalid mark validation.

## 5. Project Structure

```text
StudentResult/
├── go.mod
├── main.go
└── result/
    └── result.go
```

## 6. Grading System

| Average | Grade |
|---|---|
| 90–100 | A+ |
| 80–89 | A |
| 70–79 | B |
| 60–69 | C |
| 50–59 | D |
| Below 50 | E |
| Any subject below 40 | F (Fail) |

## 7. Execution Instructions

1. Install Go.
2. Download or clone this repository.
3. Open the project folder in Visual Studio Code.
4. Open the terminal.
5. Execute:

```bash
go mod tidy
go run .
```

## 8. Sample Output

```text
Student Name : Deeksha
Roll Number  : 15

Tamil                : 99
English              : 91
Maths                : 100
Science              : 98
Computer Science     : 95

Total Marks  : 483
Average      : 96.60
Grade        : A+
Result       : PASS
```

## 9. Go Concepts Used
- Variables and data types
- Functions
- Arrays and slices
- Conditional statements
- Switch statements
- For loops
- Packages and modules
- Input and output operations

## 10. Version Control
Git is used to track changes in the project. GitHub is used to store and share the source code.

## 11. Conclusion
The Student Result Processing System was successfully developed using Go. It automates student result calculations and demonstrates the use of Go fundamentals, functions, packages, modular programming and GitHub version control.
