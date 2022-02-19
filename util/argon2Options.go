package util

import "github.com/marlonmp/argon2"

var Argon2Options = argon2.Options{
	Algorithm:   &argon2.Argon2ID{},
	Memory:      16,
	Iterations:  3,
	Parallelism: 3,
	HashLength:  32,
	SaltLength:  32,
}
