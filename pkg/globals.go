package lem

var (
	Ways          = make(map[string][]string)
	Emptyroom     = make(map[string]bool)
	Rooms         = []string{}
	Start, End    string
	Ants          int
	Graphoverview []byte
)
