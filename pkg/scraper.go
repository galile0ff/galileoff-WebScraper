package pkg

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// Result, bir web sayfasından kazınan verileri tutar.
type Result struct {
	HTML       string
	Title      string
	FinalURL   string
	Screenshot []byte
	Links      []string
	HTTPStatus int
}

// Scrape, verilen URL'yi hedefler, sayfayı kazır ve bir Result nesnesi döndürür.
// Tarayıcı otomasyonu, sayfa içeriğini, ekran görüntüsünü ve linkleri almayı içerir.
func Scrape(targetURL string, infoLog, debugLog, errorLog *log.Logger) (*Result, error) {
	edgePath := FindEdge()
	if edgePath == "" {
		return nil, fmt.Errorf("microsoft Edge tarayıcısı sistemde bulunamadı")
	}
	infoLog.Println("Kullanılan tarayıcı:", edgePath)

	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.ExecPath(edgePath),
			chromedp.Flag("headless", true),
			chromedp.Flag("disable-gpu", true),
		)...,
	)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(func(string, ...interface{}) {}),
		chromedp.WithDebugf(func(string, ...interface{}) {}),
		chromedp.WithErrorf(func(string, ...interface{}) {}),
	)
	defer cancel()

	var response *network.Response
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			if e.Type == network.ResourceTypeDocument {
				debugLog.Println("Request:", e.Request.URL)
			}
		case *network.EventResponseReceived:
			if e.Type == network.ResourceTypeDocument {
				response = e.Response
				infoLog.Printf(
					"Response | Status=%d | MIME=%s | Protocol=%s",
					int(e.Response.Status),
					e.Response.MimeType,
					e.Response.Protocol,
				)
			}
		case *network.EventLoadingFailed:
			errorLog.Printf(
				"Yükleme hatası | Reason: %s",
				e.ErrorText,
			)
		}
	})

	// Kazınacak veriler için değişkenler
	var html, title, finalURL string
	var screenshot []byte
	var links []string

	// chromedp görevlerini çalıştır
	err := chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(targetURL),
		chromedp.WaitReady("body"),
		chromedp.Location(&finalURL),
		chromedp.Title(&title),
		chromedp.OuterHTML("html", &html),
		chromedp.FullScreenshot(&screenshot, 90),
		chromedp.Evaluate(`
			[...new Set(Array.from(document.querySelectorAll("a"))
			.map(a => a.href)
			.filter(h => h))]
		`, &links),
	)
	if err != nil {
		return nil, fmt.Errorf("sayfa yükleme veya kazıma işlemi başarısız: %v", err)
	}

	var httpStatus int
	if response != nil {
		httpStatus = int(response.Status)
	}

	// Sonuçları paketle
	result := &Result{
		HTML:       html,
		Title:      title,
		FinalURL:   finalURL,
		Screenshot: screenshot,
		Links:      links,
		HTTPStatus: httpStatus,
	}

	return result, nil
}
