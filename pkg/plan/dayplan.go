package plan

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bonaysoft/engra/apis/client"
	"github.com/bonaysoft/engra/apis/graph/model"
	"github.com/saltbo/learn-eng/glossaries"
	"github.com/samber/lo"
)

type DayPlan struct {
	Date  time.Time
	Words []Word

	numsPerDay int
	glossaries []string
}

func NewDayPlan() *DayPlan {
	return &DayPlan{}
}

func (d *DayPlan) Run() (err error) {
	words := make([]string, 0)
	for _, glossaryName := range d.glossaries {
		words = append(words, glossaries.Get(glossaryName)...)
	}

	masteredWords := glossaries.GetMasteredWords()
	unMasteredWords, _ := lo.Difference(words, masteredWords)
	d.Date = time.Now()
	d.Words, err = d.findHighPriorityWords(unMasteredWords, masteredWords, d.numsPerDay)
	fmt.Println(len(unMasteredWords), len(words))
	return

}

func (d *DayPlan) SetNumsPerDay(numsPerDay int) {
	d.numsPerDay = numsPerDay
}

func (d *DayPlan) SelectGlossary(glossary ...string) {
	d.glossaries = append(d.glossaries, glossary...)
}

func (d *DayPlan) findHighPriorityWords(words []string, masteredWords []string, nums int) ([]Word, error) {
	// TODO: 待实现，先按顺序返回

	// endpoint := "https://engra.saltbo.cn/query"
	endpoint := "http://localhost:8081/query"
	ecdict := client.NewClient(endpoint)
	resp, err := client.Lookup(ecdict.WithContext(context.Background()), words[:nums])
	if err != nil {
		return nil, err
	}

	return lo.Map(resp.Vocabularies, func(v model.Vocabulary, index int) Word {
		return Word{
			Name:      v.Name,
			Phonetic:  v.Phonetic,
			Meaning:   strings.ReplaceAll(v.Meaning, "\\n", "\n\n"),
			Roots:     v.Roots,
			RootGraph: v.Mnemonic,
		}
	}), nil
}
