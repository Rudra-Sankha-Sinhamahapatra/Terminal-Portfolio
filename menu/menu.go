package menu

import (
	"strings"

	"terminal-portfolio/menu/pages"
	"terminal-portfolio/menu/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	cursor       int
	items        []string
	width        int
	height       int
	scrollOffset int
	pages        []pages.Page
}

func InitialMenuModel() menuModel {
	return menuModel{
		cursor:       0,
		items:        []string{"Home", "About", "Projects", "Contact"},
		width:        80,
		height:       24,
		scrollOffset: 0,
		pages: []pages.Page{
			pages.NewHomePage(),
			pages.NewAboutPage(),
			pages.NewProjectsPage(),
			pages.NewContactPage(),
		},
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			if m.cursor > 0 {
				m.cursor--
				m.scrollOffset = 0
			}
		case "right":
			if m.cursor < len(m.items)-1 {
				m.cursor++
				m.scrollOffset = 0
			}
		case "up":
			if m.scrollOffset > 0 {
				m.scrollOffset--
			}
		case "down":
			maxScroll := m.getMaxScroll()
			if m.scrollOffset < maxScroll {
				m.scrollOffset++
			}
		case "home":
			m.scrollOffset = 0
		case "end":
			m.scrollOffset = m.getMaxScroll()
		}
	}
	return m, nil
}

func (m menuModel) getMaxScroll() int {
	var fixedWidth, fixedHeight int

	if m.width < 90 {
		fixedWidth = 70
		fixedHeight = m.height - 6
	} else if m.width < 120 {
		fixedWidth = 85
		fixedHeight = 28
	} else {
		fixedWidth = 100
		fixedHeight = 30
	}

	if fixedWidth < 60 {
		fixedWidth = 60
	}
	if fixedHeight < 18 {
		fixedHeight = 18
	}

	content := m.getContent(fixedWidth)
	contentLines := strings.Split(content, "\n")
	availableHeight := fixedHeight - 8

	if len(contentLines) <= availableHeight {
		return 0
	}
	return len(contentLines) - availableHeight
}

func (m menuModel) View() string {
	var fixedWidth, fixedHeight int

	if m.width < 90 {
		fixedWidth = 70
		fixedHeight = m.height - 6
	} else if m.width < 120 {
		fixedWidth = 85
		fixedHeight = 28
	} else {
		fixedWidth = 100
		fixedHeight = 30
	}

	if fixedWidth < 60 {
		fixedWidth = 60
	}
	if fixedHeight < 18 {
		fixedHeight = 18
	}

	styles := styles.NewMenuStyles(m.width, m.height, fixedWidth, fixedHeight)
	title := styles.TitleStyle.Render("Rudra | Portfolio")

	var navItems []string
	for i, item := range m.items {
		if i == m.cursor {
			navItems = append(navItems, styles.SelectedNavStyle.Render(item))
		} else {
			navItems = append(navItems, styles.NavItemStyle.Render(item))
		}
	}

	navigation := lipgloss.JoinHorizontal(lipgloss.Left, navItems...)

	var spacing string
	if fixedWidth < 80 {
		spacing = strings.Repeat(" ", 2)
	} else if fixedWidth < 95 {
		spacing = strings.Repeat(" ", 5)
	} else {
		spacing = strings.Repeat(" ", 10)
	}

	header := lipgloss.JoinHorizontal(lipgloss.Left, title, spacing, navigation)

	content := m.getContent(fixedWidth)
	contentLines := strings.Split(content, "\n")

	availableHeight := fixedHeight - 8
	startLine := m.scrollOffset
	endLine := startLine + availableHeight

	if endLine > len(contentLines) {
		endLine = len(contentLines)
	}

	var visibleContent []string
	if startLine < len(contentLines) {
		visibleContent = contentLines[startLine:endLine]
	}

	scrolledContent := strings.Join(visibleContent, "\n")

	footerText := "Press 'q' to quit | ←→ menu | ↑↓ scroll"
	footer := styles.FooterStyle.Render(footerText)

	body := lipgloss.JoinVertical(lipgloss.Left,
		header,
		strings.Repeat("-", fixedWidth-4),
		scrolledContent,
		"",
		footer,
	)

	borderedContent := styles.BorderStyle.Render(body)

	return styles.CenterStyle.Render(borderedContent)
}

func (m menuModel) getContent(fixedWidth int) string {
	styles := styles.NewStyles(fixedWidth)

	if m.cursor >= 0 && m.cursor < len(m.pages) {
		return m.pages[m.cursor].GetContent(fixedWidth, styles)
	}

	return styles.ContentStyle.Render("Select a menu item to view content.")
}
