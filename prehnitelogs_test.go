package prehnitelogs

import (
	"fmt"
	"testing"
)

func TestLogs(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test logs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunTestLogs()
		})
	}
}

func TestScript(t *testing.T) {
	// Simple function to test things!

	fmt.Println("Test Script")

	AddType("TEST", " - TEST : ")

	err := Log("SimpleLog!")

	fmt.Println("Error - Simple ", err)

	err = CustomLog("TEST", "Wow much custom logging!", false)
	fmt.Println("Error - Custom ", err)

	// Testing the Make Custom functions
	AddType("Multiple", " - Multiple - So Lazy! : ")
	l := GetCustomLogMethod("Multiple", false)
	err = l("Wow much custom logging!")
	fmt.Println("Error - Custom MAGIC! ", err)
	err = l("MORE CUSTOM LOGGING!")
	fmt.Println("Error - MORE MAGIC! ", err)

	RunTestLogs()
}
