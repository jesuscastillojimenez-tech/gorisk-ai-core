package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// AIResponse defines the strict JSON structure we expect from Gemini
type AIResponse struct {
	Score  float64 `json:"score"`
	Status string  `json:"status"` // APPROVED or REJECTED
}

// EvaluateRisk calls the Gemini API to analyze the credit application with exponential retry logic
func EvaluateRisk(income float64, requestedAmount float64) AIResponse {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Println("WARNING: GEMINI_API_KEY is not set. Simulating AI response...")
		return AIResponse{Score: 50.0, Status: "PENDING"}
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("ERROR: Failed to connect to Gemini API: %v", err)
		return AIResponse{Score: 0, Status: "REJECTED"}
	}
	defer client.Close()

	// Using the updated Gemini 3.5 Flash model available in 2026
	model := client.GenerativeModel("gemini-3.5-flash")
	model.ResponseMIMEType = "application/json" // Enforce JSON output

	// Strict financial prompt in English
	prompt := fmt.Sprintf(`Act as a senior bank risk analyst.
The client has a monthly income of $%.2f and is requesting a credit amount of $%.2f.
Evaluate the financial risk. You must return ONLY a valid JSON object with this exact structure, with no markdown formatting or extra text:
{"score": [number from 0 to 100], "status": ["APPROVED" or "REJECTED"]}
Strict business rule: If the requested amount exceeds 10 times the monthly income, the status MUST be REJECTED and the score MUST be under 30.`, income, requestedAmount)

	// Retry configuration for Free Tier rate limit management
	const maxRetries = 3
	retryDelay := 10 * time.Second

	var resp *genai.GenerateContentResponse

	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err = model.GenerateContent(ctx, genai.Text(prompt))
		if err == nil {
			// Request succeeded, exit retry loop
			break
		}

		log.Printf("WARNING: AI evaluation attempt %d/%d failed: %v", attempt, maxRetries, err)

		if attempt < maxRetries {
			log.Printf("INFO: Rate limit reached. Retrying automatically in %v...", retryDelay)
			time.Sleep(retryDelay)
			retryDelay *= 2 // Exponential backoff: double wait time for the next retry (10s, 20s...)
		}
	}

	// If all retries fail, return a fallback response
	if err != nil {
		log.Printf("ERROR: AI evaluation failed after %d retries: %v", maxRetries, err)
		return AIResponse{Score: 0, Status: "REJECTED"}
	}

	var result AIResponse
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		part := resp.Candidates[0].Content.Parts[0]
		if textPart, ok := part.(genai.Text); ok {
			jsonStr := strings.TrimSpace(string(textPart))

			// Attempt to unmarshal the JSON response
			if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
				log.Printf("ERROR: Failed to parse Gemini JSON output: %v. Raw output: %s", err, jsonStr)
			}
		}
	}

	// Fallback to ensure status validity
	if result.Status != "APPROVED" && result.Status != "REJECTED" {
		result.Status = "REJECTED"
	}

	log.Printf("SUCCESS: AI Evaluation completed. Score: %.2f, Status: %s", result.Score, result.Status)
	return result
}
