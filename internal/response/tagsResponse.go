package response

import "awesomeProject/internal/model"

type TagsResponse struct {
	Tags []string `json:"tags"`
}

func NewTagsResponse(t []model.Tag) *TagsResponse {
	tags := make([]string, len(t))
	for _, tag := range t {
		tags = append(tags, tag.Tag)
	}

	return &TagsResponse{Tags: tags}
}
