package goexamplepackage
import "fmt"
import foo "github.com/digitallumens/goexamplepackage/foo"

func DoExample(boo int) int {
     fmt.Println("Example:", boo)
     return foo.DoBar()
}
