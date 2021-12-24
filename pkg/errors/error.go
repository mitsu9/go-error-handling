package errors

import (
	"github.com/labstack/echo/v4"
)

type (
	AppError struct {
		// 返却するHTTPステータスコード
		HttpStatus int

		// アプリ独自のエラーコード
		Code int

		// エラーメッセージ
		Message string

		// 元のエラー
		err error
	}

	ErrorResponse struct {
		Result  int    `json:"result"`
		Message string `json:"message"`
	}
)

func New(httpStatus, code int, message string) *AppError {
	return &AppError{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    message,
	}
}

// error interface
func (e *AppError) Error() string {
	return e.Message
}

// echo error response
func (e *AppError) Response(c echo.Context) error {
	return c.JSON(e.HttpStatus, &ErrorResponse{
		Result:  e.Code,
		Message: e.Message,
	})
}

/*
infra層で外部APIの呼び出しエラーを独自エラーでWrapして使う想定

err := dynamo.ErrNotFound
if err != nil {
    return ERR_ROOM_NOT_FOUND.Wrap(err)
}
*/
func (e *AppError) Wrap(next error) *AppError {
	e.err = next
	return e
}
