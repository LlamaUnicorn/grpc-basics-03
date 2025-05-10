package note

import (
	"github.com/llamaunicorn/grpc-basics-03/internal/service"
	desc "github.com/llamaunicorn/grpc-basics-03/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
