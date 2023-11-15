package package1
import "fmt"
func package1Function() int {
    fmt.Println("Package 1 side-effect")
    return 1
}
var globalVariable = package1Function()
func init() {
    fmt.Println("Package 1 init side effect")
}