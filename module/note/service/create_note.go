package service

import (
	"context"
	"ghkd/kitex_gen/note"
	"ghkd/module/note/model/db"
)

type CreateNoteService struct {
	ctx context.Context
}

// NewCreateNoteService new CreateNoteService
func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx}
}

func (s *CreateNoteService) CreateNote(req *note.CreateNoteRequest) error {
	noteModel := &db.Note{
		UserID: req.UserId,
		Title: req.Title,
		Content: req.Content,
	}
	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
