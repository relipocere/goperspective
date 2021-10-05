package goperspective

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PerpectiveTSuite struct {
	suite.Suite
	c *Client
}

func (ps *PerpectiveTSuite) SetupSuite() {
	client := NewClient(os.Getenv("TOKEN"))
	ps.c = client
}

func (ps *PerpectiveTSuite) TestAnalyze() {
	data := AnalyzeRequest{
		Comment: AnalyzeRequestComment{
			Text: "You are fucking stupid, aren't you?",
		},
		ReqAttr: map[Attribute]AnalyzeRequestAttr{
			Toxicity:       {},
			SevereToxicity: {},
		},
	}

	r, err := ps.c.AnalyzeComment(data)
	ps.NoError(err)
	ps.NotEmpty(r.AttributeScores[Toxicity])
	ps.NotEmpty(r.AttributeScores[SevereToxicity])
}

func (ps *PerpectiveTSuite) TestAnalyzeError() {
	data := AnalyzeRequest{
		Comment: AnalyzeRequestComment{
			Text: "You are fucking stupid, aren't you?",
		},
	}

	r, err := ps.c.AnalyzeComment(data)
	ps.Error(err)
	ps.EqualValues(400, r.Error.Code)
}

func (ps *PerpectiveTSuite) TestSuggest() {
	data := SuggestRequest{
		Comment: AnalyzeRequestComment{
			Text: "You are fucking stupid, aren't you?",
		},
		AttributeScores: map[Attribute]AttributeScore{
			Toxicity: {
				SummaryScore: Score{
					Value: 1.0,
				},
			},
		},
	}

	_, err := ps.c.SuggestCommentScore(data)
	ps.NoError(err)
}

func (ps *PerpectiveTSuite) TestSuggestError() {
	data := SuggestRequest{
		AttributeScores: map[Attribute]AttributeScore{
			Toxicity: {
				SummaryScore: Score{
					Value: 1.0,
				},
			},
		},
	}

	r, err := ps.c.SuggestCommentScore(data)
	ps.Error(err)
	ps.EqualValues(400, r.Error.Code)
}

func TestRun(t *testing.T) {
	suite.Run(t, new(PerpectiveTSuite))
}
