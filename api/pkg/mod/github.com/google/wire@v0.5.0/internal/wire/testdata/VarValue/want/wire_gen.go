// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func injectedMessage() string {
	string2 := _wireStringValue
	return string2
}

var (
	_wireStringValue = msg
)