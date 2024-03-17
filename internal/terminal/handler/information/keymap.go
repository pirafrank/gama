package information

import (
	"fmt"

	teakey "github.com/charmbracelet/bubbles/key"
	pkgconfig "github.com/termkit/gama/pkg/config"
)

type keyMap struct {
	SwitchTabRight teakey.Binding
	Quit           teakey.Binding
}

func (k keyMap) ShortHelp() []teakey.Binding {
	return []teakey.Binding{k.SwitchTabRight, k.Quit}
}

func (k keyMap) FullHelp() [][]teakey.Binding {
	return [][]teakey.Binding{
		{k.SwitchTabRight},
		{k.Quit},
	}
}

var keys = func() keyMap {
	cfg, err := pkgconfig.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	var switchTabRight = fmt.Sprintf("%s", cfg.Shortcuts.SwitchTabRight)

	return keyMap{
		SwitchTabRight: teakey.NewBinding(
			teakey.WithKeys(""), // help-only binding
			teakey.WithHelp(switchTabRight, "next tab"),
		),
		Quit: teakey.NewBinding(
			teakey.WithKeys("q", cfg.Shortcuts.Quit),
			teakey.WithHelp("q", "quit"),
		),
	}
}()

func (m *ModelInfo) ViewHelp() string {
	return m.Help.View(m.Keys)
}
