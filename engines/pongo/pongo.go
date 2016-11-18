package pongo

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/ghetzel/diecast/engines"
	"io"
	"os"
)

var isInitialized bool

type PongoTemplate struct {
	engines.Template
	template *pongo2.Template
}

func Initialize() error {
	if !isInitialized {
		for name, funcdef := range GetBaseFunctions() {
			pongo2.RegisterFilter(name, funcdef)
		}

		isInitialized = true
	}

	return nil
}

func New() engines.ITemplate {
	return &PongoTemplate{}
}

func (self *PongoTemplate) Load(key string) error {
	tplPath := fmt.Sprintf("%s/%s.pongo", self.GetTemplateDir(), key)

	if _, err := os.Stat(tplPath); err == nil {
		if tpl, err := pongo2.FromFile(tplPath); err == nil {
			self.template = tpl
			return nil
		} else {
			return err
		}
		return nil
	} else {
		return err
	}
}

func (self *PongoTemplate) Render(output io.Writer, payload map[string]interface{}) error {
	if self.template != nil {
		return self.template.ExecuteWriter(pongo2.Context(payload), output)
	} else {
		return fmt.Errorf("Cannot execute nil template")
	}
}