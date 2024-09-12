package utils_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ecshreve/dndgen/internal/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCleanString tests the CleanString method.
func TestCleanString(t *testing.T) {
	assert.Equal(t, "index1", utils.CleanString("Index1"))
	assert.Equal(t, "index_1", utils.CleanString("Index-1"))
	assert.Equal(t, "index1", utils.CleanString("Index'1"))
	assert.Equal(t, "index1", utils.CleanString("Index,1"))
	assert.Equal(t, "index_1", utils.CleanString("Index 1"))
}

// TestIntOrDefault tests the IntOrDefault method.
func TestIntOrDefault(t *testing.T) {
	testInt := 1
	assert.Equal(t, 1, utils.IntOrDefault(&testInt, 1))
	assert.Equal(t, 0, utils.IntOrDefault(nil, 0))
	assert.Equal(t, 1, utils.IntOrDefault(nil, 1))
}

func TestReadJSONFile(t *testing.T) {
	testcases := []struct {
		name        string
		fileContent string
		fileName    string
		expected    interface{}
		expectError bool
	}{
		{
			name:        "Valid JSON",
			fileContent: `{"key": "value"}`,
			expected: map[string]interface{}{
				"key": "value",
			},
			expectError: false,
		},
		{
			name:        "Invalid JSON",
			fileContent: `{"key": "value"`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty JSON",
			fileContent: `{}`,
			expected:    map[string]interface{}{},
			expectError: false,
		},
		{
			name:        "Non-existent file",
			fileContent: "",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Non-JSON file",
			fileContent: "This is not a JSON file",
			fileName:    "testfile.txt",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			fileName := "testfile.json"
			if tc.fileName != "" {
				fileName = tc.fileName
			}
			// Create a temporary file for the test
			tempFile, err := os.CreateTemp("", fileName)
			require.NoError(t, err)
			defer os.Remove(tempFile.Name()) // Clean up the temporary file

			// Write content to the temp file if it's part of the test case
			if tc.fileContent != "" {
				_, err = tempFile.Write([]byte(tc.fileContent))
				require.NoError(t, err)

				// Close the file after writing
				require.NoError(t, tempFile.Close())
			} else {
				// Delete the file to simulate a non-existent file case
				os.Remove(tempFile.Name())
			}

			// Prepare a variable to hold the result
			var result interface{}

			// Execute the function under test
			err = utils.LoadJSONFile(tempFile.Name(), &result)

			// Check for expected errors
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}

		})
	}
}

// Helper function to compare interfaces (used to compare JSON results)
func compareInterfaces(a, b interface{}) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
