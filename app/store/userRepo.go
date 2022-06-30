package store

import "github.com/b-o-e-v/rest-api-store/app/model"

type UserRepo struct {
	store *Store
}

func (r *UserRepo) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING ID",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepo) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}