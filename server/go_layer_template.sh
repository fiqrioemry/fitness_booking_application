from pathlib import Path

# Define target files
base_path = Path("internal")
handler_path = base_path / "handlers" / "example_handler.go"
service_path = base_path / "services" / "example_service.go"
repo_path = base_path / "repositories" / "example_repo.go"

# Define the content for each file
handler_code = '''package handlers

import (
\t"net/http"
\t"server/internal/dto"
\t"server/internal/services"
\t"server/pkg/utils"

\t"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
\tservice services.ExampleService
}

func NewExampleHandler(service services.ExampleService) *ExampleHandler {
\treturn &ExampleHandler{service}
}

func (h *ExampleHandler) CreateExample(c *gin.Context) {
\tvar req dto.CreateExampleRequest

\tif !utils.BindAndValidateForm(c, &req) {
\t\treturn
\t}

\tif req.Image != nil {
\t\timageURL, err := utils.UploadImageWithValidation(req.Image)
\t\tif err != nil {
\t\t\tutils.HandleServiceError(c, err, err.Error())
\t\t\treturn
\t\t}
\t\treq.ImageURL = imageURL
\t}

\tif len(req.Images) > 0 {
\t\timageURLs, err := utils.UploadMultipleImagesWithValidation(req.Images)
\t\tif err != nil {
\t\t\tutils.CleanupImageOnError(req.ImageURL)
\t\t\tutils.HandleServiceError(c, err, err.Error())
\t\t\treturn
\t\t}
\t\treq.ImageURLs = imageURLs
\t}

\tif err := h.service.CreateExample(req); err != nil {
\t\tutils.CleanupImageOnError(req.ImageURL)
\t\tutils.CleanupImagesOnError(req.ImageURLs)
\t\tutils.HandleServiceError(c, err, err.Error())
\t\treturn
\t}

\tc.JSON(http.StatusCreated, gin.H{"message": "Example created successfully"})
}

func (h *ExampleHandler) UpdateExample(c *gin.Context) {
\tid := c.Param("id")
\tvar req dto.UpdateExampleRequest
\tif !utils.BindAndValidateForm(c, &req) {
\t\treturn
\t}

\tif req.Image != nil {
\t\timageURL, err := utils.UploadImageWithValidation(req.Image)
\t\tif err != nil {
\t\t\tutils.HandleServiceError(c, err, err.Error())
\t\t\treturn
\t\t}
\t\treq.ImageURL = imageURL
\t}

\tif err := h.service.UpdateExample(id, req); err != nil {
\t\tutils.CleanupImageOnError(req.ImageURL)
\t\tutils.HandleServiceError(c, err, err.Error())
\t\treturn
\t}

\tc.JSON(http.StatusOK, gin.H{"message": "Example updated successfully"})
}

func (h *ExampleHandler) DeleteExample(c *gin.Context) {
\tid := c.Param("id")

\tif err := h.service.DeleteExample(id); err != nil {
\t\tutils.HandleServiceError(c, err, err.Error())
\t\treturn
\t}

\tc.JSON(http.StatusOK, gin.H{"message": "Example deleted successfully"})
}

func (h *ExampleHandler) GetExampleByID(c *gin.Context) {
\tid := c.Param("id")

\texampleResponse, err := h.service.GetExampleByID(id)
\tif err != nil {
\t\tutils.HandleServiceError(c, err, err.Error())
\t\treturn
\t}
\tc.JSON(http.StatusOK, exampleResponse)
}

func (h *ExampleHandler) GetAllExamples(c *gin.Context) {
\tvar params dto.ExampleQueryParam
\tif !utils.BindAndValidateForm(c, &params) {
\t\treturn
\t}

\texamples, pagination, err := h.service.GetAllExamples(params)
\tif err != nil {
\t\tutils.HandleServiceError(c, err, err.Error())
\t\treturn
\t}

\tc.JSON(http.StatusOK, gin.H{
\t\t"examples":   examples,
\t\t"pagination": pagination,
\t})
}
'''

