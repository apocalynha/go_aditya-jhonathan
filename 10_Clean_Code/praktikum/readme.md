# Clean Code
## Code 1

1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
    - Berapa banyak kekurangan dalam penulisan kode tersebut?
    Dalam kode tersebut terdapat kekurangan sebanyak 6.

    - Bagian mana saja terjadi kekurangan tersebut?
    Kekurangan terjadi karena kesalahan penggunaan pada tipe data, penamaan function/struct yang kurang tepat dan variabel yang tidak mendeskripsikan dengan jelas tujuannya.

    - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.

```
package main

// 1. Tipe data username dan password sebaiknya String
type user struct {
	id       int
	username int
	password int
}

// 2. Penamaan struct lebih baik menjadi userService
type userservice struct {
	// 3. Variabel t tidak mendeksripsikan lebih jelas untuk apa
	t []user
}

// 4. Penamaan fungsi lebih baik menjadi getAllUsers()
func (u userservice) getallusers() []user {
	return u.t
}

// 5. Penamaan fungsi lebih baik menjadi getUserById()
func (u userservice) getuserbyid(id int) user {
	// 6. Variabel r tidak mendeksripsikan lebih jelas untuk apa
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
```

## Code 2

2. Mengubah code menjadi lebih mudah terbaca dan rapi.

```
package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()
}
```