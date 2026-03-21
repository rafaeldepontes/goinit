package log

import "fmt"

const (
	InfoColor    = "\u001b[38;5;104m"
	ErrorColor   = "\u001b[38;5;203m"
	ArrowColor   = "\u001b[38;5;39m"
	WarningColor = "\u001b[38;5;11m"
	BannerColor  = "\u001b[38;5;218m"
	NoneColor    = "\033[0m"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg ...any) {
	fmt.Print(InfoColor)
	fmt.Print(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Infoln(msg ...any) {
	fmt.Print(InfoColor)
	fmt.Println(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Infof(msg string, args ...any) {
	fmt.Print(InfoColor)
	fmt.Printf(msg, args...)
	fmt.Print(NoneColor)
}

func (l *Logger) InfoPrefix(pref string, msg ...any) {
	fmt.Print(ArrowColor)
	fmt.Print(pref)
	fmt.Print(InfoColor)
	fmt.Print(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) InfoPrefixln(pref string, msg ...any) {
	fmt.Print(ArrowColor)
	fmt.Print(pref)
	fmt.Print(InfoColor)
	fmt.Println(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) InfoPrefixf(pref string, msg string, args ...any) {
	fmt.Print(ArrowColor)
	fmt.Print(pref)
	fmt.Print(InfoColor)
	fmt.Printf(msg, args...)
	fmt.Print(NoneColor)
}

func (l *Logger) Warning(msg ...any) {
	fmt.Print(WarningColor)
	fmt.Print(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Warningln(msg ...any) {
	fmt.Print(WarningColor)
	fmt.Println(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Warningf(msg string, args ...any) {
	fmt.Print(WarningColor)
	fmt.Printf(msg, args...)
	fmt.Print(NoneColor)
}

func (l *Logger) Error(msg ...any) {
	fmt.Print(ErrorColor)
	fmt.Print(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Errorln(msg ...any) {
	fmt.Print(ErrorColor)
	fmt.Println(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) Errorf(msg string, args ...any) {
	fmt.Print(ErrorColor)
	fmt.Printf(msg, args...)
	fmt.Print(NoneColor)
}

func (l *Logger) PrintBanner(msg ...any) {
	fmt.Print(BannerColor)
	fmt.Println(msg...)
	fmt.Print(NoneColor)
}

func (l *Logger) SInfof(msg string, args ...any) string {
	fmt.Print(NoneColor)
	return fmt.Sprintf(InfoColor+msg, args...)
}
