package noos

import (
	"internal"
	"syscall"
	"unsafe"
)

// ChanT is pseudo channel (synchronous, read only) that allows to receive once
// at or after time stored in task local variable. This is dirty hack but
// allows to implement deadlines in select statements without createing
// additional gorutines.
type chanT struct {
	event syscall.Event // Must be syscall.Alarm
}

func (c *chanT) TryRecv(_, _ unsafe.Pointer) (unsafe.Pointer, uintptr) {
	p := &tasker.tasks[tasker.curTask].at
	if *p == 0 {
		return nil, cclosed
	}
	if ns := syscall.Nanosec(); ns >= *p {
		*p = ns
		return unsafe.Pointer(p), 0
	}
	syscall.SetAlarm(*p)
	return nil, cagain
}

func (c *chanT) Recv(_ unsafe.Pointer) (p unsafe.Pointer, d uintptr) {
	for {
		p, d = c.TryRecv(nil, nil)
		if p != nil || d == cclosed {
			return
		}
		c.event.Wait()
	}
}

func (c *chanT) Done(_ uintptr) {
	tasker.tasks[tasker.curTask].at = 0
}

func (c *chanT) Len() int {
	return 0
}

func (c *chanT) Cap() int {
	return 0
}

var chanTMethods = struct {
	Send       func(c *chanT, e unsafe.Pointer) (p unsafe.Pointer, d uintptr)
	Recv       func(c *chanT, e unsafe.Pointer) (p unsafe.Pointer, d uintptr)
	TrySend    func(c *chanT, e, _ unsafe.Pointer) (p unsafe.Pointer, d uintptr)
	TryRecv    func(c *chanT, e, _ unsafe.Pointer) (p unsafe.Pointer, d uintptr)
	CancelSend func()
	CancelRecv func()
	Done       func(c *chanT, d uintptr)
	Close      func(c *chanT)
	Len        func(c *chanT) int
	Cap        func(c *chanT) int
}{
	Recv:    (*chanT).Recv,
	TryRecv: (*chanT).TryRecv,
	Done:    (*chanT).Done,
	Len:     (*chanT).Len,
	Cap:     (*chanT).Cap,
}

var timeChan = chanT{syscall.Alarm}

func makeTimeChan() internal.Chan {
	return internal.Chan{
		C: unsafe.Pointer(&timeChan),
		M: (*internal.ChanMethods)(unsafe.Pointer(&chanTMethods)),
	}
}
