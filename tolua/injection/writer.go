package injection

import "strings"

type (
	codeWriter struct {
		builder strings.Builder
	}
)

const (
	quote  = "\""
	line   = "\n"
	space  = " "
	lbrace = "{"
	rbrace = "}"
)

func newCodeWriter() *codeWriter {
	return &codeWriter{
		builder: strings.Builder{},
	}
}

func (w *codeWriter) writeLine(lines ...string) {
	for i := 0; i < len(lines); i++ {
		w.builder.WriteString(lines[i])
		w.builder.WriteString(line)
	}
}
func (w *codeWriter) writeQuote() {
	w.builder.WriteString(quote)
}
func (w *codeWriter) writeSpace() {
	w.builder.WriteString(space)
}

func (w *codeWriter) writeString(s string) {
	w.builder.WriteString(s)
}
func (w *codeWriter) writeLineEnd() {
	w.builder.WriteString(line)
}

func (w *codeWriter) len() int {
	return w.builder.Len()
}
func (w *codeWriter) write(c *codeWriter) {
	l := c.len()
	if l > 0 {
		w.builder.WriteString(c.code())
	}
}
func (w *codeWriter) code() string {
	return w.builder.String()
}
