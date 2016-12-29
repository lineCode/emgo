package opamp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type OPAMP_Periph struct {
	CSR CSR
}

func (p *OPAMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var OPAMP = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP_BASE)))

var OPAMP1 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP1_BASE)))

var OPAMP2 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP2_BASE)))

var OPAMP3 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP3_BASE)))

var OPAMP4 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP4_BASE)))

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

func (p *OPAMP_Periph) OPAMPxEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPAMPxEN)}}
}

func (p *OPAMP_Periph) FORCEVP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(FORCEVP)}}
}

func (p *OPAMP_Periph) VPSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VPSEL)}}
}

func (p *OPAMP_Periph) VMSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VMSEL)}}
}

func (p *OPAMP_Periph) TCMEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(TCMEN)}}
}

func (p *OPAMP_Periph) VMSSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VMSSEL)}}
}

func (p *OPAMP_Periph) VPSSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VPSSEL)}}
}

func (p *OPAMP_Periph) CALON() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CALON)}}
}

func (p *OPAMP_Periph) CALSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CALSEL)}}
}

func (p *OPAMP_Periph) PGGAIN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PGGAIN)}}
}

func (p *OPAMP_Periph) USERTRIM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(USERTRIM)}}
}

func (p *OPAMP_Periph) TRIMOFFSETP() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(TRIMOFFSETP)}}
}

func (p *OPAMP_Periph) TRIMOFFSETN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(TRIMOFFSETN)}}
}

func (p *OPAMP_Periph) TSTREF() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(TSTREF)}}
}

func (p *OPAMP_Periph) OUTCAL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OUTCAL)}}
}

func (p *OPAMP_Periph) LOCK() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(LOCK)}}
}
