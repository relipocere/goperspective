package goperspective

type Attribute string

//Stable attributes
const (
	Toxicity       Attribute = "TOXICITY"
	SevereToxicity Attribute = "SEVERE_TOXICITY"
	IdentityAttack Attribute = "IDENTITY_ATTACK"
	Insult         Attribute = "INSULT"
	Profanity      Attribute = "PROFANITY"
	Threat         Attribute = "THREAT"
)

//Experemental attributes
const (
	ExpToxicity         Attribute = "TOXICITY_EXPERIMENTAL"
	ExpSevereToxicity   Attribute = "SEVERE_TOXICITY_EXPERIMENTAL"
	ExpIdentityAttack   Attribute = "IDENTITY_ATTACK_EXPERIMENTAL"
	ExpInsult           Attribute = "INSULT_EXPERIMENTAL"
	ExpProfanity        Attribute = "PROFANITY_EXPERIMENTAL"
	ExpThreat           Attribute = "THREAT_EXPERIMENTAL"
	ExpSexuallyExplicit Attribute = "SEXUALLY_EXPLICIT"
	ExpFlirtation       Attribute = "FLIRTATION"
)
