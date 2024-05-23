package sflow

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSFlowDecode(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0xac, 0x10, 0x00, 0x11, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x01, 0xaa, 0x67, 0xee, 0xaa, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x04, 0x13, 0x00, 0x00, 0x08, 0x00,
		0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xaa, 0x00, 0x00, 0x04, 0x13,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x52, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x4e, 0x00, 0xff, 0x12, 0x34,
		0x35, 0x1b, 0xff, 0xab, 0xcd, 0xef, 0xab, 0x64, 0x81, 0x00, 0x00, 0x20, 0x08, 0x00, 0x45, 0x00,
		0x00, 0x3c, 0x5c, 0x07, 0x00, 0x00, 0x7c, 0x01, 0x48, 0xa0, 0xac, 0x10, 0x20, 0xfe, 0xac, 0x10,
		0x20, 0xf1, 0x08, 0x00, 0x97, 0x61, 0xa9, 0x48, 0x0c, 0xb2, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66,
		0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76,
		0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x00, 0x00,
	}
	buf := bytes.NewBuffer(data)
	var packet Packet
	assert.Nil(t, DecodeMessageVersion(buf, &packet))
}

func TestExpandedSFlowDecode(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x0f, 0xa7, 0x72, 0xc2, 0x0f, 0x76, 0x73, 0x48, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0xdc, 0x20, 0x90, 0x93, 0x26, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa4,
		0x00, 0x00, 0x3f, 0xff, 0x04, 0x38, 0xec, 0xda, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0x52, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xea, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80,
		0x08, 0xec, 0xf5, 0x2a, 0x8f, 0xbe, 0x74, 0x83, 0xef, 0x30, 0x65, 0xb7, 0x81, 0x00, 0x00, 0x1e,
		0x08, 0x00, 0x45, 0x00, 0x05, 0xd4, 0x3b, 0xba, 0x40, 0x00, 0x3f, 0x06, 0xbd, 0x99, 0xb9, 0x3b,
		0xdc, 0x93, 0x58, 0xee, 0x4e, 0x13, 0x01, 0xbb, 0xcf, 0xd6, 0x45, 0xb7, 0x1b, 0xc0, 0xd5, 0xb8,
		0xff, 0x24, 0x80, 0x10, 0x00, 0x04, 0x01, 0x55, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xc8, 0xc8,
		0x56, 0x95, 0x00, 0x34, 0xf6, 0x0f, 0xe8, 0x1d, 0xbd, 0x41, 0x45, 0x92, 0x4c, 0xc2, 0x71, 0xe0,
		0xeb, 0x2e, 0x35, 0x17, 0x7c, 0x2f, 0xb9, 0xa8, 0x05, 0x92, 0x0e, 0x03, 0x1b, 0x50, 0x53, 0x0c,
		0xe5, 0x7d, 0x86, 0x75, 0x32, 0x8a, 0xcc, 0xe2, 0x26, 0xa8, 0x90, 0x21, 0x78, 0xbf, 0xce, 0x7a,
		0xf8, 0xb5, 0x8d, 0x48, 0xe4, 0xaa, 0xfe, 0x26, 0x34, 0xe0, 0xad, 0xb9, 0xec, 0x79, 0x74, 0xd8,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0xdc, 0x20, 0x90, 0x93, 0x27, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x3f, 0xff, 0x04, 0x39, 0x2c, 0xd9, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0x4b,
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x17,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xca, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x80, 0xda, 0xb1, 0x22, 0xfb, 0xd9, 0xcf, 0x74, 0x83, 0xef, 0x30, 0x65, 0xb7,
		0x81, 0x00, 0x00, 0x17, 0x08, 0x00, 0x45, 0x00, 0x05, 0xb4, 0xe2, 0x28, 0x40, 0x00, 0x3f, 0x06,
		0x15, 0x0f, 0xc3, 0xb5, 0xaf, 0x26, 0x05, 0x92, 0xc6, 0x9e, 0x00, 0x50, 0x0f, 0xb3, 0x35, 0x8e,
		0x36, 0x02, 0xa1, 0x01, 0xed, 0xb0, 0x80, 0x10, 0x00, 0x3b, 0xf7, 0xd4, 0x00, 0x00, 0x01, 0x01,
		0x08, 0x0a, 0xd2, 0xe8, 0xac, 0xbe, 0x00, 0x36, 0xbc, 0x3c, 0x37, 0x36, 0xc4, 0x80, 0x3f, 0x66,
		0x33, 0xc5, 0x50, 0xa6, 0x63, 0xb2, 0x92, 0xc3, 0x6a, 0x7a, 0x80, 0x65, 0x0b, 0x22, 0x62, 0xfe,
		0x16, 0x9c, 0xab, 0x55, 0x03, 0x47, 0xa6, 0x54, 0x63, 0xa5, 0xbc, 0x17, 0x8e, 0x5a, 0xf6, 0xbc,
		0x24, 0x52, 0xe9, 0xd2, 0x7b, 0x08, 0xe8, 0xc2, 0x6b, 0x05, 0x1c, 0xc0, 0x61, 0xb4, 0xe0, 0x43,
		0x59, 0x62, 0xbf, 0x0a, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0xdc, 0x04, 0x12, 0xa0, 0x65,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa8, 0x00, 0x00, 0x3f, 0xff, 0xa4, 0x06, 0x9f, 0x9b,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa8, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10,
		0x00, 0x00, 0x05, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x39, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xf2,
		0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0x74, 0x83, 0xef, 0x30, 0x65, 0xb7, 0x28, 0x99,
		0x3a, 0x4e, 0x89, 0x27, 0x81, 0x00, 0x05, 0x39, 0x08, 0x00, 0x45, 0x18, 0x05, 0xdc, 0x8e, 0x5c,
		0x40, 0x00, 0x3a, 0x06, 0x53, 0x77, 0x89, 0x4a, 0xcc, 0xd5, 0x59, 0xbb, 0xa9, 0x55, 0x07, 0x8f,
		0xad, 0xdc, 0xf2, 0x9b, 0x09, 0xb4, 0xce, 0x1d, 0xbc, 0xee, 0x80, 0x10, 0x75, 0x40, 0x58, 0x02,
		0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0xb0, 0x18, 0x5b, 0x6f, 0xd7, 0xd6, 0x8b, 0x47, 0xee, 0x6a,
		0x03, 0x0b, 0x9b, 0x52, 0xb1, 0xca, 0x61, 0x4b, 0x84, 0x57, 0x75, 0xc4, 0xb2, 0x18, 0x11, 0x39,
		0xce, 0x5d, 0x2a, 0x38, 0x91, 0x29, 0x76, 0x11, 0x7d, 0xc1, 0xcc, 0x5c, 0x4b, 0x0a, 0xde, 0xbb,
		0xa8, 0xad, 0x9d, 0x88, 0x36, 0x8b, 0xc0, 0x02, 0x87, 0xa7, 0xa5, 0x1c, 0xd9, 0x85, 0x71, 0x85,
		0x68, 0x2b, 0x59, 0xc6, 0x2c, 0x3c, 0x84, 0x0c, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0xdc,
		0x20, 0x90, 0x93, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x3f, 0xff,
		0x04, 0x39, 0x6c, 0xd8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa4,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0x4b, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x03, 0xe9,
		0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x17,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x05, 0xf2, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80, 0xda, 0xb1, 0x22, 0xfb,
		0xd9, 0xcf, 0x74, 0x83, 0xef, 0x30, 0x65, 0xb7, 0x81, 0x00, 0x00, 0x17, 0x08, 0x00, 0x45, 0x00,
		0x05, 0xdc, 0x7e, 0x42, 0x40, 0x00, 0x3f, 0x06, 0x12, 0x4d, 0xb9, 0x66, 0xdb, 0x43, 0x67, 0xc2,
		0xa9, 0x20, 0x63, 0x75, 0x57, 0xae, 0x6d, 0xbf, 0x59, 0x7c, 0x93, 0x71, 0x09, 0x67, 0x80, 0x10,
		0x00, 0xeb, 0xfc, 0x16, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x40, 0x96, 0x88, 0x38, 0x36, 0xe1,
		0x64, 0xc7, 0x1b, 0x43, 0xbc, 0x0e, 0x1f, 0x81, 0x6d, 0x39, 0xf6, 0x12, 0x0c, 0xea, 0xc0, 0xea,
		0x7b, 0xc1, 0x77, 0xe2, 0x92, 0x6a, 0xbf, 0xbe, 0x84, 0xd9, 0x00, 0x18, 0x57, 0x49, 0x92, 0x72,
		0x8f, 0xa3, 0x78, 0x45, 0x6f, 0xc6, 0x98, 0x8f, 0x71, 0xb0, 0xc5, 0x52, 0x7d, 0x8a, 0x82, 0xef,
		0x52, 0xdb, 0xe9, 0xdc, 0x0a, 0x52, 0xdb, 0x06, 0x51, 0x80, 0x80, 0xa9, 0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0xdc, 0x20, 0x90, 0x93, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa4,
		0x00, 0x00, 0x3f, 0xff, 0x04, 0x39, 0xac, 0xd7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0f, 0x42, 0xa4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x42, 0xa5, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x03, 0xe9, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x03, 0xbd, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x03, 0xbd, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x90,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x05, 0xf2, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x80,
		0x90, 0xe2, 0xba, 0x89, 0x21, 0xad, 0x74, 0x83, 0xef, 0x30, 0x65, 0xb7, 0x81, 0x00, 0x03, 0xbd,
		0x08, 0x00, 0x45, 0x00, 0x05, 0xdc, 0x76, 0xa2, 0x40, 0x00, 0x38, 0x06, 0xac, 0x75, 0x33, 0x5b,
		0x74, 0x6c, 0xc3, 0xb5, 0xae, 0x87, 0x1f, 0x40, 0x80, 0x68, 0xab, 0xbb, 0x2f, 0x90, 0x01, 0xee,
		0x3a, 0xaf, 0x80, 0x10, 0x00, 0xeb, 0x8e, 0xf4, 0x00, 0x00, 0x01, 0x01, 0x08, 0x0a, 0x34, 0xc0,
		0xff, 0x26, 0xac, 0x90, 0xd5, 0xc4, 0xcc, 0xd7, 0xa4, 0xa5, 0x5b, 0xa3, 0x79, 0x33, 0xc1, 0x25,
		0xcd, 0x84, 0xdc, 0xaa, 0x37, 0xc9, 0xe3, 0xab, 0xc6, 0xb4, 0xeb, 0xe3, 0x8d, 0x72, 0x06, 0xd1,
		0x5a, 0x1f, 0x9a, 0x8b, 0xe9, 0x9a, 0xf7, 0x33, 0x35, 0xe5, 0xca, 0x67, 0xba, 0x04, 0xf9, 0x3c,
		0x27, 0xff, 0xa3, 0xca, 0x5e, 0x90, 0xf9, 0xc7, 0xd1, 0xe4, 0xf8, 0xf5, 0x7a, 0x14, 0xdc, 0x1c,
		0xb1, 0xde, 0x63, 0x75, 0xb2, 0x65, 0x27, 0xf0, 0x0d, 0x29, 0xc5, 0x56, 0x60, 0x4a, 0x50, 0x10,
		0x00, 0x77, 0xc0, 0xef, 0x00, 0x00, 0x74, 0xcf, 0x8a, 0x79, 0x87, 0x77, 0x75, 0x64, 0x75, 0xeb,
		0xa4, 0x56, 0xb4, 0xd8, 0x70, 0xca, 0xe6, 0x11, 0xbb, 0x9f, 0xa1, 0x63, 0x95, 0xa1, 0xb4, 0x81,
		0x8d, 0x50, 0xe0, 0xd5, 0xa9, 0x2c, 0xd7, 0x8f, 0xfe, 0x78, 0xce, 0xff, 0x5a, 0xa6, 0xb6, 0xb9,
		0xf1, 0xe9, 0x5f, 0xda, 0xcb, 0xf3, 0x62, 0x61, 0x5f, 0x2b, 0x32, 0x95, 0x5d, 0x96, 0x2e, 0xef,
		0x32, 0x04, 0xff, 0xcc, 0x76, 0xba, 0x49, 0xab, 0x92, 0xa7, 0xf1, 0xcc, 0x52, 0x68, 0xde, 0x94,
		0x90, 0xdb, 0x1b, 0xa0, 0x28, 0x8a, 0xf8, 0x64, 0x55, 0x9c, 0x9b, 0xf6, 0x9c, 0x44, 0xd9, 0x68,
		0xc0, 0xe5, 0x2c, 0xe1, 0x3d, 0x29, 0x19, 0xef, 0x8b, 0x0c, 0x9d, 0x0a, 0x7e, 0xcd, 0xc2, 0xe9,
		0x85, 0x6b, 0x85, 0xb3, 0x97, 0xbe, 0xc6, 0x26, 0xd2, 0xe5, 0x2e, 0x90, 0xa9, 0xac, 0xe3, 0xd8,
		0xef, 0xbd, 0x7b, 0x40, 0xf8, 0xb7, 0xe3, 0xc3, 0x8d, 0xa7, 0x38, 0x0f, 0x87, 0x7a, 0x50, 0x62,
		0xc8, 0xb8, 0xa4, 0x52, 0x6e, 0xdc, 0x92, 0x7f, 0xe6, 0x8d, 0x45, 0x39, 0xfd, 0x06, 0x6e, 0xd9,
		0xb5, 0x65, 0xac, 0xae, 0x2b, 0x8d, 0xea, 0xcf, 0xa2, 0x98, 0x0b, 0xc6, 0x43, 0x2e, 0xa7, 0x71,
		0x99, 0x2b, 0xea, 0xc3, 0x9c, 0x27, 0x74, 0x9e, 0xd5, 0x11, 0x60, 0x7a, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x7b, 0xd6, 0x2a, 0x39, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	buf := bytes.NewBuffer(data)
	var packet Packet
	assert.Nil(t, DecodeMessageVersion(buf, &packet))
}

