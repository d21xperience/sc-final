package service

import (
	"sc-service/data/request"
	"sc-service/data/response"

	"github.com/google/uuid"
)

type IjazahService interface {
	Create(ij request.CreateTagsRequest)
	Update(ij request.UpdateTagsRequest)
	Delete(ijID uuid.UUID)
	FindAll(thnLulus int) []response.TagsResponse
}
