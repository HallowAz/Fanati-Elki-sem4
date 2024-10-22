package problem

type Problem struct {
	ID               uint32
	Title            string
	Description      string
	SpecificLocation string
	Category         string
	Media            []string
	VoteCount        uint16
	Lat              string
	Long             string
}
