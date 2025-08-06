package navigator

import tea "github.com/charmbracelet/bubbletea"

// Jump is used by NModels to signal the navigator that it needs to jump
// to the previous or next model. The Prev and Next values indicates which is
// the target model; NextArg will be passed to the function that returns the
// next model.
type Jump struct {
	Prev    bool
	Next    bool
	NextArg string
}

// NModel is a bubbletea model that also implements NUpdate, which can return
// a jump, alongside the model and the command.
type NModel interface {
	tea.Model
	NUpdate(tea.Msg) (tea.Model, tea.Cmd, Jump)
}

// TryNUpdate tries to execute the NUpdate method on a model. If that model
// is not an NModel, it executes the Update method and returns a void jump.
func TryNUpdate(mod tea.Model, msg tea.Msg) (m tea.Model, c tea.Cmd, j Jump) {
	if nmod, ok := mod.(NModel); ok {
		m, c, j = nmod.NUpdate(msg)
	} else {
		m, c = mod.Update(msg)
	}

	return m, c, j
}
