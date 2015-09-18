package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var times, temp int
    var result int64
    times = read()
    for i := 0; i < times; i++ {
        temp = read()
        result += int64(temp)
    }
    fmt.Println(result)
}

func read() int {
    var n int
    _, err := fmt.Scan(&n)

    if err != nil {
        panic(err)
    }

    return n
}
