package problem

import "fe-sem4/internal/models/problem"

type ProblemRow struct {
	ID               int      `db:"id"`
	Title            string   `db:"title"`
	Description      string   `db:"description"`
	SpecificLocation string   `db:"specific_location"`
	Category         string   `db:"category"`
	Media            []string `db:"media"`
	VoteCount        int      `db:"vote_count"`
	Lat              string   `db:"lat"`
	Long             string   `db:"long"`
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
	}
}
