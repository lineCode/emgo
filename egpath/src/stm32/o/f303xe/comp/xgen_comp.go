package comp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type COMP_Periph struct {
	CSR CSR
}

func (p *COMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var COMP = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP_BASE)))

var COMP1 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP1_BASE)))

var COMP2 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP2_BASE)))

var COMP3 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP3_BASE)))

var COMP4 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP4_BASE)))

var COMP5 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP5_BASE)))

var COMP6 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP6_BASE)))

var COMP7 = (*COMP_Periph)(unsafe.Pointer(uintptr(mmap.COMP7_BASE)))

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *COMP_Periph) COMPxEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxEN)}}
}

func (p *COMP_Periph) COMP1SW1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMP1SW1)}}
}

func (p *COMP_Periph) COMPxMODE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxMODE)}}
}

func (p *COMP_Periph) COMPxINSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxINSEL)}}
}

func (p *COMP_Periph) COMPxNONINSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxNONINSEL)}}
}

func (p *COMP_Periph) COMPxWNDWEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxWNDWEN)}}
}

func (p *COMP_Periph) COMPxOUTSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxOUTSEL)}}
}

func (p *COMP_Periph) COMPxPOL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxPOL)}}
}

func (p *COMP_Periph) COMPxHYST() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxHYST)}}
}

func (p *COMP_Periph) COMPxBLANKING() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxBLANKING)}}
}

func (p *COMP_Periph) COMPxBLANKING_2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxBLANKING_2)}}
}

func (p *COMP_Periph) COMPxINSEL_3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxINSEL_3)}}
}

func (p *COMP_Periph) COMPxOUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxOUT)}}
}

func (p *COMP_Periph) COMPxLOCK() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(COMPxLOCK)}}
}
