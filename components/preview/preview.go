package preview

import (
	"books/components/navigator"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/mosaic"
	"image/jpeg"
	"os"
)

// Model contains a preview jpeg image on the left and a nested model on the right side
type Model struct {
	preview string
	model   tea.Model
}

// New returns a new Model
func New(path string, width int, model tea.Model) tea.Model {
	mosaic := mosaic.New().Width(width).Height(mosaic.Fit)

	var m Model

	// Loads the image
	f, err := os.Open(path)
	defer f.Close()
	if err == nil {
		img, _ := jpeg.Decode(f)
		if err != nil {
			m.preview = mosaic.Render(img)
		}
	}

	m.model = model
	return m
}

// Init is used by bubbletea
func (m Model) Init() tea.Cmd {
	return m.model.Init()
}

// NUpdate is used by navigator and bubbletea
func (m Model) NUpdate(msg tea.Msg) (nm tea.Model, c tea.Cmd, j navigator.Jump) {
	m.model, c, j = navigator.TryNUpdate(m.model, msg)
	return m, c, j
}

// Update is used by bubbletea
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	nm, c, _ := m.NUpdate(msg)
	return nm, c
}

// View is used by bubbletea
func (m Model) View() string {
	return lipgloss.NewStyle().
		Padding(3, 3).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				preview,
				m.model.View(),
			),
		)
}
