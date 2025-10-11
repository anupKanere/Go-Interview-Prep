# Go Coding Templates

## 📁 Template Files

### `problem_template.go`
A comprehensive template for solving coding problems with:
- Problem description and constraints
- Multiple solution approaches
- Test cases with table-driven tests
- Benchmarking for performance testing
- Example usage in main function

### `go.mod`
Go module file for dependency management

## 🚀 Usage

1. Copy `problem_template.go` to your problem directory
2. Fill in the problem details
3. Implement your solution(s)
4. Add comprehensive test cases
5. Run tests: `go test -v`
6. Run benchmarks: `go test -bench=.`

## 📝 Best Practices

### Code Style
- Use meaningful variable names
- Add comments for complex logic
- Follow Go naming conventions
- Keep functions small and focused

### Testing
- Write test cases for edge cases
- Test with empty inputs
- Test with single element
- Test with maximum constraints

### Performance
- Consider time and space complexity
- Use benchmarks to measure performance
- Optimize based on constraints
- Document complexity in comments

### Problem Solving
1. **Understand**: Read problem carefully
2. **Plan**: Think about approach before coding
3. **Code**: Implement solution
4. **Test**: Verify with test cases
5. **Optimize**: Improve if needed
6. **Explain**: Be ready to explain your approach

## 🎯 Interview Tips

### Before Coding
- Ask clarifying questions
- Discuss approach with interviewer
- Consider edge cases
- Estimate time/space complexity

### While Coding
- Think out loud
- Write clean, readable code
- Handle edge cases
- Test your solution

### After Coding
- Walk through your solution
- Discuss time/space complexity
- Mention potential optimizations
- Be ready for follow-up questions

## 📚 Common Patterns

### Two Pointers
```go
left, right := 0, len(nums)-1
for left < right {
    // Process elements
    left++
    right--
}
```

### Sliding Window
```go
left := 0
for right := 0; right < len(nums); right++ {
    // Expand window
    for windowIsInvalid() {
        // Shrink window
        left++
    }
}
```

### Hash Map
```go
freq := make(map[int]int)
for _, num := range nums {
    freq[num]++
}
```

### Binary Search
```go
left, right := 0, len(nums)-1
for left <= right {
    mid := left + (right-left)/2
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        left = mid + 1
    } else {
        right = mid - 1
    }
}
```

Happy coding! 🚀
