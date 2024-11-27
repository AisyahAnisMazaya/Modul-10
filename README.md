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
