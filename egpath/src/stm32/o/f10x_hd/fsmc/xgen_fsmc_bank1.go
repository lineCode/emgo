package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank1_Periph struct {
	BTCR [8]BTCR
}

var FSMC_Bank1 = (*FSMC_Bank1_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank1_R_BASE)))

type BTCR_Bits uint32

func (b BTCR_Bits) Field(mask BTCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BTCR_Bits) J(v int) BTCR_Bits {
	return BTCR_Bits(bits.Make32(v, uint32(mask)))
}

type BTCR struct{ mmio.U32 }

func (r *BTCR) Bits(mask BTCR_Bits) BTCR_Bits { return BTCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BTCR) StoreBits(mask, b BTCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BTCR) SetBits(mask BTCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BTCR) ClearBits(mask BTCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BTCR) Load() BTCR_Bits               { return BTCR_Bits(r.U32.Load()) }
func (r *BTCR) Store(b BTCR_Bits)             { r.U32.Store(uint32(b)) }

type BTCR_Mask struct{ mmio.UM32 }

func (rm BTCR_Mask) Load() BTCR_Bits   { return BTCR_Bits(rm.UM32.Load()) }
func (rm BTCR_Mask) Store(b BTCR_Bits) { rm.UM32.Store(uint32(b)) }