service_code = '''package services

import (
\t"server/internal/dto"
\t"server/internal/models"
\t"server/internal/repositories"
\tcustomErr "server/pkg/errors"
\t"server/pkg/utils"
)

type ExampleService interface {
\tCreateExample(req dto.CreateExampleRequest) error
\tUpdateExample(id string, req dto.UpdateExampleRequest) error
\tDeleteExample(id string) error
\tGetExampleByID(id string) (*dto.ExampleDetailResponse, error)
\tGetAllExamples(params dto.ExampleQueryParam) ([]dto.ExampleResponse, *dto.PaginationResponse, error)
}

type exampleService struct {
\trepo repositories.ExampleRepository
}

func NewExampleService(repo repositories.ExampleRepository) ExampleService {
\treturn &exampleService{repo}
}

func (s *exampleService) CreateExample(req dto.CreateExampleRequest) error {
\texample := models.Example{
\t\tExample1: req.Example1,
\t\tExample2: req.Example2,
\t\tExample3: req.Example3,
\t\tImage:    req.ImageURL,
\t}
\treturn s.repo.CreateExample(&example)
}

func (s *exampleService) UpdateExample(id string, req dto.UpdateExampleRequest) error {
\texample, err := s.repo.GetExampleByID(id)
\tif err != nil {
\t\treturn customErr.NewNotFound("example not found")
\t}

\texample.Example1 = req.Example1
\texample.Example2 = req.Example2
\texample.Example3 = req.Example3

\tif req.ImageURL != "" {
\t\t_ = utils.DeleteFromCloudinary(example.Image)
\t\texample.Image = req.ImageURL
\t}

\treturn s.repo.UpdateExample(example)
}

func (s *exampleService) DeleteExample(id string) error {
\texample, err := s.repo.GetExampleByID(id)
\tif err != nil {
\t\treturn customErr.NewNotFound("example not found")
\t}

\tif example.Image != "" {
\t\t_ = utils.DeleteFromCloudinary(example.Image)
\t}

\treturn s.repo.DeleteExample(id)
}

func (s *exampleService) GetExampleByID(id string) (*dto.ExampleDetailResponse, error) {
\texample, err := s.repo.GetExampleByID(id)
\tif err != nil {
\t\treturn nil, customErr.NewNotFound("example not found")
\t}

\treturn &dto.ExampleDetailResponse{
\t\tID:        example.ID.String(),
\t\tExample1:  example.Example1,
\t\tExample2:  example.Example2,
\t\tExample3:  example.Example3,
\t\tImage:     example.Image,
\t\tCreatedAt: example.CreatedAt.Format("2006-01-02"),
\t}, nil
}

func (s *exampleService) GetAllExamples(params dto.ExampleQueryParam) ([]dto.ExampleResponse, *dto.PaginationResponse, error) {
\texamples, total, err := s.repo.GetAllExamples(params)
\tif err != nil {
\t\treturn nil, nil, err
\t}

\tvar results []dto.ExampleResponse
\tfor _, e := range examples {
\t\tresults = append(results, dto.ExampleResponse{
\t\t\tID:        e.ID.String(),
\t\t\tExample1:  e.Example1,
\t\t\tExample2:  e.Example2,
\t\t\tExample3:  e.Example3,
\t\t\tImage:     e.Image,
\t\t\tCreatedAt: e.CreatedAt.Format("2006-01-02"),
\t\t})
\t}

\tpagination := utils.Paginate(total, params.Page, params.Limit)
\treturn results, pagination, nil
}
'''

repository_code = '''package repositories

import (
\t"server/internal/dto"
\t"server/internal/models"

\t"gorm.io/gorm"
)

type ExampleRepository interface {
\tCreateExample(example *models.Example) error
\tUpdateExample(example *models.Example) error
\tDeleteExample(id string) error
\tGetExampleByID(id string) (*models.Example, error)
\tGetAllExamples(params dto.ExampleQueryParam) ([]models.Example, int64, error)
}

type exampleRepository struct {
\tdb *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
\treturn &exampleRepository{db}
}

func (r *exampleRepository) CreateExample(example *models.Example) error {
\treturn r.db.Create(example).Error
}

func (r *exampleRepository) UpdateExample(example *models.Example) error {
\treturn r.db.Save(example).Error
}

func (r *exampleRepository) DeleteExample(id string) error {
\treturn r.db.Delete(&models.Example{}, \"id = ?\", id).Error
}

func (r *exampleRepository) GetExampleByID(id string) (*models.Example, error) {
\tvar example models.Example
\terr := r.db.First(&example, \"id = ?\", id).Error
\treturn &example, err
}

func (r *exampleRepository) GetAllExamples(params dto.ExampleQueryParam) ([]models.Example, int64, error) {
\tvar examples []models.Example
\tvar count int64

\tdb := r.db.Model(&models.Example{})

\tif params.Q != \"\" {
\t\tlike := \"%%\" + params.Q + \"%%\"
\t\tdb = db.Where(\"example1 LIKE ? OR example2 LIKE ?\", like, like)
\t}

\tdb.Count(&count)

\toffset := (params.Page - 1) * params.Limit
\tdb = db.Limit(params.Limit).Offset(offset).Order(\"created_at desc\")

\terr := db.Find(&examples).Error
\treturn examples, count, err
}
'''

# Write the files
handler_path.parent.mkdir(parents=True, exist_ok=True)
handler_path.write_text(handler_code)

service_path.parent.mkdir(parents=True, exist_ok=True)
service_path.write_text(service_code)

repo_path.parent.mkdir(parents=True, exist_ok=True)
repo_path.write_text(repository_code)

import ace_tools as tools; tools.display_dataframe_to_user(name="Generated Example Files", dataframe=None)
