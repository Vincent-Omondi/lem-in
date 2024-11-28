// pathfinder.go
package pkg

import (
	"reflect"
	"testing"
)

func TestFindPaths(t *testing.T) {
	tests := []struct {
		name  string
		setup func() // Optional: Function to set up test-specific data
		want  [][]string
	}{
		{
			name: "No room connections",
			setup: func() {
				// Mock StartRoom and RoomConnections
				StartRoom = "RoomA"
				RoomConnections = map[string][]string{}
				ValidPaths = [][]string{}
			},
			want: [][]string{},
		},
		{
			name: "Single valid path",
			setup: func() {
				// Mock StartRoom and RoomConnections
				StartRoom = "RoomA"
				RoomConnections = map[string][]string{
					"RoomA": {"RoomB"},
					"RoomB": {},
				}
				ValidPaths = [][]string{{"RoomA", "RoomB"}}
			},
			want: [][]string{{"RoomA", "RoomB"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			if got := FindPaths(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPathRatings(t *testing.T) {
	tests := []struct {
		name       string
		validPaths [][]string
		want       map[int]int
	}{
		{
			name:       "Empty ValidPaths",
			validPaths: [][]string{},
			want:       map[int]int{},
		},
		{
			name: "Single Path",
			validPaths: [][]string{
				{"A", "B", "C"},
			},
			want: map[int]int{
				0: 3,
			},
		},
		{
			name: "Multiple Paths",
			validPaths: [][]string{
				{"A", "B"},
				{"C", "D", "E"},
				{"F"},
			},
			want: map[int]int{
				0: 2,
				1: 3,
				2: 1,
			},
		},
		{
			name: "Paths of same length",
			validPaths: [][]string{
				{"X", "Y"},
				{"A", "B"},
			},
			want: map[int]int{
				0: 2,
				1: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ValidPaths = tt.validPaths
			PathRatings = make(map[int]int)

			SetPathRatings()

			if !reflect.DeepEqual(PathRatings, tt.want) {
				t.Errorf("SetPathRatings() = %v, want %v", PathRatings, tt.want)
			}
		})
	}
}

func TestArePathsIdentical(t *testing.T) {
	type args struct {
		path1 []string
		path2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Identical paths",
			args: args{
				path1: []string{"A", "B", "C"},
				path2: []string{"A", "B", "C"},
			},
			want: true,
		},
		{
			name: "Paths with different elements",
			args: args{
				path1: []string{"A", "B", "C"},
				path2: []string{"A", "D", "C"},
			},
			want: false,
		},
		{
			name: "Paths with different lengths",
			args: args{
				path1: []string{"A", "B"},
				path2: []string{"A", "B", "C"},
			},
			want: true,
		},
		{
			name: "Empty paths",
			args: args{
				path1: []string{},
				path2: []string{},
			},
			want: true,
		},
		{
			name: "Both paths are nil",
			args: args{
				path1: nil,
				path2: nil,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArePathsIdentical(tt.args.path1, tt.args.path2); got != tt.want {
				t.Errorf("ArePathsIdentical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDepthFirstSearch(t *testing.T) {
	tests := []struct {
		name            string
		startRoom       string
		endRoom         string
		roomConnections map[string][]string
		wantValidPaths  [][]string
	}{
		{
			name:            "Single room is start and end",
			startRoom:       "A",
			endRoom:         "A",
			roomConnections: map[string][]string{"A": {}},
			wantValidPaths:  [][]string{{"A"}},
		},
		{
			name:      "Linear path",
			startRoom: "A",
			endRoom:   "C",
			roomConnections: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {},
			},
			wantValidPaths: [][]string{{"A", "B", "C"}},
		},
		{
			name:      "Branching paths",
			startRoom: "A",
			endRoom:   "D",
			roomConnections: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {"D"},
				"D": {},
			},
			wantValidPaths: [][]string{
				{"A", "B", "D"},
				{"A", "C", "D"},
			},
		},
		{
			name:      "No valid path",
			startRoom: "A",
			endRoom:   "D",
			roomConnections: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {},
			},
			wantValidPaths: [][]string{},
		},
		{
			name:      "Graph with loops",
			startRoom: "A",
			endRoom:   "C",
			roomConnections: map[string][]string{
				"A": {"B", "C"},
				"B": {"A", "C"},
				"C": {"A"},
			},
			wantValidPaths: [][]string{
				{"A", "B", "C"},
				{"A", "C"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock data
			StartRoom = tt.startRoom
			EndRoom = tt.endRoom
			RoomConnections = tt.roomConnections
			TraversalStack = []string{}
			VisitedRooms = make(map[string]bool)
			ValidPaths = [][]string{}

			// Call the function
			DepthFirstSearch(StartRoom)

			// Verify the result
			if !reflect.DeepEqual(ValidPaths, tt.wantValidPaths) {
				t.Errorf("ValidPaths = %v, want %v", ValidPaths, tt.wantValidPaths)
			}
		})
	}
}

func TestTraverseGraph(t *testing.T) {
	tests := []struct {
		name            string
		startNode       string
		endRoom         string
		roomConnections map[string][]string
		wantValidPaths  [][]string
		wantReturn      bool
	}{
		{
			name:            "Single node is start and end",
			startNode:       "A",
			endRoom:         "A",
			roomConnections: map[string][]string{"A": {}},
			wantValidPaths:  [][]string{{"A"}},
			wantReturn:      false,
		},
		{
			name:      "Linear graph",
			startNode: "A",
			endRoom:   "C",
			roomConnections: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {},
			},
			wantValidPaths: [][]string{{"A", "B", "C"}},
			wantReturn:     false,
		},
		{
			name:      "Branching graph",
			startNode: "A",
			endRoom:   "D",
			roomConnections: map[string][]string{
				"A": {"B", "C"},
				"B": {"D"},
				"C": {"D"},
				"D": {},
			},
			wantValidPaths: [][]string{
				{"A", "B", "D"},
			},
			wantReturn: false,
		},
		{
			name:      "Disconnected graph",
			startNode: "A",
			endRoom:   "D",
			roomConnections: map[string][]string{
				"A": {"B"},
				"B": {"C"},
				"C": {},
			},
			wantValidPaths: [][]string{},
			wantReturn:     true,
		},
		{
			name:      "Graph with cycles",
			startNode: "A",
			endRoom:   "C",
			roomConnections: map[string][]string{
				"A": {"B", "C"},
				"B": {"A", "C"},
				"C": {"A"},
			},
			wantValidPaths: [][]string{
				{"A", "C"},
			},
			wantReturn: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock data
			StartRoom = tt.startNode
			EndRoom = tt.endRoom
			RoomConnections = tt.roomConnections
			VisitedRooms = make(map[string]bool)
			ValidPaths = [][]string{}

			// Call the function
			got := TraverseGraph(tt.startNode)

			// Validate the return value
			if got != tt.wantReturn {
				t.Errorf("TraverseGraph() = %v, want %v", got, tt.wantReturn)
			}

			// Validate ValidPaths
			if !reflect.DeepEqual(ValidPaths, tt.wantValidPaths) {
				t.Errorf("ValidPaths = %v, want %v", ValidPaths, tt.wantValidPaths)
			}
		})
	}
}

func Test_sortSolutions(t *testing.T) {
	tests := []struct {
		name       string
		validPaths [][]string
		want       [][]string
	}{
		{
			name:       "Different path lengths",
			validPaths: [][]string{{"A", "B", "C"}, {"A", "B"}, {"A"}},
			want:       [][]string{{"A"}, {"A", "B"}, {"A", "B", "C"}},
		},
		{
			name:       "Same length paths, compare used for tie-breaking",
			validPaths: [][]string{{"A", "C"}, {"A", "B"}},
			want:       [][]string{{"A", "C"}, {"A", "B"}},
		},
		{
			name:       "Mixed lengths and ties",
			validPaths: [][]string{{"A", "D"}, {"A", "B", "C"}, {"A", "B"}},
			want:       [][]string{{"A", "D"}, {"A", "B"}, {"A", "B", "C"}},
		},
		{
			name:       "Empty paths",
			validPaths: [][]string{},
			want:       [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up global ValidPaths
			ValidPaths = tt.validPaths

			// Call the function
			sortSolutions()

			// Validate the result
			if !reflect.DeepEqual(ValidPaths, tt.want) {
				t.Errorf("ValidPaths = %v, want %v", ValidPaths, tt.want)
			}
		})
	}
}

