# Go Lang
- Go uses the garbage collecttion for the memory management. This automaticalyy handle the memory allocation and delocation but can introduce the latency due to GC pauses .
- TGo provide the goroutine for the lightweight concurrent executions and channels for the communication between them . 
- These two commands are the most important command for the execution of the things . 
- go build filename.go
- go run filename.go 
 This is the most important thing in the go lang to start the code execution .
 # Go 
- MVC stands for the Model-View- Controller it is the design pattern used in software engineering to seperate the concerns in the development of applications . Model is data or businees logic of the application , View is the user interface and cotroller as as intermediary between model and the view .
- Handler 
- Router (servermux in Go terminology)
- Web server One of the great things about Go is that you can establish a web server and listen for incoming requests as part of application itself .
- Each time web server listening new HTTP request it will pass the request on the servermux in turn the servermux will check the URL path and dispatch the request to the matching handler .\
- The TCP network address that you pass to the http.ListenAndServe() should be in the format "host:port". If we omit the host then the server will listen on all your computers available .
- We use log.Println() and log.Fatal() functions to output log message.Both these function output messages via GO 'standards' .
- log.fatal () function will also call os.exit(1) after writing the message causing the application to immediately exit 
- Go servermux supports two different types of URL patterns: fixed paths and the subtree paths .
   - Fixed path dont end with the trailing slash whereas subtree paths do end with a trailing slash .
   

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

# API Development with the GO lang
-  Go Fiber comes with the zero memory allocation and prformance in mind .It main aim is for the high performance and low latecy .
- `Fiber` It represents the fiber package where you start to create an instance .
- `App` The app instance conventially denotes the fiber application .
- `Ctx` The Ctx struct holds the Http request and response .
- `Handler` are the first component of the api it is responsible for executing the application logic and for writting HTTP response headers and bodies .
- `Router` This is also known as the servermux this stores a mapping between the URL patterns for your application and the corresponding handllers . Usually one servermux is per application controll whole of routes .
-` webserver`  it listen incoming request as part of your application itself.
# The defaultserveMux
- `http.Handle () and http.HandleFunc()` These allow you to register routes without declaring servermux . This is not recomended for the production application 
```GO
func main() {
http.HandleFunc("/", home)
http.HandleFunc("/snippet", showSnippet)
http.HandleFunc("/snippet/create", createSnippet)
log.Println("Starting server on :4000")
err := http.ListenAndServe(":4000", nil)
log.Fatal(err)
}
```
- Servermux will always dispatches the pattern that have the longest URL if there is any conflicts among the multiple requests .
- The http.Error Shortcut `http.Error` is lightweight helper function takes given message and status code then calls the w.WriteHeader() w.write () methods behind the scenes for us .
# Manipulating the Header Map
- `w.Header().Set` to add a new header to the response header map .
- `w.Header().Add('')` to add new cache-control
- `w.Header().Del('')` to Delete all values for the cache-control
- `w.Header().Get('')` to retirive the first value for the cache-control.
- Go will send the automatically set of three system-generated header for you date and content-length and the content type .
