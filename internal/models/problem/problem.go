package problem

type Problem struct {
	ID               uint32
	Title            string
	Description      string
	SpecificLocation string
	Category         string
	Media            []string
	MediaFiles       [][]byte
	VoteCount        uint16
	Lat              string
	Long             string
	Status           string
}
