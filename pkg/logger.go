package pkg

import (
	"fmt"
	"os"
	"strings"
)

// PrintInfo, mavi renkte bilgi mesajı
func PrintInfo(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("\033[34m[ℹ]\033[0m %s\n", msg)
}

// PrintSuccess, yeşil renkte başarı mesajı
func PrintSuccess(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("\033[32m[✔]\033[0m %s\n", msg)
}

// PrintError, kırmızı renkte hata mesajı
func PrintError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("\033[31m[✖]\033[0m %s\n", msg)
}

// FatalError, programı durdurur ve hata mesajı basar
func FatalError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Print("\r\033[K")
	fmt.Printf("\033[31;1m[FATAL]\033[0m %s\n", msg)
	os.Exit(1)
}

// PrintBanner, kutu içinde metin basar
func PrintBanner(text string) {
	line := strings.Repeat("─", len(text)+2)
	fmt.Printf("\n\033[35m┌%s┐\n", line)
	fmt.Printf("│ %s │\n", text)
	fmt.Printf("└%s┘\033[0m\n", line)
}

// FormatBytes, byte boyutunu okunabilir hale getirir
func FormatBytes(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	kb := float64(size) / 1024.0
	if kb < 1024 {
		return fmt.Sprintf("%.1f KB", kb)
	}
	mb := kb / 1024.0
	return fmt.Sprintf("%.1f MB", mb)
}
