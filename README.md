# API QURAN INDONESIA

Api Quran Bahasa Indonesia adalah API untuk mendapatkan data-data Al-Quran dalam terjemahan bahasa Indonesia yang ditulis dalam bahasa pemrograman Golang

## Cara Install

## Mengubah .env
ubah PORT pada .env menggunakan port yang diinginkan, serta ubah GIN_MODE menjadi release

### Menggunakan Clone

clone repo ini menggunakan perintah

    git clone github.com/hdkef/api-quran-indonesia

compile ke binary dengan menggunakan perintah (pastikan present working directory pada folder target)

    go build server.go

jalankan file binary

    ./server

PERHATIAN!

aplikasi akan berjalan secara default pada PORT 8080 kecuali PORT pada .env telah diubah

### Menggunakan Docker

jalankan perintah berikut di terminal

    docker run -d -p 1010:8080 --name quran-indonesia 081218068401/quran-indonesia

aplikasi akan berjalan pada localhost:1010 dengan nama container api-lokasi-indonesia

## Cara Menggunakan

### Mendapatkan List Surat

masukkan perintah berikut

    HOST:PORT/sura

### Mendapatkan Semua Ayat

masukkan perintah berikut

    HOST:PORT/aya/[id surat]/all

misalnya ingin mendapatkan semua ayat dari surat al fatihah

    HOST:PORT/aya/1/all

response yang didapat tampak seperti ini

[{"id":"1","sura":"1","aya":"1","indonesia":"Dengan nama Allah Yang Maha Pengasih, Maha Penyayang.","footnotes":"","arabic":"بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ"},{"id":"2","sura":"1","aya":"2","indonesia":"Segala puji bagi Allah, Tuhan seluruh alam,","footnotes":"","arabic":"الْحَمْدُ لِلَّهِ رَبِّ الْعَالَمِينَ"},....dst]

### Mendapatkan Ayat Dimulai dari Ayat Tertentu dengan Range Tertentu

#### Misal Mengambil Satu Ayat Pada Surat Tertentu

masukkan perintah berikut

    HOST:PORT/aya/[id surat]/take/[ambil berapa ayat]/from/[dimulai dari ayat berapa]

misal

    HOST:PORT/aya/1/take/1/from/5

maka akan mendapatkan respon seperti berikut

[{"id":"5","sura":"1","aya":"5","indonesia":"Hanya kepada Engkaulah kami menyembah dan hanya kepada Engkaulah kami mohon pertolongan.","footnotes":"","arabic":"إِيَّاكَ نَعْبُدُ وَإِيَّاكَ نَسْتَعِينُ"}]

#### Misal Mengambil Beberapa Ayat Pada Surat Tertentu

masukkan perintah berikut

    HOST:PORT/aya/[id surat]/take/[ambil berapa ayat]/from/[dimulai dari ayat berapa]

misal

    HOST:PORT/aya/1/take/3/from/2

maka akan mendapatkan respon seperti berikut

[{"id":"2","sura":"1","aya":"2","indonesia":"Segala puji bagi Allah, Tuhan seluruh alam,","footnotes":"","arabic":"الْحَمْدُ لِلَّهِ رَبِّ الْعَالَمِينَ"},{"id":"3","sura":"1","aya":"3","indonesia":"Yang Maha Pengasih, Maha Penyayang,","footnotes":"","arabic":"الرَّحْمَٰنِ الرَّحِيمِ"},{"id":"4","sura":"1","aya":"4","indonesia":"Pemilik hari pembalasan.1)","footnotes":"*1) Yaumiddīn (hari pembalasan), hari waktu manusia menerima pembalasan amalnya, baik atau buruk. Disebut juga yaumul qiyāmah, yaumul hisab dan sebagainya.","arabic":"مَالِكِ يَوْمِ الدِّينِ"}]