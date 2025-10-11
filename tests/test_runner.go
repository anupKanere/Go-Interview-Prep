package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
Test Runner for FAANG Interview Prep
This utility helps run tests across all practice directories
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run test_runner.go [command]")
		fmt.Println("Commands:")
		fmt.Println("  test-all     - Run all tests")
		fmt.Println("  test-dir <dir> - Run tests in specific directory")
		fmt.Println("  benchmark     - Run all benchmarks")
		fmt.Println("  coverage      - Run tests with coverage")
		fmt.Println("  help          - Show this help")
		return
	}

	command := os.Args[1]

	switch command {
	case "test-all":
		runAllTests()
	case "test-dir":
		if len(os.Args) < 3 {
			fmt.Println("Please specify directory: go run test_runner.go test-dir <directory>")
			return
		}
		runDirectoryTests(os.Args[2])
	case "benchmark":
		runBenchmarks()
	case "coverage":
		runCoverage()
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}

func runAllTests() {
	fmt.Println("🚀 Running all tests...")

	directories := []string{
		"01-arrays-strings",
		"02-hash-tables",
		"03-linked-lists",
		"04-stacks-queues",
		"05-trees",
		"06-heaps",
		"07-graphs",
		"08-dynamic-programming",
		"09-greedy",
		"10-system-design",
	}

	totalPassed := 0
	totalFailed := 0

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("⏭️  Skipping %s (directory not found)\n", dir)
			continue
		}

		fmt.Printf("\n📁 Testing %s...\n", dir)
		passed, failed := runTestsInDirectory(dir)
		totalPassed += passed
		totalFailed += failed
	}

	fmt.Printf("\n📊 Summary:\n")
	fmt.Printf("✅ Passed: %d\n", totalPassed)
	fmt.Printf("❌ Failed: %d\n", totalFailed)
	fmt.Printf("📈 Success Rate: %.1f%%\n", float64(totalPassed)/float64(totalPassed+totalFailed)*100)
}

func runDirectoryTests(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("❌ Directory %s not found\n", dir)
		return
	}

	fmt.Printf("🚀 Running tests in %s...\n", dir)
	passed, failed := runTestsInDirectory(dir)

	fmt.Printf("\n📊 Results for %s:\n", dir)
	fmt.Printf("✅ Passed: %d\n", passed)
	fmt.Printf("❌ Failed: %d\n", failed)
}

func runTestsInDirectory(dir string) (int, int) {
	passed := 0
	failed := 0

	// Check if directory has test files or subdirectories with tests
	testFiles, err := filepath.Glob(filepath.Join(dir, "*_test.go"))
	if err != nil {
		fmt.Printf("❌ Error reading directory %s: %v\n", dir, err)
		return 0, 1
	}

	// Check for subdirectories with tests
	subdirs, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		fmt.Printf("❌ Error reading directory %s: %v\n", dir, err)
		return 0, 1
	}

	hasTests := len(testFiles) > 0

	// Check subdirectories for test files
	for _, subdir := range subdirs {
		if stat, err := os.Stat(subdir); err == nil && stat.IsDir() {
			subTestFiles, _ := filepath.Glob(filepath.Join(subdir, "*_test.go"))
			if len(subTestFiles) > 0 {
				hasTests = true
				break
			}
		}
	}

	if !hasTests {
		fmt.Printf("  ⏭️  No test files found in %s\n", dir)
		return 0, 0
	}

	fmt.Printf("  🧪 Running tests in %s...\n", dir)

	// Run go test on the directory (includes subdirectories)
	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("    ❌ Failed: %s\n", strings.TrimSpace(string(output)))
		failed++
	} else {
		fmt.Printf("    ✅ Passed\n")
		passed++
	}

	return passed, failed
}

func runBenchmarks() {
	fmt.Println("🏃 Running benchmarks...")

	directories := []string{
		"01-arrays-strings",
		"02-hash-tables",
		"03-linked-lists",
		"04-stacks-queues",
		"05-trees",
		"06-heaps",
		"07-graphs",
		"08-dynamic-programming",
		"09-greedy",
	}

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		fmt.Printf("\n📁 Benchmarking %s...\n", dir)

		// Run benchmarks
		cmd := exec.Command("go", "test", "-bench=.", "-benchmem", dir)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("❌ Error running benchmarks in %s: %v\n", dir, err)
		} else {
			fmt.Printf("%s\n", string(output))
		}
	}
}

func runCoverage() {
	fmt.Println("📊 Running tests with coverage...")

	directories := []string{
		"01-arrays-strings",
		"02-hash-tables",
		"03-linked-lists",
		"04-stacks-queues",
		"05-trees",
		"06-heaps",
		"07-graphs",
		"08-dynamic-programming",
		"09-greedy",
	}

	totalCoverage := 0.0
	dirCount := 0

	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			continue
		}

		fmt.Printf("\n📁 Coverage for %s...\n", dir)

		// Run coverage
		cmd := exec.Command("go", "test", "-cover", dir)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("❌ Error running coverage in %s: %v\n", dir, err)
		} else {
			fmt.Printf("%s\n", string(output))
			// Parse coverage percentage (simplified)
			dirCount++
		}
	}

	if dirCount > 0 {
		fmt.Printf("\n📈 Average coverage: %.1f%%\n", totalCoverage/float64(dirCount))
	}
}

func showHelp() {
	fmt.Println("🧪 FAANG Interview Prep Test Runner")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  test-all     - Run all tests across all directories")
	fmt.Println("  test-dir <dir> - Run tests in specific directory")
	fmt.Println("  benchmark     - Run all benchmarks")
	fmt.Println("  coverage      - Run tests with coverage")
	fmt.Println("  help          - Show this help")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run test_runner.go test-all")
	fmt.Println("  go run test_runner.go test-dir 01-arrays-strings")
	fmt.Println("  go run test_runner.go benchmark")
	fmt.Println("  go run test_runner.go coverage")
}
