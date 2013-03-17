// bitvector - Simple BitVector for Go
// 
// Copyright (c) 2013 - Thiago Cardoso <thiagoncc@gmail.com>
// 
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met: 
// 
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer. 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution. 
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package bitvector

import "testing"

// Test cases for blockIndex and bitIndex function. It checks if the
// bit position in the data slice is correctly calculated.
var blockIndexCases = []struct {
    index int // Index of the bit
    block int // Expected block index
    pos uint // Expected position inside block
}{
    {0, 0, 0},
    {1, 0, 1},
    {_BITS_PER_BLOCK-1, 0, uint(_BITS_PER_BLOCK-1)},
    {_BITS_PER_BLOCK, 1, 0},
    {_BITS_PER_BLOCK+1, 1, 1},
    {2*_BITS_PER_BLOCK, 2, 0},
}

// TestIndexToBlock checks if the bit's position in the data slice is
// being correctly obtained from the bit index.
func TestIndexToBlock (t *testing.T) {
    for _, c := range blockIndexCases {
        b := blockIndex(c.index)
        if c.block != b {
            t.Errorf("wrong block. index= %d expected= %d actual= %d",
            c.index, c.block, b)
        }
        p := bitIndex(c.index)
        if c.pos != p {
            t.Errorf("wrong position. index= %d expected= %d actual= %d",
            c.index, c.pos, p)
        }
    }
}

// TestSimpleGetAndSet checks if bits are being correctly setted and
// retrieved.
func TestSimpleGetAndSet (t *testing.T) {
    bv := New(10)

    err := bv.Set(2, true)
    if err != nil { t.Errorf("set should be ok") }
    err = bv.Set(5, true)
    if err != nil { t.Errorf("set should be ok") }
    err = bv.Set(9, true)
    if err != nil { t.Errorf("set should be ok") }
    err = bv.Set(10, true)
    if err == nil { t.Errorf("set should not be ok") }

    v, err := bv.Get(0)
    if err != nil { t.Errorf("get should be ok") }
    if v { t.Errorf("wrong value for index 0") }
    v, err = bv.Get(1)
    if err != nil { t.Errorf("get should be ok") }
    if v { t.Errorf("wrong value for index 1") }
    v, err = bv.Get(2)
    if err != nil { t.Errorf("get should be ok") }
    if !v { t.Errorf("wrong value for index 2") }
    v, err = bv.Get(4)
    if err != nil { t.Errorf("get should be ok") }
    if v { t.Errorf("wrong value for index 4") }
    v, err = bv.Get(5)
    if err != nil { t.Errorf("get should be ok") }
    if !v { t.Errorf("wrong value for index 5") }
    v, err = bv.Get(9)
    if err != nil { t.Errorf("get should be ok") }
    if !v { t.Errorf("wrong value for index 9") }
    v, err = bv.Get(10)
    if err == nil { t.Errorf("get should not be ok") }

    err = bv.Set(2, false)
    if err != nil { t.Errorf("set should be ok") }
    v, err = bv.Get(2)
    if err != nil { t.Errorf("get should be ok") }
    if v { t.Errorf("wrong value for index 2") }
}

// TestOr checks if the OR operation between two BitVectors is being
// correctly calculated.
func TestOr (t *testing.T) {
    trueBits := []int { 10, 15, 20, 25, 30, 35, 64, 65 }
    falseBits := []int { 0, 1, 2, 3, 4, 11, 12, 24, 27, 34, 62, 63 }

    bv1 := New(66)
    bv1.Set(10, true)
    bv1.Set(20, true)
    bv1.Set(30, true)
    bv1.Set(64, true)
    bv1.Set(65, true)

    bv2 := New(66)
    bv2.Set(15, true)
    bv2.Set(25, true)
    bv2.Set(35, true)

    and, err := bv1.Or(bv2)
    if err != nil {
        t.Errorf("And should not return error")
    }

    for _, i := range trueBits {
        v, err := and.Get(i)
        if err != nil {
            t.Errorf("get for index %d should not fail", i)
        }
        if !v {
            t.Errorf("get for index %d should return true", i)
        }
    }
    for _, i := range falseBits {
        v, err := and.Get(i)
        if err != nil {
            t.Errorf("get for index %d should not fail", i)
        }
        if v {
            t.Errorf("get for index %d should return false", i)
        }
    }
}

// TestAnd checks if the And operation between two BitVectors is being
// correctly calculated.
func TestAnd (t *testing.T) {
    trueBits := []int { 10, 20, 30, 64, 65 }
    falseBits := []int { 0, 1, 2, 3, 4, 11, 12, 24, 27, 34, 62, 63 }

    bv1 := New(66)
    bv1.Set(10, true)
    bv1.Set(20, true)
    bv1.Set(30, true)
    bv1.Set(64, true)
    bv1.Set(65, true)

    bv1.Set(0, true)
    bv1.Set(1, true)
    bv1.Set(2, true)
    bv1.Set(3, true)
    bv1.Set(4, true)

    bv2 := New(66)
    bv2.Set(10, true)
    bv2.Set(20, true)
    bv2.Set(30, true)
    bv2.Set(64, true)
    bv2.Set(65, true)

    bv2.Set(15, true)
    bv2.Set(25, true)
    bv2.Set(35, true)

    and, err := bv1.And(bv2)
    if err != nil {
        t.Errorf("And should not return error")
    }

    for _, i := range trueBits {
        v, err := and.Get(i)
        if err != nil {
            t.Errorf("get for index %d should not fail", i)
        }
        if !v {
            t.Errorf("get for index %d should return true", i)
        }
    }
    for _, i := range falseBits {
        v, err := and.Get(i)
        if err != nil {
            t.Errorf("get for index %d should not fail", i)
        }
        if v {
            t.Errorf("get for index %d should return false", i)
        }
    }
}

