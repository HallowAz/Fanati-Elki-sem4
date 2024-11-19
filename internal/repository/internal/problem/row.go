package problem

import "fe-sem4/internal/models/problem"

type ProblemRow struct {
	ID               uint32   `db:"id"`
	Title            string   `db:"title"`
	Description      string   `db:"description"`
	SpecificLocation string   `db:"specific_location"`
	Category         string   `db:"category"`
	Media            []string `db:"media"`
	MediaFiles       [][]byte `db:"mediaFiles"`
	VoteCount        uint16   `db:"vote_count"`
	Lat              string   `db:"lat"`
	Long             string   `db:"long"`
	Status           string   `db:"status"`
}

func NewProblemRow(model problem.Problem) ProblemRow {
	return ProblemRow{
		ID:               model.ID,
		Title:            model.Title,
		Description:      model.Description,
		SpecificLocation: model.SpecificLocation,
		Category:         model.Category,
		Media:            model.Media,
		VoteCount:        model.VoteCount,
		Lat:              model.Lat,
		Long:             model.Long,
		Status:           model.Status,
	}
}

func (r *ProblemRow) ToModel() problem.Problem {
	return problem.Problem{
		ID:               r.ID,
		Title:            r.Title,
		Description:      r.Description,
		SpecificLocation: r.SpecificLocation,
		Category:         r.Category,
		Media:            r.Media,
		MediaFiles:       r.MediaFiles,
		VoteCount:        r.VoteCount,
		Lat:              r.Lat,
		Long:             r.Long,
		Status:           r.Status,
	}
}
