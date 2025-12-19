package cli

import (
	"errors"
	"math/rand"
	"time"
)

type Art struct {
	Name  string
	Color string
	Data  string
}

// Renk kodları
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorPurple = "\033[35m"
	ColorWhite  = "\033[37m"
	ColorReset  = "\033[0m"
)

var arts = []Art{
	{"banner1", ColorCyan, banner1},
	{"banner2", ColorRed, banner2},
	{"banner3", ColorGreen, banner3},
	{"banner4", ColorYellow, banner4},
	{"banner5", ColorPurple, banner5},
	{"banner6", ColorBlue, banner6},
	{"banner7", ColorWhite, banner7},
	{"banner8", ColorGreen, banner8},
	{"banner9", ColorRed, banner9},
	{"banner10", ColorCyan, banner10},
	{"banner11", ColorYellow, banner11},
	{"banner12", ColorPurple, banner12},
	{"banner13", ColorBlue, banner13},
	{"banner14", ColorWhite, banner14},
	{"banner15", ColorRed, banner15},
	{"banner16", ColorYellow, banner16},
	{"banner17", ColorRed, banner17},
	{"banner18", ColorWhite, banner18},
	{"banner19", ColorCyan, banner19},
	{"banner20", ColorGreen, banner20},
	{"banner21", ColorRed, banner21},
	{"banner22", ColorWhite, banner22},
	{"banner23", ColorBlue, banner23},
	{"banner24", ColorWhite, banner24},
	{"banner25", ColorRed, banner25},
	{"banner26", ColorYellow, banner26},
}

func RandomArt() Art {
	rand.Seed(time.Now().UnixNano())
	return arts[rand.Intn(len(arts))]
}

func GetArt(name string) (Art, error) {
	for _, a := range arts {
		if a.Name == name {
			return a, nil
		}
	}
	return RandomArt(), errors.New("ascii bulunamadı")
}

const banner1 = `
      (   )   (   )        
       \ /     \ /             LITTLE LITTLE INTO THE MIDDLE
     __.'._____.'__        -------------------------------------
    |              |        "Veriyi ortaya karışık yaptım."
    |  _ _ _ _ _   |       
    | ( ) ( ) ( )  |        "Hani marjinal bizdik?"
    |  _ _ _ _ _   |       
    |______________|        (Tabak tabak veri geliyo abi masaya)
`

const banner2 = `
      .---.        .___________
     /     \  __  /    ------
    / /     \(..)/    -----
   //////   ' \/ '   ---
  //// / // :    : ---
 // /   /  / '  '
//          //..\\
=============UU==UU=======================================================
                             G A L I L E O F F   W E B   K A Z I Y I C I
==========================================================================
`

const banner3 = `
 ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄     ▄▄▄ ▄▄▄     ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ 
█       █       █   █   █   █   █   █       █       █       █       █
█   ▄▄▄▄█   ▄   █   █   █   █   █   █    ▄▄▄█   ▄   █    ▄▄▄█    ▄▄▄█
█  █  ▄▄█  █▄█  █   █   █   █   █   █   █▄▄▄█  █ █  █   █▄▄▄█   █▄▄▄ 
█  █ █  █       █   █▄▄▄█   █   █▄▄▄█    ▄▄▄█  █▄█  █    ▄▄▄█    ▄▄▄█
█  █▄▄█ █   ▄   █       █   █       █   █▄▄▄█       █   █   █   █    
█▄▄▄▄▄▄▄█▄▄█ █▄▄█▄▄▄▄▄▄▄█▄▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄▄▄▄▄█▄▄▄█   █▄▄▄█    
             >> SİSTEM KEŞİF ve İSTİHBARAT ARACI <<
`

const banner4 = `
      _______________________
     /                      /|  
    /   H A K L A R I M    / |  
   /________VAR!__________/  |    "Anayasa diyor ki her site kazınabilir!"
  |  ___   ___   ___   ___|  |                        - galileoff Goodman
  | |   | |   | |   | |   |  |
  | |___| |___| |___| |___|  |    "Başı dertte bir veri tabanı mı?"
  |_______________________| /   
                           /       Better Call galileoff!
`

