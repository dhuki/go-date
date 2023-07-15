package domain

import "errors"

var (
	KeyMismatchPassword = "failed.login.attempt"
	KeyTemporarySuspend = "failed.login.issuspend"

	ErrUserAlreadyExist   = errors.New("username sudah dipakai, mohon menggunakan username lain")
	ErrTooManyFailedLogin = errors.New("akun anda sementara terkunci, mohon tunggu 15 menit")
	ErrWrongPassword      = errors.New("username dan password salah, silahkan coba beberapa saat lagi")
)
