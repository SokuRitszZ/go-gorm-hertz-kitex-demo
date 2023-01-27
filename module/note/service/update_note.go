package service

import (
	"context"
	"ghkd/kitex_gen/note"
	"ghkd/module/note/model/db"
)

type UpdateNoteService struct {
	ctx context.Context
}

func NewUpdateNoteService(ctx context.Context) *UpdateNoteService {
	return &UpdateNoteService{ctx}
}

func (s *UpdateNoteService) UpdateNote(req *note.UpdateNoteRequest) error {
	return db.UpdateNote(s.ctx, req.NoteId, req.UserId, req.Title, req.Content)
}
