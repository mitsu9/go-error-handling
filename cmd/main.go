package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	apperr "github.com/mitsu9/go-error-handling/pkg/errors"
)

// domain
type (
	User struct {
		ID   int
		Name string
	}
	Team struct {
		ID   int
		Name string
	}
)

// infrastructure
type DBErrNotFound struct{}

func (e *DBErrNotFound) Error() string {
	return "mock db error"
}

func getUser(id int) (*User, error) {
	// sample: always fail
	return nil, &DBErrNotFound{}
}

func getTeam(id int) (*Team, error) {
	// sample: always fail
	return nil, &DBErrNotFound{}
}

// usecase
func getUserUsecase(userID int) (*User, error) {
	user, err := getUser(userID)
	if err != nil {
		return nil, apperr.ERR_USER_NOT_FOUND.Wrap(err)
	}
	return user, nil
}

func getTeamUsecase(teamID int) (*Team, error) {
	team, err := getTeam(teamID)
	if err != nil {
		// err handling を忘れたケース
		return nil, err
	}
	return team, nil
}

// handler
func getErrorResponse(c echo.Context, err error) error {
	var appError *apperr.AppError
	if errors.As(err, &appError) {
		return appError.Response(c)
	} else {
		return apperr.ERR_UNKNOWN.Wrap(err).Response(c)
	}
}

func getUserHandler(c echo.Context) error {
	userID := 1 // sample
	user, err := getUserUsecase(userID)
	if err != nil {
		return getErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, user)
}

func getTeamHandler(c echo.Context) error {
	teamID := 1 // sample
	team, err := getTeamUsecase(teamID)
	if err != nil {
		return getErrorResponse(c, err)
	}
	return c.JSON(http.StatusOK, team)
}

// main
func main() {
	e := echo.New()
	e.GET("/users/:id", getUserHandler)
	e.GET("/teams/:id", getTeamHandler)
	e.Logger.Fatal(e.Start(":80"))
}
