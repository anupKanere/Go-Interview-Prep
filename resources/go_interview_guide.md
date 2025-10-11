# Go Interview Preparation Guide

## 🎯 Go-Specific Interview Topics

### 📚 Core Go Concepts

#### 1. **Goroutines & Concurrency**
```go
// Goroutine basics
go func() {
    fmt.Println("Hello from goroutine!")
}()

// Channels
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch

// Select statement
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
default:
    fmt.Println("No message received")
}
```

#### 2. **Memory Management**
```go
// Pointers
var p *int
i := 42
p = &i
fmt.Println(*p) // 42

// Garbage collection
// Go has automatic garbage collection
// No manual memory management needed
```

#### 3. **Interfaces**
```go
type Writer interface {
    Write([]byte) (int, error)
}

type FileWriter struct {
    filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}
```

#### 4. **Error Handling**
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
```

### 🚀 Go Interview Questions

#### **Easy Level**
1. **What is a goroutine?**
   - Lightweight thread managed by Go runtime
   - Can run concurrently with other goroutines
   - Much cheaper than OS threads

2. **What is a channel?**
   - Communication mechanism between goroutines
   - Can be buffered or unbuffered
   - Provides synchronization

3. **What is the difference between `make` and `new`?**
   - `make`: Creates slices, maps, channels
   - `new`: Allocates memory, returns pointer

4. **What is a slice?**
   - Dynamic array in Go
   - Has length and capacity
   - Can grow dynamically

#### **Medium Level**
1. **How do you handle panics in Go?**
   ```go
   defer func() {
       if r := recover(); r != nil {
           fmt.Println("Recovered from panic:", r)
       }
   }()
   ```

2. **What is the difference between `var` and `:=`?**
   - `var`: Declares variable with zero value
   - `:=`: Declares and initializes variable

3. **How do you implement a worker pool?**
   ```go
   func workerPool(jobs <-chan int, results chan<- int) {
       for j := range jobs {
           results <- j * 2
       }
   }
   ```

4. **What is the difference between `sync.Mutex` and `sync.RWMutex`?**
   - `Mutex`: Exclusive lock
   - `RWMutex`: Multiple readers, single writer

#### **Hard Level**
1. **How do you implement a rate limiter?**
   ```go
   type RateLimiter struct {
       tokens chan struct{}
       ticker *time.Ticker
   }
   ```

2. **How do you implement a connection pool?**
   ```go
   type Pool struct {
       connections chan *Connection
       maxSize     int
   }
   ```

3. **How do you implement a circuit breaker?**
   ```go
   type CircuitBreaker struct {
       state       State
       failureCount int
       threshold   int
   }
   ```

### 🏗️ System Design with Go

#### **Microservices Architecture**
```go
// Service interface
type UserService interface {
    GetUser(id string) (*User, error)
    CreateUser(user *User) error
}

// HTTP handler
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.userService.GetUser(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

#### **Database Operations**
```go
// Repository pattern
type UserRepository interface {
    FindByID(id string) (*User, error)
    Save(user *User) error
}

// Implementation
type userRepository struct {
    db *sql.DB
}

func (r *userRepository) FindByID(id string) (*User, error) {
    row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
    user := &User{}
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    return user, err
}
```

#### **Caching**
```go
// Cache interface
type Cache interface {
    Get(key string) (interface{}, bool)
    Set(key string, value interface{}, ttl time.Duration)
    Delete(key string)
}

// Redis implementation
type redisCache struct {
    client *redis.Client
}
```

### 🧪 Testing in Go

#### **Unit Tests**
```go
func TestUserService(t *testing.T) {
    tests := []struct {
        name     string
        userID   string
        expected *User
        wantErr  bool
    }{
        {
            name:   "valid user",
            userID: "123",
            expected: &User{ID: "123", Name: "John"},
            wantErr: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            user, err := userService.GetUser(tt.userID)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(user, tt.expected) {
                t.Errorf("GetUser() = %v, want %v", user, tt.expected)
            }
        })
    }
}
```

#### **Integration Tests**
```go
func TestUserAPI(t *testing.T) {
    // Setup test server
    server := httptest.NewServer(http.HandlerFunc(userHandler))
    defer server.Close()

    // Test request
    resp, err := http.Get(server.URL + "/users/123")
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}
```

### 🚀 Performance Optimization

#### **Profiling**
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    // Your application code
}
```

#### **Benchmarking**
```go
func BenchmarkUserService(b *testing.B) {
    service := NewUserService()
    for i := 0; i < b.N; i++ {
        service.GetUser("123")
    }
}
```

#### **Memory Optimization**
```go
// Use object pools
var userPool = sync.Pool{
    New: func() interface{} {
        return &User{}
    },
}

func getUser() *User {
    user := userPool.Get().(*User)
    user.Reset()
    return user
}

func putUser(user *User) {
    userPool.Put(user)
}
```

### 🎯 Interview Tips

#### **Before the Interview**
1. **Review Go basics**: Goroutines, channels, interfaces
2. **Practice coding**: Use Go Playground
3. **Study Go patterns**: Error handling, concurrency
4. **Prepare examples**: Real-world Go projects

#### **During the Interview**
1. **Explain your approach**: Think out loud
2. **Handle errors**: Show proper error handling
3. **Use Go idioms**: Channels, goroutines, interfaces
4. **Discuss trade-offs**: Performance vs readability

#### **Common Go Interview Questions**
1. **How do you handle concurrent access to shared data?**
2. **What is the difference between a goroutine and a thread?**
3. **How do you implement a timeout in Go?**
4. **What is the difference between `sync.WaitGroup` and channels?**
5. **How do you implement a rate limiter in Go?**

### 📚 Go Resources

#### **Documentation**
- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Blog](https://blog.golang.org/)

#### **Books**
- **The Go Programming Language** by Donovan & Kernighan
- **Concurrency in Go** by Katherine Cox-Buday
- **Go in Action** by Manning Publications

#### **Practice**
- [Go Playground](https://play.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)

### 🎯 Success Metrics

#### **Technical Skills**
- [ ] Understand goroutines and channels
- [ ] Can implement interfaces
- [ ] Know error handling patterns
- [ ] Can write concurrent code
- [ ] Understand Go memory model

#### **Problem Solving**
- [ ] Can solve DSA problems in Go
- [ ] Can implement system design in Go
- [ ] Can write clean, idiomatic Go code
- [ ] Can handle edge cases and errors
- [ ] Can optimize Go code for performance

Good luck with your Go interview preparation! 🚀
