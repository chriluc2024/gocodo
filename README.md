| Counter | Concept                         | Explanation                                                                                  | Demo                                                                                          |
|---------|---------------------------------|----------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| 1       | Hello World                     | Basic program to print "Hello, World!"                                                       | ```go fmt.Println("Hello, World!")```                                                         |
| 2       | Values                          | Basic values and their types in Go                                                            | ```go fmt.Println("Go" + "Lang")```                                                           |
| 3       | Variables                       | Declaring and using variables                                                                 | ```go var a = "initial" fmt.Println(a)```                                                     |
| 4       | Constants                       | Declaring and using constants                                                                 | ```go const Pi = 3.14```                                                                      |
| 5       | For                             | The only looping construct in Go                                                              | ```go for i := 0; i < 10; i++ { fmt.Println(i) }```                                           |
| 6       | If/Else                         | Conditional statements                                                                        | ```go if num := 9; num < 0 { fmt.Println(num, "is negative") } else { fmt.Println(num) }```   |
| 7       | Switch                          | Multiple conditional statements                                                               | ```go switch time.Now().Weekday() { case time.Saturday, time.Sunday: fmt.Println("Weekend") }```|
| 8       | Arrays                          | Fixed-length sequences of elements                                                            | ```go var a [5]int fmt.Println(a)```                                                          |
| 9       | Slices                          | Dynamically-sized, flexible view into the elements of an array                                | ```go s := []int{2, 3, 5, 7, 11, 13} fmt.Println(s)```                                        |
| 10      | Maps                            | Collection of key/value pairs                                                                 | ```go m := make(map[string]int) m["k1"] = 7 fmt.Println(m)```                                 |
| 11      | Range                           | Iterates over elements in a variety of data structures                                        | ```go for i, v := range s { fmt.Println(i, v) }```                                            |
| 12      | Functions                       | Grouping code into reusable blocks                                                            | ```go func add(a int, b int) int { return a + b }```                                          |
| 13      | Multiple Return Values          | Functions that return multiple values                                                         | ```go func vals() (int, int) { return 3, 7 }```                                               |
| 14      | Variadic Functions              | Functions that accept a variable number of arguments                                          | ```go func sum(nums ...int) { fmt.Println(nums) }```                                          |
| 15      | Closures                        | Functions that reference variables outside their scope                                        | ```go func intSeq() func() int { i := 0 return func() int { i++ return i } }```               |
| 16      | Recursion                       | Functions that call themselves                                                                | ```go func fact(n int) int { if n == 0 { return 1 } return n * fact(n-1) }```                 |
| 17      | Pointers                        | Variables that store memory addresses                                                         | ```go var p *int i := 42 p = &i fmt.Println(*p)```                                            |
| 18      | Strings and Runes               | Text handling in Go                                                                           | ```go str := "Hello" fmt.Println(len(str))```                                                 |
| 19      | Structs                         | Collections of fields                                                                         | ```go type person struct { name string age int }```                                           |
| 20      | Methods                         | Functions associated with struct types                                                        | ```go func (p person) greet() string { return "Hello, " + p.name }```                         |
| 21      | Interfaces                      | Abstractions for different types                                                              | ```go type animal interface { sound() string }```                                             |
| 22      | Struct Embedding                | Including one struct in another to inherit its fields and methods                             | ```go type base struct { num int }```                                                         |
| 23      | Generics                        | Writing functions and data structures that can work with any type                             | ```go func printSlice[T any](s []T) { for _, v := range s { fmt.Println(v) } }```             |
| 24      | Errors                          | Error handling in Go                                                                          | ```go _, err := strconv.Atoi("non-int") if err != nil { fmt.Println(err) }```                 |
| 25      | Custom Errors                   | Defining custom error types                                                                   | ```go type MyError struct { msg string }```                                                   |
| 26      | Goroutines                      | Lightweight threads managed by the Go runtime                                                 | ```go go func() { fmt.Println("in goroutine") }()```                                          |
| 27      | Channels                        | Pipes that connect concurrent goroutines                                                      | ```go ch := make(chan int) go func() { ch <- 1 }()```                                         |
| 28      | Channel Buffering               | Channels with capacity to store values                                                        | ```go ch := make(chan int, 2) ch <- 1```                                                      |
| 29      | Channel Synchronization         | Synchronizing execution between goroutines                                                    | ```go done := make(chan bool) go func() { fmt.Println("done") done <- true }() <-done```      |
| 30      | Channel Directions              | Restricting channel use to send-only or receive-only                                           | ```go func ping(pings chan<- string) { pings <- "ping" }```                                   |
| 31      | Select                          | Handling multiple channel operations                                                          | ```go select { case msg := <-ch: fmt.Println(msg) }```                                        |
| 32      | Timeouts                        | Setting time limits on operations                                                             | ```go select { case <-ch: case <-time.After(time.Second): fmt.Println("timeout") }```         |
| 33      | Non-Blocking Channel Operations | Attempting channel operations without blocking                                                | ```go select { case msg := <-ch: fmt.Println(msg) default: fmt.Println("no message") }```     |
| 34      | Closing Channels                | Closing a channel to indicate that no more values will be sent                                 | ```go close(ch)```                                                                            |
| 35      | Range over Channels             | Iterating over values received from a channel                                                 | ```go for v := range ch { fmt.Println(v) }```                                                 |
| 36      | Timers                          | Representing a single event in the future                                                     | ```go timer := time.NewTimer(2 * time.Second) <-timer.C fmt.Println("Timer fired")```         |
| 37      | Tickers                         | Repeatedly triggering events at regular intervals                                             | ```go ticker := time.NewTicker(time.Millisecond * 500)```                                     |
| 38      | Worker Pools                    | Managing a group of worker goroutines                                                         | ```go jobs := make(chan int, 100)```                                                          |
| 39      | WaitGroups                      | Waiting for a collection of goroutines to finish                                              | ```go var wg sync.WaitGroup```                                                                |
| 40      | Rate Limiting                   | Controlling how frequently events are allowed to happen                                       | ```go limiter := time.Tick(time.Millisecond * 200)```                                         |
| 41      | Atomic Counters                 | Using atomic operations for safe access to shared variables                                   | ```go var ops uint64 atomic.AddUint64(&ops, 1)```                                             |
| 42      | Mutexes                         | Ensuring exclusive access to shared resources                                                 | ```go var mu sync.Mutex mu.Lock()```                                                          |
| 43      | Stateful Goroutines             | Managing state in a goroutine                                                                 | ```go type readOp struct { key int resp chan int }```                                         |
| 44      | Sorting                         | Sorting built-in and custom types                                                             | ```go sort.Ints([]int{7, 2, 4})```                                                            |
| 45      | Sorting by Functions            | Custom sorting using functions                                                                | ```go sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })```     |
| 46      | Panic                           | Handling unexpected errors                                                                    | ```go panic("an error occurred")```                                                           |
| 47      | Defer                           | Ensuring a function call is performed later in the program’s execution                         | ```go defer fmt.Println("world") fmt.Println("hello")```                                      |
| 48      | Recover                         | Regaining control of a panicking goroutine                                                    | ```go defer func() { if r := recover(); r != nil { fmt.Println("Recovered", r) } }()```       |
| 49      | String Functions                | Common string operations                                                                      | ```go fmt.Println(strings.Contains("test", "es"))```                                          |
| 50      | String Formatting               | Formatting strings                                                                            | ```go fmt.Printf("pi: %v\n", 3.14)```                                                         |
| 51      | Text Templates                  | Creating and executing templates to generate formatted text                                   | ```go tmpl, _ := template.New("test").Parse("Value is {{.}}") tmpl.Execute(os.Stdout, "42")``` |
| 52      | Regular Expressions             | Matching patterns in strings                                                                  | ```go matched, _ := regexp.MatchString("p([a-z]+)ch", "peach")```                             |
| 53      | JSON                            | Encoding and decoding JSON                                                                    | ```go type Response struct { Page int } json.Marshal(&Response{Page: 1})```                   |
| 54      | XML                             | Encoding and decoding XML                                                                     | ```go type Plant struct { XMLName xml.Name `




