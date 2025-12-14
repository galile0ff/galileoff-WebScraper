package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	// 1. URL'i komut satırı argümanından alıyoruz
	if len(os.Args) < 2 {
		log.Fatal("Lütfen komut satırı argümanı olarak bir URL girin.")
	}
	url := os.Args[1]

	// 2. chromedp için bir bağlam oluşturuyoruz
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			// Chrome kullanmayan cihazlar için Edge'in yolunu belirtmekte fayda var.
			chromedp.ExecPath(`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`),
		)...,
	)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Yanıtı dinlemek için bir listener kuruyoruz
	var response *network.Response
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*network.EventResponseReceived); ok {
			if ev.Type == network.ResourceTypeDocument {
				response = ev.Response
			}
		}
	})

	// 3. HTML içeriğini ve ekran görüntüsünü alıyoruz burada
	var htmlContent string
	var screenshot []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &htmlContent),
		chromedp.FullScreenshot(&screenshot, 90),
	)
	if err != nil {
		log.Fatalf("URL'ye bağlanırken veya veri çekerken hata oluştu: %v", err)
	}

	// HTTP durum kodunu kontrol ediyoruz
	if response != nil && response.Status != 200 {
		log.Fatalf("Başarılı bir durum kodu alınamadı (HTTP %d %s)", response.Status, response.StatusText)
	}

	// 4. HTML içeriğini dosyaya kaydediyoruz
	err = os.WriteFile("output.html", []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("HTML içeriği dosyaya kaydedilirken hata oluştu: %v", err)
	}
	fmt.Println("HTML içeriği başarıyla output.html dosyasına kaydedildi.")

	// 5. Ekran görüntüsünü dosyaya kaydediyoruz
	err = os.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		log.Fatalf("Ekran görüntüsü dosyaya kaydedilirken hata oluştu: %v", err)
	}
	fmt.Println("Ekran görüntüsü başarıyla screenshot.png dosyasına kaydedildi.")

	// URL'leri listeliyoruz
	var links []string
	err = chromedp.Run(ctx,
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll('a')).map(a => a.href)
		`, &links),
	)
	if err != nil {
		log.Fatalf("URL'ler alınırken hata oluştu: %v", err)
	}

	fmt.Println("\nSayfada bulunan URL'ler:")
	for _, link := range links {
		fmt.Println(link)
	}
}
