package main

import (
	"go-learn/day13/person"
	"unsafe"
)

//#include "./person/person_capi.h"
import "C"

//export person_new
func person_new(name *C.char, age C.int) C.person_handle_t {
	id := person.NewObjectId(person.NewPerson(C.GoString(name), int(age)))
	return C.person_handle_t(id)
}

//export person_delete
func person_delete(h C.person_handle_t) {
	id := person.ObjectId(h)
	id.Free()
}

//export person_set
func person_set(h C.person_handle_t, name *C.char, age C.int) {
	p := person.ObjectId(h).Get().(*person.Person)
	p.Set(C.GoString(name), int(age))
}

//export person_get_name
func person_get_name(h C.person_handle_t, buf *C.char, size C.int) *C.char {
	p := person.ObjectId(h).Get().(*person.Person)
	name, _ := p.Get()

	n := int(size) - 1
	bufSlice := ((*[1 << 31]byte)(unsafe.Pointer(buf)))[0:n:n]
	n = copy(bufSlice, []byte(name))
	bufSlice[n] = 0

	return buf
}

//export person_get_age
func person_get_age(h C.person_handle_t) C.int {
	p := person.ObjectId(h).Get().(*person.Person)
	_, age := p.Get()
	return C.int(age)
}

func main() {

}
