package repository

import "sinarmas/model"

type UserRepositoryInterface interface {
	RequestOtp(model.RequestUser) model.ResponseUser
	CheckOtp(model.RequestUser) model.ResponseUser
}
