# Go Learning Notes

A personal reference for learning Go — clear notes, runnable examples, and ASCII illustrations covering core concepts through REST APIs and databases.

## Table of Contents

1. [Introduction](#introduction)
2. [Quick Start](#quick-start)
3. [Syntax & Hello World](#syntax--hello-world)
4. [Variables & Constants](#variables--constants)
5. [Pointers & Addresses](#pointers--addresses)
6. [Output Formatting](#output-formatting-fmt)
7. [Data Types](#data-types)
8. [Arrays & Slices](#arrays--slices)
9. [Maps](#maps)
10. [Operators](#operators)
11. [Conditions & Switch](#conditions--switch)
12. [Loops](#loops)
13. [Functions](#functions)
14. [Structs](#structs)
15. [Goroutines & Concurrency](#goroutines--concurrency)
16. [Working with Databases (MySQL)](#working-with-databases-mysql)
17. [REST API with Gin](#rest-api-with-gin)
18. [Key Rules & Gotchas](#key-rules--gotchas)
19. [References](#references)

---

## Introduction

Go (Golang) is a statically typed, compiled language built by Google. It prioritises:

- **Simplicity** — minimal keywords, clean syntax
- **Concurrency** — goroutines and channels are first-class citizens
- **Performance** — compiles directly to native machine code
- **Scalability** — designed for multi-core, distributed systems

Common use cases: web servers, microservices, CLI tools, cloud-native applications.

---

## Quick Start

```bash
# Check installed Go version
go version

# Run a single file
go run first.go

# Run the booking app
cd booking-app && go run main.go

# Download a dependency
go get github.com/gin-gonic/gin
```

---

## Syntax & Hello World

Every Go program requires:
- A **package declaration** — executable programs use `package main`
- **Imports** — packages the file uses
- A **`main` function** — the program's entry point

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

> **Note:** Statements end at a newline — no semicolons needed (the compiler inserts them). The opening brace `{` must be on the same line as the `func`, `if`, `for`, etc.

---

## Variables & Constants

### Declaration styles

```go
// Explicit type
var age int = 30

// Type inferred from the assigned value
var name = "Alice"

// Short declaration — only valid inside a function
city := "Nairobi"

// Constants — value can never change after declaration
const Pi      = 3.14159
const AppName = "MyApp"
```

### Multiple declarations

```go
// Single line
var a, b, c int = 1, 2, 3

// Block form — cleaner when declaring several variables at once
var (
    host  string  = "localhost"
    port  int     = 8080
    debug bool    = false
)
```

### Naming conventions

| Style      | Used for                        | Example          |
|------------|---------------------------------|------------------|
| camelCase  | local variables, parameters     | `userCount`      |
| PascalCase | exported (public) identifiers   | `UserService`    |
| UPPER_CASE | constants (community convention)| `MAX_RETRIES`    |

> **Zero values:** Uninitialized variables receive their type's zero value — `0` for numbers, `""` for strings, `false` for booleans, `nil` for pointers/maps/slices.

> **Rule:** A variable name cannot start with a digit, contain spaces, or shadow a Go keyword (`func`, `if`, `for`, `var`, …).

---

## Pointers & Addresses

A **pointer** stores the memory address of another variable, not its value.

```
  Variable v                   Pointer p
  ┌──────────────┐             ┌──────────────┐
  │  value: 5    │◄────────────│ addr: 0xc000 │
  │  addr: 0xc000│             └──────────────┘
  └──────────────┘
        ▲
        └── *p dereferences p and gives you the value here
```

| Operator | Meaning                     | Example              |
|----------|-----------------------------|----------------------|
| `&v`     | address-of variable `v`     | `p := &v`            |
| `*p`     | dereference — value at `p`  | `fmt.Println(*p)`    |

```go
func increment(n *int) {
    *n++  // modifies the original variable through the pointer
}

func main() {
    count := 5
    increment(&count)
    fmt.Println(count) // 6
}
```

Pointers are useful when:
- You need to modify a value inside a function without returning it
- You want to avoid copying a large struct
- Reading user input: `fmt.Scan(&variableName)`

---

## Output Formatting (`fmt`)

```go
name := "Alice"
age  := 30
pi   := 3.14159

fmt.Print("no newline at end")
fmt.Println("appends a newline")
fmt.Printf("name=%s  age=%d  pi=%.2f\n", name, age, pi)
// name=Alice  age=30  pi=3.14
```

### Format verbs reference

| Verb   | Output                        | Example input → output        |
|--------|-------------------------------|-------------------------------|
| `%v`   | default format                | `42` → `42`                   |
| `%#v`  | Go-syntax representation      | `42` → `42`                   |
| `%T`   | type of the value             | `42` → `int`                  |
| `%d`   | integer (decimal)             | `42` → `42`                   |
| `%b`   | integer (binary)              | `5` → `101`                   |
| `%f`   | float                         | `3.14` → `3.140000`           |
| `%.2f` | float, 2 decimal places       | `3.14159` → `3.14`            |
| `%s`   | string                        | `"hi"` → `hi`                 |
| `%q`   | quoted string                 | `"hi"` → `"hi"`               |
| `%t`   | boolean                       | `true` → `true`               |
| `%%`   | literal percent sign          | → `%`                         |

---

## Data Types

### Integers

| Go type  | Java equivalent  | Signed? | Range (approx)            |
|----------|------------------|---------|---------------------------|
| `int8`   | `byte`           | yes     | -128 to 127               |
| `int16`  | `short`          | yes     | -32,768 to 32,767         |
| `int32`  | `int`            | yes     | ±2 billion                |
| `int64`  | `long`           | yes     | ±9.2 quintillion          |
| `uint8`  | unsigned byte    | no      | 0 to 255                  |
| `uint32` | —                | no      | 0 to 4.3 billion          |
| `int`    | —                | yes     | 32- or 64-bit (platform)  |

### Floats

```go
var price float32 = 9.99          // 32-bit, ~6-7 significant digits
var pi    float64 = 3.14159265    // 64-bit, ~15-16 significant digits — prefer this
```

### Booleans & Strings

```go
var active bool   = true
var greeting string = "Hello, Go!"
fmt.Println(len(greeting)) // 10 — length in bytes
```

---

## Arrays & Slices

### Arrays — fixed size, same type

```go
var primes = [5]int{2, 3, 5, 7, 11}
colors    := [3]string{"red", "green", "blue"}
inferred  := [...]int{1, 2, 3, 4}  // compiler counts the length

fmt.Println(primes[0])  // 2   (zero-indexed)
primes[0] = 100          // update in place
fmt.Println(len(primes)) // 5
```

### Slices — dynamic, flexible

A slice is a **window into a backing array**. It carries three pieces of information:

```
Backing array:  [10][11][12][13][14][15]
                  0   1   2   3   4   5

mySlice := arr[2:4]
                      ┌────────────┐
                      │ 12 │  13  │    len = 2
                      └────────────┘
               cap = 4  (index 2 up to end of array)
```

```go
arr := [6]int{10, 11, 12, 13, 14, 15}
s   := arr[2:4]

fmt.Println(s)       // [12 13]
fmt.Println(len(s))  // 2
fmt.Println(cap(s))  // 4
```

### Creating slices from scratch

```go
// Literal
nums := []int{1, 2, 3, 4, 5}

// make(type, length, capacity)
s := make([]int, 5, 10)
// s = [0 0 0 0 0],  len=5, cap=10
```

### Appending to a slice

```go
s := []int{1, 2, 3}
s  = append(s, 4, 5)               // add individual values
s  = append(s, []int{6, 7}...)     // spread another slice

// When len reaches cap, Go allocates a new (larger) backing array
```

```
Before (cap=3):  [1][2][3]
After  (cap=6):  [1][2][3][4][5]   ← new backing array allocated
```

> Always reassign: `s = append(s, x)` — `append` may return a new slice header.

### Copying a slice (avoids holding a large array in memory)

```go
large := make([]int, 1_000_000)
small := large[0:3]   // still references the million-element array!

dest := make([]int, 3)
copy(dest, small)     // new small backing array — large can be GC'd
```

### Slice function cheat sheet

| Expression        | Description                              |
|-------------------|------------------------------------------|
| `len(s)`          | number of elements currently in slice    |
| `cap(s)`          | max elements before reallocation needed  |
| `append(s, x)`    | returns slice with x added               |
| `copy(dst, src)`  | copies min(len(dst), len(src)) elements  |
| `s[low:high]`     | sub-slice — low inclusive, high exclusive|

---

## Maps

Maps are unordered key-value collections backed by a hash table.

```
map[string]string
┌─────────┬──────────┐
│   Key   │  Value   │
├─────────┼──────────┤
│ "brand" │ "Ford"   │
│ "model" │ "Mustang"│
│ "year"  │ "1964"   │
└─────────┴──────────┘
Order is NOT guaranteed when iterating
```

### Creating maps

```go
// Map literal
cars := map[string]string{
    "brand": "Ford",
    "model": "Mustang",
    "year":  "1964",
}

// make — start empty, add entries later
cities := make(map[string]int)
cities["Oslo"]   = 1
cities["Bergen"] = 2
```

### Accessing elements

```go
brand := cars["brand"]  // "Ford"

// Safe access — always check existence on lookup
value, ok := cars["color"]
if !ok {
    fmt.Println("key not found")
}
```

### Modifying maps

```go
cars["year"]  = "1970"       // update existing key
cars["color"] = "red"        // insert new key
delete(cars, "model")        // remove a key
fmt.Println(len(cars))       // count keys
```

### Valid key types

Keys must support the `==` operator.

| Valid key types              | Invalid key types |
|------------------------------|-------------------|
| bool, int, float, string     | slice             |
| array, pointer, struct       | map               |
| interface (if dynamic type supports ==) | func |

---

## Operators

### Arithmetic

| Op  | Name           | Example       |
|-----|----------------|---------------|
| `+` | Addition       | `3 + 2 = 5`   |
| `-` | Subtraction    | `5 - 2 = 3`   |
| `*` | Multiplication | `3 * 2 = 6`   |
| `/` | Division       | `7 / 2 = 3`   |
| `%` | Modulus        | `7 % 3 = 1`   |
| `++`| Increment      | `x++` adds 1  |
| `--`| Decrement      | `x--` subtracts 1 |

### Assignment

| Op   | Equivalent to   |
|------|-----------------|
| `=`  | assign          |
| `+=` | `x = x + n`    |
| `-=` | `x = x - n`    |
| `*=` | `x = x * n`    |
| `/=` | `x = x / n`    |
| `%=` | `x = x % n`    |

### Comparison & Logical

| Op   | Meaning                    |
|------|----------------------------|
| `==` | equal                      |
| `!=` | not equal                  |
| `<`  | less than                  |
| `>`  | greater than               |
| `<=` | less than or equal         |
| `>=` | greater than or equal      |
| `&&` | AND (short-circuits on false) |
| `\|\|`| OR (short-circuits on true)  |
| `!`  | NOT                        |

---

## Conditions & Switch

```go
score := 85

if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

### Switch statement

```go
day := "Monday"

switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown day")
}
```

> Unlike C/Java, Go's `switch` does **not fall through** by default — no `break` needed after each case. Use the explicit `fallthrough` keyword only when you need the next case to execute.

---

## Loops

Go has exactly one loop construct: `for`. It covers every iteration pattern.

### Classic C-style loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
// 0 1 2 3 4
```

### While-style loop (condition only)

```go
n := 1
for n < 100 {
    n *= 2
}
fmt.Println(n) // 128
```

### Range loop — iterate over slices, arrays, maps

```go
fruits := []string{"apple", "banana", "cherry"}

for idx, fruit := range fruits {
    fmt.Printf("%d: %s\n", idx, fruit)
}

// Omit index or value with _
for _, fruit := range fruits {
    fmt.Println(fruit)
}

// Map iteration
scores := map[string]int{"Alice": 90, "Bob": 75}
for name, score := range scores {
    fmt.Printf("%s scored %d\n", name, score)
}
```

### Loop control: continue & break

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue  // skip i==3, move to next iteration
    }
    if i == 7 {
        break     // exit the loop entirely
    }
    fmt.Println(i)
}
// Output: 0 1 2 4 5 6
```

### Nested loops

```go
adj    := []string{"big", "tasty"}
fruits := []string{"apple", "orange", "banana"}

for _, a := range adj {
    for _, f := range fruits {
        fmt.Printf("%s %s\n", a, f)
    }
}
// big apple, big orange, big banana, tasty apple, tasty orange, tasty banana
```

---

## Functions

### Basic function

```go
func greet(name string) string {
    return "Hello, " + name
}

fmt.Println(greet("Alice")) // Hello, Alice
```

### Multiple return values

A distinctive Go feature — functions can return multiple values. Use this for returning a result alongside an error.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 3)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%.2f\n", result) // 3.33
```

### Named return values

Name the return variables in the signature. A bare `return` sends them back.

```go
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums {
        if n < min { min = n }
        if n > max { max = n }
    }
    return  // returns min and max
}
```

### Ignoring a return value

```go
result, _ := divide(10, 3)  // discard the error with _
```

### Recursion

A function that calls itself until a base condition is reached.

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5)) // 120
```

```
Call stack for factorial(4):
  factorial(4)
    └─ 4 × factorial(3)
             └─ 3 × factorial(2)
                      └─ 2 × factorial(1)
                               └─ 1  (base case, returns 1)
                      = 2 × 1 = 2
             = 3 × 2 = 6
    = 4 × 6 = 24
```

---

## Structs

Structs group related fields of different types into a single named type — similar to a class in other languages, but without inheritance.

```go
type Employee struct {
    Name   string
    Age    int
    Job    string
    Salary float64
}
```

### Creating and accessing instances

```go
// Named-field literal (recommended — order-independent)
emp := Employee{
    Name:   "Jane Doe",
    Age:    32,
    Job:    "Engineer",
    Salary: 85000,
}

fmt.Println(emp.Name)  // Jane Doe
emp.Salary = 90000      // update a field with dot notation
```

### Passing structs to functions

```go
// By value — function receives a copy, original is unchanged
func display(e Employee) {
    fmt.Printf("%s (%d) — %s\n", e.Name, e.Age, e.Job)
}

// By pointer — function can modify the original
func giveRaise(e *Employee, amount float64) {
    e.Salary += amount
}

giveRaise(&emp, 5000)
fmt.Println(emp.Salary) // 95000
```

```
Pass by value:                Pass by pointer:
┌──────────────┐              ┌──────────────┐
│  emp (orig)  │              │  emp (orig)  │◄─── modified here
└──────────────┘              └──────────────┘
       │ copied                       ▲
       ▼                             │ &emp
┌──────────────┐              ┌──────────────┐
│  e  (copy)   │              │  e  (*ptr)   │
└──────────────┘              └──────────────┘
```

### Struct tags (used for JSON / database mapping)

```go
type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

The backtick tags tell Go's JSON encoder/decoder what field names to use in the serialised output. Without tags, Go uses the struct field name as-is.

---

## Goroutines & Concurrency

Go was designed for concurrency from the ground up. A **goroutine** is a lightweight thread managed by the Go runtime — thousands can run simultaneously with minimal memory overhead.

```
Sequential:                        Concurrent (goroutines):
──────────────────────────         ──────────────────────────────
  Task A starts                      goroutine 1: Task A ──────►
  Task A finishes                    goroutine 2: Task B ──────►
  Task B starts                      goroutine 3: Task C ──────►
  Task B finishes                    all running at the same time
  Task C starts
  Task C finishes
  (total: A + B + C time)            (total: ~max(A, B, C) time)
```

### Starting a goroutine

```go
go myFunction()        // run myFunction concurrently
go func() {            // anonymous goroutine
    fmt.Println("running in background")
}()
```

### Coordinating with `sync.WaitGroup`

Without coordination, `main` exits before goroutines finish. `WaitGroup` fixes this.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func sendTicket(name string) {
    defer wg.Done()  // decrement counter when this function returns
    time.Sleep(2 * time.Second)
    fmt.Printf("Ticket sent to %s\n", name)
}

func main() {
    names := []string{"Alice", "Bob", "Carol"}

    for _, name := range names {
        wg.Add(1)            // increment counter for each goroutine
        go sendTicket(name)
    }

    wg.Wait()  // block here until counter reaches 0
    fmt.Println("All tickets sent")
}
```

```
Timeline:
main:          ─── Add×3 ─── Wait() ──────────────────────────► Done
goroutine 1:   ─── sendTicket("Alice") ────────────────────►
goroutine 2:     ─── sendTicket("Bob")  ────────────────────►
goroutine 3:       ─── sendTicket("Carol")──────────────────►
                                           ↑ all call Done() → Wait() unblocks
```

> See [booking-app/main.go](booking-app/main.go) for a real-world example.

---

## Working with Databases (MySQL)

Source: [goMysqlTemplate.go](goMysqlTemplate.go)

Uses:
- `database/sql` — Go's standard database interface
- `github.com/go-sql-driver/mysql` — MySQL-specific driver

### Connection setup

```go
cfg := mysql.NewConfig()
cfg.User   = os.Getenv("DBUSER")   // read from env, never hardcode credentials
cfg.Passwd = os.Getenv("DBPASS")
cfg.Net    = "tcp"
cfg.Addr   = "127.0.0.1:3307"
cfg.DBName = "recordings"

db, err = sql.Open("mysql", cfg.FormatDSN())
if err != nil { log.Fatal(err) }

// sql.Open only validates args — it does NOT connect yet
// db.Ping() makes the actual connection and verifies it works
if err = db.Ping(); err != nil { log.Fatal(err) }
fmt.Println("Connected!")
```

### Querying multiple rows

```go
func albumsByArtist(name string) ([]Album, error) {
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil { return nil, err }
    defer rows.Close()  // always close when done — releases the connection

    var albums []Album
    for rows.Next() {
        var a Album
        if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
            return nil, err
        }
        albums = append(albums, a)
    }
    return albums, rows.Err()  // check for errors that occurred during iteration
}
```

### Querying a single row

```go
func albumByID(id int64) (Album, error) {
    var a Album
    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
    if err == sql.ErrNoRows {
        return a, fmt.Errorf("no album with id %d", id)
    }
    return a, err
}
```

### Inserting a row

```go
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec(
        "INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
        alb.Title, alb.Artist, alb.Price,
    )
    if err != nil { return 0, err }
    return result.LastInsertId()  // ID assigned by the database
}
```

### SQL method summary

| Method           | Use for                       | Returns        |
|------------------|-------------------------------|----------------|
| `db.Query()`     | SELECT — multiple rows        | `*sql.Rows`    |
| `db.QueryRow()`  | SELECT — at most one row      | `*sql.Row`     |
| `db.Exec()`      | INSERT / UPDATE / DELETE      | `sql.Result`   |
| `rows.Scan()`    | map columns into Go variables | `error`        |
| `rows.Close()`   | release the connection        | `error`        |

> Always use `?` placeholders for values. Never concatenate user input into queries — that is a SQL injection vulnerability.

---

## REST API with Gin

Source: [restApiTemplate.go](restApiTemplate.go)

Uses `github.com/gin-gonic/gin` — a fast HTTP framework for Go.

### Request flow

```
Client                     Gin Router               Handler Function
  │                             │                         │
  ├── GET /albums ──────────────►── getAlbums(c) ─────────►│
  │                             │                         │  c.IndentedJSON(200, albums)
  │◄── 200 JSON [{...},{...}] ──│◄────────────────────────┤
  │                             │                         │
  ├── GET /albums/2 ────────────►── getAlbumByID(c) ──────►│
  │                             │   c.Param("id") = "2"   │
  │◄── 200 JSON {id:"2",...} ───│◄────────────────────────┤
  │                             │                         │
  └── POST /albums ─────────────►── postAlbums(c) ────────►│
      Body: {"id":"4",...}      │   c.BindJSON(&album)     │
  ◄── 201 JSON {id:"4",...} ────│◄────────────────────────┘
```

### Data model

```go
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

### Router setup

```go
func main() {
    router := gin.Default()  // Default includes Logger and Recovery middleware

    router.GET("/albums",     getAlbums)     // list all
    router.GET("/albums/:id", getAlbumByID)  // get one — :id is a URL parameter
    router.POST("/albums",    postAlbums)    // create one

    router.Run("localhost:8080")
}
```

### Handler: GET all albums

```go
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)  // 200 + pretty-printed JSON body
}
```

```bash
curl http://localhost:8080/albums
```

### Handler: GET album by ID

```go
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")  // extracts the :id segment from the URL path

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
    // gin.H is shorthand for map[string]any — useful for one-off JSON objects
}
```

```bash
curl http://localhost:8080/albums/1
```

### Handler: POST — create an album

```go
func postAlbums(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        return  // Gin automatically sends 400 Bad Request on bind failure
    }
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)  // 201 Created
}
```

```bash
curl -X POST http://localhost:8080/albums \
  -H "Content-Type: application/json" \
  -d '{"id":"4","title":"Somethin Else","artist":"Cannonball Adderley","price":17.99}'
```

### HTTP status codes used

| Code | Constant                | Meaning                        |
|------|-------------------------|--------------------------------|
| 200  | `http.StatusOK`         | Success — body contains result |
| 201  | `http.StatusCreated`    | Resource successfully created  |
| 400  | `http.StatusBadRequest` | Invalid input (Gin handles this automatically on `BindJSON` failure) |
| 404  | `http.StatusNotFound`   | Requested resource does not exist |

---

## Key Rules & Gotchas

| Rule | Detail |
|------|--------|
| Unused imports = compile error | Remove any import you aren't using |
| Unused variables = compile error | Go will not compile if a declared variable is never read |
| `:=` is for functions only | Cannot use short declaration at package level |
| Exported names are PascalCase | `UserService` is public; `userService` is package-private |
| No mixed types in a map | All keys must share one type; all values must share one type |
| Structs allow mixed types | Unlike maps, struct fields can each be a different type |
| Always reassign after `append` | `s = append(s, x)` — the returned slice may have a new backing array |
| `sql.Open` does not connect | Call `db.Ping()` immediately after to verify the connection works |
| Goroutines need coordination | Without `WaitGroup` or channels, `main` may exit before goroutines finish |
| Use `?` placeholders in SQL | Never concatenate user input into queries — it opens SQL injection attacks |

---

## References

- [Official Go documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [A Tour of Go (interactive)](https://tour.golang.org/)
- [Gin web framework](https://gin-gonic.com/docs/)
- [database/sql tutorial](https://go.dev/doc/database/open-handle)
