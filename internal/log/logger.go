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
	fmt.Print(ColorInfo, msg, ColorNone)
}

func (l *Logger) Infoln(msg ...any) {
	fmt.Println(ColorInfo, msg, ColorNone)
}

func (l *Logger) Infof(msg string, args ...any) {
	fmt.Printf(ColorInfo+msg+ColorNone, args...)
}

func (l *Logger) InfoPrefix(pref string, msg ...any) {
	fmt.Print(ColorArrow, pref, ColorInfo, msg, ColorNone)
}

func (l *Logger) InfoPrefixln(pref string, msg ...any) {
	fmt.Println(ColorArrow, pref, ColorInfo, msg, ColorNone)
}

func (l *Logger) InfoPrefixf(pref string, msg string, args ...any) {
	fmt.Printf(ColorArrow+pref+ColorInfo+msg+ColorNone, args...)
}

func (l *Logger) Warning(msg ...any) {
	fmt.Print(ColorWarning, msg, ColorNone)
}

func (l *Logger) Warningln(msg ...any) {
	fmt.Println(ColorWarning, msg, ColorNone)
}

func (l *Logger) Warningf(msg string, args ...any) {
	fmt.Printf(ColorWarning+msg+ColorNone, args...)
}

func (l *Logger) Error(msg ...any) {
	fmt.Print(ColorError, msg, ColorNone)
}

func (l *Logger) Errorln(msg ...any) {
	fmt.Println(ColorError, msg, ColorNone)
}

func (l *Logger) Errorf(msg string, args ...any) {
	fmt.Printf(ColorError+msg+ColorNone, args...)
}

func (l *Logger) PrintBanner(msg ...any) {
	fmt.Println(ColorBanner, msg, ColorNone)
}
