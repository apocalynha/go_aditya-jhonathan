package main

// Tipe data username dan password sebaiknya String
type user struct {
	id       int
	username int
	password int
}

// Penamaan struct lebih baik menjadi userService
type userservice struct {
	// Variabel t tidak mendeksripsikan lebih jelas untuk apa
	t []user
}

// Penamaan fungsi lebih baik menjadi getAllUsers()
func (u userservice) getallusers() []user {
	return u.t
}

// Penamaan fungsi lebih baik menjadi getUserById()
func (u userservice) getuserbyid(id int) user {
	// Variabel r tidak mendeksripsikan lebih jelas untuk apa
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}