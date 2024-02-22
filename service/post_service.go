package service

import (
	"egardev-gin-socmed/dto"
	"egardev-gin-socmed/entity"
	"egardev-gin-socmed/errorHandler"
	"egardev-gin-socmed/repository"
)

type PostService interface {
	Create(req *dto.PostRequest) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) *postService {
	return &postService{
		repository: r,
	}
}

func (s *postService) Create(req *dto.PostRequest) error {
	post := entity.Post{
		UserID: req.UserID,
		Tweet:  req.Tweet,
	}

	if req.Picture != nil {
		post.PictureUrl = &req.Picture.Filename
	}

	if err := s.repository.Create(&post); err != nil {
		return &errorHandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
