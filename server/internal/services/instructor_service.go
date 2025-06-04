package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"

	"github.com/google/uuid"
)

type InstructorService interface {
	DeleteInstructor(id string) error
	GetAllInstructors() ([]dto.InstructorResponse, error)
	CreateInstructor(req dto.CreateInstructorRequest) error
	GetInstructorByID(id string) (*dto.InstructorResponse, error)
	UpdateInstructor(id string, req dto.UpdateInstructorRequest) error
}

type instructorService struct {
	repo repositories.InstructorRepository
	user repositories.UserRepository
}

func NewInstructorService(repo repositories.InstructorRepository, user repositories.UserRepository) InstructorService {
	return &instructorService{repo, user}
}

func (s *instructorService) DeleteInstructor(id string) error {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return customErr.NewNotFound("Instructor not found")
	}

	user, err := s.user.GetUserByID(instructor.UserID.String())
	if err == nil {
		user.Role = "customer"
		_ = s.user.UpdateUser(user)
	}

	if err := s.repo.DeleteInstructor(id); err != nil {
		return customErr.ErrDeleteFailed
	}
	return nil
}

func (s *instructorService) GetAllInstructors() ([]dto.InstructorResponse, error) {
	instructors, err := s.repo.GetAllInstructors()
	if err != nil {
		return nil, customErr.ErrInternalServer
	}

	var result []dto.InstructorResponse
	for _, i := range instructors {
		result = append(result, dto.InstructorResponse{
			ID:             i.ID.String(),
			UserID:         i.UserID.String(),
			Fullname:       i.User.Fullname,
			Avatar:         i.User.Avatar,
			Experience:     i.Experience,
			Specialties:    i.Specialties,
			Certifications: i.Certifications,
			Rating:         i.Rating,
			TotalClass:     i.TotalClass,
		})
	}
	return result, nil
}

func (s *instructorService) CreateInstructor(req dto.CreateInstructorRequest) error {
	user, err := s.user.GetUserByID(req.UserID)
	if err != nil {
		return customErr.NewNotFound("User not found")
	}

	user.Role = "instructor"
	if err := s.user.UpdateUser(user); err != nil {
		return customErr.ErrUpdateFailed
	}

	instructor := models.Instructor{
		UserID:         uuid.MustParse(req.UserID),
		Experience:     req.Experience,
		Specialties:    req.Specialties,
		Certifications: req.Certifications,
	}

	if err := s.repo.CreateInstructor(&instructor); err != nil {
		return customErr.ErrCreateFailed
	}

	return nil
}

func (s *instructorService) GetInstructorByID(id string) (*dto.InstructorResponse, error) {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return nil, customErr.NewNotFound("Instructor not found")
	}

	return &dto.InstructorResponse{
		ID:             instructor.ID.String(),
		UserID:         instructor.UserID.String(),
		Fullname:       instructor.User.Fullname,
		Avatar:         instructor.User.Avatar,
		Experience:     instructor.Experience,
		Specialties:    instructor.Specialties,
		Certifications: instructor.Certifications,
		Rating:         instructor.Rating,
		TotalClass:     instructor.TotalClass,
	}, nil
}

func (s *instructorService) UpdateInstructor(id string, req dto.UpdateInstructorRequest) error {
	instructor, err := s.repo.GetInstructorByID(id)
	if err != nil {
		return customErr.NewNotFound("Instructor not found")
	}

	if instructor.UserID.String() != req.UserID {
		return customErr.ErrForbidden
	}

	if req.Experience != 0 {
		instructor.Experience = req.Experience
	}
	if req.Specialties != "" {
		instructor.Specialties = req.Specialties
	}
	if req.Certifications != "" {
		instructor.Certifications = req.Certifications
	}

	if err := s.repo.UpdateInstructor(instructor); err != nil {
		return customErr.ErrUpdateFailed
	}

	return nil
}
