package pages

import (
	"terminal-portfolio/menu/styles"

	"github.com/charmbracelet/lipgloss"
)

type ContactPage struct{}

type ContactInfo struct {
	Label string
	Value string
}

func NewContactPage() *ContactPage {
	return &ContactPage{}
}

func (c *ContactPage) getContactInfo() []ContactInfo {
	return []ContactInfo{
		{Label: "Email", Value: "workforrudra24@gmail.com"},
		{Label: "GitHub", Value: "Rudra-Sankha-Sinhamahapatra"},
		{Label: "LinkedIn", Value: "https://www.linkedin.com/in/rudra-sankha-sinhamahapatra-6311aa1bb/"},
		{Label: "Twitter", Value: "@RudraSankha"},
	}
}

func (c *ContactPage) createContactCard(info ContactInfo, styles *styles.Styles) string {
	return styles.ContactItemStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.ContactLabelStyle.Render(info.Label),
			styles.ContactValueStyle.Render(info.Value),
		),
	)
}

func (c *ContactPage) GetContent(fixedWidth int, styles *styles.Styles) string {
	contactTitle := styles.SectionTitleStyle.Render("# Get In Touch")

	contactInfo := c.getContactInfo()
	var contactCards []string

	for _, info := range contactInfo {
		card := c.createContactCard(info, styles)
		contactCards = append(contactCards, card)
	}

	var contactRows []string
	if fixedWidth >= 80 {
		for i := 0; i < len(contactCards); i += 2 {
			if i+1 < len(contactCards) {
				row := lipgloss.JoinHorizontal(lipgloss.Left, contactCards[i], contactCards[i+1])
				contactRows = append(contactRows, row)
			} else {
				contactRows = append(contactRows, contactCards[i])
			}
		}
	} else {
		contactRows = contactCards
	}

	availabilityTitle := styles.ContactHeadingStyle.Render("## Current Status")
	availabilityText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#90EE90")).
		Bold(true).
		Render("Open to opportunities")

	responseTitle := styles.ContactHeadingStyle.Render("## Response Time")
	responseText := styles.ContactValueStyle.Render("Usually respond within 24 hours")

	callToAction := styles.CallToActionStyle.Render("»  Feel free to reach out if you are hiring or just to say hi!")

	contactElements := []string{contactTitle, ""}
	contactElements = append(contactElements, contactRows...)
	contactElements = append(contactElements,
		"",
		availabilityTitle,
		availabilityText,
		"",
		responseTitle,
		responseText,
		"",
		callToAction,
	)

	return styles.ContentStyle.Render(lipgloss.JoinVertical(lipgloss.Left, contactElements...))
}
