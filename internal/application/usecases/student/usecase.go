package student

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/beeploop/our-srts/internal/application/interfaces"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/domain/repositories"
	"github.com/beeploop/our-srts/internal/pkg/utils"
)

type UseCase struct {
	studentRepo      repositories.StudentRepository
	documentRepo     repositories.DocumentRepository
	documentTypeRepo repositories.DocumentTypeRepository
	fs               interfaces.Storage
}

func NewUseCase(
	studentRepo repositories.StudentRepository,
	documentRepo repositories.DocumentRepository,
	documentTypeRepo repositories.DocumentTypeRepository,
	fs interfaces.Storage,
) *UseCase {
	return &UseCase{
		studentRepo:      studentRepo,
		documentRepo:     documentRepo,
		documentTypeRepo: documentTypeRepo,
		fs:               fs,
	}
}

func (u *UseCase) AddStudent(ctx context.Context, student *entities.Student) error {
	if err := student.Validate(); err != nil {
		return err
	}

	if _, err := u.studentRepo.Create(ctx, student); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) Search(ctx context.Context, params url.Values) ([]*entities.Student, error) {
	if !params.Has("query") || params.Get("query") == "" {
		return make([]*entities.Student, 0), nil
	}

	if !params.Has("program") {
		params.Set("program", "all")
	}

	if !params.Has("type") {
		params.Set("type", string(repositories.SEARCH_BY_FIRSTNAME))
	}

	filter := repositories.StudentFilter{
		Query:      params.Get("query"),
		SearchType: repositories.SearchType(params.Get("type")),
		ProgramID:  utils.Ternary(params.Get("program") != "all", params.Get("program"), ""),
	}

	return u.studentRepo.Search(ctx, filter)
}

func (u *UseCase) GetStudent(ctx context.Context, controlNumber string) (*entities.Student, error) {
	if controlNumber == "" {
		return nil, errors.New("invalid control number")
	}

	return u.studentRepo.FindByControlNumber(ctx, controlNumber)
}

func (u *UseCase) UpdateStudent(ctx context.Context, student *entities.Student) error {
	if err := student.Validate(); err != nil {
		return err
	}

	existing, err := u.studentRepo.FindByControlNumber(ctx, student.ControlNumber)
	if err != nil {
		return err
	}

	if err := existing.FullUpdate(student); err != nil {
		return err
	}

	return u.studentRepo.Save(ctx, existing)
}

func (u *UseCase) UploadDocument(ctx context.Context, studentControlNumber, docType, filename string, content *multipart.FileHeader) error {
	if studentControlNumber == "" {
		return errors.New("invalid control number")
	}
	if docType == "" {
		return errors.New("invalid document type")
	}

	file, err := content.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	student, err := u.studentRepo.FindByControlNumber(ctx, studentControlNumber)
	if err != nil {
		return err
	}

	documentType, err := u.documentTypeRepo.FindByTitle(ctx, docType)
	if err != nil {
		return err
	}

	{
		ext := filepath.Ext(content.Filename)
		if documentType.Title == "picture" {
			validFiles := []string{".png", ".jpg", ".jpeg"}

			if !slices.Contains(validFiles, ext) {
				return errors.New("invalid file type")
			}
		} else {
			if ext != ".pdf" {
				return errors.New("invalid file type")
			}
		}
	}

	filename = fmt.Sprintf("%s%s", filename, filepath.Ext(content.Filename))
	folder := fmt.Sprintf("%s_%s", student.ControlNumber, utils.WhiteSpaceToUnderscore(student.LastName))
	filepath := u.fs.ConstructPath(ctx, folder, filename)
	tempPath, err := u.fs.Save(ctx, filepath, file)
	if err != nil {
		return err
	}
	if err := u.fs.Delete(ctx, tempPath); err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		filepath = strings.ReplaceAll(filepath, "\\", "/")
	}

	document := entities.NewDocument(*documentType, filename, filepath)
	if _, err := u.studentRepo.UploadDocument(ctx, document, &student.Envelope); err != nil {
		return err
	}

	return nil
}

func (u *UseCase) ReuploadDocument(ctx context.Context, filename, documentID string, content *multipart.FileHeader) error {
	if documentID == "" {
		return errors.New("invalid document ID")
	}

	file, err := content.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	existingDocument, err := u.documentRepo.FindByID(ctx, documentID)
	if err != nil {
		return err
	}

	if filename != "" {
		filename = fmt.Sprintf("%s%s", filename, filepath.Ext(content.Filename))
		if err := existingDocument.UpdateFilename(filename); err != nil {
			return err
		}
	}

	tempPath, err := u.fs.Save(ctx, existingDocument.StoragePath, file)
	if err != nil {
		return err
	}
	if err := u.fs.Delete(ctx, tempPath); err != nil {
		return err
	}

	if err := u.documentRepo.Save(ctx, existingDocument); err != nil {
		return err
	}

	return nil
}
