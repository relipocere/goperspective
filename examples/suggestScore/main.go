package main

import (
	"log"
	"os"

	gp "github.com/exsocial/goperspective"
)

func main() {
	client := gp.NewClient(os.Getenv("TOKEN"))

	data := gp.SuggestRequest{
		Comment: gp.AnalyzeRequestComment{
			Text: "You are fucking stupid, aren't you?",
		},
		AttributeScores: map[gp.Attribute]gp.AttributeScore{
			gp.Toxicity: {
				SummaryScore: gp.Score{
					Value: 1.0,
				},
			},
			gp.Threat: {
				SummaryScore: gp.Score{
					Value: 0,
				},
			},
		},
	}

	_, err := client.SuggestCommentScore(data)
	//If there is no error, then suggestion is accepted.
	if err != nil {
		log.Fatal(err)
	}

}
