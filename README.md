# Projekt KRYS

---
Kod projektu jest udukomentowany przy użyciu komentarzy w plikach *.go

---

## Skład zespołu:

- Matysiak Piotr
- Nawrocki Jan
- Sobczak Wojciech

## Szyfr Trivium

---

*Przedstawiona tu implementacja szyfru `Trivium` nie jest bezpieczna i nie nadaje się do 
użytku innego niż edukacyjny.*

---

*This implementation of `Trivium` cipher isn't safe and is intended for educational purposes only.*

---
## Użytkowanie

### Prerequisites

Zanim przejdziesz dalej, upewnij się, że masz zainstalowany [język go](https://go.dev/doc/install) w swoim środowisku testowym.

Szyfr został przetestowany na wersji `go1.16.6`. Jeżeli nie posiadasz GO i nie chcesz posiadać, użyj
skomilowanego pliku, dostępnego w repozytorium `Trivium\Cipher\cmd\trivium.exe` (Tylko windows!)

Upewnij się, że GO został poprawnie zainstalowany wpisując `go version` i przejdź do sekcji [Budowanie](#Budowanie)

### Budowanie

Aby zbudować plik wykonywalny w folderze `Trivium\Cipher\cmd` wpisz:

#### a) WINDOWS
```bash
go build -o trivium.exe .
```

#### b) LINUX / macOS

```bash
go build -o trivium .
```

### Przykłady użycia

*W dalszych przykładach używaj pliku `trivium.exe` - jeżeli pracujesz na windows 
lub `trivium`, jeżeli pracujesz na Linux/macOS*

#### Dozwolone flagi: -mode -msg -cipher -key -iv
```bash
# Usage of .\trivium.exe:
  -cipher string
        Ciphertext to decrypt (HEX)
  -iv string
        An IV to encrypt/decrypt message (80 bit, HEX)
  -key string
        A KEY to encrypt/decrypt message (80 bit, HEX)
  -mode string
        e - encrypt, d - decrypt output HEX, dp - decrypt - output HEX and PLAINTEXT (default "null")
  -msg string
        Message to be encrypted/decrypted (HEX))

```
#### Szyfrowanie
Aby zaszyfrować wiadomość (msg), należy wywołać flagi: **-mode** ("e"), **-msg** (hex), **-key** (hex, 80 bit) oraz **-iv** (hex, 80 bit).

***Przykłady:***
```bash
# Example 1
.\trivium.exe -mode "e" -msg "00000000000000000000000000000000000000" -key "80000000000000000000" -iv "00000000000000000000"

# Example 2
.\trivium.exe -mode "e" -msg "416A666F6E20637A79207369616A6F6D693F" -key "54656c636f52756c6573" -iv "4e696d6f6d706f6a6563"
```

#### Deszyfrowanie

Aby zdeszyfrować kryptogram (cipher), należy wywołać flagi: **-mode** ("d" lub "dp"), **-cipher** (hex), **-key** (hex, 80 bit) oraz **-iv** (hex, 80 bit).

***Przykłady:***
```bash
# Example 1
.\trivium.exe -mode "d" -cipher "38EB86FF730D7A9CAF8DF13A4420540DBB7B65" -key "80000000000000000000" -iv "00000000000000000000"

.\trivium.exe -mode "dp" -cipher "38EB86FF730D7A9CAF8DF13A4420540DBB7B65" -key "80000000000000000000" -iv "00000000000000000000"

# Example 2

.\trivium.exe -mode "d" -cipher "D89507972D45AD8E9FF26175B49D512A86F9" -key "54656c636f52756c6573" -iv "4e696d6f6d706f6a6563"

.\trivium.exe -mode "dp" -cipher "D89507972D45AD8E9FF26175B49D512A86F9" -key "54656c636f52756c6573" -iv "4e696d6f6d706f6a6563"
```
