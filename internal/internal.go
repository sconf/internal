// Copyright 2015 The Sconf Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal // import "gopkg.in/sconf/internal.v0/internal"

import (
	"gopkg.in/sconf/internal.v0/internal/gcfg"
)

type Struct struct {
	Ptr interface{}
}

func (s Struct) Read(readers ...IdemReader) error {
	for _, r := range readers {
		b, err := r.Bytes()
		if err != nil {
			return err
		}
		err = gcfg.ReadStringInto(s.Ptr, string(b))
		if err != nil {
			return err
		}
	}
	return nil
}

type IdemReader interface {
	Bytes() ([]byte, error)
}

type ErrIdemReader struct {
	Err error
}

func (e ErrIdemReader) Bytes() ([]byte, error) {
	return nil, e.Err
}

type BytesIdemReader []byte

func (b BytesIdemReader) Bytes() ([]byte, error) {
	return b, nil
}
