package pages

import (
	"terminal-portfolio/menu/styles"
	"terminal-portfolio/menu/utils"

	"github.com/charmbracelet/lipgloss"
)

type ProjectsPage struct{}

type Project struct {
	Title        string
	Description  string
	Technologies []string
}

func NewProjectsPage() *ProjectsPage {
	return &ProjectsPage{}
}

func (p *ProjectsPage) getProjects() []Project {
	return []Project{
		{
			Title:        "## Terminal Portfolio",
			Description:  "Interactive SSH-based portfolio application built with modern terminal UI",
			Technologies: []string{"Go", "Wish", "Bubbletea", "Lipgloss", "SSH"},
		},
		{
			Title:        "## AI PPT Generator",
			Description:  "AI-powered presentation generator with modern web interface and backend processing",
			Technologies: []string{"Next.js", "TypeScript", "Go", "Redis", "PostgreSQL", "Prisma", "Gemini AI", "Pub Sub", "Node.js"},
		},
		{
			Title:        "## Adda AI",
			Description:  "Full-stack AI chat application enabling conversations with AI characters",
			Technologies: []string{"React", "Next.js", "Nest.js", "Docker", "PostgreSQL", "Prisma", "Turborepo"},
		},
		{
			Title:        "## Real-time Chat",
			Description:  "WebSocket-based chat application with pub/sub messaging architecture",
			Technologies: []string{"WebSocket", "Redis", "Pub/Sub", "Node.js", "React"},
		},
	}
}

func (p *ProjectsPage) GetContent(fixedWidth int, styles *styles.Styles) string {
	projectTitle := styles.SectionTitleStyle.Render("# My Projects")

	projects := p.getProjects()
	var projectSections []string

	projectSections = append(projectSections, projectTitle, "")

	for _, project := range projects {
		projectTitleSection := styles.ProjectHeadingStyle.Render(project.Title)
		projectDesc := styles.ProjectDescStyle.Render(project.Description)
		techLabel := styles.TechLabelStyle.Render("Technologies:")
		techRow := utils.ArrangeSkillBoxes(project.Technologies, styles.SkillBoxStyle, fixedWidth-10)

		projectSections = append(projectSections,
			projectTitleSection,
			projectDesc,
			techLabel,
			techRow,
			"",
		)
	}

	return styles.ContentStyle.Render(lipgloss.JoinVertical(lipgloss.Left, projectSections...))
}
