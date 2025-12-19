package pkg

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Typewriter, metni daktilo efektiyle yazıyo
func Typewriter(text string, delayMs int) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}

// PrintInfo, mavi renkte bilgi mesajı
func PrintInfo(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Print("\033[36m➜ \033[0m")
	Typewriter(msg+"\n", 30)
}

// PrintSuccess, yeşil renkte başarı mesajı
func PrintSuccess(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Print("\033[1;32m✔ \033[0m")
	Typewriter(msg+"\n", 30)
}

// PrintError, kırmızı renkte hata mesajı
func PrintError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Print("\033[1;31m✖ \033[0m")
	Typewriter(msg+"\n", 30)
}

// PrintStep
func PrintStep(current, total int, message string) {
	fmt.Printf("\033[33m[%d/%d]\033[0m %s...\n", current, total, message)
	time.Sleep(500 * time.Millisecond)
}

// PrintKeyValue
func PrintKeyValue(key, value string) {
	const padding = 35
	keyLen := len([]rune(key))
	if keyLen > padding {
		keyLen = padding
	}
	dots := strings.Repeat(".", padding-keyLen)
	fmt.Printf("\033[36m%s\033[0m%s: \033[1;37m%s\033[0m\n", key, dots, value)
	time.Sleep(150 * time.Millisecond)
}

// FatalError, programı durdurur ve hata mesajı basar
// GetFriendlyErrorMessage, teknik hata mesajını anlaşılabilir hale getirir
func GetFriendlyErrorMessage(msg string) string {
	if strings.Contains(msg, "ERR_NAME_NOT_RESOLVED") || strings.Contains(msg, "no such host") {
		return "Siteye ulaşılamadı. URL yanlış olabilir veya internet bağlantısı yok."
	} else if strings.Contains(msg, "ERR_CONNECTION_REFUSED") {
		return "Bağlantı reddedildi. Site güvenliği veya güvenlik duvarı engelliyor."
	} else if strings.Contains(msg, "403") {
		return "Erişim engellendi (403). WAF veya güvenlik önlemlerine takıldık."
	} else if strings.Contains(msg, "404") {
		return "Sayfa bulunamadı (404). Link kırık veya sayfa kaldırılmış."
	} else if strings.Contains(msg, "500") || strings.Contains(msg, "502") || strings.Contains(msg, "503") {
		return "Karşı sunucu hatası (5xx). Site şu an hizmet veremiyor."
	} else if strings.Contains(msg, "deadline exceeded") || strings.Contains(msg, "timeout") || strings.Contains(msg, "ERR_CONNECTION_TIMED_OUT") {
		return "İşlem zaman aşımına uğradı. Site çok yavaş yanıt veriyor."
	} else if strings.Contains(msg, "certificate") || strings.Contains(msg, "x509") || strings.Contains(msg, "ERR_CERT_") {
		return "Güvenlik sertifikası (SSL) hatası. Bağlantı güvenli değil veya tarih/saat yanlış."
	} else if strings.Contains(msg, "executable file not found") {
		return "Chrome/Chromium veya MS Edge tarayıcısı sistemde bulunamadı."
	} else if strings.Contains(msg, "ERR_INTERNET_DISCONNECTED") {
		return "İnternet bağlantısı yok."
	}
	return "Beklenmedik teknik bir hata oluştu."
}

// PrintScrapeError, hatayı fatal olmadan ekrana basar
func PrintScrapeError(err error) {
	msg := err.Error()
	friendlyMsg := GetFriendlyErrorMessage(msg)

	fmt.Print("\r\033[K")
	fmt.Println()

	fmt.Printf("\033[41;1;37m %-58s \033[0m\n", "HATA OLUŞTU")
	fmt.Printf("\033[90m│\033[0m \033[31m%s\033[0m\n", friendlyMsg)
	fmt.Printf("\033[90m│ Teknik Detay: %s\033[0m\n", msg)
	fmt.Println()
}

// FatalError, programı durdurur ve hata mesajı basar
func FatalError(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	friendlyMsg := GetFriendlyErrorMessage(msg)

	fmt.Print("\r\033[K")
	fmt.Println()

	fmt.Printf("\033[41;1;37m %-58s \033[0m\n", "KRİTİK HATA")
	fmt.Printf("\033[90m│\033[0m \033[31m%s\033[0m\n", friendlyMsg)

	fmt.Printf("\033[90m│ Teknik Detay: %s\033[0m\n", msg)

	GracefulExit(1)
}

// GracefulExit, programı kapatma animasyonu için
func GracefulExit(code int) {
	fmt.Println()

	msg := "[!] Program kapatılıyor..."
	byeMsg := "[!] Kendine cici bak."

	if code == 0 {
		msg = "➜ İşlemler tamamlandı, kapatılıyor..."
		byeMsg = "➜ Kendine cici bak."
	}

	fmt.Printf("\033[1;36m%s\033[0m", msg)
	time.Sleep(3000 * time.Millisecond)
	fmt.Print("\r\033[K")
	fmt.Printf("\033[1;36m%s\033[0m\n", byeMsg)
	time.Sleep(2000 * time.Millisecond)
	os.Exit(code)
}

// PrintBanner
func PrintBanner(title string) {
	contentWidth := len(title) + 4
	topBorder := "╔" + strings.Repeat("═", contentWidth) + "╗"
	botBorder := "╚" + strings.Repeat("═", contentWidth) + "╝"

	fmt.Println()
	fmt.Println("\033[1;35m" + topBorder + "\033[0m")
	fmt.Printf("\033[1;35m║  \033[1;37m%s\033[1;35m  ║\033[0m\n", title)
	fmt.Println("\033[1;35m" + botBorder + "\033[0m")
	fmt.Println()
}

// PrintBox, detaylı bilgi kutusu
func PrintBox(title string, items map[string]string) {
	width := 60
	paddingLen := (width - len(title)) / 2
	padding := strings.Repeat(" ", paddingLen)
	extraPad := ""
	if (width-len(title))%2 != 0 {
		extraPad = " "
	}

	fmt.Println()
	fmt.Printf("\033[46;1;37m%s%s%s%s\033[0m\n", padding, title, padding, extraPad)

	fmt.Println("\033[90m┌" + strings.Repeat("─", width-2) + "┐\033[0m")

	for k, v := range items {
		lineContent := fmt.Sprintf(" %s: %s", k, v)
		if len(lineContent) > width-4 {
			lineContent = lineContent[:width-7] + "..."
		}

		padLen := width - 2 - len([]rune(lineContent)) // UTF-8 fix
		if padLen < 0 {
			padLen = 0
		}

		fmt.Printf("\033[90m│\033[0m\033[36m%s\033[0m%s\033[90m│\033[0m\n",
			lineContent, strings.Repeat(" ", padLen))
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\033[90m└" + strings.Repeat("─", width-2) + "┘\033[0m")
}

// PrintTreeList, dosya listesini ağaç yapısında terminale basıyo
func PrintTreeList(title string, items map[string]string) {
	fmt.Printf("\n\033[1;36m➜ %s\033[0m\n", title)

	i := 0
	total := len(items)
	for k, v := range items {
		i++
		prefix := "├──"
		if i == total {
			prefix = "└──"
		}

		fmt.Printf(" \033[90m%s\033[0m \033[36m%s\033[0m \033[90m(%s)\033[0m\n", prefix, k, v)
		time.Sleep(150 * time.Millisecond)
	}
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
