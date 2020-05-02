Latihan ini menunjukkan bagaimana menghubungkan interface dengan struct yang berbeda file. 
Ada 2 struct berbeda yang disimpan di `main.go` dan `Rect.go`. Interface yang didefinisikan di `Rect.go` kemudian bisa digunakan di `main.go`. Ini karena keduanya berada di package yang sama, yakni `main`.

Untuk menjalankan, lakukan `go run *.go`