const banner5 = `
      .-------.
     /   _ _   \       LOS DATOS HERMANOS
    |   (o_o)   |    -----------------------
    |    (_)    |    "Tavuklarımız lezzetli,
    |   __|__   |     verilerimiz taze."
   /   /  |  \   \   
  /   /   |   \   \  (Gustavo Fring Onaylı)
 |    |---|---|    |
      |   |   |      "Biz sadece tavuk satmıyoruz..."
`

const banner6 = `
   G   A   L   I   L   E   O - F F   [+] Hedef: Belirlendi  [+] Mod: Gizli
 __________________________________________________________________________
 \________________________________________________________________________/
`

const banner7 = `
 ██████╗  █████╗ ██╗     ██╗██╗     ███████╗ ██████╗ ███████╗███████╗
██╔════╝ ██╔══██╗██║     ██║██║     ██╔════╝██╔═══██╗██╔════╝██╔════╝
██║  ███╗███████║██║     ██║██║     █████╗  ██║   ██║█████╗  █████╗  
██║   ██║██╔══██║██║     ██║██║     ██╔══╝  ██║   ██║██╔══╝  ██╔══╝  
╚██████╔╝██║  ██║███████╗██║███████╗███████╗╚██████╔╝██║     ██║     
 ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝╚══════╝╚══════╝ ╚═════╝ ╚═╝     ╚═╝     
                       - SİBER İSTİHBARAT -
`

const banner8 = `
   _____      _ _ _             __  __ 
  / ____|    | (_) |           / _|/ _|
 | |  __  __ | |_| | ___   ___| |_| |_ 
 | | |_ |/ _|| | | |/ _ \ / _ \  _|  _|
 | |__| | (_ | | | |  __/| (_) | | | |  
  \_____|\__,_|_|_|_|\___| \___/|_| |_|  
                                       
         --[ WEB KAZIYICI ]--
`

const banner9 = `
      _.-"   "-._        
    .'   .   .   '.      ---------------------------------------
   /  .  .   .  .  \     "Azizim, burası Vahşi Batı..."
   |  .  .   .  .  |
   \    (o) (o)    /     "Burada robots.txt sökmez!"
    \      ^      /      
     |    ___    |       "Bana bunu yapma Jerry, verileri ver!"
      \_________ /
        \  *  /          (Şerif Yıldızı Veri İçin Parlıyor)
         \___/
`

const banner10 = `
         ___
        /   \                                    SÖYLESENE!!
       /_____\         --------------------------------------
      (   _   )        "Sen 18'de ne yapıyordun?"
       \_____/         "19'da ne yapıyordun? 20'de?"
          |
         _|_           "Ben o portlarda scraping yapıyordum!"
        (___)          "Anladın mı? Scraping!"
`

const banner11 = `
      |\__/,|   ('"` + "`" + `\                 G A L I L E O F F
     _.|o o  |_   ) )           ----------------------------------
   -(((---(((--------            "Tırmalamadık site bırakma knk."
 `

const banner12 = `
         ( (            SAKİN OL
          ) )              VE
        ........     KAZIMAYA BAŞLA
        |      |]    - GALILEOFF -
        \      /    
         '----'      
`

const banner13 = `
      _______
     /       \         
    |  _   _  |        ----------------------------------
    | (o) (o) |        "Bu alemde gece mahkum olan,
    |    <    |         gündüz veri çekemez!"
     \  ___  /         
      \_____/          "Sonunu düşünen kahraman olamaz,
                        veriyi çeken kral olur."
    
    (Polat Alemdar Onaylı)
`

const banner14 = `
       .           		               G A L I L E O F F
      ":"          		 --------------------------------
    ___:____     |"\/"|
  ,'        '.    \  /      "Vurgun yeme, veri ye..."
  |  O        \___/  |
~^~^~^~^~^~^~^~^~^~^~^~^~
 `

const banner15 = `
          _
        _( )_
       |     |                            HOKKAĞBAZ MI OYNUYO BURDA
       |_____|           -------------------------------------------
         | |             "Biz hacker değiliz İskender abi..."
       ( ) ( )
      (       )          "Biz scraperız! El çabukluğu marifet."
       \     /
        |   |            "Yoksa veri orada duruyor, almamak ayıp."
`

