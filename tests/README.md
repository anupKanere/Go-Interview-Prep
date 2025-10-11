# Testing Framework

## 🧪 Test Runner

The `test_runner.go` utility helps you run tests across all practice directories efficiently.

## 🚀 Usage

### Run All Tests
```bash
go run tests/test_runner.go test-all
```

### Run Tests in Specific Directory
```bash
go run tests/test_runner.go test-dir 01-arrays-strings
```

### Run Benchmarks
```bash
go run tests/test_runner.go benchmark
```

### Run Coverage
```bash
go run tests/test_runner.go coverage
```

## 📊 Test Structure

Each problem file should include:
- **Test cases** with table-driven tests
- **Benchmarks** for performance testing
- **Edge cases** for comprehensive coverage
- **Multiple solutions** for comparison

## 🎯 Best Practices

### Writing Tests
1. **Test all edge cases**: Empty inputs, single elements, maximum constraints
2. **Use table-driven tests**: Easy to add new test cases
3. **Test multiple solutions**: Compare different approaches
4. **Include benchmarks**: Measure performance improvements
5. **Test error conditions**: Handle invalid inputs gracefully

### Test Naming
- Use descriptive test names
- Include the scenario being tested
- Group related tests together
- Use subtests for complex scenarios

### Example Test Structure
```go
func TestSolution(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        {
            name:     "Example 1",
            input:    []int{1, 2, 3},
            expected: 6,
        },
        // Add more test cases
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := solution(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %d, got %d", tc.expected, result)
            }
        })
    }
}
```

## 📈 Performance Testing

### Benchmarks
```go
func BenchmarkSolution(b *testing.B) {
    input := []int{1, 2, 3, 4, 5}
    for i := 0; i < b.N; i++ {
        solution(input)
    }
}
```

### Memory Testing
```go
func BenchmarkSolution(b *testing.B) {
    b.ReportAllocs()
    // Your benchmark code
}
```

## 🔍 Coverage Goals

- **Line Coverage**: > 90%
- **Branch Coverage**: > 80%
- **Function Coverage**: 100%

## 🚀 Running Tests

### Individual File
```bash
go test -v 01-arrays-strings/two_sum.go
```

### All Files in Directory
```bash
go test -v ./01-arrays-strings/...
```

### With Coverage
```bash
go test -cover ./01-arrays-strings/...
```

### With Benchmarks
```bash
go test -bench=. ./01-arrays-strings/...
```

## 📝 Test Checklist

- [ ] All edge cases covered
- [ ] Multiple solutions tested
- [ ] Benchmarks included
- [ ] Error conditions handled
- [ ] Test names are descriptive
- [ ] Table-driven tests used
- [ ] Coverage > 90%

## 🎯 Interview Practice

### Test-Driven Development
1. Write test cases first
2. Implement solution
3. Refactor and optimize
4. Add more test cases
5. Repeat

### Mock Interview Practice
- Explain your test strategy
- Discuss edge cases
- Show test coverage
- Demonstrate debugging skills

Happy testing! 🚀
