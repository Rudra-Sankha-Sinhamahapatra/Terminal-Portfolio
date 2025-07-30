package pages

import (
	"terminal-portfolio/banner"
	"terminal-portfolio/menu/styles"

	"github.com/charmbracelet/lipgloss"
)

type HomePage struct{}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (h *HomePage) GetContent(fixedWidth int, styles *styles.Styles) string {
	responsiveBannerStyle := styles.BannerStyle.Width(fixedWidth - 10)
	bannerText := responsiveBannerStyle.Render(banner.Banner())

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Align(lipgloss.Center).
		Bold(true).
		Padding(1, 0).
		Width(fixedWidth - 10)

	subtitle := subtitleStyle.Render("21 | Developer")
	return lipgloss.JoinVertical(lipgloss.Center, bannerText, subtitle)
}
