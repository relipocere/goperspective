package main

import (
	"fmt"
	"log"
	"os"

	gp "github.com/exsocial/goperspective"
)

func main() {
	client := gp.NewClient(os.Getenv("TOKEN"))

	data := gp.AnalyzeRequest{
		Comment: gp.AnalyzeRequestComment{
			Text: "You are fucking stupid, aren't you?",
		},
		ReqAttr: map[gp.Attribute]gp.AnalyzeRequestAttr{
			gp.Toxicity: {
				ScoreThreshold: 0.5,
			},
			gp.SevereToxicity: {},
			gp.Threat:         {},
			gp.IdentityAttack: {},
		},
	}

	r, err := client.AnalyzeComment(data)
	if err != nil {
		log.Fatal(err)
	}

	//Retrieving values
	for name, as := range r.AttributeScores {
		fmt.Println(name, as.SummaryScore.Value)
	}
	//Prints:
	// IDENTITY_ATTACK 0.54661906
	// THREAT 0.34904236
	// SEVERE_TOXICITY 0.89362836
	// TOXICITY 0.9897058
}
