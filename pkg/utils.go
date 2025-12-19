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

	// Hostname içinde nokta yoksa veya son kısım yaygın bir TLD değilse .com ekle.
	// TLD'yi nasıl tamamlıyo:
	// 1. Nokta hiç yoksa -> kesin ekle (örn: "galileoff" -> "galileoff.com")
	// 2. Nokta var ama son kısım (TLD) çok uzun veya bilinenlerden değilse -> ekle (örn: "sub.domain" -> "sub.domain.com")

	hostname := u.Hostname()
	parts := strings.Split(hostname, ".")

	shouldAppend := false
	if len(parts) == 1 {
		shouldAppend = true
	} else if len(parts) > 1 {
		tld := parts[len(parts)-1]
		// Yaygın TLD kontrolü
		// Buraya en sık kullanılan TLD'leri ekledim.
		commonTLDs := map[string]bool{
			"com": true, "net": true, "org": true, "edu": true, "gov": true, "mil": true,
			"io": true, "co": true, "ai": true, "app": true, "dev": true, "tr": true,
			"uk": true, "de": true, "fr": true, "ru": true, "jp": true, "it": true,
			"xyz": true, "info": true, "biz": true, "tv": true, "me": true,
		}

		if !commonTLDs[strings.ToLower(tld)] {
			shouldAppend = true
		}
	}

	if shouldAppend {
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
