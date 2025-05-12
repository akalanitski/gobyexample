package markdown

import (
	"regexp"
	"strings"
)

// cleanup replace sequnces of spaces to one space and sequence of EOL more the
// 2 in a row by 2 line breaks
func cleanup(input string) string {
	multiSpace := regexp.MustCompile(`\s+`)
	multiBreak := regexp.MustCompile(`\n{3,}`)
	ouput := multiBreak.ReplaceAllString(input, "\n\n")
	output := multiSpace.ReplaceAllString(ouput, " ")
	return output
}

func escapeTex(input string) string {
	e := regexp.MustCompile(`[\\{}#%&^~]`)
	return e.ReplaceAllString(input, "\\$0")
}

func excapeMarkdownInCodeHtml(input string) string {
	e := regexp.MustCompile("`(.*)(\\*)(.*)`")
	for e.MatchString(input) {
		input = e.ReplaceAllString(input, "`$1&ast;$3`")
	}
	return input
}
func excapeMarkdownInCodeTex(input string) string {
	e := regexp.MustCompile("`(.*)(\\*)(.*)`")
	for e.MatchString(input) {
		input = e.ReplaceAllString(input, "`$1\\char\"42$3`")
	}
	return input
}

func MarkdownToTex(input string) string {
	input = cleanup(input)
	input = escapeTex(input)
	input = excapeMarkdownInCodeTex(input)
	boldItalicPattern := regexp.MustCompile(`[*_]{3}(.+?)[*_]{3}`)
	boldPattern := regexp.MustCompile(`[*_]{2}(.+?)[*_]{2}`)
	italicPattern := regexp.MustCompile(`\*(.+?)\*`)
	italicPattern2 := regexp.MustCompile(`\b_(.+?)_\b`)
	codePattern := regexp.MustCompile("`(.+?)`")
	urlPattern := regexp.MustCompile(`\[(.+?)\]\((https?://.+?)\)`)
	refPattern := regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)
	underscorePattern := regexp.MustCompile(`_`)

	output := input
	output = boldItalicPattern.ReplaceAllString(output, "\\textbf{\\textit{$1}}")
	output = boldPattern.ReplaceAllString(output, "\\textbf{$1}")
	output = italicPattern.ReplaceAllString(output, "\\textit{$1}")
	output = italicPattern2.ReplaceAllString(output, "\\textit{$1}")
	output = codePattern.ReplaceAllString(output, "\\texttt{$1}")
	output = urlPattern.ReplaceAllString(output, "\\href{$2}{$1}")
	output = refPattern.ReplaceAllString(output, "\\hyperref[sec:$2]{\\textcolor{blue}{$1}}") //\hyperref[sec:strings-and-runes]{Strings and Runes}
	output = underscorePattern.ReplaceAllString(output, "\\$0")
	return output
}

func MarkdownToHTML(input string) string {
	input = cleanup(input)
	input = excapeMarkdownInCodeHtml(input)
	blocks := strings.Split(input, "\n\n")
	boldItalicPattern := regexp.MustCompile(`[*_]{3}(.+?)[*_]{3}`)
	boldPattern := regexp.MustCompile(`[*_]{2}(.+?)[*_]{2}`)
	italicPattern := regexp.MustCompile(`\*(.+?)\*`)
	italicPattern2 := regexp.MustCompile(`\b_(.+?)_\b`)
	codePattern := regexp.MustCompile("`(.+?)`")
	urlPattern := regexp.MustCompile(`\[(.+?)\]\((https?://.+?)\)`)
	refPattern := regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)

	var output []string
	for _, p := range blocks {
		p = boldItalicPattern.ReplaceAllString(p, "<em><strong>$1</strong></em>")
		p = boldPattern.ReplaceAllString(p, "<strong>$1</strong>")
		p = italicPattern.ReplaceAllString(p, "<em>$1</em>")
		p = italicPattern2.ReplaceAllString(p, "<em>$1</em>")
		p = codePattern.ReplaceAllString(p, "<code>$1</code>")
		p = urlPattern.ReplaceAllString(p, "<a href='$2' class='external' target='_blank'>$1</a>")
		p = refPattern.ReplaceAllString(p, "<a href='$2'>$1</a>")
		output = append(output, "<p>"+p+"</p>")
	}

	return strings.Join(output, "\n")
}
