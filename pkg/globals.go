package pkg

var (
	RoomConnections = make(map[string][]string)
	EmptyRoom       = make(map[string]bool)
	PathRatings     = make(map[int]int)
	RoomList        []string
	StartRoom 		string
	EndRoom 		string
	AntsCount       int
	Graphoverview   []byte
)

type PathInfo struct {
	Rating int
	Index  int
}

var (
	VisitedRooms   = make(map[string]bool)
	ValidPaths     [][]string
	TraversalStack []string
)
