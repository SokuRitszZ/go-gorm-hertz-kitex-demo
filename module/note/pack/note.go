package pack

import (
	"ghkd/kitex_gen/note"
	"ghkd/module/note/model/db"
)

// Note pack note info
func Note(m *db.Note) *note.Note {
	if m == nil {
		return nil
	}

	return &note.Note{
		ID: int64(m.ID),
		UserId: m.UserID,
		Title: m.Title,
		Content: m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}

// Notes pack list of note info
func Notes(ms []*db.Note) []*note.Note {
	notes := make([]*note.Note, 0)
	for _, m := range ms {
		if n := Note(m); n != nil {
			notes = append(notes, n)
		}
	}
	return notes
}

// UserIds pack UserIds
func UserIds(ms []*db.Note) []int64 {
	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserID] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
