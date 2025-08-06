package preview

import (
	"books/components/navigator"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/mosaic"
	"image"
)

// Model is a bubbletea model composed of a preview image on the left (rendered
// with mosaic) and another model on the right side.
type Model struct {
	preview string
	model   tea.Model
}

// New returns a new Model. If img is nil the image won't be rendered.
func New(img image.Image, model tea.Model) tea.Model {
	m := Model{model: model}

	if img != nil {
		mosaic := mosaic.New().Width(Width).Height(calcHeight(img))
		m.preview = mosaic.Render(img)
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.model.Init()
}

func (m Model) NUpdate(msg tea.Msg) (nm tea.Model, c tea.Cmd, j navigator.Jump) {
	m.model, c, j = navigator.TryNUpdate(m.model, msg)
	return m, c, j
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Padding(3, 3).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.preview,
				m.model.View(),
			),
		)
}
