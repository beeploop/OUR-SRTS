package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"slices"

	"github.com/beeploop/our-srts/internal/application/usecases/program"
	"github.com/beeploop/our-srts/internal/application/usecases/student"
	"github.com/beeploop/our-srts/internal/domain/entities"
	"github.com/beeploop/our-srts/internal/infrastructure/http/viewmodel"
	"github.com/beeploop/our-srts/internal/infrastructure/session"
	"github.com/beeploop/our-srts/internal/pkg/contextkeys"
	"github.com/beeploop/our-srts/internal/pkg/utils"
	"github.com/beeploop/our-srts/web/views/pages/app"
	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	studentUseCase *student.UseCase
	programUseCase *program.UseCase
	sessionManager *session.SessionManager
}

func NewStudentHandler(
	studentUseCase *student.UseCase,
	programUseCase *program.UseCase,
	sessionManager *session.SessionManager,
) *studentHandler {
	return &studentHandler{
		studentUseCase: studentUseCase,
		programUseCase: programUseCase,
		sessionManager: sessionManager,
	}
}

func (h *studentHandler) RenderSearch(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	students, err := h.studentUseCase.Search(ctx, c.QueryParams())
	if err != nil {
		slog.Error("Search Student Failed", "error", err.Error())
	}

	studentModels := slices.AppendSeq(
		make([]viewmodel.StudentListItem, 0),
		utils.Map(students, func(student *entities.Student) viewmodel.StudentListItem {
			return viewmodel.StudentItemFromDomain(student)
		}),
	)

	programs, err := h.programUseCase.GetProgramList(ctx)
	if err != nil {
		slog.Error("Get Program List Failed", "error", err.Error())
	}

	programModels := slices.AppendSeq(
		make([]viewmodel.Program, 0),
		utils.Map(programs, func(program *entities.Program) viewmodel.Program {
			return viewmodel.ProgramFromDomain(program)
		}),
	)

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if ok {
		ctx = context.WithValue(ctx, contextkeys.ToastKey, toast)
	}

	page := app.SearchPage(admin, programModels, studentModels)
	return page.Render(ctx, c.Response().Writer)
}

func (h *studentHandler) RenderStudentPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	control_number := c.Param("controlNumber")

	student, err := h.studentUseCase.GetStudent(ctx, control_number)
	if err != nil {
		slog.Error("Get Student Failed", "error", err.Error())
		page := app.StudentPage(admin, viewmodel.Student{}, make([]viewmodel.ProgramWithMajors, 0))
		return page.Render(ctx, c.Response().Writer)
	}

	studentModel := viewmodel.StudentFromDomain(student)

	programs, err := h.programUseCase.GetProgramList(ctx)
	if err != nil {
		slog.Error("Get Program List Failed", "error", err.Error())
	}

	programModels := slices.AppendSeq(
		make([]viewmodel.ProgramWithMajors, 0),
		utils.Map(programs, func(program *entities.Program) viewmodel.ProgramWithMajors {
			return viewmodel.ProgramWithMajors{
				Program: viewmodel.ProgramFromDomain(program),
				Majors: slices.AppendSeq(
					make([]viewmodel.Major, 0),
					utils.Map(program.Majors, func(major entities.Major) viewmodel.Major {
						return viewmodel.MajorFromDomain(&major)
					}),
				),
			}
		}),
	)

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if ok {
		ctx = context.WithValue(ctx, contextkeys.ToastKey, toast)
	}

	page := app.StudentPage(admin, studentModel, programModels)
	return page.Render(ctx, c.Response().Writer)
}