func TestSFlowDecodeDropEgressQueue(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0xc0, 0xa8, 0x77, 0xb8, 0x00, 0x01, 0x86, 0xa0,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x30, 0x7e, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x0c, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x2a,
	}

	buf := bytes.NewBuffer(data)
	var packet Packet
	assert.NoError(t, DecodeMessageVersion(buf, &packet))
	assert.Len(t, packet.Samples, 1)
	assert.NotNil(t, packet.Samples[0])
	sample, ok := packet.Samples[0].(DropSample)
	assert.True(t, ok)
	assert.Len(t, sample.Records, 1)
	assert.Equal(t, EgressQueue{Queue: 42}, sample.Records[0].Data)
}

func TestSFlowDecodeDropExtendedACL(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0xc0, 0xa8, 0x77, 0xb8, 0x00, 0x01, 0x86, 0xa0,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x30, 0x7e, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x38, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x0d, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x2a,
		0x00, 0x00, 0x00, 0x04, 0x66, 0x6f, 0x6f, 0x21, 0x00, 0x00, 0x00, 0x02,
	}

	buf := bytes.NewBuffer(data)
	var packet Packet
	assert.NoError(t, DecodeMessageVersion(buf, &packet))
	assert.Len(t, packet.Samples, 1)
	assert.NotNil(t, packet.Samples[0])
	sample, ok := packet.Samples[0].(DropSample)
	assert.True(t, ok)
	assert.Len(t, sample.Records, 1)
	assert.Equal(t, ExtendedACL{Number: 42, Name: "foo!", Direction: 2}, sample.Records[0].Data)
}

func TestSFlowDecodeDropExtendedFunction(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0xc0, 0xa8, 0x77, 0xb8, 0x00, 0x01, 0x86, 0xa0,
		0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x30, 0x7e, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x32, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x04, 0x0e, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x06,
		0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72,
	}

	buf := bytes.NewBuffer(data)
	var packet Packet
	assert.NoError(t, DecodeMessageVersion(buf, &packet))
	assert.Len(t, packet.Samples, 1)
	assert.NotNil(t, packet.Samples[0])
	sample, ok := packet.Samples[0].(DropSample)
	assert.True(t, ok)
	assert.Len(t, sample.Records, 1)
	assert.Equal(t, ExtendedFunction{Symbol: "foobar"}, sample.Records[0].Data)
}
