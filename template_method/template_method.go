package template_method

type printer interface {
	open() string
	print() string
	close() string
}

type AbstractDisplay struct {
}

func (self *AbstractDisplay) Display(printer printer) string {
	result := printer.open()
	for i := 0; i < 5; i++ {
		result += printer.print()
	}
	result += printer.close()
	return result
}

type CharDisplay struct {
	*AbstractDisplay
	Char rune
}

func (self *CharDisplay) open() string {
	return "<<"
}
func (self *CharDisplay) print() string {
	return string(self.Char)
}
func (self *CharDisplay) close() string {
	return ">>"
}

type StringDisplay struct {
	*AbstractDisplay
	Str string
}

func (self *StringDisplay) open() string {
	return self.printLine()
}
func (self *StringDisplay) print() string {
	return "| " + self.Str + " |\n"
}
func (self *StringDisplay) close() string {
	return self.printLine()
}

func (self *StringDisplay) printLine() string {
	line := "+-"
	for _, _ = range self.Str {
		line += "-"
	}
	line += "-+\n"
	return line
}