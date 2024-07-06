# Go Lang
- Go uses the garbase collecttion for the memory management. This automaticalyy handle the memory allocation and delocation but can introduce the latency due to GC pauses .
- TGo provide the goroutine for the lightweight concurrent executions and channels for the communication between them . 
- These two commands are the most important command for the execution of the things . 
- go build filename.go
- go run filename.go 
 This is the most important thing in the go lang to start the code execution .

```Go
package main 
import ( 

)
func main () {

}

```
# Priniting the input 
- The below command is use to print the value of the go . 
```Go
fmt.Println ("")
```
# ARRAY
- Arrays are considered values rather than pointers and represent the entirety of the array. Whenever an array is passed to a function, a copy is created, resulting in additional memory usage. To avoid this, it is possible to pass a pointer to an array, or use slices instead. The size of the array is constant and it must be known at compile time .
- 

```go

package main
import "fmt"

func main () {
 
var nu =[3]int{1,3,4,};
nu2 := [3]init {2,4,5};
fmt.Println (nu);
fmt.Println(nu2);
};
```




# FUNCTION 
# LINKED LIST
# QUEUE

# API Development with the 
-  Go Fiber comes with the zero memory allocation and prformance in mind .It main aim is for the high performance and low latecy .
- `Fiber` It represents the fiber package where you start to create an instance .
- `App` The app instance conventially denotes the fiber application .
- ` Ctx` The Ctx struct holds the Http request and response .

