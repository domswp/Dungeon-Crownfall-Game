## Dungeon/Penjara Crownfall

adalah sebuah game rogue like,text based turn base rpg classic, dengan tema dark fantasy, game ini dibuat dengan bahasa golang, dan dapat di mainkan di Terminal (CLI)
inspirasi game ini adalah dari game dark souls dan elden ring (souls like).

---

## Overview

   Versi: Alpha 0.1  
   Dibangun dengan: Go (Golang)

### Gameplay Utama

- **Pilih Karakter & Kelas**
  - Ksatria (Knight)
  - Pencuri (Rogue)
  - Penyihir (Sorcerer)

- **Peta Dungeon Acak**
  - 7x7 area penjara bawah tanah
  - Ruangan acak dengan musuh atau loot

- **Pertarungan Turn-Based**
  - Serangan dasar
  - Menggunakan item
  - Lari dari musuh

- **Sistem Musuh**
  - Goblin: licik dan cepat
  - Skeleton: tangguh dan kutukan
  - Dark Mage: penyihir bayangan yang mematikan

- **Sistem XP & Level Up**
  - Dapat XP dari musuh
  - Naik level otomatis saat XP cukup
  - Max HP meningkat saat level naik

- **Loot & Inventaris**
  - Dapat item seperti Ramuan Penyembuh dari dungeon atau musuh
  - Gunakan item saat bertarung

- **Status Pemain Real-time**
  - Tampilkan HP, XP, Level, jumlah item setiap giliran

---

## 🗺️ Fitur yang Akan Datang (Planned)

- Save system
- Mini Map (ASCII Map Viewer)
- Sistem Equipment (Senjata & Armor)
- Efek Status (Poison, Burn, Buff)
- Skill Khusus (Fireball, Heal, dsb)
- Boss Room
- More Monster

----

## Cara Menjalankan

Pastikan kamu sudah menginstall Go.

```bash
go run main.go
```

atau compile menjadi binary dengan cara

```bash
go build -o crownfall
./crownfall
```

## Notes

Game ini masil dalam tahap awal, mungkin akan menemukan beberapa monster tetapi tidak spawn lagi, atau spawnnya terlalu random, tapi hal ini menjadi notice kedepannya
maaf kalau development gamenya terkesan lambat karena, memang banyak kesibukan diluar kerjaan dan ini hanyalah proyek iseng dari saya, class yang tersedia juga masih
sekedar pajangan saja, belum ada pendetailan lebih lanjut

## Contribution

jika memang ingin membantu dalam menambahkan fitur ataupun memperbaiki bug yang ada boleh saja PR ide, apappun di terima kok

## Lisensi
MIT License © 2025 Domas Wiladatu
