*NOMOR 1*
Deskripsi Program
Program ini ditulis dalam bahasa Go (Golang) dan bertujuan untuk mencari berat terkecil (minimum) dan berat terbesar (maksimum) dari sejumlah berat anak kelinci yang dimasukkan oleh pengguna.
Penjelasan Fungsi dan Alur Program
Deklarasi Fungsi cariMinMax
Fungsi ini menerima dua parameter:
kelinci: array berisi berat anak kelinci.
n: jumlah anak kelinci yang datanya akan diproses.
Fungsi ini:
Menginisialisasi nilai minimum (min) dan maksimum (max) dengan elemen pertama array.
Mengiterasi array untuk membandingkan setiap elemen dengan nilai min dan max:
Jika elemen lebih kecil dari min, elemen tersebut menjadi nilai minimum baru.
Jika elemen lebih besar dari max, elemen tersebut menjadi nilai maksimum baru.
Mengembalikan nilai min dan max.
Fungsi main
Meminta input jumlah anak kelinci (n) dari pengguna.
Memvalidasi apakah nilai n berada di antara 1 hingga 1000:
Jika tidak valid, program memberikan pesan kesalahan dan berhenti.
Meminta input berat masing-masing kelinci sebanyak n kali.
Memanggil fungsi cariMinMax untuk menghitung berat terkecil dan terbesar.
Menampilkan hasil pencarian berupa berat terkecil dan terbesar.
Fitur Program
Validasi Input: Memastikan jumlah kelinci berada dalam rentang 1 hingga 1000.
Dukungan Format Float: Memungkinkan input berat kelinci dengan angka desimal.
Hasil Akurat: Menampilkan berat terkecil dan terbesar dengan format dua angka di belakang desimal.

*NOMOR 2*
Deskripsi Program
Program ini ditulis dalam bahasa Go (Golang) dan digunakan untuk menghitung total berat ikan yang ditempatkan dalam beberapa wadah dengan kapasitas tertentu. Program juga menghitung rata-rata berat ikan per wadah.
Penjelasan Fungsi dan Struktur
Struktur Data Wadah
Struktur ini digunakan untuk menyimpan informasi tentang wadah:
berat: array yang menyimpan total berat ikan di setiap wadah.
jumlah: jumlah wadah yang digunakan.
rataRata: rata-rata berat ikan per wadah.
Fungsi hitungTotalBerat Fungsi ini bertugas menghitung total berat ikan yang dimasukkan ke dalam setiap wadah, berdasarkan jumlah ikan dan kapasitas wadah yang diberikan:
Parameter:
w: pointer ke struktur Wadah yang akan diperbarui.
ikan: array berisi berat ikan.
x: jumlah total ikan.
y: kapasitas maksimum setiap wadah.
Logika:
Memproses ikan secara bertahap, memasukkan ikan ke dalam wadah hingga mencapai kapasitas maksimum (y), atau hingga ikan habis.
Menambahkan total berat tiap wadah ke array w.berat.
Menghitung rata-rata berat ikan per wadah dengan membagi total berat seluruh wadah dengan jumlah wadah yang digunakan.
Fungsi main
Meminta input jumlah ikan (x) dan kapasitas maksimum wadah (y).
Memvalidasi input (x tidak boleh lebih dari 1000).
Meminta input berat masing-masing ikan.
Memanggil fungsi hitungTotalBerat untuk memproses data ikan ke dalam wadah.
Menampilkan total berat di setiap wadah dan rata-rata berat per wadah.
Fitur Program
Pemisahan Ikan ke dalam Wadah:
Program secara otomatis membagi ikan ke dalam wadah berdasarkan kapasitas maksimum yang ditentukan.
Validasi Input:
Memastikan jumlah ikan tidak melebihi batas maksimum (1000).
Output Terstruktur:
Menampilkan total berat ikan di setiap wadah dan rata-rata berat ikan secara rapi dengan dua angka desimal.
