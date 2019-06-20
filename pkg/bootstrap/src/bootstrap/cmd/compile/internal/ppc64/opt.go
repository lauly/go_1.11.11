// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/ubuntu/go_1.11.11/go/src/cmd/compile/internal/ppc64/opt.go

//line /home/ubuntu/go_1.11.11/go/src/cmd/compile/internal/ppc64/opt.go:1
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

// Many Power ISA arithmetic and logical instructions come in four
// standard variants. These bits let us map between variants.
const (
	V_CC = 1 << 0 // xCC (affect CR field 0 flags)
	V_V  = 1 << 1 // xV (affect SO and OV flags)
)
