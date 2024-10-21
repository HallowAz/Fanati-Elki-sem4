package problem

type Problem struct {
	ID               int
	Title            string
	Description      string
	SpecificLocation string
	Category         string
	Media            []string
	VoteCount        int
	Lat              string
	Long             string
}
