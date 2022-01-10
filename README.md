# Projekt KRYS

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

#### Dozwolone flagi: -d -e -i -k -m
```bash
# .\trivium.exe -h
  -d    Decrypt encrypted message in the fly. Disallowed without '-e' flag
  -e    Encrypt message
  -i string
        An iv to encrypt/decrypt message (80 bit, hex)
  -k string
        A key to encrypt/decrypt message (80 bit, hex)
  -m string
        Message to be encrypted/decrypted
```
#### Kodowanie
Zakoduj `-e` wiadomość `-m` przy użyciu 80-bitowego klucza `-k` oraz wektora inicjalizacji `-i`
```bash
 .\trivium.exe -e -m "Hello darkness my old friend" -k "576557416e744d417850" -i "486954726976756d2121"
```

#### Dekodowanie

Dekodowanie wymaga usprawnienia. Aktualnie możliwe jest zdekodowanie w locie aktualnie kodowanej wiadmości. 
Poniższa operacja zwróci `Error`

```bash
 .\trivium.exe -d -k "576557416e744d417850" -i "486954726976756d2121" -c "55fa555dce7457a22fcae990ceb35baa9325dbc156189d8fd3a50584"
```

#### Dekodowanie aktualnie kodowanej wiadomości

Zakoduj `-e` wiadomość `-m`, a następnie zdekoduj `-d` otrzymany szyfrogram przy użyciu 80-bitowego
klucza `-k` oraz wektora inicjalizacji `-i`

```bash
 .\trivium.exe -e -d -m "Hello darkness my old friend" -k "576557416e744d417850" -i "486954726976756d2121"
```

## TODO List

- [x] Przygotować raport
- [x] Napisać własną implementację wybranego szyfru
