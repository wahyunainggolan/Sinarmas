package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"sinarmas/model"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{Db: db}
}

// RequestOtp implements MangaRepositoryInterface
func (m *UserRepository) RequestOtp(post model.RequestUser) model.ResponseUser {
	query, err := m.Db.Query("SELECT * FROM users WHERE user_id = $1", post.UserId)
	if err != nil {
		log.Println(err)

		return model.ResponseUser{}
	}
	var result model.ResponseUser
	var count int

	for query.Next() {
		count++
	}

	t := fmt.Sprint(time.Now().Nanosecond())
	newOtp := t[:5]
	now := time.Now()

	if count == 0 {
		_, err := m.Db.Exec(
			"INSERT INTO users(user_id, otp, start_date_otp) VALUES ($1,$2,$3)",
			post.UserId,
			newOtp,
			now,
		)
		if err != nil {
			log.Println(err)

			return model.ResponseUser{}
		}

		result = model.ResponseUser{UserId: post.UserId, Otp: newOtp}
	} else {
		_, err := m.Db.Exec("UPDATE users SET otp = $1, start_date_otp = $2 WHERE user_id = $3",
			newOtp,
			now,
			post.UserId,
		)
		if err != nil {
			log.Println(err)

			return model.ResponseUser{}
		}

		result = model.ResponseUser{UserId: post.UserId, Otp: newOtp}
	}

	return result
}

func (m *UserRepository) CheckOtp(post model.RequestUser) model.ResponseUser {
	query, err := m.Db.Query("SELECT * FROM users WHERE user_id = $1 and otp = $2",
		post.UserId,
		post.Otp,
	)
	if err != nil {
		log.Println(err)

		return model.ResponseUser{}
	}

	var count int

	for query.Next() {
		count++
	}

	if count > 0 {
		return model.ResponseUser{UserId: post.UserId, Otp: post.Otp}
	} else {
		return model.ResponseUser{}
	}
}
