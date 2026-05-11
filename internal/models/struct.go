package models

type Room struct {
	Name string
	X    int
	Y    int
}

type PathInfo struct {
	Path   []string
	AntNbr int
}

type Ant struct {
	ID    int
	Path  []string
	Index int
}
