// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,386 darwin,amd64

package runtime

import "unsafe"

// The *_trampoline functions convert from the Go calling convention to the C calling convention
// and then call the underlying libc function.  They are defined in sys_darwin_$ARCH.s.

//go:nosplit
//go:cgo_unsafe_args
func pthread_attr_init(attr *pthreadattr) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(pthread_attr_init_trampoline)), unsafe.Pointer(&attr))
}
func pthread_attr_init_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func pthread_attr_setstacksize(attr *pthreadattr, size uintptr) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(pthread_attr_setstacksize_trampoline)), unsafe.Pointer(&attr))
}
func pthread_attr_setstacksize_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func pthread_attr_setdetachstate(attr *pthreadattr, state int) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(pthread_attr_setdetachstate_trampoline)), unsafe.Pointer(&attr))
}
func pthread_attr_setdetachstate_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func pthread_create(attr *pthreadattr, start uintptr, arg unsafe.Pointer) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(pthread_create_trampoline)), unsafe.Pointer(&attr))
}
func pthread_create_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func pthread_kill(thread pthread, sig int) (errno int32) {
	return asmcgocall(unsafe.Pointer(funcPC(pthread_kill_trampoline)), unsafe.Pointer(&thread))
}
func pthread_kill_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func pthread_self() (t pthread) {
	asmcgocall(unsafe.Pointer(funcPC(pthread_self_trampoline)), unsafe.Pointer(&t))
	return
}
func pthread_self_trampoline()

func mmap(addr unsafe.Pointer, n uintptr, prot, flags, fd int32, off uint32) (unsafe.Pointer, int) {
	args := struct {
		addr            unsafe.Pointer
		n               uintptr
		prot, flags, fd int32
		off             uint32
		ret1            unsafe.Pointer
		ret2            int
	}{addr, n, prot, flags, fd, off, nil, 0}
	asmcgocall(unsafe.Pointer(funcPC(mmap_trampoline)), unsafe.Pointer(&args))
	return args.ret1, args.ret2
}
func mmap_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func munmap(addr unsafe.Pointer, n uintptr) {
	asmcgocall(unsafe.Pointer(funcPC(munmap_trampoline)), unsafe.Pointer(&addr))
}
func munmap_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func madvise(addr unsafe.Pointer, n uintptr, flags int32) {
	asmcgocall(unsafe.Pointer(funcPC(madvise_trampoline)), unsafe.Pointer(&addr))
}
func madvise_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func read(fd int32, p unsafe.Pointer, n int32) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(read_trampoline)), unsafe.Pointer(&fd))
}
func read_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func closefd(fd int32) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(close_trampoline)), unsafe.Pointer(&fd))
}
func close_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func exit(code int32) {
	asmcgocall(unsafe.Pointer(funcPC(exit_trampoline)), unsafe.Pointer(&code))
}
func exit_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func usleep(usec uint32) {
	asmcgocall(unsafe.Pointer(funcPC(usleep_trampoline)), unsafe.Pointer(&usec))
}
func usleep_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func write(fd uintptr, p unsafe.Pointer, n int32) int32 {
	return asmcgocall(unsafe.Pointer(funcPC(write_trampoline)), unsafe.Pointer(&fd))
}
func write_trampoline()

//go:nosplit
//go:cgo_unsafe_args
func open(name *byte, mode, perm int32) (ret int32) {
	return asmcgocall(unsafe.Pointer(funcPC(open_trampoline)), unsafe.Pointer(&name))
}
func open_trampoline()

// Not used on Darwin, but must be defined.
func exitThread(wait *uint32) {
}

// Tell the linker that the libc_* functions are to be found
// in a system library, with the libc_ prefix missing.

//go:cgo_import_dynamic libc_pthread_attr_init pthread_attr_init "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_pthread_attr_setstacksize pthread_attr_setstacksize "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_pthread_attr_setdetachstate pthread_attr_setdetachstate "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_pthread_create pthread_create "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_exit exit "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_pthread_kill pthread_kill "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_pthread_self pthread_self "/usr/lib/libSystem.B.dylib"

//go:cgo_import_dynamic libc_open open "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_close close "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_read read "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_write write "/usr/lib/libSystem.B.dylib"

//go:cgo_import_dynamic libc_mmap mmap "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_munmap munmap "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_madvise madvise "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_error __error "/usr/lib/libSystem.B.dylib"
//go:cgo_import_dynamic libc_usleep usleep "/usr/lib/libSystem.B.dylib"

// Magic incantation to get libSystem actually dynamically linked.
// TODO: Why does the code require this?  See cmd/compile/internal/ld/go.go:210
//go:cgo_import_dynamic _ _ "/usr/lib/libSystem.B.dylib"
