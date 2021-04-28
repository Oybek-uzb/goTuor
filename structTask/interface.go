package main

import (

)

type PersonI interface {
	Add(plist *Person)
	Update(id uint64, plist *Person) error
	Delete(id uint64) error
	Get(id uint64) *Person
	Print()
}
