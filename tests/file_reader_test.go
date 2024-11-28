package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/Vincent-Omondi/lem-in/pkg"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name    string
		args    struct{ inputFile *os.File }
		want    []string
		wantErr bool
	}{
		{
			name: "Valid file with multiple lines",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "line 1\nline 2\nline 3\n"),
			},
			want:    []string{"line 1", "line 2", "line 3"},
			wantErr: false,
		},
		{
			name: "Valid file with single line",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "single line\n"),
			},
			want:    []string{"single line"},
			wantErr: false,
		},
		{
			name: "File does not exist",
			args: struct{ inputFile *os.File }{
				inputFile: nil, // Simulating a file that does not exist
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pkg.ReadFile(tt.args.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to create temporary test files
func createTestFile(t *testing.T, content string) *os.File {
	tmpFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write content to temp file: %v", err)
	}
	// Rewind file pointer to the start for reading
	tmpFile.Seek(0, 0)
	return tmpFile
}
