package utils

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"server/internal/config"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

const MaxFileSize = 2 * 1024 * 1024

var AllowedImageTypes = []string{"image/jpeg", "image/png", "image/gif", "image/webp"}

func UploadToCloudinary(file io.Reader) (string, error) {
	ctx := context.Background()

	folder := os.Getenv("CLOUDINARY_FOLDER_NAME")

	uploadResult, err := config.Cloud.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder:         folder,
		Transformation: "w_500,h_500,c_limit,f_webp",
	})

	if err != nil {
		log.Printf("failed to upload file to Cloudinary %v :", err)
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func DeleteFromCloudinary(imageURL string) error {
	ctx := context.Background()

	publicID, err := extractPublicID(imageURL)
	if err != nil {
		return err
	}

	deleteResult, err := config.Cloud.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: os.Getenv("CLOUDINARY_FOLDER_NAME") + "/" + publicID,
	})
	if err != nil {
		log.Printf("failed to delete file from Cloudinary: %v", err)
		return errors.New("failed to delete asset from Cloudinary")
	}

	if deleteResult.Result != "ok" {
		return errors.New("failed to delete asset from Cloudinary: not found")
	}

	return nil
}

func extractPublicID(imageURL string) (string, error) {
	parts := strings.Split(imageURL, "/")
	if len(parts) == 0 {
		return "", errors.New("invalid image URL")
	}

	fileName := parts[len(parts)-1]

	publicID := strings.Split(fileName, ".")[0]
	if publicID == "" {
		return "", errors.New("failed to extract public_id from image URL")
	}

	return publicID, nil
}

func ValidateImageFile(fileHeader *multipart.FileHeader) error {
	if fileHeader.Size > MaxFileSize {
		return errors.New("file size is too large, maximum 1MB")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return err
	}

	mimeType := http.DetectContentType(buffer)
	if !isAllowedImageType(mimeType) {
		return fmt.Errorf("invalid file format: %s. Only JPG, PNG, GIF, and WEBP are allowed", mimeType)
	}

	return nil
}

func isAllowedImageType(fileType string) bool {
	for _, allowedType := range AllowedImageTypes {
		if strings.EqualFold(fileType, allowedType) {
			return true
		}
	}
	return false
}

func UploadImageWithValidation(fileHeader *multipart.FileHeader) (string, error) {
	if fileHeader == nil {
		return "", errors.New("no image file provided")
	}

	if fileHeader.Size > MaxFileSize {
		return "", errors.New("file size is too large, maximum 1MB")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	if !isAllowedImageType(mimeType) {
		return "", fmt.Errorf("invalid file format: %s. Only JPG, PNG, GIF, and WEBP are allowed", mimeType)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	return UploadToCloudinary(file)
}

func CleanupImageOnError(imageURL string) {
	if imageURL != "" {
		_ = DeleteFromCloudinary(imageURL)
	}
}

func UploadMultipleImagesWithValidation(fileHeaders []*multipart.FileHeader) ([]string, error) {
	var uploadedURLs []string

	for _, fileHeader := range fileHeaders {
		if fileHeader == nil {
			return nil, errors.New("one of the images is missing")
		}

		if err := ValidateImageFile(fileHeader); err != nil {
			return nil, err
		}

		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		imageURL, err := UploadToCloudinary(file)
		if err != nil {
			return nil, err
		}

		uploadedURLs = append(uploadedURLs, imageURL)
	}

	return uploadedURLs, nil
}

func CleanupImagesOnError(imageURLs []string) {
	for _, url := range imageURLs {
		if url != "" {
			_ = DeleteFromCloudinary(url)
		}
	}
}

func UploadToLocal(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	uploadPath := "./uploads/"
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
	filePath := filepath.Join(uploadPath, filename)

	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		return "", err
	}

	_, err = out.ReadFrom(file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
