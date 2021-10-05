package goperspective

//SuggestRequst is the data for the API's analyze method.
type AnalyzeRequest struct {
	Comment AnalyzeRequestComment            `json:"comment"`
	ReqAttr map[Attribute]AnalyzeRequestAttr `json:"requestedAttributes"`

	// Context map[string]interface{} has no use yet `json:"context,omitempty"`

	//Indicates if the request should return spans that describe the scores for each part of the text (currently done at per-sentence level).
	SpanAnnotations bool `json:"spanAnnotations"`
	//Stored comments will be used for future research
	DoNotStore bool `json:"doNotStore"`
	//(optional) Example: "en", "es", "fr", "de", etc. If unspecified, the API will auto-detect the comment language.
	Languages []string `json:"languages,omitempty"`
	//(optional) An opaque token that is echoed back in the response. (Note: This is not the clientId, which is automatically set through Google Cloud. This is a field users can set to help them keep track of their requests.)
	ClientToken string `json:"clientToken,omitempty"`
	//(optional) An opaque session ID. This should be set for authorship experiences by the client side so that groups of requests can be grouped together into a session.
	SessionID string `json:"sessionId,omitempty"`
	//(optional) An opaque identifier associating this comment with a particular community within your platform.
	CommunityID string `json:"communityId,omitempty"`
}

//AnalyzeRequestComment is a type for AnalyzeRequest struct.
type AnalyzeRequestComment struct {
	Text string `json:"text"`
	//(optional) Type must be either "PLAIN_TEXT" or "HTML". Currently only "PLAIN_TEXT" is supported.
	Type string `json:"type,omitempty"`
}

//AnalyzeRequestAttr is a type for AnalyzeRequest struct.
type AnalyzeRequestAttr struct {
	//(optional) Currently, only "PROBABILITY" is supported.
	ScoreType string `json:"scoreType,omitempty"`
	//(optional)The API won't return scores that are below this threshold for this attribute.
	ScoreTreshold string `json:"scoreTreshold,omitempty"`
}

//AnalyzeResponse is the response data from the API's analyze method.
type AnalyzeResponse struct {
	//The attribute names will mirror the request's requestedAttributes.
	AttributeScores map[Attribute]AttributeScore `json:"attributeScores"`
	//Mirrors the request's languages. If no languages were specified, the API returns the auto-detected language.
	Languages []string `json:"languages"`
	Error     ErrorRes `json:"error"`
}

type AttributeScore struct {
	SummaryScore Score       `json:"summaryScore"`
	SpanScores   []SpanScore `json:"spanScores"`
}

type SpanScore struct {
	Begin int   `json:"begin"`
	End   int   `json:"end"`
	Score Score `json:"score"`
}

type Score struct {
	Value int64  `json:"value"`
	Type  string `json:"type"`
}

//SuggestRequst is the data for the API's suggest method.
type SuggestRequest struct {
	Comment AnalyzeRequestComment `json:"comment"`

	//Context map[string]interface{} has no use yet `json:"context,omitempty"`

	//This holds the attribute scores that the client believes the comment should have. It has the same format as the attributeScores.
	ReqAttr map[string]AnalyzeRequestAttr `json:"requestedAttributes"`
	//(optional) Example: "en", "es", "fr", "de", etc. If unspecified, the API will auto-detect the comment language.
	Languages []string `json:"languages,omitempty"`
	//(optional) An opaque identifier associating this comment with a particular community within your platform.
	CommunityID string `json:"communityId,omitempty"`
	//(optional) An opaque token that is echoed back in the response. (Note: This is not the clientId, which is automatically set through Google Cloud. This is a field users can set to help them keep track of their requests.)
	ClientToken string `json:"clientToken,omitempty"`
}

//AnalyzeResponse is the response data from the API's suggest method.
type SuggestResposne struct {
	ClientToken string   `json:"clientToken"`
	Error       ErrorRes `json:"error"`
}

//ErrorRes is the mapped API error response.
type ErrorRes struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Status  string              `json:"status"`
	Details []map[string]string `json:"details"`
}
