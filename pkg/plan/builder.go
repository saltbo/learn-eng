package plan

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
)

type Builder struct {
	plan Plan
}

func NewBuilder(plan Plan) *Builder {
	return &Builder{
		plan: plan,
	}
}

// Build builds a plan
func (p *Builder) Build() error {
	err := p.plan.Run()
	if err != nil {
		return err
	}

	return p.generate()
}

func (p *Builder) generate() error {
	now := time.Now()
	filePath := now.Format("plans/0601")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.MkdirAll(filePath, 0755); err != nil {
			return err
		}
	}

	f, err := os.Create(filepath.Join(filePath, fmt.Sprintf("%d.md", now.Day())))
	if err != nil {
		return err
	}

	t, err := template.New("plan.md").Funcs(sprig.TxtFuncMap()).ParseFiles("pkg/plan/plan.md")
	if err != nil {
		return err
	}

	return t.Execute(f, p.plan)
}
