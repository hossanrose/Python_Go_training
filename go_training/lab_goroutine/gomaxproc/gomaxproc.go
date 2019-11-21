package
main
import
(
    "runtime"
    "fmt"
)
func getGOMAXPROCS() int {
    return runtime.GOMAXPROCS(0)
}

func test(msg string) {
        fmt.Println(msg)
    }

func main() {
    fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
    for i:=0;i<10;i++{
	    go test("Hello")
    }
    runtime.GOMAXPROCS(2)
    fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
}

