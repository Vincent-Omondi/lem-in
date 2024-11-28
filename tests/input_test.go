// input.go
package tests

import (
	"os"
	"testing"

	"github.com/Vincent-Omondi/lem-in/pkg"
)

func TestProcessInputFile(t *testing.T) {
	tests := []struct {
		name    string
		args    struct{ inputFile *os.File }
		want    string
		wantErr bool
	}{
		{
			name: "Valid file with start and end room",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##start\nroom1 0 0\n##end\nroom2 1 1\nroom1-room2\n"),
			},
			want:    "dfs",
			wantErr: false,
		},
		{
			name: "File with invalid special command",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##start\nroom1 0 0\n##invalid\nroom2 1 1\n"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Missing room after ##start",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##start\n"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Missing room after ##end",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##end\n"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Duplicate connection between rooms",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##start\nroom1 0 0\n##end\nroom2 1 1\nroom1-room2\nroom1-room2\n"),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Room name starts with L",
			args: struct{ inputFile *os.File }{
				inputFile: createTestFile(t, "10\n##start\nLroom 0 0\n##end\nroom2 1 1\n"),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pkg.ProcessInputFile(tt.args.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessInputFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProcessInputFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
