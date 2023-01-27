package errno

import (
	"errors"
	"fmt"
	"ghkd/kitex_gen/user"
)

type ErrNo struct {
	ErrCode int64
	ErrMsg string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg: msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success = NewErrNo(int64(user.ErrCode_SuccessCode), "Success")
	ServiceErr = NewErrNo(int64(user.ErrCode_ServiceErrCode), "ServiceErr")
	ParamErr = NewErrNo(int64(user.ErrCode_ParamErrCode), "ParamErr")
	UserAlreadyExistErr = NewErrNo(int64(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationErr = NewErrNo(int64(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
