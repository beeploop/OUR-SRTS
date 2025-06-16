package handlers

import (
	"fmt"
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

type appHandler struct {
	sm             *session.SessionManager
	studentUseCase *student.UseCase
	programUseCase *program.UseCase
}

func NewAppHandler(
	sm *session.SessionManager,
	studentUseCase *student.UseCase,
	programUseCase *program.UseCase,
) *appHandler {
	return &appHandler{
		sm:             sm,
		studentUseCase: studentUseCase,
		programUseCase: programUseCase,
	}
}

func (h *appHandler) RenderSearch(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	students, err := h.studentUseCase.Search(ctx, c.QueryParams())
	if err != nil {
		fmt.Println("error search student: ", err.Error())
	}

	studentModels := slices.AppendSeq(
		make([]viewmodel.StudentListItem, 0),
		utils.Map(students, func(student *entities.Student) viewmodel.StudentListItem {
			return viewmodel.StudentItemFromDomain(student)
		}),
	)

	programs, err := h.programUseCase.GetProgramList(ctx)
	if err != nil {
		fmt.Println("error get program list: ", err.Error())
	}

	programModels := slices.AppendSeq(
		make([]viewmodel.Program, 0),
		utils.Map(programs, func(program *entities.Program) viewmodel.Program {
			return viewmodel.ProgramFromDomain(program)
		}),
	)

	page := app.SearchPage(admin, programModels, studentModels)
	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (h *appHandler) RenderAddStudentPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	programs, err := h.programUseCase.GetProgramList(ctx)
	if err != nil {
		fmt.Println("error get program list: ", err.Error())
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

	page := app.AddStudentPage(admin, programModels)
	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (h *appHandler) HandleAddStudent(c echo.Context) error {
	ctx := c.Request().Context()

	lastname := c.FormValue("lastname")
	firstname := c.FormValue("firstname")
	middlename := c.FormValue("middlename")
	control_number := c.FormValue("controlNumber")
	archive_location := c.FormValue("fileLocaton")
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
		archive_location,
	)

	if err := h.studentUseCase.AddStudent(ctx, student); err != nil {
		return c.Redirect(http.StatusSeeOther, "/app/add-student")
	}

	return c.Redirect(http.StatusSeeOther, "/app/add-student")
}

func (h *appHandler) RenderManageStaffPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	page := app.ManageStaffPage(admin)
	return page.Render(c.Request().Context(), c.Response().Writer)
}

func (h *appHandler) RenderRequestsPage(c echo.Context) error {
	ctx := c.Request().Context()

	admin, ok := ctx.Value(contextkeys.SessionKey).(viewmodel.Admin)
	if !ok {
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	page := app.RequestsPage(admin)
	return page.Render(c.Request().Context(), c.Response().Writer)
}