| Pitfall                                | Description                                                                                       | Correct Approach                                                                                             |
|----------------------------------------|---------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------|
| Ignoring Errors                        | Ignoring returned errors can cause unexpected behavior and make debugging difficult.               | Always check and handle errors immediately. Use meaningful error messages.                                    |
| Resource Leaks                         | Not closing resources like file handles or HTTP response bodies can lead to resource exhaustion.   | Use `defer` to ensure resources are closed properly.                                                          |
| Panic and Recover Misuse                | Using `panic` for error handling can lead to crashes and difficult-to-debug programs.              | Use `panic` for truly exceptional situations. Use proper error handling for recoverable errors.            |
| Concurrency Issues                     | Improper use of goroutines and channels can lead to race conditions and deadlocks.                 | Use synchronization primitives (`sync.Mutex`, `sync.WaitGroup`, channels) to manage concurrency.            |
| Improper Use of `defer`                | Incorrect usage of `defer` can lead to unexpected behaviors and resource leaks.                    | Place `defer` statements right after the resource is acquired to ensure it gets executed.                  |
| Ignoring `recover`                     | Ignoring `recover` in deferred functions can cause the program to crash.                           | Use `recover` in deferred functions to handle panics gracefully.                                             |
| Hardcoding Values                      | Hardcoding configuration values makes the code less flexible and harder to maintain.               | Use configuration files, environment variables, or flags for configuration.                                 |
| Inconsistent Naming                    | Inconsistent naming can make the code harder to read and understand.                               | Follow Go naming conventions: camelCase for variables and functions, initialisms should be uppercase.       |
| Global Variables                       | Using global variables can lead to unexpected behaviors and make the code less testable.           | Avoid global variables. Use dependency injection to pass required dependencies.                             |
| Not Using `context`                    | Ignoring the use of `context` for long-running operations can make it difficult to manage timeouts. | Use `context` to handle timeouts, cancellations, and deadlines.                                              |
| Ignoring Go Vet                        | Not using `go vet` can lead to unnoticed potential issues in the code.                             | Run `go vet` regularly to catch common mistakes and improve code quality.                                   |
| Not Using `go fmt`                     | Inconsistent formatting can make the code harder to read.                                          | Use `go fmt` to automatically format the code according to Go standards.                                    |
| Using `interface{}` Unnecessarily      | Overusing `interface{}` can lead to type assertion issues and less type safety.                     | Use specific types wherever possible. Use `interface{}` only when absolutely necessary.                      |
| Misusing `sync.Pool`                   | Incorrectly using `sync.Pool` can lead to performance degradation.                                  | Use `sync.Pool` for temporary objects that are frequently allocated and deallocated. Ensure objects are reset before reuse. |





    1. Ignoring Errors
    
    Pitfall:
    
    go
    
    out, _ := os.Create(fileName)
    
    Correct:
    
    go
    
    out, err := os.Create(fileName)
    if err != nil {
        return fmt.Errorf("failed to create file %s: %w", fileName, err)
    }
    
    2. Resource Leaks
    
    Pitfall:
    
    go
    
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    // resp.Body.Close() is missing
    _, err = io.Copy(out, resp.Body)
    
    Correct:
    
    go
    
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }
    
    3. Panic and Recover Misuse
    
    Pitfall:
    
    go
    
    if err != nil {
        panic(err)
    }
    
    Correct:
    
    go
    
    if err != nil {
        log.Fatalf("Fatal error: %s", err)
    }
    
    4. Concurrency Issues
    
    Pitfall:
    
    go
    
    var counter int
    go func() {
        counter++
    }()
    
    Correct:
    
    go
    
    var mu sync.Mutex
    var counter int
    go func() {
        mu.Lock()
        defer mu.Unlock()
        counter++
    }()
    
    5. Improper Use of defer
    
    Pitfall:
    
    go
    
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    processFile(f)
    f.Close() // If processFile panics, this won't be called
    
    Correct:
    
    go
    
    f, err := os.Open(file)
    if err != nil {
        return err
    }
    defer f.Close()
    processFile(f)
    
    6. Ignoring recover
    
    Pitfall:
    
    go
    
    func mightPanic() {
        panic("something went wrong")
    }
    
    Correct:
    
    go
    
    func safeCall() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Recovered from panic: %v", r)
            }
        }()
        mightPanic()
    }
    
    7. Hardcoding Values
    
    Pitfall:
    
    go
    
    port := "8080"
    
    Correct:
    
    go
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    8. Inconsistent Naming
    
    Pitfall:
    
    go
    
    func Fetch_data() {}
    
    Correct:
    
    go
    
    func fetchData() {}
    
    9. Global Variables
    
    Pitfall:
    
    go
    
    var config Config
    
    Correct:
    
    go
    
    func main() {
        config := loadConfig()
        runApp(config)
    }
    
    10. Not Using context
    
    Pitfall:
    
    go
    
    func fetchData() error {
        // long running operation
    }
    
    Correct:
    
    go
    
    func fetchData(ctx context.Context) error {
        // use ctx for cancellation and timeouts
    }
    
    11. Ignoring Go Vet
    
    Pitfall:
    
    go
    
    // Not running go vet
    
    Correct:
    
    sh
    
    go vet ./...
    
    12. Not Using go fmt
    
    Pitfall:
    
    go
    
    func main(){fmt.Println("Hello, world!")}
    
    Correct:
    
    go
    
    func main() {
        fmt.Println("Hello, world!")
    }
    
    13. Using interface{} Unnecessarily
    
    Pitfall:
    
    go
    
    func process(data interface{}) error {
        // process data
    }
    
    Correct:
    
    go
    
    func process(data []byte) error {
        // process data
    }
    
    14. Misusing sync.Pool
    
    Pitfall:
    
    go
    
    var pool = sync.Pool{
        New: func() interface{} {
            return new(bytes.Buffer)
        },
    }
    // Not resetting the buffer before putting it back
    
    Correct:
    
    go
    
    var bufPool = sync.Pool{
        New: func() interface{} {
            return new(bytes.Buffer)
        },
    }
    buf := bufPool.Get().(*bytes.Buffer)
    buf.Reset() // Reset before using
    defer bufPool.Put(buf)
    // Use buf
    
        enter code here