func (h *studentHandler) RenderAddStudentPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	programs, err := h.programUseCase.GetProgramList(ctx)
	if err != nil {
		slog.Error("Get Program List Failed", "error", err.Error())
	}

	programModels := slices.AppendSeq(
		make([]viewmodel.ProgramWithMajors, 0),
		utils.Map(programs, func(program *entities.Program) viewmodel.ProgramWithMajors {
			return viewmodel.ProgramWithMajors{
				Program: viewmodel.ProgramFromDomain(program),
				Majors: slices.AppendSeq(
					make([]viewmodel.Major, 0),
					utils.Map(program.Majors, func(major entities.Major) viewmodel.Major {
						return viewmodel.MajorFromDomain(&major)
					}),
				),
			}
		}),
	)

	toast, ok := h.sessionManager.GetFlash(c.Response().Writer, c.Request())
	if ok {
		ctx = context.WithValue(ctx, contextkeys.ToastKey, toast)
	}

	page := app.AddStudentPage(admin, programModels)
	return page.Render(ctx, c.Response().Writer)
}

func (h *studentHandler) HandleAddStudent(c echo.Context) error {
	ctx := c.Request().Context()

	lastname := c.FormValue("lastname")
	firstname := c.FormValue("firstname")
	middlename := c.FormValue("middlename")
	control_number := c.FormValue("controlNumber")
	file_location := c.FormValue("fileLocation")
	student_type := c.FormValue("type")
	civil_status := c.FormValue("civilStatus")
	program_id := c.FormValue("program")
	major_id := c.FormValue("major")

	student := entities.NewStudent(
		control_number,
		firstname,
		middlename,
		lastname,
		"",
		entities.StudentType(student_type),
		entities.CivilStatus(civil_status),
		program_id,
		major_id,
		file_location,
	)

	if err := h.studentUseCase.AddStudent(ctx, student); err != nil {
		slog.Error("Add Student Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/app/add-student")
	}

	toast := viewmodel.NewSuccessToast("student added to the system")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		slog.Error("Flash Message", "error", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/app/add-student")
}

func (h *studentHandler) HandleUpdateStudent(c echo.Context) error {
	ctx := c.Request().Context()

	lastname := c.FormValue("lastname")
	firstname := c.FormValue("firstname")
	middlename := c.FormValue("middlename")
	control_number := c.FormValue("controlNumber")
	file_location := c.FormValue("fileLocation")
	student_type := c.FormValue("type")
	civil_status := c.FormValue("civilStatus")
	program_id := c.FormValue("program")
	major_id := c.FormValue("major")

	student := entities.NewStudent(
		control_number,
		firstname,
		middlename,
		lastname,
		"",
		entities.StudentType(student_type),
		entities.CivilStatus(civil_status),
		program_id,
		major_id,
		file_location,
	)

	if err := h.studentUseCase.UpdateStudent(ctx, student); err != nil {
		slog.Error("Update Student Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
	}

	toast := viewmodel.NewSuccessToast("student data updated")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		slog.Error("Flash Message", "error", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
}

func (h *studentHandler) HandleUploadDocument(c echo.Context) error {
	ctx := c.Request().Context()

	controlNumber := c.FormValue("controlNumber")
	documentType := c.FormValue("documentType")
	filename := utils.Ternary(c.FormValue("filename") == "", documentType, c.FormValue("filename"))

	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("Extracting Form File Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
	}

	if err := h.studentUseCase.UploadDocument(ctx, controlNumber, documentType, filename, file); err != nil {
		slog.Error("Upload Document Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Upload Document Failed", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
	}

	toast := viewmodel.NewSuccessToast("upload successful")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
}

func (h *studentHandler) HandleReuploadDocument(c echo.Context) error {
	ctx := c.Request().Context()

	documentID := c.FormValue("documentID")
	filename := c.FormValue("filename")

	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("Extracting Form File Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, c.Request().Referer())
	}

	if err := h.studentUseCase.ReuploadDocument(ctx, filename, documentID, file); err != nil {
		slog.Error("Reupload Document Failed", "error", err.Error())
		toast := viewmodel.NewErrorToast(err.Error())
		if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
			slog.Error("Flash Message", "error", err.Error())
		}
		return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
	}

	toast := viewmodel.NewSuccessToast("reupload successful")
	if err := h.sessionManager.SetFlash(c.Response().Writer, c.Request(), toast.ToJson()); err != nil {
		fmt.Println("error setting flash: ", err.Error())
	}

	return c.Redirect(http.StatusSeeOther, utils.StripQueryParams(c.Request().Referer()))
}
