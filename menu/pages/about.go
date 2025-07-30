package pages

import (
	"terminal-portfolio/menu/styles"
	"terminal-portfolio/menu/utils"

	"github.com/charmbracelet/lipgloss"
)

type AboutPage struct{}

func NewAboutPage() *AboutPage {
	return &AboutPage{}
}

func (a *AboutPage) GetContent(fixedWidth int, styles *styles.Styles) string {
	aboutTitle := styles.SectionTitleStyle.Render("# About Me")
	aboutText := "21 | Software Engineer. Explore the\nopen-source projects and libraries I maintain on GitHub"

	techTitle := styles.SectionTitleStyle.Render("# Tools & Technologies")

	langTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## Languages")
	languages := []string{"Python", "C", "JavaScript", "TypeScript", "Go"}
	langRow := utils.ArrangeSkillBoxes(languages, styles.SkillBoxStyle, fixedWidth-10)

	fwTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## Frameworks & Libraries")
	frameworks := []string{"ReactJS", "NextJS", "ExpressJS", "Nest Js", "Gin", "GraphQL"}
	fwRow := utils.ArrangeSkillBoxes(frameworks, styles.SkillBoxStyle, fixedWidth-10)

	stateTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## State Management & Styling")
	stateTools := []string{"Zustand", "Recoil", "Tailwind CSS"}
	stateRow := utils.ArrangeSkillBoxes(stateTools, styles.SkillBoxStyle, fixedWidth-10)

	toolsTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## Development Tools")
	devTools := []string{"Git", "Jest", "Vitest"}
	devRow := utils.ArrangeSkillBoxes(devTools, styles.SkillBoxStyle, fixedWidth-10)

	dbTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## Databases & Cloud")
	databases := []string{"PostgreSQL", "Mongo DB", "Redis", "Supabase", "Prisma", "Drizzle"}
	dbRow := utils.ArrangeSkillBoxes(databases, styles.SkillBoxStyle, fixedWidth-10)

	monTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## Monitoring & Messaging")
	monitoring := []string{"Prometheus", "Grafana", "Kafka"}
	monRow := utils.ArrangeSkillBoxes(monitoring, styles.SkillBoxStyle, fixedWidth-10)

	devopsTitle := lipgloss.NewStyle().Foreground(lipgloss.Color("#CCCCCC")).Bold(true).Render("## DevOps")
	devopsTools := []string{"Docker", "Kubernetes", "CI/CD", "Nginx", "AWS EC2", "AWS S3", "AWS ECS"}
	devopsRow := utils.ArrangeSkillBoxes(devopsTools, styles.SkillBoxStyle, fixedWidth-10)

	return styles.ContentStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
		aboutTitle,
		aboutText,
		"",
		techTitle,
		"",
		langTitle,
		langRow,
		"",
		fwTitle,
		fwRow,
		"",
		stateTitle,
		stateRow,
		"",
		toolsTitle,
		devRow,
		"",
		dbTitle,
		dbRow,
		"",
		monTitle,
		monRow,
		"",
		devopsTitle,
		devopsRow,
	))
}
