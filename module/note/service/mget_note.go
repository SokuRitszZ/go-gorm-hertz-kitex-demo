package service

import (
	"context"
	"ghkd/kitex_gen/note"
	"ghkd/kitex_gen/user"
	"ghkd/module/note/model/db"
	"ghkd/module/note/pack"
	"ghkd/module/note/rpc"
)

type MGetNoteService struct {
	ctx context.Context
}

func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx}
}

func (s *MGetNoteService) MGetNote(req *note.MGetNoteRequest) ([]*note.Note, error)  {
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}
	uIds := pack.UserIds(noteModels)
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.Name
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
