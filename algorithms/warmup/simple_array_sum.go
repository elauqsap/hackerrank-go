package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    var size,total int
    fmt.Scanf("%d", &size)
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        for _, sub := range strings.Split(scanner.Text(), " ") {
            i, _ := strconv.Atoi(sub)
            total = total + i
        }
    }
    fmt.Printf("%d\n", total)
}