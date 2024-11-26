package pkg

var (
	Graphoverview    []byte                 // Stores the overview of the graph as a byte array
	AntsCount        int                    // Stores the number of ants
	RoomList         []string               // Stores the list of room names
	RoomConnections  map[string][]string    // Stores room connections as an adjacency list
	StartRoom        string                 // Stores the start room name
	EndRoom          string                 // Stores the end room name
)