package service

import (
	"context"
	"ghkd/kitex_gen/note"
	"ghkd/module/note/model/db"
)

type DelNoteService struct {
	ctx context.Context
}

func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{ctx}
}

func (s *DelNoteService) DelNote(req *note.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
