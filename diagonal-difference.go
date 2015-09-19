package main
import "fmt"
import "math"

func main() {
    var n int
    fmt.Scan(&n)
    var total, tmp int
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            fmt.Scan(&tmp)
            if (i == j) {
                if i + j + 1 == n {
                    continue
                }
                total += tmp
            } else if (math.Abs(float64(i - j)) + 1 == float64(n)) {
                total -= tmp
            } else {
                continue
            }
        }
    }
    fmt.Println(math.Abs(float64(total)))
}
