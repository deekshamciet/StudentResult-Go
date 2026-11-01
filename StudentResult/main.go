package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "studentresult/result"
)

func main() {

    reader := bufio.NewReader(os.Stdin)

    var s result.Student

    // Get student details
    fmt.Print("Enter Student Name: ")
    s.Name, _ = reader.ReadString('\n')
    s.Name = strings.TrimSpace(s.Name)

    fmt.Print("Enter Roll Number: ")
    s.RollNo, _ = reader.ReadString('\n')
    s.RollNo = strings.TrimSpace(s.RollNo)

    // Subject names
    subjects := []string{
        "Tamil",
        "English",
        "Maths",
        "Science",
        "Computer Science",
    }

    // Get subject marks
    for _, subject := range subjects {

        for {
            fmt.Printf("Enter %s Mark: ", subject)

            input, _ := reader.ReadString('\n')
            input = strings.TrimSpace(input)

            mark, err := strconv.Atoi(input)

            if err == nil && mark >= 0 && mark <= 100 {
                s.Marks = append(s.Marks, mark)
                break
            }

            fmt.Println("Invalid mark! Enter between 0 and 100.")
        }
    }

    // Calculate results
    total := result.CalculateTotal(s.Marks)

    average := result.CalculateAverage(s.Marks)

    passed := result.IsPassed(s.Marks)

    grade := result.GetGrade(average, passed)

    // Display student result
    fmt.Println("\n================================")
    fmt.Println("       STUDENT RESULT")
    fmt.Println("================================")

    fmt.Println("Student Name :", s.Name)
    fmt.Println("Roll Number  :", s.RollNo)

    fmt.Println("--------------------------------")

    for i, subject := range subjects {
        fmt.Printf("%-20s : %d\n", subject, s.Marks[i])
    }

    fmt.Println("--------------------------------")

    fmt.Println("Total Marks  :", total)

    fmt.Printf("Average      : %.2f\n", average)

    fmt.Println("Grade        :", grade)

    if passed {
        fmt.Println("Result       : PASS")
    } else {
        fmt.Println("Result       : FAIL")
    }

    fmt.Println("================================")
}