package glossaries

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed *
var Fs embed.FS

type Glossary []string

func Get(name string) Glossary {
	content, err := Fs.ReadFile(fmt.Sprintf("%s.json", name))
	if err != nil {
		return nil
	}

	var glo Glossary
	if err := json.Unmarshal(content, &glo); err != nil {
		return nil
	}

	return glo
}

func GetMasteredWords() []string {
	v, err := os.ReadFile("/Users/saltbo/Develop/bogit/learn-eng/glossaries/mastered/mastered.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var words []string
	if err := json.Unmarshal(v, &words); err != nil {
		fmt.Println(err)
		return nil
	}

	return words
}
