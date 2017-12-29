package main
import "fmt"
import "time"

func main() {
  fmt.Println(time.Now())
  fmt.Println(time.Now().Truncate(24 * time.Hour))		// Truncate works in UTC
  fmt.Println(time.Now().UTC().Truncate(24 * time.Hour))	// Truncate works in UTC

  t := time.Now()
  truncated := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
  fmt.Println("Now local:", t)
  fmt.Println("Truncated now local:", truncated)

  fmt.Println("Yesterday:", t.Add(-1 * 24 * time.Hour))
  fmt.Println("Truncated yesterday:", truncated.Add(-1 * 24 * time.Hour))

  fmt.Println("Yesterday:", t.AddDate(0, 0, -1))
  fmt.Println("Truncated yesterday:", truncated.AddDate(0, 0, -1))
}