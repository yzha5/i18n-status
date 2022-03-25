# i18n-status
version:`1.0.0_alpha.0`


`i18n-status` has a [web front-end version](//github.com/yzha5/i18n-status-web), using the combination of the two, you can easily process messages

## Install
```
go get -u github.com/yzha5/i18n-status
```

## Usage
### Base usage

```go
package main

import (
    "fmt"
    "github.com/yzha5/i18n-status"
)

func main()
    /*
    s := status.Init()
    fmt.Println(s)
    */
    
    var s = status.InitWith(status.ZH, map[string]string{
    	"key": "值",
    })
	
    fmt.Println(s)//code 0 message: 
    fmt.Println(s.CurrentLanguage())//zh
    s.SetLanguage(status.EN)
    fmt.Println(s.CurrentLanguage())//en
    fmt.Println(s.T("key"))//值
    err := s.New(status.Server, "test error message")
    fmt.Println(err)//code 3 message: test error message
    err = s.New(status.Cache, s.T("key"))
    fmt.Println(err)//code 7 message: 值
    fmt.Println(s.Code())//7
    fmt.Println(s.Msg())//值
    err = s.Newf(status.Gateway, "%s--%d", "string", 123)
    fmt.Println(err)//code 2 message: string--123
    err = s.FromError(status.Redirect, err)
    fmt.Println(err)//code 9 message: string--123
}
```

## types

### type Status struct {}
### type Lang string
### type Code uint32

## APIs

### status.Init() *Status
Set the default language to EN
### status.InitWith(lang Lang, messages map[string]string) *Status
Init locale messages
### func (s *Status) SetLanguage(l Lang)
Set a language for current
### func (s *Status) SetMessage(message map[string]string)
Set message for current
### func (s *Status) CurrentLanguage() Lang
Get current language
### func (s *Status) T(key string) string
Translate message
### func (s *Status) Code() Code
Get status code
### func (s *Status) Msg() string
Get Status's message text
### func (s *Status) New(code Code, text string) error
Make a new error from string
### func (s *Status) Newf(code Code, format string, a ...interface{}) error
Make a new error from format string
### func (s *Status) FromError(code Code, err error) error
Make a new error from error

## ISO Code

| Language                         | Code |
|:---------------------------------|:-----|
| Afar                             | AA   |
| Abkhazian                        | AB   |
| Afrikaans                        | AF   |
| Amharic                          | AM   |
| Arabic                           | AR   |
| Assamese                         | AS   |
| Aymara                           | AY   |
| Azerbaijani                      | AZ   |
| Bashkir                          | BA   |
| Byelorussian (Belarusian)        | BE   |
| Bulgarian                        | BG   |
| Bihari                           | BH   |
| Bislama                          | BI   |
| Bengali (Bangla)                 | BN   |
| Tibetan                          | BO   |
| Breton                           | BR   |
| Catalan                          | CA   |
| Corsican                         | CO   |
| Czech                            | CS   |
| Welsh                            | CY   |
| Danish                           | DA   |
| German                           | DE   |
| Bhutani                          | DZ   |
| Greek                            | EL   |
| English                          | EN   |
| Esperanto                        | EO   |
| Spanish                          | ES   |
| Estonian                         | ET   |
| Basque                           | EU   |
| Farsi                            | FA   |
| Finnish                          | FI   |
| Fiji                             | FJ   |
| Faeroese                         | FO   |
| French                           | FR   |
| Frisian                          | FY   |
| Irish                            | GA   |
| Gaelic (Scottish)                | GD   |
| Galician                         | GL   |
| Guarani                          | GN   |
| Gujarati                         | GU   |
| Gaelic (Manx)                    | GV   |
| Hausa                            | HA   |
| Hebrew                           | HE   |
| Hindi                            | HI   |
| Croatian                         | HR   |
| Hungarian                        | HU   |
| Armenian                         | HY   |
| Interlingua                      | IA   |
| Indonesian                       | ID   |
| Interlingue                      | IE   |
| Inupiak                          | IK   |
| Icelandic                        | IS   |
| Italian                          | IT   |
| Inuktitut                        | IU   |
| Japanese                         | JA   |
| Javanese                         | JV   |
| Georgian                         | KA   |
| Kazakh                           | KK   |
| Greenlandic                      | KL   |
| Cambodian                        | KM   |
| Kannada                          | KN   |
| Korean                           | KO   |
| Kashmiri                         | KS   |
| Kurdish                          | KU   |
| Kirghiz                          | KY   |
| Latin                            | LA   |
| Limburgish ( Limburger)          | LI   |
| Lingala                          | LN   |
| Laothian                         | LO   |
| Lithuanian                       | LT   |
| Latvian (Lettish)                | LV   |
| Malagasy                         | MG   |
| Maori                            | MI   |
| Macedonian                       | MK   |
| Malayalam                        | ML   |
| Mongolian                        | MN   |
| Moldavian                        | MO   |
| Marathi                          | MR   |
| Malay                            | MS   |
| Maltese                          | MT   |
| Burmese                          | MY   |
| Nauru                            | NA   |
| Nepali                           | NE   |
| Dutch                            | NL   |
| Norwegian                        | NO   |
| Occitan                          | OC   |
| Oromo (Afan, Galla)              | OM   |
| Oriya                            | OR   |
| Punjabi                          | PA   |
| Polish                           | PL   |
| Pashto (Pushto)                  | PS   |
| Portuguese                       | PT   |
| Quechua                          | QU   |
| Rhaeto-Romance                   | RM   |
| Kirundi (Rundi)                  | RN   |
| Romanian                         | RO   |
| Russian                          | RU   |
| Kinyarwanda (Ruanda)             | RW   |
| Sanskrit                         | SA   |
| Sindhi                           | SD   |
| Sangro                           | SG   |
| Serbo-Croatian                   | SH   |
| Sinhalese                        | SI   |
| Slovak                           | SK   |
| Slovenian                        | SL   |
| Samoan                           | SM   |
| Shona                            | SN   |
| Somali                           | SO   |
| Albanian                         | SQ   |
| Serbian                          | SR   |
| Siswati                          | SS   |
| Sesotho                          | ST   |
| Sundanese                        | SU   |
| Swedish                          | SV   |
| Swahili (Kiswahili)              | SW   |
| Tamil                            | TA   |
| Telugu                           | TE   |
| Tajik                            | TG   |
| Thai                             | TH   |
| Tigrinya                         | TI   |
| Turkmen                          | TK   |
| Tagalog                          | TL   |
| Setswana                         | TN   |
| Tonga                            | TO   |
| Turkish                          | TR   |
| Tsonga                           | TS   |
| Tatar                            | TT   |
| Twi                              | TW   |
| Uighur                           | UG   |
| Ukrainian                        | UK   |
| Urdu                             | UR   |
| Uzbek                            | UZ   |
| Vietnamese                       | VI   |
| Volapük                          | VO   |
| Wolof                            | WO   |
| Xhosa                            | XH   |
| Yiddish                          | YI   |
| Yoruba                           | YO   |
| Chinese (Simplified/Traditional) | ZH   |
| Zulu                             | ZU   |

