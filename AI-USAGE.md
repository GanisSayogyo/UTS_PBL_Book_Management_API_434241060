\# AI Usage Documentation



\## 1. Tujuan Penggunaan AI



AI digunakan sebagai alat bantu selama pengembangan project Book Management API untuk mata kuliah Backend. Penggunaan AI bertujuan membantu memahami materi perkuliahan, merancang struktur aplikasi, menyusun implementasi awal, menganalisis kesalahan, merencanakan pengujian, dan menyiapkan dokumentasi project.



AI digunakan melalui diskusi dan instruksi bertahap. Implementasi disesuaikan dengan kebutuhan tugas, teknologi yang dipelajari di kelas, dan batas waktu pengerjaan.



\## 2. Tools yang Digunakan



\- \*\*ChatGPT:\*\* Membantu penjelasan konsep, perancangan, contoh kode, debugging, penyusunan perintah pengujian, dan dokumentasi.

\- \*\*Go dan Fiber v2:\*\* Teknologi utama untuk membangun REST API.

\- \*\*PostgreSQL dan pgx:\*\* Penyimpanan data dan komunikasi aplikasi dengan database.

\- \*\*PowerShell/CMD:\*\* Menjalankan aplikasi, menguji endpoint, dan mengelola project.

\- \*\*Git dan GitHub:\*\* Pengelolaan versi kode dan penyimpanan repository.



\## 3. Bagian Project yang Dibantu AI



\### 3.1 Perancangan Project dan Struktur Kode



AI membantu menentukan ruang lingkup Book Management API dan menyusun struktur project yang memisahkan model, repository, service, handler, konfigurasi, database, middleware, helper, dan route.



Pemisahan ini digunakan agar tanggung jawab setiap bagian lebih jelas dan implementasi lebih mudah dipahami serta dikembangkan.



\### 3.2 Konfigurasi Aplikasi dan Environment



AI membantu menjelaskan pengelolaan konfigurasi aplikasi melalui environment variable, penggunaan file .env dan .env.example, serta pengaturan port server dan koneksi database.



AI juga membantu mengingatkan agar file yang berisi kredensial tidak dimasukkan ke repository GitHub.



\### 3.3 Database dan Repository Pattern



AI membantu menyusun model data pengguna dan buku, migrasi tabel PostgreSQL, koneksi menggunakan pgx connection pool, serta repository untuk mengakses database.



Bantuan mencakup penggunaan query berparameter, penanganan data yang tidak ditemukan, penanganan duplikasi data, dan pemisahan operasi database dari logika bisnis.



\### 3.4 Authentication



AI membantu menjelaskan dan menyusun implementasi registrasi serta login pengguna, hashing password menggunakan bcrypt, pembuatan JWT, dan validasi token melalui middleware.



AI juga membantu menjelaskan penggunaan header Authorization dengan skema Bearer untuk mengakses endpoint yang memerlukan autentikasi.



\### 3.5 Authorization dan Role-Based Access Control



AI membantu menyusun mekanisme pembatasan akses berdasarkan role user dan admin.



Dalam implementasi ini, registrasi publik membuat akun dengan role user. Pengguna harus melewati autentikasi sebelum pemeriksaan role, sedangkan operasi pembuatan, pembaruan, dan penghapusan buku dibatasi untuk admin.



AI membantu menjelaskan perbedaan status HTTP 401 Unauthorized dan 403 Forbidden dalam konteks pengujian akses.



\### 3.6 Implementasi CRUD Buku



AI membantu menyusun dan menjelaskan implementasi endpoint untuk membuat, melihat, memperbarui, dan menghapus data buku.



Bantuan juga mencakup pembaruan seluruh data menggunakan PUT, pembaruan sebagian data menggunakan PATCH, pencarian, pagination, validasi input, dan penanganan respons HTTP.



\### 3.7 Validasi dan Error Handling



AI membantu menentukan penggunaan validator untuk memeriksa input serta menjelaskan penanganan kondisi seperti request tidak valid, autentikasi gagal, akses ditolak, data tidak ditemukan, dan data duplikat.



AI juga membantu menyesuaikan respons endpoint dengan status HTTP yang sesuai dengan kondisi pengujian.



\### 3.8 Debugging dan Pemecahan Masalah



AI digunakan untuk membantu menganalisis kesalahan selama implementasi, memahami ketidaksesuaian pemanggilan fungsi dan signature method, serta memberikan langkah perbaikan.



AI juga membantu menyusun perintah terminal agar proses menjalankan aplikasi, memeriksa konfigurasi, dan menguji endpoint dapat dilakukan secara bertahap.



\### 3.9 Pengujian API



AI membantu menyusun urutan pengujian manual menggunakan PowerShell, meliputi:



\- Registrasi pengguna.

\- Login dan penerimaan token JWT.

\- Pengambilan profil pengguna.

\- Pengambilan daftar buku.

\- Pembuatan buku oleh admin.

\- Pembaruan data menggunakan PUT.

\- Pembaruan sebagian data menggunakan PATCH.

\- Penghapusan buku dan pemeriksaan setelah penghapusan.

\- Penolakan akses admin terhadap pengguna biasa.

\- Pemeriksaan endpoint health.

\- Eksekusi automated test menggunakan perintah Go.



Hasil pengujian diperiksa melalui respons yang dihasilkan aplikasi. Pengujian automated test dijalankan menggunakan perintah `go test ./... -count=1`.



\### 3.10 Dokumentasi Project



AI membantu menyusun dan memperbaiki dokumentasi README, termasuk deskripsi project, teknologi, struktur direktori, konfigurasi environment, langkah menjalankan aplikasi, daftar endpoint, aturan akses, dan instruksi pengujian.



AI juga membantu menyusun dokumentasi penggunaan AI ini dan merapikan materi laporan pengujian.



\## 4. Batasan Penggunaan AI



AI digunakan sebagai alat bantu, sehingga saran dan kode yang dihasilkan tidak dianggap benar secara otomatis. Implementasi tetap perlu diperiksa terhadap kebutuhan tugas, dijalankan, dan diuji.



Penjelasan AI digunakan untuk membantu memahami tujuan fungsi, alur request, komunikasi dengan database, autentikasi, authorization, serta respons API. Jika ditemukan kesalahan selama implementasi atau pengujian, kode perlu diperbaiki sebelum dianggap selesai.



\## 5. Tanggung Jawab Mahasiswa



Mahasiswa bertanggung jawab atas repository dan hasil akhir project yang dikumpulkan, termasuk konfigurasi, implementasi, pengujian, keamanan kredensial, serta dokumentasi.



Mahasiswa perlu memahami fungsi setiap bagian kode dan mampu menjelaskan alur aplikasi, alasan penggunaan teknologi, mekanisme autentikasi dan authorization, serta hasil pengujian kepada dosen.



\## 6. Kesimpulan



ChatGPT digunakan sebagai pendamping pembelajaran dan pengembangan selama pengerjaan Book Management API, terutama dalam perancangan, implementasi, debugging, pengujian, dan dokumentasi. Penggunaan AI ditujukan untuk membantu proses belajar dan mempercepat penyelesaian kendala teknis, sementara tanggung jawab terhadap pemahaman, pemeriksaan, dan hasil akhir project tetap berada pada mahasiswa.



