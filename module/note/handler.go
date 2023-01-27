package main

import (
	"context"
	note "ghkd/kitex_gen/note"
	"ghkd/module/note/pack"
	"ghkd/module/note/service"
	"ghkd/pkg/errno"
)

// NoteServiceImpl implements the last service interface defined in the IDL.
type NoteServiceImpl struct{}

// CreateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) CreateNote(ctx context.Context, req *note.CreateNoteRequest) (resp *note.CreateNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.CreateNoteResponse)
	
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateNoteService(ctx).CreateNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return
}

// MGetNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) MGetNote(ctx context.Context, req *note.MGetNoteRequest) (resp *note.MGetNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.MGetNoteResponse)

	if err = resp.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}

	notes, err := service.NewMGetNoteService(ctx).MGetNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ServiceErr)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	return
}

// DeleteNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) DeleteNote(ctx context.Context, req *note.DeleteNoteRequest) (resp *note.DeleteNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.DeleteNoteResponse)

	if err = resp.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewDelNoteService(ctx).DelNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// QueryNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) QueryNote(ctx context.Context, req *note.QueryNoteRequest) (resp *note.QueryNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.QueryNoteResponse)

	if err = resp.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return 
	}

	notes, total, err := service.NewQueryNoteService(ctx).QueryNote(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ServiceErr)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Notes = notes
	resp.Total = total
	return
}

// UpdateNote implements the NoteServiceImpl interface.
func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) (resp *note.UpdateNoteResponse, err error) {
	// TODO: Your code here...
	resp = new(note.UpdateNoteResponse)
	if err = resp.IsValid(); resp != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	if err = service.NewUpdateNoteService(ctx).UpdateNote(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ServiceErr)
		return 
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
