//qn: print args[0]

package main
import "os"
import "fmt"

func main()  {
	fmt.Printf("the first arg is %s", os.Args[0])
}