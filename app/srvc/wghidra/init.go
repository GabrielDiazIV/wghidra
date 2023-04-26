package wghidra

import (
	"github.com/gabrieldiaziv/wghidra/app/bo/defs"
	"github.com/gabrieldiaziv/wghidra/app/bo/iface"
)

type wghidra struct {
	dokr  iface.Dokr
	store iface.Store
}

func NewWGhidra(dokr iface.Dokr, store iface.Store) defs.WGhidra {
	return &wghidra{
		dokr:  dokr,
		store: store,
	}
}
