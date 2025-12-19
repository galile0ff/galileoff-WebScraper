<div align="center">

# 🕷️ galileoff-WebScraper

![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Maintained](https://img.shields.io/badge/Maintained-Yes-blue?style=for-the-badge)

**Modern, Hızlı ve Güçlü Web Kazıma Aracı**

*Siber Vatan Programı Yıldız CTI Ekibi görevi kapsamında geliştirilmiştir.*

[Özellikler](#-özellikler) • [Kurulum](#-kurulum) • [Kullanım](#-kullanım) • [Teknolojiler](#-teknolojiler) • [Destek](#-destek)

</div>

---

## 📖 Hakkında

**galileoff-WebScraper**, Go dili ve `chromedp` kütüphanesi kullanılarak geliştirilmiş, gelişmiş bir web kazıma (web scraping) aracıdır. Modern web sitelerinin dinamik içeriklerini (JavaScript ile yüklenen veriler dahil) yakalayabilir, ekran görüntüsü alabilir ve sayfa üzerindeki tüm bağlantıları analiz edebilir.

Kullanıcı dostu CLI (Komut Satırı Arayüzü) ve görsel geri bildirimleri ile siber güvenlik araştırmacıları ve geliştiriciler için pratik bir çözüm sunar.

## ✨ Özellikler

- 🚀 **Headless Browser**: Görünmez bir tarayıcı (headless Chrome/MS Edge) kullanarak JavaScript tabanlı siteleri eksiksiz tarar.
- 📸 **Otomatik Screenshot**: Hedef sitenin tam sayfa ekran görüntüsünü alır ve kaydeder.
- 🔗 **Link Çıkarma**: Sayfadaki tüm bağlantıları (href) toplayarak raporlar.
- 💾 **HTML Dökümü**: Sayfanın işlenmiş son HTML halini kaydeder.
- 🎨 **Etkileşimli CLI**: ASCII bannerlar, ilerleme çubukları (spinner) ve renkli terminal çıktıları.
- 📂 **Organize Çıktı**: Her tarama için siteye özel klasörler oluşturur ve logları tutar.
- 🔄 **URL Normalizasyonu**: Girilen URL'leri otomatik olarak düzeltir ve standart formata getirir.

## 🛠 Kurulum

Projeyi yerel makinenizde çalıştırmak için aşağıdaki adımları izleyin:

### Gereksinimler
- [Go](https://go.dev/dl/) (1.23 veya üzeri)
- Google Chrome veya MS Edge (Chromedp için gereklidir)

### Adım 1: Depoyu Klonlayın
```bash
git clone https://github.com/galile0ff/galileoff-WebScraper.git
cd galileoff-WebScraper
```

### Adım 2: Bağımlılıkları Yükleyin
```bash
go mod tidy
```

## 🚀 Kullanım

Projeyi çalıştırmak için terminalde aşağıdaki komutu kullanın:

```bash
go run main.go
```

Program başladığında sizi karşılayan menüden sonra hedef URL'yi girin (örn: `galileoff.com`). Araç otomatik olarak:
1. Siteye bağlanır.
2. İçeriği analiz eder.
3. Sonuçları (HTML, Screenshot, Linkler) site adıyla oluşturulan klasöre kaydeder.

### Kontroller
- **F**: Yeni bir tarama başlatır.
- **SPACE**: Programdan çıkış yapar.

## 💻 Teknolojiler

Bu proje aşağıdaki açık kaynak teknolojiler kullanılarak oluşturulmuştur:

- **[Go (Golang)](https://go.dev/)**: Ana programlama dili.
- **[Chromedp](https://github.com/chromedp/chromedp)**: Chrome DevTools Protocol ile tarayıcı otomasyonu.
- **[Spinner](https://github.com/briandowns/spinner)**: Terminal ilerleme göstergesi.
- **[Term](https://golang.org/x/term)**: Terminal raw mod ve giriş işlemleri.

## ☕ Destek

Bu projeyi beğendiyseniz ve geliştirmeme destek olmak isterseniz bana bir kahve ısmarlayabilirsiniz!

<div align="center">
<a href="https://www.buymeacoffee.com/galile0ff" target="_blank">
<img src="https://cdn.buymeacoffee.com/buttons/v2/default-red.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" >
</a>
</div>

## 📄 Lisans

Bu proje MIT Lisansı altında dağıtılmaktadır. Detaylar için `LICENSE` dosyasına bakınız.

---

<div align="center">
Made with ❤️ by <a href="https://github.com/galile0ff">galile0ff</a>
</div>
