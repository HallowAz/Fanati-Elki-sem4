package problem

import "fe-sem4/internal/models/problem"

type CreateProblemRequest struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	SpecificLocation string   `json:"specificLocation"`
	Category         string   `json:"category"`
	Media            []string `json:"media"`
	Lat              string   `json:"lat"`
	Long             string   `json:"long"`
}

func (c *CreateProblemRequest) ToModel() problem.Problem {
	return problem.Problem{
		Title:            c.Title,
		Description:      c.Description,
		SpecificLocation: c.SpecificLocation,
		Category:         c.Category,
		Media:            c.Media,
		Lat:              c.Lat,
		Long:             c.Long,
	}
}
