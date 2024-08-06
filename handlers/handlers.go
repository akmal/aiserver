package handlers

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/akmal/aiserver/ollama"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//go:embed templates/prompt.tmpl
var promptTemplate string

/**
 * PromptHandler handles the /prompt endpoint.
 */
func PromptHandler(c *gin.Context, debug bool) {
	type PromptRequest struct {
		Participants []json.RawMessage `json:"participants"`
		Interview    []json.RawMessage `json:"interview"`
	}

	var req PromptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if debug {
		fmt.Println("Participants:")
		for _, participant := range req.Participants {
			var participantMap map[string]interface{}
			if err := json.Unmarshal(participant, &participantMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			prettyJSON, err := json.MarshalIndent(participantMap, "", "  ")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			fmt.Println(string(prettyJSON))
		}

		fmt.Println("Interview:")
		for _, question := range req.Interview {
			var questionMap map[string]interface{}
			if err := json.Unmarshal(question, &questionMap); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			prettyJSON, err := json.MarshalIndent(questionMap, "", "  ")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			fmt.Println(string(prettyJSON))
		}
	}

	tmpl, err := template.New("prompt").Parse(promptTemplate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	participantsStr := make([]string, len(req.Participants))
	for i, participant := range req.Participants {
		var participantMap map[string]interface{}
		if err := json.Unmarshal(participant, &participantMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		participantJSON, err := json.Marshal(participantMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		participantsStr[i] = string(participantJSON)
	}

	interviewStr := make([]string, len(req.Interview))
	for i, question := range req.Interview {
		var questionMap map[string]interface{}
		if err := json.Unmarshal(question, &questionMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		questionJSON, err := json.Marshal(questionMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		interviewStr[i] = string(questionJSON)
	}

	data := struct {
		Participants template.HTML
		Interview    template.HTML
	}{
		Participants: template.HTML(fmt.Sprintf("%v", participantsStr)),
		Interview:    template.HTML(fmt.Sprintf("%v", interviewStr)),
	}

	var promptBuffer bytes.Buffer
	if err := tmpl.Execute(&promptBuffer, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	prompt := promptBuffer.String()
	fmt.Printf("Prompt being passed to Ollama: %s\n", prompt)

	response, err := ollama.QueryOllama(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