const banner16 = `
   ___________________________________________________________
  |                                                           |
  |  "SİZİN DE HAKLARINIZ VAR! ANAYASA DİYOR Kİ VAR!"         |
  |  "BEN İNANIYORUM Kİ AKSİ İSPATLANANA KADAR..."            |
  |  HER VERİ MASUMDUR!                                       |
  |___________________________________________________________|
           \
            \     ( ͡° ͜ʖ ͡°)
                 <|     |>  
                  |     |
`

const banner17 = `
       _______
      /      //      
     /______//         -------------------------------------------
    (______(/          "Bizim barı açıyoruz Altan..."
    /      \
   /  ____  \          "Bilemiyorum Altan..."
  |  (    )  |
   \  \__/  /          "Verileri çekince bir rahatlama geliyor."
    \______/           "En azından HTML temiz Altan."
`

const banner18 = `
         ,---.
        (     )        
         \   /         --------------------------------------------
          | |          "Mavi Fransız Kornası'nı senin için çaldım."
      ___/   \___      (Şaka şaka, verilerini çaldım.)
     /           \
    |             |    g a l i l e o f f
     \___________/     
     
     "Sanırım bu veriye aşık oldum."
`

const banner19 = `
  _______________________
 /                       \
|     BU ABİ BİRİ Mİ?     |
| GALILEOFF KİM Kİ ZATEN? |
 \_______________________/
          \
           \   (o)__(o)
              (   ..   )
              (  ____  )
              (_______)
`

const banner20 = `
        _.---._
      .'       '.        
     /           \       -----------------------------------------
    |   O  O  O   |      "Komutan Logar, bir cisim yaklaşıyor..."
     \___________/       "- Veriymiş, veri!"
        / / \ \          
       /_/   \_\         "Seni seçtim çünkü sen farklısın,
                         Japon yapıştırıcısı gibisin."
`

const banner21 = `
      .-""""""-.
    .'          '.        
   /   O      O   \       --------------------------------------------
  :           '    :      "Sen hata mı gördün?"
  |                |      
  :    .------.    :      "Kodu çalıştırırken hata mı gördün?"
   \  '        '  /       
    '.          .'        "Hayır yani, hata varsa söyle düzeltelim."
      '-......-'
`

const banner22 = `
      _________
     |   BAR   |                     mukadderat diyelim.
   __|_________|__       --------------------------------
  |               |      "Barı açıyorum..."
  |  ( )     ( )  |      
  |   |       |   |      "Ayla'yla aramı düzeltiyorum."
  |  _|_     _|_  |      
  |_______________|      "Babamı da yanıma alıyorum!"
                         
`

const banner23 = `
      ,----------.
     /  _      _  \          
    |  (o)    (x)  |         -------------------------
    |   \____/     |         "Ben kodunu beğenmezsem,
    |    |__|      |          bu uzman görüşü olur."
   /     /  \       \        
  |     /    \       |       "Sen beğenmezsen..."
 /     /      \       \      "Çekememezlik olur!"
`

const banner24 = `
         .---.
        /     \                     Bekir'in de Kaderi Uğursuzdu
       | (o) (o)         ---------------------------------------
       |    |            "Bak site..."
      /|   ___           "Ben bu scraping'i bırakırım dedim,"
     / |  \___/          "Bırakamadım..."
    /  |_______|         "İmkansız olduğunu bile bile... 
      /   |  | \          403 yiyeceğimi bile bile..."
`

const banner25 = `
       .---.
      /  _  \          
     |  (o)  |         ---------------------------
     |   |   |         "Nereye gidiyorsun Bekir?"
    /   -^-   \        
   /     |     \       "- Verinin gittiği yere."
  /      |      \      
`

const banner26 = `
      ________________
     | |__|__|__|__|__|      
     |  _      _      |      ----------------------------------------
     | |o|    |o|     |      "Dönüp dolaşıp aynı siteye geliyorsun."
     |_|_|____|_|_____|      
      (o)      (o)           "Kaderin bu senin Bekir,
                             bu HTML'i parse etmek..."
`
