package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"gitlab.com/d3vus/projects/go/noteai/model"
)

type NoteHandler struct {

	store *model.NoteStore
	ai *openai.Client
}


func NewNoteHandler(store *model.NoteStore, openAiKey string) *NoteHandler{

	return &NoteHandler{

		store: store,
		ai: openai.NewClient(openAiKey),

	}

}

func (n *NoteHandler) CreateNewNote(c *gin.Context) {

	// Check the request from the User
	var note model.NoteModel

	// The request submitted by user should match the all the Info you needed
	if err := c.ShouldBindJSON(&note) ; err != nil {

		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()} )
		return
	}

	// Now generate summary using Open AI

	summary, err := n.generateSummary(note.Content)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})

	}

	note.Summary = summary

	// Generate suggestions based on the Content

	suggestion, err := n.generateSuggestions(note.Content)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
	}

	note.Suggestions = suggestion

	n.store.SaveNote(&note)
}

func (n *NoteHandler) generateSummary(content string) (string, error){

	resp, err := n.ai.CreateChatCompletion(

		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleUser,
					Content: "Please summarize this content in 2-3 sentences:"+content,
				},
			},
		},

	)

	if err != nil {

		return "",err
	}

	return resp.Choices[0].Message.Content, nil

}

func (n *NoteHandler) generateSuggestions(content string) ([]string, error) {


	resp, err := n.ai.CreateChatCompletion(

		context.Background(),
		openai.ChatCompletionRequest{

			Model:  openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				
				{
					Role: openai.ChatMessageRoleUser,
					Content: "Based on this note, suggest 3 related topics to write:"+content,
				},
			},
		},
	)

	if err != nil {

		return nil,err
	}

	suggestions := []string{resp.Choices[0].Message.Content}

	return suggestions, nil

}
 
