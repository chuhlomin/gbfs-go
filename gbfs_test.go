package gbfs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalLanguageFeeds(t *testing.T) {
	jsonPart := `{
		"en": {
			"feeds": [
				{
					"name": "system_information",
					"url": "https://www.example.com/gbfs/1/en/system_information"
				},
				{
					"name": "station_information",
					"url": "https://www.example.com/gbfs/1/en/station_information"
				}
			]
		},
		"fr" : {
			"feeds": [
				{
					"name": "system_information",
					"url": "https://www.example.com/gbfs/1/fr/system_information"
				},
				{
					"name": "station_information",
					"url": "https://www.example.com/gbfs/1/fr/station_information"
				}
			]
		}
	}`

	var actual LanguageFeeds
	if err := json.Unmarshal([]byte(jsonPart), &actual); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	var expected LanguageFeeds = map[string]DataFeeds{
		"en": {
			Feeds: []Feed{
				{
					Name: "system_information",
					URL:  "https://www.example.com/gbfs/1/en/system_information",
				},
				{
					Name: "station_information",
					URL:  "https://www.example.com/gbfs/1/en/station_information",
				},
			},
		},
		"fr": {
			Feeds: []Feed{
				{
					Name: "system_information",
					URL:  "https://www.example.com/gbfs/1/fr/system_information",
				},
				{
					Name: "station_information",
					URL:  "https://www.example.com/gbfs/1/fr/station_information",
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestUnmarshalLanguageFeedsLegacy(t *testing.T) {
	body := `{
		"feeds": [
			{
				"name": "system_information",
				"url": "https://www.example.com/gbfs/1/en/system_information"
			},
			{
				"name": "station_information",
				"url": "https://www.example.com/gbfs/1/en/station_information"
			}
		]
	}`

	var actual LanguageFeeds
	if err := json.Unmarshal([]byte(body), &actual); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	var expected LanguageFeeds = map[string]DataFeeds{
		"default": {
			Feeds: []Feed{
				{
					Name: "system_information",
					URL:  "https://www.example.com/gbfs/1/en/system_information",
				},
				{
					Name: "station_information",
					URL:  "https://www.example.com/gbfs/1/en/station_information",
				},
			},
		},
	}

	assert.Equal(t, expected, actual)
}
