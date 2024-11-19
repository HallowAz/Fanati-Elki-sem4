package problem

import (
	"fe-sem4/internal/models/problem"
	"io"
	"net/http"
)

type createProblemRequest struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	SpecificLocation string   `json:"specificLocation"`
	Category         string   `json:"category"`
	MediaFiles       [][]byte `json:"mediaFiles"`
	Lat              string   `json:"lat"`
	Long             string   `json:"long"`
	Status           string   `json:"status"`
}

func (c *createProblemRequest) toModel() problem.Problem {
	return problem.Problem{
		Title:            c.Title,
		Description:      c.Description,
		SpecificLocation: c.SpecificLocation,
		Category:         c.Category,
		MediaFiles:       c.MediaFiles,
		Lat:              c.Lat,
		Long:             c.Long,
		Status:           c.Status,
	}
}

func parseFormCreateProblem(r *http.Request) (*createProblemRequest, error) {
	var c createProblemRequest

	c.Title = r.FormValue("title")
	c.Description = r.FormValue("description")
	c.SpecificLocation = r.FormValue("specificLocation")
	c.Category = r.FormValue("category")
	c.Lat = r.FormValue("lat")
	c.Long = r.FormValue("long")

	files := r.MultipartForm.File["mediaFiles"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}

		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		c.MediaFiles = append(c.MediaFiles, bytes)

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}

type getProblemsResponse struct {
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
	Status           string   `json:"status"`
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
		Status:           model.Status,
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
	Status           string   `json:"status"`
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
		Status:           u.Status,
	}
}
