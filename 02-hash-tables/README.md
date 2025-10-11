# Hash Tables Practice

## 📚 Topics Covered
- Hash table operations
- Map data structures
- Set operations
- Frequency counting
- Sliding window with hash tables

## 🎯 Learning Objectives
- Master Go map operations
- Understand hash table collision handling
- Learn frequency counting patterns
- Practice sliding window with hash tables
- Build intuition for hash table problems

## 📝 Practice Problems

### Week 1: Basic Hash Tables
1. **Two Sum** (Easy) - `two_sum.go`
2. **Contains Duplicate** (Easy) - `contains_duplicate.go`
3. **Valid Anagram** (Easy) - `valid_anagram.go`
4. **Group Anagrams** (Medium) - `group_anagrams.go`
5. **First Unique Character** (Easy) - `first_unique_char.go`

### Week 2: Advanced Hash Tables
1. **Longest Substring Without Repeating Characters** (Medium) - `longest_substring.go`
2. **Minimum Window Substring** (Hard) - `min_window_substring.go`
3. **Find All Anagrams** (Medium) - `find_anagrams.go`
4. **Subarray Sum Equals K** (Medium) - `subarray_sum_k.go`
5. **Longest Consecutive Sequence** (Medium) - `longest_consecutive.go`

## 🚀 Getting Started

1. Copy the template: `cp ../templates/problem_template.go group_anagrams.go`
2. Fill in the problem details
3. Implement your solution
4. Add test cases
5. Run tests: `go test -v`

## 📊 Progress Tracking

- [ ] Week 1: Basic Hash Tables (5 problems)
- [ ] Week 2: Advanced Hash Tables (5 problems)

## 🎯 Key Patterns to Master

### Frequency Counting
```go
freq := make(map[int]int)
for _, num := range nums {
    freq[num]++
}
```

### Two Sum Pattern
```go
numMap := make(map[int]int)
for i, num := range nums {
    complement := target - num
    if index, exists := numMap[complement]; exists {
        return []int{index, i}
    }
    numMap[num] = i
}
```

### Sliding Window with Hash
```go
charCount := make(map[byte]int)
left := 0
for right := 0; right < len(s); right++ {
    charCount[s[right]]++
    // Shrink window if needed
    for windowIsInvalid() {
        charCount[s[left]]--
        left++
    }
}
```

## 💡 Tips for Success

1. **Choose the right data structure**: Map vs Set vs Array
2. **Handle edge cases**: Empty inputs, single elements
3. **Optimize space**: Use arrays for limited character sets
4. **Practice explaining**: Be ready to explain hash table operations
5. **Time yourself**: Aim for 15-25 minutes per medium problem

## 🔗 Additional Resources

- [LeetCode Hash Table](https://leetcode.com/tag/hash-table/)
- [Go Maps Tutorial](https://golang.org/doc/effective_go.html#maps)
- [Go Maps Package](https://golang.org/pkg/maps/)

Happy coding! 🚀
