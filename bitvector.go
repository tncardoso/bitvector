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

import(
    "errors"
)

// Type that should be used for each block.
type Block int64

// Number of bits in each block
var _BITS_PER_BLOCK = 64

// BitVector data structure. Contains the current size and the bits in
// a data slice.
type BitVector struct {
    size    int     // Current size of the BitVector.
    data    []Block // Stored bits.
}

// Returns the position of the block in the data slice for the bit
// with index 'index'.
func blockIndex (index int) int {
    return index / _BITS_PER_BLOCK
}

// Returns the position inside the block for the bit with index
// 'index'.
func bitIndex (index int) uint {
    return uint(index % _BITS_PER_BLOCK)
}

// Create and initialize a new BitVector with given initial size.
func New (size int) *BitVector {
    bv := new(BitVector)
    bv.size = size
    ds := blockIndex(size) + 1
    bv.data = make([]Block, ds, ds)
    return bv
}

// Get bit value of given position.
func (bv *BitVector) Get (index int) (bool, error) {
    if index >= bv.size {
        // Bit outside range.
        return false, errors.New("Bit outside BitVector range.")
    }

    blockI := blockIndex(index)
    block := bv.data[blockI]

    bitI := bitIndex(index)
    mask := Block(1 << bitI)

    return (block & mask) > 0, nil
}

// Set bit value of given position.
func (bv *BitVector) Set (index int, value bool) (error) {
    if index >= bv.size {
        // Bit outside range.
        return errors.New("Bit outside BitVector range.")
    }


    blockI := blockIndex(index)
    block := bv.data[blockI]

    bitI := bitIndex(index)
    mask := Block(1 << bitI)

    if value {
        bv.data[blockI] = block | mask
    } else {
        bv.data[blockI] = block & ^mask
    }

    return nil
}

// Bitwise AND of two BitVectors.
func (bv *BitVector) And (other *BitVector) (*BitVector, error) {
    if bv.size != other.size {
        // BitVectors have different number of bits.
        return nil, errors.New("Can't And BitVectors with different sizes")
    }

    res := New(bv.size)
    for i, _ := range bv.data {
        res.data[i] = bv.data[i] & other.data[i]
    }

    return res, nil
}

// Bitwise OR of two BitVectors.
func (bv *BitVector) Or (other *BitVector) (*BitVector, error) {
    if bv.size != other.size {
        // BitVectors have different number of bits.
        return nil, errors.New("Can't Or BitVectors with different sizes")
    }

    res := New(bv.size)
    for i, _ := range bv.data {
        res.data[i] = bv.data[i] | other.data[i]
    }

    return res, nil
}

