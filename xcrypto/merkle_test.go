// tokucore
//
// Copyright (c) 2018 TokuBlock
// BSD License

package xcrypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerkle(t *testing.T) {
	tests := []struct {
		name  string
		root  []byte
		hashs [][]byte
	}{
		{
			name:  "nil",
			root:  nil,
			hashs: [][]byte{nil},
		},
		{
			name:  "A",
			root:  []byte{'A'},
			hashs: [][]byte{{'A'}},
		},
		{
			name:  "AB",
			root:  []byte{0x5d, 0x60, 0xdc, 0x7f, 0xf4, 0xb5, 0xa3, 0x86, 0x1f, 0x98, 0x67, 0x19, 0xc6, 0x4, 0x6a, 0x66, 0xbb, 0xdb, 0x3a, 0x76, 0xe, 0x81, 0xe4, 0xe6, 0xfe, 0x8e, 0xc5, 0x9e, 0xd4, 0x58, 0x8c, 0xbe},
			hashs: [][]byte{{'A'}, {'B'}},
		},
		{
			name:  "CD",
			root:  []byte{0xdf, 0x13, 0xc2, 0xe, 0x5a, 0xbc, 0x89, 0xfd, 0x71, 0xd5, 0xdd, 0x16, 0x61, 0x1f, 0xf3, 0x30, 0xfe, 0xbb, 0x1, 0xf3, 0x6d, 0x78, 0x8e, 0x5c, 0xfe, 0xd7, 0x4d, 0xde, 0xbc, 0xab, 0x86, 0x4},
			hashs: [][]byte{{'C'}, {'D'}},
		},
		{
			name: "2",
			root: []byte{0x42, 0x2d, 0x90, 0x49, 0x41, 0xb2, 0xcf, 0x4e, 0x59, 0x2c, 0x7d, 0x81, 0x5c, 0xe5, 0xa2, 0xca, 0x52, 0xd1, 0x98, 0x9e, 0xc2, 0x94, 0xd0, 0xe, 0x16, 0x48, 0xcd, 0x81, 0x62, 0x4c, 0x2c, 0xfb},
			hashs: [][]byte{
				{0x5d, 0x60, 0xdc, 0x7f, 0xf4, 0xb5, 0xa3, 0x86, 0x1f, 0x98, 0x67, 0x19, 0xc6, 0x4, 0x6a, 0x66, 0xbb, 0xdb, 0x3a, 0x76, 0xe, 0x81, 0xe4, 0xe6, 0xfe, 0x8e, 0xc5, 0x9e, 0xd4, 0x58, 0x8c, 0xbe},
				{0xdf, 0x13, 0xc2, 0xe, 0x5a, 0xbc, 0x89, 0xfd, 0x71, 0xd5, 0xdd, 0x16, 0x61, 0x1f, 0xf3, 0x30, 0xfe, 0xbb, 0x1, 0xf3, 0x6d, 0x78, 0x8e, 0x5c, 0xfe, 0xd7, 0x4d, 0xde, 0xbc, 0xab, 0x86, 0x4},
			},
		},
		{
			name:  "ABC",
			root:  []byte{0x3a, 0xe5, 0x59, 0x9f, 0xda, 0x6a, 0xd7, 0x6b, 0x74, 0xc2, 0xee, 0x34, 0x1, 0x3a, 0x18, 0x62, 0xae, 0xef, 0x64, 0x51, 0xe2, 0x3c, 0xe5, 0x8e, 0xd5, 0xa4, 0x9a, 0xba, 0xc8, 0x28, 0xfd, 0x35},
			hashs: [][]byte{{'A'}, {'B'}, {'C'}},
		},
		{
			name:  "ABCD",
			root:  []byte{0x42, 0x2d, 0x90, 0x49, 0x41, 0xb2, 0xcf, 0x4e, 0x59, 0x2c, 0x7d, 0x81, 0x5c, 0xe5, 0xa2, 0xca, 0x52, 0xd1, 0x98, 0x9e, 0xc2, 0x94, 0xd0, 0xe, 0x16, 0x48, 0xcd, 0x81, 0x62, 0x4c, 0x2c, 0xfb},
			hashs: [][]byte{{'A'}, {'B'}, {'C'}, {'D'}},
		},
	}

	for _, test := range tests {
		merkle := NewMerkle(test.hashs)
		assert.Equal(t, test.root, merkle.Hashs[len(merkle.Hashs)-1])
	}
}