package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	ContentStyle        lipgloss.Style
	SectionTitleStyle   lipgloss.Style
	ProjectHeadingStyle lipgloss.Style
	ProjectDescStyle    lipgloss.Style
	TechLabelStyle      lipgloss.Style
	SkillBoxStyle       lipgloss.Style
	BannerStyle         lipgloss.Style
	ContactHeadingStyle lipgloss.Style
	ContactItemStyle    lipgloss.Style
	ContactLabelStyle   lipgloss.Style
	ContactValueStyle   lipgloss.Style
	CallToActionStyle   lipgloss.Style
}

type MenuStyles struct {
	TitleStyle       lipgloss.Style
	NavItemStyle     lipgloss.Style
	SelectedNavStyle lipgloss.Style
	BorderStyle      lipgloss.Style
	FooterStyle      lipgloss.Style
	CenterStyle      lipgloss.Style
}

func NewStyles(fixedWidth int) *Styles {
	return &Styles{
		ContentStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(1, 0),

		SectionTitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Bold(true).
			Padding(0, 0, 1, 0),

		ProjectHeadingStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Bold(true).
			Padding(0, 0, 0, 0),

		ProjectDescStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			Padding(0, 0, 1, 0),

		TechLabelStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Italic(true),

		SkillBoxStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#666666")).
			Padding(0, 1).
			Margin(0, 1).
			Foreground(lipgloss.Color("#FFFFFF")),

		BannerStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Align(lipgloss.Center).
			Bold(true),

		ContactHeadingStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Bold(true).
			Padding(0, 0, 0, 0),

		ContactItemStyle: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF6B35")).
			Padding(1, 2).
			Margin(0, 1, 1, 0).
			Foreground(lipgloss.Color("#FFFFFF")).
			Width(fixedWidth/2 - 8),

		ContactLabelStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Bold(true),

		ContactValueStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			Padding(0, 0, 0, 1),

		CallToActionStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#FF6B35")).
			Padding(1, 2).
			Margin(1, 0).
			Bold(true).
			Align(lipgloss.Center).
			Width(fixedWidth - 10),
	}
}

func NewMenuStyles(width, height, fixedWidth, fixedHeight int) *MenuStyles {
	return &MenuStyles{
		TitleStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true).
			Padding(0, 1),

		NavItemStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			Margin(0, 1),

		SelectedNavStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#FF6B35")).
			Padding(0, 2).
			Margin(0, 1).
			Bold(true),

		BorderStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FF6B35")).
			Padding(1).
			Width(fixedWidth).
			Height(fixedHeight),

		FooterStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Align(lipgloss.Center).
			Width(fixedWidth - 4),

		CenterStyle: lipgloss.NewStyle().
			Width(width).
			Height(height).
			Align(lipgloss.Center, lipgloss.Center),
	}
}
