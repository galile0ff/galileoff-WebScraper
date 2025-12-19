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
	"golang.org/x/term"
)

func main() {
	// ---- CLI / ASCII ----
	opts := cli.Parse()
	cli.PrintASCII(os.Stdout, opts)

	pkg.PrintBanner("g a l i l e o f f .   W E B   S C R A P E R")

	pkg.PrintStep(1, 3, "Sistem Başlatılıyor")
	time.Sleep(800 * time.Millisecond)
	pkg.PrintStep(2, 3, "Modüller Yükleniyor")
	time.Sleep(800 * time.Millisecond)
	pkg.PrintStep(3, 3, "Hazır")
	time.Sleep(500 * time.Millisecond)

	for {
		runScraper()

		fmt.Println("\n\033[1;36m[?] İşlem seçiniz:\033[0m")
		fmt.Println("    \033[1;37m[ SPACE ]\033[0m : Çıkış")
		fmt.Println("    \033[1;37m[   F   ]\033[0m : Yeni Tarama")

		// Raw Mod
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			pkg.PrintError("Terminal raw moda geçemedi, çıkılıyor: %v", err)
			break
		}

		var char byte
		buf := make([]byte, 1)
		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				break
			}
			char = buf[0]
			// Space (32), F (102), f (70), Ctrl+C (3)
			if char == ' ' || char == 'f' || char == 'F' || char == 3 {
				break
			}
		}
		_ = term.Restore(int(os.Stdin.Fd()), oldState)

		if char == ' ' || char == 3 {
			pkg.GracefulExit(0)
			break
		} else if char == 'f' || char == 'F' {
			fmt.Println("\n\033[1;32m➜ Yeni tarama başlatılıyor...\033[0m")
			time.Sleep(500 * time.Millisecond)
			continue
		}
	}
}

func runScraper() {
	start := time.Now()
	spin := bspinner.New(bspinner.CharSets[9], 100*time.Millisecond)

	// ---- ARGÜMAN KONTROL ----
	if len(cli.Args()) > 0 {
		pkg.PrintError("Doğrudan URL yazmanız kabul edilmez.")
		return
	}

	// ---- URL AL ----
	fmt.Println()
	fmt.Print("\033[1;36m➜ Hedef URL (örn: galileoff.com / galileoff): \033[0m")

	time.Sleep(200 * time.Millisecond)

	var rawURL string
	fmt.Scanln(&rawURL)

	if strings.TrimSpace(rawURL) == "" {
		pkg.PrintError("URL'i boş bırakma.")
		return
	}

	targetURL, normalized := pkg.NormalizeURL(rawURL)
	parsed, err := url.Parse(targetURL)
	if err != nil {
		pkg.PrintError("Geçersiz URL formatı: %s", targetURL)
		return
	}

	fmt.Println()
	pkg.PrintKeyValue("Hedef", targetURL)
	pkg.PrintKeyValue("Durum", "Analiz Ediliyor")
	time.Sleep(1000 * time.Millisecond)

	// ---- KLASÖR ----
	siteName := strings.ReplaceAll(parsed.Hostname(), ".", "_")
	baseDir := filepath.Join(".", siteName)

	if err := pkg.RecreateDir(baseDir); err != nil {
		pkg.PrintError("Klasör hatası: %v", err)
		return
	}

	pkg.PrintKeyValue("Çalışma Alanı", baseDir)
	time.Sleep(800 * time.Millisecond)

	// ---- LOG ----
	logFile, err := os.OpenFile(
		filepath.Join(baseDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		pkg.PrintError("Log dosyası açılamadı: %v", err)
		return
	}
	defer logFile.Close()

	infoLog := log.New(logFile, "[INFO] ", log.Ldate|log.Ltime)
	debugLog := log.New(logFile, "[DEBUG] ", log.Ldate|log.Ltime)
	errorLog := log.New(logFile, "[ERROR] ", log.Ldate|log.Ltime)

	logInit(infoLog, targetURL, normalized)

	// ---- SCRAPE ----
	fmt.Println()
	pkg.PrintInfo("Headless tarayıcı başlatılıyor...")
	time.Sleep(1500 * time.Millisecond)

	spin.Prefix = "\033[36m[...]\033[0m İnceleniyor: "
	spin.Start()
	navStart := time.Now()

	result, err := pkg.Scrape(targetURL, infoLog, debugLog, errorLog)

	spin.Stop()
	navDuration := time.Since(navStart)

	if err != nil {
		pkg.PrintError("Kazıma sırasında hata oluştu: %v", err)
		return
	}

	fmt.Println()
	pkg.PrintSuccess("Erişim sağlandı ve veri çekildi. (%.2fs)", navDuration.Seconds())
	time.Sleep(500 * time.Millisecond)

	// ---- SONUÇLAR ----
	pkg.PrintInfo("Veriler işleniyor...")
	time.Sleep(1500 * time.Millisecond)

	logScrapeInfo(infoLog, result)
	saveResults(baseDir, result, errorLog)

	time.Sleep(500 * time.Millisecond)
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
	summaryData := map[string]string{
		"Taranan URL":   r.FinalURL,
		"HTTP Durumu":   fmt.Sprintf("%d", r.HTTPStatus),
		"Sayfa Başlığı": r.Title,
		"Link Sayısı":   fmt.Sprintf("%d", len(r.Links)),
		"Toplam Süre":   fmt.Sprintf("%.2fs", total.Seconds()),
		"Çıktı Yolu":    baseDir,
	}

	pkg.PrintBox(" T A R A M A   R A P O R U ", summaryData)

	// app.log boyutu
	logPath := filepath.Join(baseDir, "app.log")
	logInfo, err := os.Stat(logPath)
	logSize := "0 B"
	if err == nil {
		logSize = pkg.FormatBytes(logInfo.Size())
	}

	filesMap := map[string]string{
		"output.html":    pkg.FormatBytes(int64(len(r.HTML))),
		"screenshot.png": pkg.FormatBytes(int64(len(r.Screenshot))),
		"links.txt":      pkg.FormatBytes(int64(len(r.Links))),
		"app.log":        logSize,
	}

	pkg.PrintTreeList("Oluşturulan Dosyalar:", filesMap)
}
