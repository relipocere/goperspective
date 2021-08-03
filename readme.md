# GO Perspective API Client
This library is an implimentation of [Perspective API](https://developers.perspectiveapi.com/s/) client in Go.
[Attributes and supported languages](https://developers.perspectiveapi.com/s/about-the-api-attributes-and-languages).
Methods and fields can be found [here](https://developers.perspectiveapi.com/s/about-the-api-methods).

## Initizalizing client
```go
c := goperspective.NewClient("YOUR TOKEN HERE")

```

## Analyzing Comments

```go

data := map[string]interface{}{
		"comment": map[string]interface{}{
			"text": "Get the fuck out of here",
		},
		"requestedAttributes": map[string]interface{}{
			"TOXICITY": map[string]interface{}{
				"scoreThreshold": 0,
			},
			"IDENTITY_ATTACK": map[string]interface{}{
				"scoreThreshold": 0,
			},
			"INSULT": map[string]interface{}{
				"scoreThreshold": 0,
			},
			"THREAT": map[string]interface{}{
				"scoreThreshold": 0,
			},
		},
	}

	obj, err := c.AnalyzeComment(data)
	if err != nil {
		log.Fatal(err)
	}

	//Retrieving values
	attrScores := obj["attributeScores"].(map[string]interface{})
	for name, attr := range attrScores {
		sumScore := attr.(map[string]interface{})["summaryScore"]
		value := sumScore.(map[string]interface{})["value"]
		fmt.Println(name, value)
	}

```
### Prints
```

IDENTITY_ATTACK 0.37591976
INSULT 0.7148725
THREAT 0.46167082
TOXICITY 0.9328236

```
## Suggesting comment score
```go

data := map[string]interface{}{
		"comment": map[string]interface{}{
			"text": "I will break your knee caps",
		},
		"attributeScores": map[string]interface{}{
			"THREAT": map[string]interface{}{
				"summaryScore": map[string]interface{}{
					"value": 1.0,
					"type":  "PROBABILITY",
				},
			},
		},
	}

	_, err := c.SuggestCommentScore(data)
	if err != nil {
		log.Fatal(err)
	}

```
