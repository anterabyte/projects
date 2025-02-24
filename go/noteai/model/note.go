package model

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type NoteModel struct{

	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
	Suggestions []string `json:"suggestions"`
	CreatedAt time.Time `json:"createdat"`

}

type NoteStore struct {

	store map[string]NoteModel

}

func NewNoteStore() *NoteStore {

	return &NoteStore{

		store: make(map[string]NoteModel),
	}

}

func (ns *NoteStore) SaveNote(note *NoteModel) string {

	// Set ID set first, if not already present 
	if note.ID == "" {

		note.ID = uuid.New().String()

	}

	// Set Creation Time, if not already given

	if note.CreatedAt.IsZero() {

		note.CreatedAt = time.Now()
		
	}

	// Now Store the note in the Note Store

	ns.store[note.ID] = *note

	return fmt.Sprintf("Your note is saved and here is the ID generated %s",note.ID)

}

func (ns *NoteStore) GetNote(id string) NoteModel {

	// Check if Note exist with that ID
	note, err := ns.store[id]

	if !err {

		fmt.Println("The note could not be found",err)
		os.Exit(1)
	}

	return note
}

func (ns *NoteStore) ListNote() []NoteModel {

	noteList :=  make([]NoteModel,0,len(ns.store))
	
	for _,v := range ns.store {

		noteList = append(noteList, v)

	}

	return noteList
}

func (ns *NoteStore) DeleteNotes(id string) {

	// Check if Note with ID Exist

	_, err := ns.store[id]

	if !err {

		fmt.Println("Note with following ID does not exist")
		os.Exit(1)
	}

	delete(ns.store, id)

}
