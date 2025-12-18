package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"galileoff-WebScraper/pkg"
	"galileoff-WebScraper/pkg/cli"

	bspinner "github.com/briandowns/spinner"
)

func main() {
	start := time.Now()
	spin := bspinner.New(bspinner.CharSets[9], 100*time.Millisecond)

	// ---- CLI / ASCII ----
	opts := cli.Parse()
	cli.PrintASCII(os.Stdout, opts)

	pkg.PrintInfo("galileoff. Web Scraper")
	pkg.PrintInfo("Hazır.")

	// ---- ARGÜMAN KONTROL ----
	if len(cli.Args()) > 0 {
		pkg.PrintError("URL komut satırından verilmez.")
		pkg.PrintInfo("Kullanım: go run main.go")
		pkg.PrintInfo("URL sizden istendiğinde yazın.")
		return
	}

	// ---- URL AL ----
	fmt.Print("\033[33m[?] İncelenecek URL'yi giriniz (örn: galileoff.com): \033[0m")
	var rawURL string
	fmt.Scanln(&rawURL)

	if strings.TrimSpace(rawURL) == "" {
		pkg.PrintError("Boş URL girildi.")
		return
	}

	targetURL, normalized := pkg.NormalizeURL(rawURL)
	parsed, err := url.Parse(targetURL)
	if err != nil {
		pkg.FatalError("Geçersiz URL formatı: %s", targetURL)
	}

	pkg.PrintInfo("Hedef URL: %s", targetURL)

	// ---- KLASÖR ----
	siteName := strings.ReplaceAll(parsed.Hostname(), ".", "_")
	baseDir := filepath.Join(".", siteName)

	if err := pkg.RecreateDir(baseDir); err != nil {
		pkg.FatalError("Sonuç klasörü oluşturulamadı: %v", err)
	}

	pkg.PrintInfo("Sonuçlar '%s' klasörüne kaydedilecek.", baseDir)

	// ---- LOG ----
	logFile, err := os.OpenFile(
		filepath.Join(baseDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		pkg.FatalError("Log dosyası açılamadı: %v", err)
	}
	defer logFile.Close()

	infoLog := log.New(logFile, "[INFO] ", log.Ldate|log.Ltime)
	debugLog := log.New(logFile, "[DEBUG] ", log.Ldate|log.Ltime)
	errorLog := log.New(logFile, "[ERROR] ", log.Ldate|log.Ltime)

	logInit(infoLog, targetURL, normalized)

	// ---- SCRAPE ----
	spin.Start()
	navStart := time.Now()

	result, err := pkg.Scrape(targetURL, infoLog, debugLog, errorLog)

	spin.Stop()
	navDuration := time.Since(navStart)

	if err != nil {
		pkg.FatalError("Kazıma sırasında hata oluştu: %v", err)
	}

	pkg.PrintSuccess("Sayfa başarıyla kazındı. (%.2fs)", navDuration.Seconds())

	// ---- SONUÇLAR ----
	logScrapeInfo(infoLog, result)
	saveResults(baseDir, result, errorLog)
	printSummary(baseDir, result, time.Since(start))

	logFinal(infoLog, time.Since(start))
}

// ---- LOG FONKSİYONLARI ----
func logInit(infoLog *log.Logger, targetURL string, normalized bool) {
	infoLog.Println("===================================")
	infoLog.Println("Program başlatıldı")
	infoLog.Println("Hedef URL:", targetURL)
	if normalized {
		infoLog.Println("URL normalize edildi")
	}
	infoLog.Println("===================================")
}

func logScrapeInfo(infoLog *log.Logger, r *pkg.Result) {
	infoLog.Println("HTTP Durumu:", r.HTTPStatus)
	infoLog.Println("Taranan URL:", r.FinalURL)
	infoLog.Println("Sayfa Başlığı:", r.Title)
	infoLog.Println("HTML Boyutu:", len(r.HTML), "byte")
	infoLog.Println("Screenshot Boyutu:", len(r.Screenshot), "byte")
	infoLog.Println("Toplam Link Sayısı:", len(r.Links))
}

func logFinal(infoLog *log.Logger, totalDuration time.Duration) {
	infoLog.Println("Toplam Süre:", totalDuration)
	infoLog.Println("===================================")
	infoLog.Println("Program sonlandı")
	infoLog.Println("===================================")
}

// ---- DOSYA YAZMA ----
func saveResults(baseDir string, r *pkg.Result, errorLog *log.Logger) {
	pkg.PrintInfo("Sonuçlar diske yazılıyor...")

	files := map[string][]byte{
		"output.html":    []byte(r.HTML),
		"screenshot.png": r.Screenshot,
		"links.txt":      []byte(strings.Join(r.Links, "\n")),
	}

	for name, content := range files {
		path := filepath.Join(baseDir, name)
		if err := os.WriteFile(path, content, 0644); err != nil {
			errorLog.Printf("%s yazılamadı: %v", name, err)
		}
	}
}

// ---- ÖZET ----
func printSummary(baseDir string, r *pkg.Result, total time.Duration) {
	pkg.PrintBanner("Tarama Tamamlandı!")

	pkg.PrintInfo("Taranan URL: %s", r.FinalURL)
	pkg.PrintInfo("HTTP Durumu: %d", r.HTTPStatus)
	pkg.PrintInfo("Sayfa Başlığı: %s", r.Title)
	pkg.PrintInfo("Bulunan Link Sayısı: %d", len(r.Links))
	pkg.PrintInfo("Toplam Süre: %.2fs", total.Seconds())

	fmt.Println()
	pkg.PrintSuccess("Sonuçlar '%s' klasörüne kaydedildi:", baseDir)

	fmt.Printf(`
    ├── output.html    (%s)
    ├── screenshot.png (%s)
    ├── links.txt      (%s)
    └── app.log
`,
		pkg.FormatBytes(int64(len(r.HTML))),
		pkg.FormatBytes(int64(len(r.Screenshot))),
		pkg.FormatBytes(int64(len(r.Links))),
	)
}
