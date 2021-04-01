Ingat fundamental gunkan .env .gitignore .ymal atau ainnya dan buat reusable code

1. read database in index (done)
   a. html untuk UI (GET)
   b. backend connection ke database
   c. buat backend read database
   d. render ke html
   e. read barang dengan idnya
2. login 
   a. login html untuk UI
   b. backed login untuk verify user
   c. authentic login
   d. jika berhasil render dasbord jika gagal kembali ke login dan diberi tahu errornya
3. register
   a. sama seperti login
4. dasbord
   a. dasbord UI 
   b. pilihan edit input delete
   c. backend dasbord
   d. logout
5. input barang
   a. input UI
   b. penerimaan data dari HTML ke backend (POST)
   c. middleware prosess
   d. jika berhasil berikan tanda berhasil jika gagal berikan tanda gagal
6. edit barang
   a. edit UI
      i. tampilkan semua data dan berikan satu tombol tambahan yaitu edit  
      ii. value sudah ada isinya yaitu data yang akan di edit dan berikan bisa di edit dan yang tidak bisa
   b. pengiriman data dari HTML ke backend (PUT) / kalo bisa jika stuck gunakan post
   c. middleware prosess edit
   d. sama seperti di prosesn input 5d
7. delete barang
   a. alternatifnya gabungin saja dengan edit
   b. pengiriman delete dari html ke backend(DELETE)
8. edit profil
   a. edit profil UI
   b. pengiriman data dari html ke backend(PUT)
   c. pemberitahuan status bisa tidaknya
9. checkout
   TBA