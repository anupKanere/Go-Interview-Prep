package main

import (
	"fmt"
	"os"
	"os/exec"
)

/*
FAANG Interview Prep Setup Script
This script helps set up the Go environment for FAANG interview preparation
*/

func main() {
	fmt.Println("🚀 Setting up FAANG Interview Prep Environment...")

	// Check if Go is installed
	if !isGoInstalled() {
		fmt.Println("❌ Go is not installed. Please install Go first:")
		fmt.Println("   https://golang.org/doc/install")
		return
	}

	// Initialize Go module
	fmt.Println("📦 Initializing Go module...")
	initGoModule()

	// Create go.mod file
	fmt.Println("📝 Creating go.mod file...")
	createGoMod()

	// Run initial tests
	fmt.Println("🧪 Running initial tests...")
	runTests()

	// Display setup completion
	fmt.Println("\n✅ Setup completed successfully!")
	fmt.Println("\n📚 Next steps:")
	fmt.Println("1. Review the study plan in STUDY_PLAN.md")
	fmt.Println("2. Start with 01-arrays-strings directory")
	fmt.Println("3. Use templates in templates/ directory")
	fmt.Println("4. Run tests with: go run tests/test_runner.go test-all")
	fmt.Println("5. Practice daily for 3-4 months")

	fmt.Println("\n🎯 Good luck with your FAANG preparation! 🚀")
}

func isGoInstalled() bool {
	cmd := exec.Command("go", "version")
	err := cmd.Run()
	return err == nil
}

func initGoModule() {
	cmd := exec.Command("go", "mod", "init", "faang-interview-prep")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("⚠️  Go module already initialized or error: %v\n", err)
	}
}

func createGoMod() {
	content := `module faang-interview-prep

go 1.21

require (
	// Add dependencies as needed
	// github.com/stretchr/testify v1.8.4
)`

	err := os.WriteFile("go.mod", []byte(content), 0644)
	if err != nil {
		fmt.Printf("⚠️  Error creating go.mod: %v\n", err)
	}
}

func runTests() {
	// Test if we can run tests
	cmd := exec.Command("go", "test", "-v", "./01-arrays-strings/...")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("⚠️  Some tests failed (this is normal for initial setup):\n%s\n", string(output))
	} else {
		fmt.Printf("✅ Tests passed:\n%s\n", string(output))
	}
}

func checkDirectoryStructure() {
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
		"templates",
		"tests",
		"resources",
		"mock-interviews",
	}

	fmt.Println("📁 Checking directory structure...")
	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("❌ Directory %s not found\n", dir)
		} else {
			fmt.Printf("✅ Directory %s exists\n", dir)
		}
	}
}

func displayProjectStructure() {
	fmt.Println("\n📁 Project Structure:")
	fmt.Println("├── 01-arrays-strings/          # Arrays, Strings, Two Pointers")
	fmt.Println("├── 02-hash-tables/             # Hash Tables, Sliding Window")
	fmt.Println("├── 03-linked-lists/            # Linked Lists")
	fmt.Println("├── 04-stacks-queues/           # Stacks and Queues")
	fmt.Println("├── 05-trees/                   # Binary Trees, BST")
	fmt.Println("├── 06-heaps/                   # Heaps and Priority Queues")
	fmt.Println("├── 07-graphs/                  # Graph Algorithms")
	fmt.Println("├── 08-dynamic-programming/     # DP Problems")
	fmt.Println("├── 09-greedy/                  # Greedy Algorithms")
	fmt.Println("├── 10-system-design/           # System Design Basics")
	fmt.Println("├── templates/                  # Go coding templates")
	fmt.Println("├── tests/                      # Test files")
	fmt.Println("├── resources/                  # Study resources")
	fmt.Println("├── mock-interviews/            # Practice interviews")
	fmt.Println("├── README.md                   # Project overview")
	fmt.Println("├── STUDY_PLAN.md              # 3-4 month study plan")
	fmt.Println("└── setup.go                    # This setup script")
}
