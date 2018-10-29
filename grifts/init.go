package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/one-tec-lab/go-app-start/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
