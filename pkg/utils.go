package pkg

import (
	"net/url"
	"os"
	"strings"
)

// NormalizeURL, standart bir formata getirmek için ham bir URL'i işler.
// Gerekirse "https://" şemasını ekler ve linkte nokta yoksa ".com" ekler.
// Değiştirilip değiştirilmediğini belirten boole değeri döndürür.
func NormalizeURL(raw string) (string, bool) {
	changed := false
	raw = strings.TrimSpace(raw)

	// http/https yoksa ekle
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "https://" + raw
		changed = true
	}

	u, err := url.Parse(raw)
	if err != nil {
		return raw, changed
	}

	// hostname içinde nokta yoksa .com ekle
	if !strings.Contains(u.Hostname(), ".") {
		u.Host = u.Host + ".com"
		changed = true
	}

	return u.String(), changed
}

// RecreateDir, belirtilen yolda bir klasörü oluşturur.
// Eğer klasör zaten varsa içindekileri siler.
func RecreateDir(path string) error {
	_ = os.RemoveAll(path)
	return os.MkdirAll(path, 0755)
}

// FindEdge, Windows makinelerde Microsoft Edge'in yolunu bulmaya çalışır.
func FindEdge() string {
	paths := []string{
		`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
		`C:\Program Files\Microsoft\Edge\Application\msedge.exe`,
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}
