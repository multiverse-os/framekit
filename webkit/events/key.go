package events

import (
	gdk "github.com/mattn/go-gtk/gdk"
)

const (
	NONE Modifier = iota
	CTRL
	FN
	HYPER
	META
	SUPER
)

type Modifier int

type KeyPressEvent struct {
	KeyVal   int
	Modifier gdk.ModifierType
}

func (kpe KeyPressEvent) GetKeyValue() int { return kpe.KeyVal }

func (kpe KeyPressEvent) GetModifier() Modifier {
	mod := kpe.Modifier
	switch {
	case mod&gdk.CONTROL_MASK != 0:
		return CTRL
	}
	return NONE
}

// TODO: Why?
func (kpe KeyPressEvent) Equals(k2 KeyPressEvent) bool {
	return kpe.GetKeyValue() == k2.GetKeyValue() && kpe.GetModifier() == k2.GetModifier()
}
