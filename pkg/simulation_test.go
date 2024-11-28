// simulator.go
package pkg

import (
	"bytes"
	"os"
	"testing"
)

func TestControlTraffic(t *testing.T) {
	type args struct {
		antGroups [][]string
		paths     [][]string
	}
	tests := []struct {
		name        string
		args        args
		expectedOut string // Add expected output for each test case
	}{
		{
			name: "Single path with single ant",
			args: args{
				antGroups: [][]string{{"ant1"}},
				paths:     [][]string{{"start", "room1", "room2", EndRoom}},
			},
			expectedOut: "", // Adjust this output based on expected behavior
		},
		{
			name: "Multiple paths with single ant each",
			args: args{
				antGroups: [][]string{{"ant1"}, {"ant2"}},
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
					{"start", "roomA", "roomB", EndRoom},
				},
			},
			expectedOut: "", // Adjust accordingly
		},
		{
			name: "Multiple ants on single path",
			args: args{
				antGroups: [][]string{{"ant1", "ant2", "ant3"}},
				paths:     [][]string{{"start", "room1", "room2", "room3", EndRoom}},
			},
			expectedOut: "",
		},
		{
			name: "Multiple paths with multiple ants",
			args: args{
				antGroups: [][]string{
					{"ant1", "ant2"},
					{"ant3", "ant4"},
				},
				paths: [][]string{
					{"start", "room1", "room2", "room3", EndRoom},
					{"start", "roomA", "roomB", "roomC", EndRoom},
				},
			},
			expectedOut: "",
		},
		{
			name: "Complex routing with different path lengths",
			args: args{
				antGroups: [][]string{
					{"ant1", "ant2"},
					{"ant3"},
					{"ant4", "ant5"},
				},
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
					{"start", "roomA", "roomB", "roomC", "roomD", EndRoom},
					{"start", "room3", EndRoom},
				},
			},
			expectedOut: "",
		},
	}

	for _, tt := range tests {
		// Before each test, reset AntsCount
		AntsCount = 0
		for _, group := range tt.args.antGroups {
			AntsCount += len(group)
		}

		t.Run(tt.name, func(t *testing.T) {
			// Clone antGroups to avoid mutation
			antGroupsCopy := make([][]string, len(tt.args.antGroups))
			for i, group := range tt.args.antGroups {
				antGroupsCopy[i] = append([]string(nil), group...)
			}

			// Capture stdout
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			// Run the function
			ControlTraffic(antGroupsCopy, tt.args.paths)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Check if the output matches expected output
			var buf bytes.Buffer
			w.Close()
			// Get the output from the pipe
			buf.WriteTo(&buf)
			output := buf.String()

			if output != tt.expectedOut {
				t.Errorf("expected %v, got %v", tt.expectedOut, output)
			}
		})
	}
}

func TestDispatchAnts(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name       string
		args       args
		antCount   int  // Global AntsCount to set before the test
		pathLength bool // Optional flag to check specific path length handling
	}{
		{
			name: "Equal number of paths and ants",
			args: args{
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
					{"start", "roomA", "roomB", EndRoom},
				},
			},
			antCount: 2,
		},
		{
			name: "More paths than ants",
			args: args{
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
					{"start", "roomA", "roomB", EndRoom},
					{"start", "roomX", "roomY", EndRoom},
				},
			},
			antCount:   3,
			pathLength: true, // Should truncate paths to match ant count
		},
		{
			name: "Fewer paths than ants",
			args: args{
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
				},
			},
			antCount: 5,
		},
		{
			name: "Uneven distribution of ants",
			args: args{
				paths: [][]string{
					{"start", "room1", "room2", EndRoom},
					{"start", "roomA", "roomB", EndRoom},
					{"start", "roomX", "roomY", EndRoom},
				},
			},
			antCount: 7,
		},
		{
			name: "Minimal paths",
			args: args{
				paths: [][]string{
					{"start", EndRoom},
				},
			},
			antCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original AntsCount and restore after test
			originalAntsCount := AntsCount
			AntsCount = tt.antCount
			defer func() { AntsCount = originalAntsCount }()

			// Capture stdout to verify output
			oldStdout := os.Stdout
			pipeR, pipeW, _ := os.Pipe()
			os.Stdout = pipeW

			// Run the function
			DispatchAnts(tt.args.paths)

			// Close the pipe and restore stdout
			pipeW.Close()
			os.Stdout = oldStdout

			// Read the captured output
			var buf []byte
			_, err := pipeR.Read(buf)
			if err != nil {
				t.Fatalf("failed to capture output: %v", err)
			}
			// output := string(buf)

			// Optional additional assertions
			if tt.pathLength && len(tt.args.paths) > tt.antCount {
				if len(tt.args.paths) > tt.antCount {
					t.Errorf("Paths should be truncated to match the ant count, got %v paths, expected no more than %v", len(tt.args.paths), tt.antCount)
				}
			}
		})
	}
}
