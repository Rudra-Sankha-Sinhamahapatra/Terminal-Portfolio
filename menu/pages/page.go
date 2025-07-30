package pages

import "terminal-portfolio/menu/styles"

type Page interface {
	GetContent(fixedWidth int, styles *styles.Styles) string
}
