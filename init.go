package main

import (
	"encoding/gob"
	. "github.com/mokyabun/go-serialization/structs"
)

func init() {
	// single
	gob.Register(Tiny{})
	gob.Register(Medium{})
	gob.Register(Big{})
	gob.Register(Huge{})
	// slices
	gob.Register([]Tiny{})
	gob.Register([]Medium{})
	gob.Register([]Big{})
	gob.Register([]Huge{})
	// slices of pointers
	gob.Register([]*Tiny{})
	gob.Register([]*Medium{})
	gob.Register([]*Big{})
	gob.Register([]*Huge{})
}
