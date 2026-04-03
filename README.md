# 🗡️ Dungeon Crownfall: CLI Roguelike RPG

Sebuah petualangan _Roguelike Text-Based RPG_ berbasis Command Line Interface (CLI) yang ditulis dalam C / **Golang**. Jelajahi ruang bawah tanah Crownfall, kalahkan berbagai monster ganas, kumpulkan loot, dan bertahanlah selama mungkin untuk melihat seberapa dalam Lantai yang bisa kamu taklukkan!

## ✨ Fitur Utama

- **🗺️ Procedural Dungeon Generation (Infinite Floors)** 
  Setiap level diacak kembali setiap kali karakter melangkah masuk ke Portal. Area akan berskala (*scaling*) naik tingkat kesulitannya sampai tak terhingga.
- **📈 Progressive Experience & Level Up HUD**
  Dapatkan Exp dari setiap musuh yang dikalahkan untuk mendongkrak status HP, Atk, dan Def. Tersedia *Heads-Up Display* (HUD) UI interaktif secara langsung di atas peta untuk memantau statusmu setiap saat.
- **👹 Sistem Pertarungan Turn-Based Acak**
  Hadapi berbagai musuh dengan nama dan statistik pertarungan gila serta kalkulasi kerusakan *turn-based*.
- **💀 Boss Battles**
  Probabilitas langka untuk memicu kehadiran "*Raja (Boss)*" di setiap kelipatan lantai ke-5 (Serangan ganda & HP monster boss 3x lipat!).
- **🧑‍⚕️ Class & Roles Pemain**
  Tersedia class dasar (Ksatria, Pencuri, Penyihir) dengan pembagian status, *power* dasar yang saling melengkapi.

## 🚀 Cara Menjalankan (How to Run)

1. Pastikan Anda telah menginstal **[Go](https://go.dev/dl/)**.
2. *Clone* repositori ini, dan masuk ke dalam folder proyek menggunakan terminal.
3. Jalankan perintah berikut untuk mengeksekusi game:
   ```bash
   cd roguelike_rpg
   go run main.go
