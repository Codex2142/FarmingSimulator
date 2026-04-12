package utils

import (
	"farming/internal/dto"
	"farming/internal/model"
)

func ToFarmResponse(f model.Farm) dto.FarmResponse {
	var leader *dto.UserSimple

	if f.Leader != nil {
		leader = &dto.UserSimple{
			ID:    f.Leader.ID,
			Name:  f.Leader.Name,
			Phone: f.Leader.Phone,
		}
	}

	return dto.FarmResponse{
		ID:       f.ID,
		Name:     f.Name,
		Location: f.Location,
		Leader:   leader,
	}
}
