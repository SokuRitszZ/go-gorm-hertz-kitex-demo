package pack

import (
	"errors"
	"ghkd/kitex_gen/note"
	"ghkd/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *note.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())	
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *note.BaseResp {
	return &note.BaseResp{
		StatusCode: err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}
