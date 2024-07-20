# Go Lang
- Go uses the garbage collecttion for the memory management. This automaticalyy handle the memory allocation and delocation but can introduce the latency due to GC pauses .
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
```Go
package main 
import "fmt"
func main () {

    n_4 := 'F';  // This will present the unicode literals .
    // To print the excact single value we can use i/o format 
    fmt.Println(n_4);      //  Writen the default the formats .
    fmt.Printf("%c\n",n_4);  // Printf rturn according to format specifier . 
}
```
# FUNCTION

- It is a designed of a code that is self block of code and designed to perform the specific task .
```Go
package main 
import "fmt"


// Global function 
 // add is function name , a and v are the two parameter then we retunring function .
func add (a float32, v float32) float32 {
    return a + v;
}
func main () {

    n3 := add(3,5);
    fmt.Println(n3);

}
```

# LINKED LIST
- In a singly list a node is the fundamental elemet that constitues the structure of the list . 
- It is the linear data structure where each element is stored in node and each node contains the two main componets .
- Data componets holds the value or information associated with that particular node .
- Next pointer It is also known as the reference or link, points to the next node in the sequence ,It establish the connection between nodes in the list creating the unidirectional flow from the node to the next . And the last node is is typically nil . 

# QUEUE
# Struct 
- It is the typicaaly collection of the fields . They are useful of grouping the data together .
```Go
package main
import "fmt"

type person struct {
    name string
    age int 
}


func main () {
fmt.Println(person{"KAusik",30})
}




```

# API Development with the 
-  Go Fiber comes with the zero memory allocation and prformance in mind .It main aim is for the high performance and low latecy .
- `Fiber` It represents the fiber package where you start to create an instance .
- `App` The app instance conventially denotes the fiber application .
- ` Ctx` The Ctx struct holds the Http request and response .

