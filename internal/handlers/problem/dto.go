package problem

import "fe-sem4/internal/models/problem"

type createProblemRequest struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	SpecificLocation string   `json:"specificLocation"`
	Category         string   `json:"category"`
	Media            []string `json:"media"`
	Lat              string   `json:"lat"`
	Long             string   `json:"long"`
}

func (c *createProblemRequest) toModel() problem.Problem {
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

type GetProblemsResponse struct {
	Problems []getProblemResponse `json:"problems"`
}

type getProblemResponse struct {
	ID               uint32   `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	SpecificLocation string   `json:"specificLocation"`
	Category         string   `json:"category"`
	VoteCount        uint16   `json:"voteCount"`
	Media            []string `json:"media"`
	Lat              string   `json:"lat"`
	Long             string   `json:"long"`
}

func newGetProblemResponse(model problem.Problem) getProblemResponse {
	return getProblemResponse{
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

type updateProblemRequest struct {
	ID               uint32   `json:"-"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	SpecificLocation string   `json:"specificLocation"`
	Category         string   `json:"category"`
	VoteCount        uint16   `json:"voteCount"`
	Media            []string `json:"media"`
	Lat              string   `json:"lat"`
	Long             string   `json:"long"`
}

func (u *updateProblemRequest) toModel() problem.Problem {
	return problem.Problem{
		ID:               u.ID,
		Title:            u.Title,
		Description:      u.Description,
		SpecificLocation: u.SpecificLocation,
		VoteCount:        u.VoteCount,
		Category:         u.Category,
		Media:            u.Media,
		Lat:              u.Lat,
		Long:             u.Long,
	}
}
