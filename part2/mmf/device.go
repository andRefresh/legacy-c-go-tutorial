package mmf

import "C"

// Also here a type alias
type Device = C.struct_Device

// C logic should be hidden in this library
func (d *Device) SetOn(on bool) {
	if on {
		d.IsActive = 1
	} else {
		d.IsActive = 0
	}
}

func (d *Device) ResetValue() {
	d.SomeValue = 0
}

func (d *Device) SetValue(someValue int) {
	d.SomeValue = C.int(someValue)
}
