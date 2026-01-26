package log

import "fmt"

const (
	ColorInfo    = "\u001b[38;5;104m"
	ColorError   = "\u001b[38;5;203m"
	ColorArrow   = "\u001b[38;5;39m"
	ColorWarning = "\u001b[38;5;11m"
	ColorBanner  = "\u001b[38;5;218m"
	ColorNone    = "\033[0m"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg ...any) {
	fmt.Print(ColorInfo)
	fmt.Print(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Infoln(msg ...any) {
	fmt.Print(ColorInfo)
	fmt.Println(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Infof(msg string, args ...any) {
	fmt.Print(ColorInfo)
	fmt.Printf(msg, args...)
	fmt.Print(ColorNone)
}

func (l *Logger) InfoPrefix(pref string, msg ...any) {
	fmt.Print(ColorArrow)
	fmt.Print(pref)
	fmt.Print(ColorInfo)
	fmt.Print(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) InfoPrefixln(pref string, msg ...any) {
	fmt.Print(ColorArrow)
	fmt.Print(pref)
	fmt.Print(ColorInfo)
	fmt.Println(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) InfoPrefixf(pref string, msg string, args ...any) {
	fmt.Print(ColorArrow)
	fmt.Print(pref)
	fmt.Print(ColorInfo)
	fmt.Printf(msg, args...)
	fmt.Print(ColorNone)
}

func (l *Logger) Warning(msg ...any) {
	fmt.Print(ColorWarning)
	fmt.Print(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Warningln(msg ...any) {
	fmt.Print(ColorWarning)
	fmt.Println(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Warningf(msg string, args ...any) {
	fmt.Print(ColorWarning)
	fmt.Printf(msg, args...)
	fmt.Print(ColorNone)
}

func (l *Logger) Error(msg ...any) {
	fmt.Print(ColorError)
	fmt.Print(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Errorln(msg ...any) {
	fmt.Print(ColorError)
	fmt.Println(msg...)
	fmt.Print(ColorNone)
}

func (l *Logger) Errorf(msg string, args ...any) {
	fmt.Print(ColorError)
	fmt.Printf(msg, args...)
	fmt.Print(ColorNone)
}

func (l *Logger) PrintBanner(msg ...any) {
	fmt.Print(ColorBanner)
	fmt.Println(msg...)
	fmt.Print(ColorNone)
}
