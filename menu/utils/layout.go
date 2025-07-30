package utils

import "github.com/charmbracelet/lipgloss"

func ArrangeSkillBoxes(skills []string, skillBoxStyle lipgloss.Style, containerWidth int) string {
	if len(skills) == 0 {
		return ""
	}

	boxesPerRow := containerWidth / 15
	if boxesPerRow < 1 {
		boxesPerRow = 1
	}
	if boxesPerRow > len(skills) {
		boxesPerRow = len(skills)
	}

	var rows []string
	for i := 0; i < len(skills); i += boxesPerRow {
		end := i + boxesPerRow
		if end > len(skills) {
			end = len(skills)
		}

		var rowBoxes []string
		for j := i; j < end; j++ {
			rowBoxes = append(rowBoxes, skillBoxStyle.Render(skills[j]))
		}

		row := lipgloss.JoinHorizontal(lipgloss.Left, rowBoxes...)
		rows = append(rows, row)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
