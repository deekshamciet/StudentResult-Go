package result

type Student struct {
    Name   string
    RollNo string
    Marks  []int
}

func CalculateTotal(marks []int) int {
    total := 0
    for _, mark := range marks {
        total += mark
    }
    return total
}

func CalculateAverage(marks []int) float64 {
    if len(marks) == 0 {
        return 0
    }
    return float64(CalculateTotal(marks)) / float64(len(marks))
}

func IsPassed(marks []int) bool {
    for _, mark := range marks {
        if mark < 40 {
            return false
        }
    }
    return len(marks) > 0
}

func GetGrade(average float64, passed bool) string {
    if !passed {
        return "F"
    }

    switch {
    case average >= 90:
        return "A+"
    case average >= 80:
        return "A"
    case average >= 70:
        return "B"
    case average >= 60:
        return "C"
    case average >= 50:
        return "D"
    default:
        return "E"
    }
}
