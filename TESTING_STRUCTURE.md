# 🧪 Proper Go Testing Structure

## ❌ **Previous Issue: Mixed Test Cases**

The original structure had test cases mixed in the main files, which caused several problems:

### **Problems with Mixed Tests:**
1. **Multiple `main` functions** - Go can't have multiple main functions in the same package
2. **Not following Go conventions** - Tests should be in separate `_test.go` files
3. **Hard to run individual tests** - Can't easily test specific functions
4. **Poor separation of concerns** - Production code mixed with test code
5. **Build failures** - Compilation errors due to conflicting main functions

## ✅ **Proper Go Testing Structure**

### **1. Separate Test Files**
```
01-arrays-strings/
├── two-sum/
│   ├── main.go          # Production code
│   └── main_test.go     # Test code
├── contains-duplicate/
│   ├── main.go          # Production code
│   └── main_test.go     # Test code
└── valid-anagram/
    ├── main.go          # Production code
    └── main_test.go     # Test code
```

### **2. Go Testing Conventions**

#### **Test File Naming**
- Test files must end with `_test.go`
- Test functions must start with `Test`
- Benchmark functions must start with `Benchmark`

#### **Example Structure**
```go
// main.go - Production code
package main

func twoSum(nums []int, target int) []int {
    // Implementation
}

func main() {
    // Example usage
}

// main_test.go - Test code
package main

import "testing"

func TestTwoSum(t *testing.T) {
    // Test cases
}

func BenchmarkTwoSum(b *testing.B) {
    // Benchmark code
}
```

### **3. Benefits of Proper Structure**

#### **Separation of Concerns**
- **Production code** in `main.go`
- **Test code** in `main_test.go`
- **Clear boundaries** between code and tests

#### **Go Tooling Support**
- `go test` works properly
- `go test -bench=.` for benchmarks
- `go test -cover` for coverage
- `go test -v` for verbose output

#### **Individual Testing**
- Can test specific functions
- Can run specific test cases
- Can benchmark individual solutions

#### **No Conflicts**
- No multiple main functions
- No import conflicts
- Clean compilation

## 🚀 **Running Tests**

### **Individual Problem Testing**
```bash
# Test specific problem
go test -v ./01-arrays-strings/two-sum/

# Test with benchmarks
go test -bench=. ./01-arrays-strings/two-sum/

# Test with coverage
go test -cover ./01-arrays-strings/two-sum/
```

### **Directory Testing**
```bash
# Test all problems in directory
go test -v ./01-arrays-strings/...

# Test with benchmarks
go test -bench=. ./01-arrays-strings/...

# Test with coverage
go test -cover ./01-arrays-strings/...
```

### **Using Test Runner**
```bash
# Test specific directory
go run tests/test_runner.go test-dir 01-arrays-strings

# Test all directories
go run tests/test_runner.go test-all

# Run benchmarks
go run tests/test_runner.go benchmark

# Run coverage
go run tests/test_runner.go coverage
```

## 📊 **Test Structure Best Practices**

### **1. Test Organization**
```go
func TestFunctionName(t *testing.T) {
    testCases := []struct {
        name     string
        input    InputType
        expected OutputType
    }{
        {
            name:     "Example 1",
            input:    exampleInput,
            expected: expectedOutput,
        },
        // More test cases...
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := functionName(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %v, got %v", tc.expected, result)
            }
        })
    }
}
```

### **2. Benchmark Organization**
```go
func BenchmarkFunctionName(b *testing.B) {
    input := createTestInput()
    for i := 0; i < b.N; i++ {
        functionName(input)
    }
}
```

### **3. Edge Case Testing**
- Empty inputs
- Single elements
- Maximum constraints
- Invalid inputs
- Boundary conditions

### **4. Multiple Solutions Testing**
```go
func TestMultipleSolutions(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        // Test cases...
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Test solution 1
            result1 := solution1(tc.input)
            if result1 != tc.expected {
                t.Errorf("Solution 1: Expected %v, got %v", tc.expected, result1)
            }
            
            // Test solution 2
            result2 := solution2(tc.input)
            if result2 != tc.expected {
                t.Errorf("Solution 2: Expected %v, got %v", tc.expected, result2)
            }
        })
    }
}
```

## 🎯 **Interview Practice Benefits**

### **1. Professional Code Structure**
- Follows Go best practices
- Clean separation of concerns
- Easy to maintain and extend

### **2. Testing Skills**
- Can write comprehensive tests
- Understands test-driven development
- Knows how to test edge cases

### **3. Performance Awareness**
- Can benchmark solutions
- Understands time/space complexity
- Can optimize based on benchmarks

### **4. Interview Readiness**
- Can explain test strategy
- Can demonstrate testing skills
- Shows attention to code quality

## 📝 **Template Usage**

### **Creating New Problems**
1. Copy template: `cp templates/problem_template.go new_problem/main.go`
2. Create test file: `cp templates/problem_template.go new_problem/main_test.go`
3. Remove main function from test file
4. Implement solution and tests
5. Run tests: `go test -v ./new_problem/`

### **Best Practices**
- Use descriptive test names
- Test all edge cases
- Include benchmarks
- Document time/space complexity
- Handle errors properly

## 🚀 **Next Steps**

1. **Follow the proper structure** for all new problems
2. **Use the test runner** for automated testing
3. **Practice writing tests** for each solution
4. **Benchmark your solutions** to compare performance
5. **Test edge cases** thoroughly

This structure ensures your code is professional, testable, and ready for FAANG interviews! 🎯
