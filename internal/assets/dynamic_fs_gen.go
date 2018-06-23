// Code generated - DO NOT EDIT.

package assets

import (
	"time"

	"github.com/richardwilkes/toolbox/xio/fs/embedded"
)

// dynamicFS holds an embedded filesystem.
var dynamicFS = embedded.NewEFS(map[string]*embedded.File{
	"/adjust_hp.html": embedded.NewFile("adjust_hp.html", time.Now(), 267, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x8c, 0x8e, 0x31, 0x4e, 0x03, 0x31,
		0x10, 0x45, 0xeb, 0xec, 0x29, 0xac, 0xa9, 0xa0, 0x81, 0x8a, 0x06, 0xad, 0x2d, 0x81, 0xa8, 0xb9,
		0xc3, 0xb0, 0xfe, 0x26, 0x86, 0xf1, 0x78, 0x25, 0x8f, 0x97, 0x44, 0xd1, 0xde, 0x1d, 0xc9, 0x80,
		0x44, 0x99, 0xfe, 0xbf, 0xf7, 0xdf, 0x1c, 0xf3, 0xe6, 0x16, 0xe1, 0xd6, 0x3c, 0x59, 0x36, 0x01,
		0x85, 0xcb, 0xe5, 0xee, 0x95, 0x0b, 0xf6, 0x7d, 0xbe, 0x8f, 0x79, 0x0b, 0xd3, 0x98, 0xe4, 0xe8,
		0x29, 0x65, 0x48, 0x6c, 0xe4, 0xaa, 0x7e, 0xe2, 0x1c, 0xeb, 0x97, 0x7a, 0x3a, 0xb2, 0x46, 0xc1,
		0x0b, 0x12, 0x77, 0xb1, 0xe7, 0x6e, 0x56, 0xf5, 0x06, 0x1b, 0xd4, 0x6e, 0x29, 0x4c, 0x87, 0xff,
		0x72, 0xe1, 0x37, 0x08, 0x85, 0xa7, 0xf8, 0xd1, 0x9b, 0x15, 0xa8, 0x3d, 0xfe, 0xea, 0x0f, 0x73,
		0xd6, 0xb5, 0xdb, 0xdf, 0x6e, 0x9c, 0x90, 0xb3, 0xf3, 0x0a, 0x4f, 0x86, 0x93, 0x91, 0x53, 0x2e,
		0xf0, 0xc4, 0x83, 0x24, 0x57, 0xf8, 0x24, 0xd0, 0x77, 0x3b, 0x7a, 0x7a, 0x20, 0xc7, 0xdd, 0xea,
		0x52, 0xcb, 0x2a, 0x30, 0x78, 0xaa, 0x29, 0x5d, 0x9b, 0x37, 0xc8, 0x54, 0x97, 0xde, 0xc2, 0xf4,
		0x53, 0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x08, 0x14, 0x15, 0xbe, 0x0b, 0x01, 0x00, 0x00,
	}),
	"/board.html": embedded.NewFile("board.html", time.Now(), 3462, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x9c, 0x56, 0xdb, 0x6e, 0xe3, 0x36,
		0x10, 0x7d, 0xde, 0xfd, 0x0a, 0x82, 0x48, 0x81, 0x04, 0x5d, 0x2b, 0x7d, 0x6e, 0x6d, 0x01, 0x69,
		0xb2, 0x41, 0x02, 0xb4, 0xbb, 0xc1, 0xc6, 0xc0, 0x3e, 0x8f, 0xc5, 0xb1, 0xc5, 0x86, 0x26, 0x55,
		0x72, 0xa4, 0x24, 0x10, 0xf4, 0xef, 0x05, 0x2f, 0xbe, 0x4a, 0xda, 0xda, 0xfb, 0x64, 0x5a, 0x3c,
		0x3c, 0x33, 0xc3, 0xb9, 0x1c, 0x4e, 0x85, 0x6c, 0x98, 0x14, 0x33, 0xbe, 0x30, 0x60, 0x05, 0x67,
		0x46, 0x0b, 0x0b, 0x2b, 0x47, 0x60, 0x69, 0xc6, 0xfd, 0xf2, 0xd9, 0x2f, 0x1f, 0x40, 0x0b, 0x85,
		0xf6, 0x12, 0x1b, 0xd4, 0x74, 0xb5, 0x41, 0xa1, 0x26, 0xb4, 0x11, 0xf5, 0xd9, 0x2f, 0x8f, 0x50,
		0x1f, 0x3f, 0x44, 0x98, 0x69, 0x36, 0xa8, 0xaf, 0x4d, 0x0f, 0x94, 0xa8, 0x14, 0x42, 0x83, 0x11,
		0xf4, 0x97, 0x5f, 0x0e, 0xa1, 0x4c, 0xe5, 0x01, 0xa6, 0x1a, 0x31, 0x83, 0x5a, 0x6c, 0x7c, 0x11,
		0x47, 0x90, 0xfc, 0xe3, 0x87, 0x10, 0x67, 0xa1, 0xc0, 0xb9, 0x19, 0x2f, 0x35, 0xac, 0x91, 0xe7,
		0xd3, 0x85, 0xcd, 0xbf, 0x97, 0x66, 0x7a, 0x2d, 0x64, 0x73, 0x8c, 0x90, 0x5a, 0x52, 0x44, 0x3c,
		0x6a, 0x49, 0x83, 0x90, 0xb2, 0x9a, 0x94, 0xc2, 0xf2, 0xfc, 0x41, 0x12, 0x7b, 0x32, 0x52, 0x93,
		0x1b, 0x83, 0x09, 0x74, 0x05, 0xcf, 0x9f, 0x09, 0xa8, 0x1e, 0xc3, 0xf0, 0xfc, 0xb6, 0xb6, 0x16,
		0xf5, 0xa8, 0xa9, 0x35, 0xbc, 0xf1, 0xfc, 0x6f, 0x78, 0x1b, 0xdc, 0x87, 0x22, 0xba, 0x72, 0x63,
		0xd7, 0xc6, 0xb2, 0x5b, 0xff, 0x75, 0x04, 0xc7, 0xf3, 0x9b, 0xdb, 0x31, 0x0a, 0x32, 0x75, 0x51,
		0xf2, 0x7c, 0xee, 0x7f, 0xc6, 0x30, 0x4b, 0x05, 0xc4, 0xf3, 0x7b, 0x05, 0xc3, 0x8e, 0x3a, 0x68,
		0xd0, 0x45, 0x5f, 0x9e, 0xa1, 0x91, 0x7a, 0xc5, 0xe6, 0xa5, 0x35, 0xaf, 0xc3, 0xde, 0x2c, 0x8d,
		0x25, 0x49, 0xb5, 0x40, 0x9e, 0xdf, 0x1b, 0x3b, 0x4c, 0x68, 0x71, 0xa9, 0xf0, 0x8d, 0xe7, 0xdf,
		0xc2, 0xef, 0x20, 0xe4, 0x55, 0x2a, 0xc5, 0xf3, 0xef, 0x52, 0xa9, 0x61, 0xa7, 0x89, 0xa0, 0x78,
		0x71, 0x31, 0x99, 0x37, 0xf1, 0xcf, 0x20, 0x50, 0x1b, 0xc2, 0x04, 0xfb, 0xe2, 0x97, 0x09, 0xd4,
		0xb6, 0x13, 0x76, 0x11, 0x9a, 0x83, 0xfd, 0x3e, 0x63, 0x59, 0xd7, 0xb5, 0xad, 0x05, 0xbd, 0x42,
		0x96, 0xdd, 0x9a, 0xf5, 0x02, 0x08, 0x34, 0xb9, 0xae, 0x3b, 0xe4, 0xf2, 0xf5, 0xd5, 0xb6, 0xf1,
		0x54, 0x96, 0x12, 0x3b, 0x87, 0x95, 0x3f, 0xcd, 0x59, 0xe1, 0xbb, 0xad, 0x6d, 0xb3, 0xc7, 0xbb,
		0xae, 0xf3, 0x95, 0x79, 0x70, 0xd2, 0xd7, 0xef, 0xa4, 0x0c, 0xd5, 0xcb, 0x99, 0xff, 0xb3, 0x82,
		0x85, 0xc2, 0x19, 0x27, 0x5b, 0x63, 0x00, 0x7f, 0x98, 0xca, 0x0d, 0x76, 0x09, 0x8e, 0x2d, 0x61,
		0x52, 0x3b, 0xb4, 0xac, 0x6d, 0xb3, 0xf9, 0x7b, 0x85, 0x9e, 0x9f, 0x24, 0xf9, 0x13, 0xe9, 0xcb,
		0x1d, 0xba, 0xc2, 0xca, 0x8a, 0xa4, 0xd1, 0xde, 0xdc, 0xf4, 0x5a, 0x06, 0x96, 0xb6, 0x95, 0x4b,
		0x96, 0x7d, 0xad, 0xa9, 0xeb, 0xa6, 0xae, 0x02, 0xbd, 0xe1, 0x34, 0x35, 0xf1, 0xbc, 0x6d, 0x51,
		0x0b, 0x1f, 0x68, 0xf6, 0x05, 0xd6, 0xe8, 0x17, 0x3b, 0xf0, 0xb5, 0x47, 0x6f, 0x10, 0xde, 0xf9,
		0x74, 0x95, 0x07, 0x51, 0x14, 0x46, 0x93, 0x35, 0xca, 0x8d, 0x78, 0xbc, 0xa8, 0x95, 0x72, 0xf8,
		0x8e, 0x6c, 0x51, 0x13, 0x19, 0x1d, 0xe8, 0x8d, 0x65, 0x97, 0xf8, 0x6f, 0xba, 0xe9, 0xec, 0x9b,
		0xa9, 0xb5, 0x60, 0xbf, 0x5d, 0xb1, 0xcb, 0xf4, 0xe1, 0xd1, 0xa5, 0x6b, 0x64, 0xd9, 0x55, 0xd7,
		0x31, 0xa9, 0x1b, 0xe9, 0xe4, 0x42, 0x61, 0x72, 0x64, 0x1b, 0xf5, 0xbd, 0x29, 0x6a, 0xc7, 0xa8,
		0x94, 0x8e, 0x15, 0x9b, 0xf4, 0xf8, 0xe9, 0x51, 0x28, 0x59, 0xbc, 0xcc, 0xf8, 0x1a, 0x5e, 0x30,
		0x11, 0x6d, 0xb3, 0x77, 0x99, 0x72, 0x71, 0xb5, 0xbb, 0x9d, 0x9e, 0xc7, 0x15, 0xea, 0x42, 0xaa,
		0x09, 0x28, 0x4a, 0x3e, 0x6f, 0x0d, 0x7e, 0x16, 0x92, 0x06, 0x4d, 0xa1, 0x90, 0xe7, 0xd9, 0x28,
		0x4c, 0xf5, 0x7e, 0xcc, 0x7e, 0x57, 0x57, 0x4a, 0x16, 0x40, 0x38, 0x68, 0x42, 0x6c, 0x76, 0xcf,
		0xb2, 0x43, 0x16, 0x5c, 0xd9, 0x33, 0x84, 0x0a, 0xc7, 0xac, 0x84, 0xad, 0x03, 0x13, 0xb1, 0x2c,
		0x3e, 0xb1, 0x9e, 0xb5, 0x6d, 0x63, 0x0d, 0x34, 0x98, 0x1f, 0xa9, 0x27, 0x36, 0x45, 0xa8, 0x88,
		0xfd, 0x52, 0x08, 0xc5, 0xe8, 0x27, 0xb1, 0x04, 0x92, 0x4d, 0x28, 0x49, 0x54, 0x2e, 0xfc, 0x56,
		0x56, 0x6a, 0x5a, 0x32, 0xfe, 0xcb, 0xaf, 0x82, 0xb3, 0x3d, 0xcc, 0x9f, 0x10, 0xf7, 0x43, 0x7d,
		0x0c, 0x36, 0x7c, 0x1c, 0xcc, 0x23, 0x2e, 0xb5, 0x6d, 0x16, 0x07, 0xf6, 0x1c, 0x56, 0x43, 0x0e,
		0xa6, 0xdd, 0x31, 0xea, 0x9f, 0xeb, 0xfe, 0xb2, 0x9a, 0x14, 0xb5, 0x0d, 0xf4, 0xf7, 0xc6, 0xae,
		0x81, 0x08, 0xc5, 0xc3, 0xd3, 0xce, 0xc6, 0x48, 0x93, 0xf5, 0x52, 0x5c, 0x22, 0xd8, 0x5e, 0xa5,
		0xde, 0x88, 0x7f, 0x6a, 0x47, 0xac, 0x94, 0xc4, 0xaa, 0x20, 0x57, 0x7b, 0x39, 0x86, 0xb0, 0xf7,
		0xf0, 0x74, 0x54, 0x3f, 0x3f, 0x4a, 0x67, 0xd4, 0xa4, 0x93, 0x13, 0x9a, 0x3d, 0x3c, 0xdd, 0xd7,
		0x4a, 0x0d, 0xdf, 0x17, 0x8c, 0x65, 0x61, 0x80, 0xe7, 0xe6, 0x76, 0x8c, 0x23, 0x0a, 0xd8, 0xe9,
		0x4c, 0x41, 0xe8, 0x7e, 0x40, 0xe7, 0xb5, 0xee, 0x74, 0x36, 0xaf, 0x89, 0x63, 0x64, 0x5b, 0xa5,
		0x3b, 0x83, 0x6e, 0x73, 0xe4, 0x19, 0x7c, 0xbd, 0x0f, 0xb1, 0x46, 0x6d, 0x3c, 0x9d, 0x32, 0x6a,
		0xe8, 0x38, 0x9f, 0x17, 0xd2, 0xd3, 0xd9, 0xbc, 0xe0, 0x8e, 0x73, 0x25, 0xd5, 0x3d, 0x23, 0xad,
		0xf1, 0xc0, 0x30, 0x5b, 0x90, 0xe6, 0x33, 0x5a, 0xaa, 0x37, 0xbe, 0x55, 0xed, 0x26, 0x85, 0xb4,
		0x85, 0xc2, 0x7e, 0x57, 0x08, 0xe6, 0xe9, 0x0f, 0x9a, 0x41, 0xf8, 0x07, 0x40, 0x7f, 0x96, 0x86,
		0xa7, 0xc0, 0x76, 0x44, 0x1e, 0x3d, 0x07, 0x2e, 0xa4, 0x16, 0xf8, 0xf6, 0x89, 0x5d, 0x78, 0xb6,
		0xb0, 0x17, 0x5e, 0x11, 0x51, 0x24, 0xf7, 0x25, 0x96, 0x60, 0xc5, 0x77, 0x1a, 0x1c, 0xe0, 0xd9,
		0x5c, 0xae, 0x51, 0xb0, 0x49, 0x00, 0xf7, 0xdd, 0x77, 0x64, 0xaa, 0x57, 0xa0, 0xa2, 0x37, 0xb5,
		0xdb, 0x36, 0x1e, 0xbf, 0x43, 0x02, 0xa9, 0xd2, 0xc4, 0xec, 0xba, 0x3f, 0x58, 0x88, 0x84, 0x91,
		0x61, 0xb1, 0xb1, 0x8f, 0x64, 0x29, 0x45, 0xb7, 0x0b, 0x25, 0xc4, 0xe9, 0xc7, 0x79, 0x0c, 0xe2,
		0x40, 0x3f, 0x7c, 0xd0, 0x7e, 0xd6, 0x8e, 0x7a, 0xf7, 0x3f, 0xda, 0x78, 0x74, 0xb9, 0x3f, 0x67,
		0x3f, 0x3d, 0x36, 0x7a, 0x17, 0x99, 0x9e, 0xd6, 0xbb, 0x7b, 0xd8, 0x7b, 0xec, 0xa4, 0xa7, 0xca,
		0xb0, 0x06, 0xca, 0x35, 0x8e, 0x55, 0x44, 0x92, 0xc2, 0x23, 0xbf, 0xa3, 0x0a, 0x9e, 0xe3, 0xf9,
		0xce, 0xfe, 0x5e, 0x08, 0x7b, 0x8f, 0xca, 0xf8, 0x25, 0x7e, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff,
		0x6f, 0x6b, 0xfb, 0x83, 0x86, 0x0d, 0x00, 0x00,
	}),
	"/detail.html": embedded.NewFile("detail.html", time.Now(), 8409, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x9c, 0x59, 0xc1, 0x6e, 0xe3, 0x38,
		0x0f, 0xbe, 0xff, 0x4f, 0x21, 0xf4, 0xf4, 0xef, 0x61, 0xba, 0xf7, 0xc5, 0x6c, 0x81, 0x34, 0x69,
		0xa7, 0x83, 0xad, 0x3b, 0x69, 0x32, 0x98, 0x9d, 0x2b, 0x63, 0xb3, 0x8e, 0x10, 0x45, 0x32, 0x24,
		0x39, 0xbb, 0x99, 0x22, 0xef, 0xbe, 0xb0, 0xa5, 0x38, 0x4e, 0x6a, 0xc9, 0x62, 0xce, 0x24, 0xbf,
		0x8f, 0xa6, 0x48, 0x8a, 0xb4, 0x3e, 0x17, 0x7c, 0xc7, 0x72, 0x01, 0xc6, 0xfc, 0x79, 0x53, 0x08,
		0x58, 0xa1, 0xb8, 0xb9, 0x7b, 0x81, 0x2d, 0xfe, 0xf1, 0xf9, 0xf7, 0x82, 0xef, 0xee, 0xfe, 0x77,
		0x26, 0x7f, 0xe3, 0x28, 0x8a, 0x9b, 0xbb, 0xf7, 0xf7, 0xdb, 0x67, 0xbe, 0xd2, 0xa0, 0xf7, 0x0f,
		0xd2, 0x72, 0xbb, 0xbf, 0x6d, 0x0c, 0x0e, 0x87, 0x21, 0x0b, 0x8f, 0x38, 0x5d, 0x90, 0xf0, 0xa6,
		0x8b, 0x28, 0xda, 0xcf, 0x39, 0x09, 0xed, 0xe7, 0xbc, 0x43, 0x7b, 0x7f, 0xff, 0xc4, 0xf8, 0x1b,
		0xbb, 0x50, 0x58, 0x40, 0x8e, 0x87, 0xc3, 0x20, 0x55, 0x23, 0x22, 0x91, 0x39, 0xac, 0x1e, 0x1d,
		0xca, 0xe2, 0x70, 0x08, 0x10, 0x4f, 0x1b, 0xb0, 0x00, 0x73, 0x2b, 0xa3, 0x45, 0xcd, 0xa1, 0x7d,
		0xe0, 0x1e, 0x42, 0x9f, 0x08, 0x5e, 0xca, 0x2d, 0x4a, 0x4b, 0x62, 0xe8, 0xac, 0xa2, 0xc7, 0xb3,
		0xe4, 0xbf, 0x68, 0x31, 0x6b, 0x0c, 0xa2, 0x88, 0xdf, 0xf7, 0x15, 0x0d, 0xb1, 0x31, 0x38, 0x1c,
		0xd8, 0x47, 0xa6, 0x7a, 0xe5, 0x44, 0x11, 0xb2, 0xaf, 0x92, 0xd3, 0xa2, 0xd2, 0x18, 0xc4, 0x03,
		0x82, 0xd2, 0x20, 0xed, 0x2c, 0x9d, 0xc9, 0x48, 0xde, 0x4e, 0x6a, 0x0d, 0xa1, 0xf3, 0xad, 0x35,
		0xd0, 0x8e, 0xb6, 0xc5, 0x4a, 0xcb, 0x9d, 0x29, 0x0d, 0x79, 0x3a, 0x74, 0x12, 0x93, 0x69, 0xa6,
		0x0a, 0x13, 0x0d, 0xdb, 0xd3, 0x58, 0x99, 0x3f, 0xcd, 0x33, 0xb4, 0x6b, 0x55, 0xdc, 0x3e, 0xcd,
		0x2f, 0x42, 0x33, 0xc4, 0xf8, 0x34, 0x1f, 0x67, 0x9c, 0x91, 0xbe, 0xec, 0x69, 0x16, 0x3f, 0x76,
		0xd8, 0x51, 0x4f, 0xbd, 0xb1, 0x18, 0x39, 0xf4, 0x19, 0xbe, 0xa1, 0x34, 0x7c, 0x87, 0x93, 0x15,
		0x17, 0xdc, 0x72, 0x0c, 0x35, 0x10, 0xa7, 0x88, 0xac, 0xd3, 0x23, 0xb9, 0x32, 0x44, 0x93, 0xd8,
		0xd5, 0x66, 0x8b, 0x90, 0x47, 0xb4, 0x5b, 0x60, 0xb6, 0x48, 0xa6, 0xfc, 0xba, 0xdd, 0xd6, 0x32,
		0xd4, 0xc3, 0x9d, 0x90, 0x56, 0xd2, 0x1e, 0x2f, 0x91, 0x7e, 0x81, 0x86, 0x1b, 0x1b, 0xba, 0x42,
		0x5a, 0x21, 0xed, 0x12, 0xf1, 0x78, 0x89, 0xf4, 0xcb, 0x50, 0xc0, 0x97, 0xb4, 0x80, 0x2f, 0xd3,
		0x03, 0xfe, 0x37, 0xc2, 0x46, 0xa2, 0x31, 0xc1, 0xec, 0x3b, 0x29, 0x90, 0x5c, 0xe8, 0xe3, 0x26,
		0x35, 0xa3, 0x65, 0x85, 0x58, 0xd0, 0xbe, 0xb2, 0xb1, 0x18, 0xbc, 0x1c, 0x1a, 0x41, 0xa6, 0x8a,
		0x91, 0x02, 0xcc, 0x50, 0x60, 0x28, 0xd5, 0x5a, 0x19, 0xc9, 0x1b, 0x8f, 0x96, 0x9a, 0x68, 0x20,
		0x4b, 0x0c, 0x85, 0xc2, 0x09, 0x89, 0xd3, 0x8a, 0xc3, 0x4b, 0x0c, 0x35, 0x75, 0x18, 0x6a, 0x2d,
		0xa2, 0x3d, 0x72, 0x81, 0x90, 0xaf, 0x89, 0xb5, 0x01, 0xf9, 0x7a, 0xe4, 0x88, 0x96, 0x15, 0xe6,
		0x1c, 0xc4, 0xc4, 0x5a, 0xc8, 0x37, 0xa1, 0x0c, 0xf5, 0x4a, 0xcc, 0x6b, 0x51, 0x73, 0xe8, 0x8c,
		0x20, 0xb5, 0x50, 0x2b, 0x14, 0xe2, 0x99, 0x6f, 0x46, 0x7b, 0x77, 0xab, 0xf8, 0x49, 0xf0, 0xcd,
		0xb5, 0xed, 0x7b, 0x88, 0x89, 0xe2, 0xa4, 0xf9, 0x4b, 0xaa, 0x7f, 0x64, 0xcc, 0x3b, 0xc3, 0x5a,
		0x15, 0xba, 0x57, 0x47, 0x68, 0x92, 0x3b, 0x73, 0x8d, 0x15, 0xe8, 0x60, 0xea, 0x7b, 0x8f, 0x8e,
		0x5a, 0x57, 0x38, 0x75, 0x22, 0xa0, 0xf8, 0x35, 0x53, 0x5b, 0xe0, 0x32, 0x7a, 0x8a, 0xcc, 0xeb,
		0xd0, 0x7d, 0xea, 0xc0, 0xd3, 0xa6, 0xb2, 0xf6, 0xa0, 0xf7, 0x6c, 0x99, 0x2b, 0x4d, 0x4c, 0x16,
		0x6f, 0xea, 0x2c, 0xa3, 0xe5, 0x7a, 0x0f, 0xcd, 0x40, 0xd1, 0x66, 0x3d, 0x89, 0xa1, 0xb1, 0x73,
		0x66, 0xf1, 0x35, 0x31, 0xbb, 0xa7, 0x6d, 0x3c, 0xd9, 0xfd, 0x08, 0x1e, 0x6d, 0xa0, 0x9b, 0x66,
		0xf1, 0x89, 0xee, 0x11, 0xc1, 0xd2, 0x42, 0xdb, 0x5a, 0xc4, 0xa7, 0xc4, 0x0d, 0x17, 0x82, 0x98,
		0x1d, 0xad, 0xc9, 0xf8, 0x52, 0xcb, 0x41, 0xb8, 0x91, 0x37, 0xb4, 0xda, 0x36, 0xed, 0xaf, 0xd1,
		0xa0, 0x6e, 0xb8, 0x1d, 0x70, 0x62, 0xad, 0x3c, 0x83, 0x2c, 0x6b, 0x28, 0x83, 0xed, 0xae, 0x93,
		0x93, 0x1c, 0xe9, 0xa1, 0xa6, 0xd6, 0xec, 0x6b, 0xa8, 0x52, 0x5f, 0x69, 0x07, 0xf0, 0x9a, 0x58,
		0x94, 0x0f, 0x72, 0xc7, 0xb5, 0xa2, 0x2f, 0xda, 0x3d, 0xbb, 0x68, 0xf2, 0x7c, 0xd3, 0x25, 0x48,
		0xfe, 0x0b, 0x2c, 0x57, 0xb4, 0x4e, 0xdc, 0x37, 0x8c, 0xaf, 0xde, 0x1a, 0xc1, 0xd4, 0x9a, 0xb8,
		0x7e, 0x7b, 0xa3, 0x28, 0xf2, 0x0c, 0x4d, 0xae, 0x79, 0xd5, 0x78, 0xc0, 0xfe, 0xff, 0x83, 0x9b,
		0x1a, 0xc4, 0x6f, 0xc4, 0x15, 0xa5, 0x03, 0x70, 0xe6, 0x89, 0x53, 0xc1, 0xf8, 0xdd, 0xeb, 0xe6,
		0x82, 0x6b, 0x2f, 0xde, 0x0b, 0x92, 0xa4, 0x44, 0xe9, 0x7d, 0xcb, 0xb5, 0x31, 0x18, 0xf9, 0xfa,
		0x2f, 0x28, 0x0b, 0xd4, 0x01, 0x7a, 0x27, 0x24, 0x31, 0x1f, 0xf1, 0x12, 0x4b, 0xef, 0x5e, 0x28,
		0x55, 0x08, 0x1e, 0x5c, 0xd2, 0x3a, 0x39, 0xed, 0x5e, 0x39, 0xa1, 0x26, 0xfa, 0x31, 0xd7, 0x6a,
		0xcd, 0x57, 0xdc, 0x62, 0xb1, 0xcc, 0xd7, 0x4a, 0x89, 0x50, 0x16, 0x9c, 0xf4, 0x98, 0x57, 0x24,
		0x39, 0x36, 0x40, 0x93, 0x1a, 0x28, 0x7c, 0x53, 0x1a, 0xa7, 0x6a, 0xbb, 0x82, 0xd0, 0x46, 0xe9,
		0x54, 0x98, 0xd3, 0xa1, 0xc5, 0xeb, 0x0c, 0x3c, 0x75, 0x9f, 0xaf, 0x35, 0x97, 0x65, 0xd4, 0x23,
		0xa7, 0x72, 0x8d, 0x47, 0xe7, 0xe0, 0x89, 0x1e, 0x65, 0x4a, 0x83, 0x08, 0xee, 0x60, 0xad, 0x90,
		0xb6, 0x84, 0x79, 0xbc, 0x44, 0xfa, 0x2f, 0x08, 0xe1, 0x42, 0x02, 0x6a, 0x19, 0x41, 0x7a, 0x11,
		0x7d, 0xb3, 0x6b, 0xd4, 0x11, 0xf6, 0x56, 0xce, 0xc8, 0x3e, 0xf4, 0x60, 0x13, 0x1d, 0xf9, 0x51,
		0x0b, 0x89, 0x1a, 0xdc, 0xdc, 0x18, 0x70, 0xe6, 0x4c, 0x87, 0xe4, 0xcf, 0x05, 0x7a, 0xa2, 0x4f,
		0x2f, 0xca, 0x86, 0x52, 0xa2, 0x11, 0xd1, 0x1e, 0x34, 0x5a, 0xac, 0x44, 0xe2, 0xef, 0xb8, 0xad,
		0x04, 0x58, 0x34, 0x93, 0xaa, 0x12, 0x3c, 0xb8, 0xa3, 0x74, 0x6a, 0xcc, 0xeb, 0xd1, 0xae, 0xd4,
		0x0f, 0x24, 0xa9, 0x39, 0xf3, 0xd6, 0xfe, 0x06, 0x8c, 0x44, 0xc7, 0x6b, 0x30, 0x72, 0x94, 0xce,
		0xa0, 0x53, 0xdb, 0x1b, 0x18, 0x5c, 0x5a, 0xb0, 0xdc, 0x58, 0x9e, 0x87, 0x9a, 0x6f, 0xbb, 0x66,
		0x9c, 0xb4, 0xc8, 0xab, 0x46, 0x9f, 0x20, 0xd1, 0xaf, 0x87, 0x7f, 0xad, 0x86, 0xdc, 0x8e, 0x2d,
		0x9a, 0x97, 0x6a, 0xb4, 0xa9, 0xee, 0x03, 0x47, 0xa2, 0x73, 0x93, 0x12, 0xa7, 0x60, 0xb1, 0x54,
		0x3a, 0x54, 0x6c, 0x93, 0x12, 0xd9, 0x51, 0x85, 0xb6, 0xfb, 0xf5, 0xa1, 0x53, 0xdb, 0xef, 0xde,
		0x58, 0x0c, 0xba, 0xe2, 0xa5, 0xb4, 0x06, 0x7c, 0x44, 0xa4, 0x3c, 0x9c, 0x4d, 0x74, 0xbe, 0x46,
		0xbb, 0xaf, 0x82, 0x93, 0x5c, 0xab, 0xc5, 0x4e, 0x6a, 0xf4, 0xd7, 0xb4, 0x3e, 0x45, 0xea, 0x80,
		0x01, 0x56, 0xab, 0xd0, 0x9f, 0x13, 0x27, 0xa4, 0x4d, 0x12, 0x1e, 0x2f, 0x91, 0xfe, 0x51, 0xe5,
		0xb5, 0x39, 0x4e, 0x1d, 0x01, 0x2f, 0xbc, 0x8e, 0x1f, 0x6c, 0x68, 0xfb, 0xec, 0x39, 0x7c, 0x6a,
		0x87, 0xd4, 0xc0, 0x6d, 0xe8, 0x94, 0x9c, 0x90, 0xb8, 0x5f, 0x38, 0xbc, 0xd4, 0xf2, 0x11, 0x16,
		0xb5, 0x04, 0x8b, 0x2f, 0xb0, 0xc5, 0x47, 0xa5, 0xb7, 0xc1, 0x47, 0x51, 0xaf, 0xc7, 0xc8, 0x2f,
		0xe0, 0x03, 0x14, 0xa9, 0x3b, 0x69, 0xd7, 0xab, 0x22, 0x2d, 0xfa, 0xa4, 0x44, 0xef, 0xd2, 0x97,
		0x04, 0xa9, 0x45, 0x1e, 0x7a, 0x54, 0xc8, 0xa0, 0xe4, 0x39, 0x73, 0x4f, 0x14, 0x20, 0x89, 0x7f,
		0x84, 0xb3, 0xa1, 0x27, 0x86, 0xc1, 0x2f, 0x56, 0xb5, 0xa6, 0xfe, 0x6d, 0x6e, 0x4d, 0x8e, 0xf8,
		0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xaa, 0x8d, 0x7e, 0xd9, 0x20, 0x00, 0x00,
	}),
	"/edit_combatant.html": embedded.NewFile("edit_combatant.html", time.Now(), 2423, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xac, 0x94, 0x5d, 0x6b, 0xdb, 0x30,
		0x14, 0x86, 0xaf, 0xb7, 0x5f, 0x21, 0x0e, 0x0c, 0x5a, 0x46, 0xbb, 0xc1, 0xc6, 0x2e, 0x46, 0x1d,
		0x48, 0xdb, 0x85, 0xee, 0x66, 0x94, 0xad, 0xb0, 0xcb, 0x72, 0x62, 0x1d, 0xd5, 0x22, 0xfa, 0x30,
		0xf6, 0x91, 0x9b, 0x60, 0xfc, 0xdf, 0x87, 0xec, 0xa6, 0x25, 0xa1, 0x2c, 0x55, 0xec, 0x9b, 0xc4,
		0x96, 0xf5, 0xea, 0x7d, 0x1e, 0x6c, 0xce, 0x85, 0xd4, 0x8d, 0xd0, 0x32, 0x03, 0xa5, 0xc9, 0xc8,
		0x1a, 0x84, 0x77, 0x2b, 0xda, 0x48, 0xff, 0xe8, 0x32, 0x28, 0xd0, 0x49, 0x43, 0xd7, 0xa4, 0x30,
		0x18, 0xbe, 0x0c, 0xcc, 0xde, 0x9d, 0x50, 0x43, 0x8e, 0x4f, 0x61, 0xf6, 0xfe, 0x5d, 0x9f, 0xcc,
		0x0d, 0xd6, 0x75, 0x06, 0x06, 0x97, 0x64, 0x60, 0xf6, 0x0b, 0x2d, 0x7d, 0xbf, 0xf8, 0x24, 0x75,
		0x13, 0x9f, 0x6b, 0x57, 0x06, 0xde, 0xee, 0xe8, 0x8f, 0x07, 0xc1, 0x9b, 0x92, 0x32, 0x60, 0x5a,
		0x33, 0x08, 0x87, 0x96, 0x32, 0x88, 0xbf, 0x20, 0x2c, 0xae, 0x0d, 0xb9, 0x07, 0x2e, 0x32, 0xf8,
		0xfa, 0x19, 0x04, 0x06, 0xf6, 0xb9, 0xb7, 0xa5, 0x21, 0xa6, 0x0c, 0xbc, 0x52, 0x20, 0x1a, 0x34,
		0x81, 0x32, 0x68, 0xdb, 0xf3, 0xd8, 0xd2, 0x75, 0x6f, 0x25, 0xed, 0xcf, 0x52, 0x3e, 0x0f, 0xf5,
		0xeb, 0xcc, 0x3f, 0x9d, 0x66, 0x8d, 0xac, 0x1b, 0x12, 0x97, 0x58, 0x27, 0xe3, 0x6b, 0xa7, 0xf9,
		0x7e, 0x89, 0xf5, 0x9e, 0xc3, 0xff, 0x15, 0xca, 0x4a, 0x3b, 0x56, 0x02, 0x3e, 0x7c, 0x94, 0x20,
		0xce, 0x5f, 0x08, 0x22, 0xc0, 0xdb, 0xcd, 0x5e, 0xf7, 0xb9, 0xd1, 0x2c, 0x6e, 0xbd, 0x76, 0x5c,
		0x8b, 0x93, 0x45, 0x30, 0xe6, 0x34, 0xd5, 0xa8, 0x28, 0xef, 0x55, 0x30, 0x66, 0xc7, 0xe7, 0xdb,
		0x81, 0x57, 0x72, 0x73, 0x1b, 0xab, 0xa6, 0x44, 0xbf, 0x23, 0x5b, 0xfa, 0x0a, 0xab, 0xcd, 0x31,
		0xfc, 0x6c, 0xcb, 0x44, 0xfc, 0xe7, 0xba, 0x29, 0x1d, 0xae, 0xd1, 0xe2, 0x03, 0x1d, 0x23, 0x20,
		0xfb, 0x64, 0xa2, 0xc3, 0x50, 0x37, 0x56, 0x60, 0x5e, 0x59, 0x5f, 0x89, 0xab, 0xb8, 0x96, 0x0a,
		0x8e, 0xf9, 0x0e, 0xf1, 0x97, 0x03, 0xc4, 0xf3, 0xab, 0xb1, 0xac, 0x77, 0x3e, 0xe4, 0x85, 0x18,
		0x45, 0x7c, 0xc6, 0xf1, 0x8c, 0x24, 0xee, 0xbe, 0x75, 0x3c, 0xfc, 0xc2, 0x20, 0x9f, 0x2d, 0xbc,
		0x67, 0x92, 0x23, 0x15, 0x94, 0x41, 0x4e, 0x32, 0x88, 0xd5, 0x13, 0x08, 0xf8, 0x8a, 0x35, 0x07,
		0x49, 0xe2, 0x0f, 0x36, 0xc9, 0x83, 0x53, 0x6d, 0xd3, 0x69, 0xe4, 0xdb, 0x54, 0xac, 0x1c, 0x2b,
		0xf0, 0x9b, 0x94, 0xa1, 0xf5, 0x51, 0xf4, 0x55, 0x1f, 0x4d, 0x42, 0x1f, 0xda, 0xa6, 0xe0, 0xfe,
		0xab, 0x8d, 0x39, 0x8a, 0xfa, 0x51, 0xef, 0xcd, 0xf5, 0x43, 0xcc, 0xb1, 0x69, 0x0a, 0xe2, 0x39,
		0x33, 0xe6, 0xab, 0xf4, 0x6f, 0x7b, 0x88, 0x1d, 0x9a, 0x23, 0xc3, 0xae, 0x34, 0xc6, 0xa2, 0xda,
		0x23, 0xcd, 0x0b, 0xca, 0x57, 0x4b, 0xbf, 0x86, 0xd9, 0x2e, 0xdc, 0xf3, 0xfa, 0x13, 0xdf, 0xcb,
		0xfd, 0xc0, 0x48, 0x8e, 0xec, 0x06, 0xda, 0x56, 0x2b, 0x71, 0xfe, 0x23, 0x5e, 0x77, 0x9d, 0xe8,
		0xf7, 0x90, 0x6c, 0x5b, 0x72, 0xb2, 0xeb, 0x66, 0xa2, 0x5f, 0x7f, 0x72, 0x1f, 0xfe, 0xfe, 0x05,
		0x00, 0x00, 0xff, 0xff, 0x73, 0x55, 0xa5, 0x31, 0x77, 0x09, 0x00, 0x00,
	}),
	"/edit_note.html": embedded.NewFile("edit_note.html", time.Now(), 1006, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xdb, 0x30,
		0x0c, 0x3d, 0x67, 0x5f, 0x41, 0x08, 0x3d, 0x6c, 0x87, 0x39, 0x3b, 0xec, 0x30, 0x14, 0x76, 0x0f,
		0x43, 0x7a, 0x28, 0x30, 0x14, 0xc3, 0xb0, 0x61, 0x67, 0xd5, 0xa2, 0x6b, 0xa1, 0x32, 0xe9, 0x59,
		0xb4, 0x9b, 0x42, 0xd0, 0xbf, 0x0f, 0x92, 0x92, 0x65, 0x49, 0x03, 0x34, 0x27, 0x02, 0xe2, 0x7b,
		0x8f, 0xef, 0x51, 0xac, 0x8d, 0x5d, 0xc0, 0x9a, 0x46, 0x75, 0x16, 0x9d, 0xf1, 0x0a, 0x98, 0x9e,
		0xf0, 0xc5, 0xf0, 0x33, 0x35, 0xaa, 0xd7, 0x64, 0x1c, 0x6e, 0xb0, 0xd3, 0xb3, 0x93, 0xaf, 0xb3,
		0x08, 0xd3, 0x7b, 0x5c, 0x90, 0xe4, 0x83, 0xba, 0x79, 0xb7, 0xca, 0xcc, 0xd6, 0x69, 0xef, 0x1b,
		0xe5, 0xf4, 0x03, 0x3a, 0x75, 0x73, 0xcf, 0x82, 0xd7, 0xf5, 0xda, 0xd8, 0x25, 0xf5, 0x2d, 0x8d,
		0xb3, 0xec, 0x11, 0x59, 0x5e, 0x81, 0xbc, 0x8c, 0xd8, 0x28, 0xc1, 0xad, 0x28, 0x20, 0x3d, 0x60,
		0xa3, 0x0c, 0xfa, 0x76, 0xb2, 0xa3, 0x58, 0x26, 0x05, 0x83, 0xde, 0x3a, 0xa4, 0x47, 0xe9, 0x1b,
		0xf5, 0xe5, 0x93, 0x02, 0x3d, 0x0b, 0xb7, 0x3c, 0x8c, 0x0e, 0x05, 0x1b, 0xc5, 0x5d, 0xa7, 0x60,
		0xd1, 0x6e, 0xc6, 0x46, 0x85, 0x50, 0xa5, 0x61, 0xd5, 0xe6, 0xc0, 0x8e, 0xf1, 0x52, 0xf3, 0x59,
		0xb7, 0xe3, 0x76, 0xf6, 0xaf, 0x6c, 0x8a, 0x1d, 0xf0, 0x9f, 0xcd, 0xb6, 0xc7, 0xf6, 0xe9, 0x81,
		0xb7, 0x7b, 0xab, 0xa5, 0x19, 0x82, 0xed, 0xa0, 0x4c, 0xff, 0x99, 0x1e, 0x62, 0x84, 0x0c, 0x44,
		0x13, 0x02, 0x92, 0x89, 0xf1, 0x64, 0x37, 0x29, 0x00, 0x7b, 0x2b, 0x98, 0x96, 0x96, 0x3b, 0xa9,
		0xae, 0x6e, 0xb7, 0xa3, 0x9d, 0xd0, 0x83, 0x16, 0x90, 0x1e, 0xd3, 0x4b, 0xed, 0xd1, 0x61, 0x2b,
		0xbb, 0x61, 0x33, 0x89, 0x75, 0x99, 0xb2, 0x5a, 0xd5, 0x9c, 0x13, 0xee, 0xc3, 0x23, 0x1d, 0xb9,
		0xf8, 0x95, 0x90, 0xb7, 0x69, 0x32, 0x14, 0x85, 0x83, 0x13, 0x24, 0x53, 0xaf, 0x0b, 0xf9, 0xac,
		0x92, 0x17, 0x3d, 0x49, 0xd1, 0x22, 0x96, 0xb7, 0xf5, 0x32, 0xfe, 0x48, 0xb1, 0x5e, 0x17, 0x4c,
		0xce, 0xb6, 0xfb, 0xf9, 0x43, 0x48, 0xee, 0x5e, 0x07, 0x7b, 0xee, 0x79, 0x17, 0x2b, 0x84, 0x8f,
		0x70, 0x45, 0x70, 0xdd, 0x94, 0xc1, 0x31, 0x86, 0x30, 0x69, 0x7a, 0x44, 0xa8, 0x7e, 0xf7, 0xfc,
		0xcd, 0x7a, 0x89, 0xf1, 0x9c, 0xe9, 0x10, 0xaa, 0xbb, 0x4d, 0x8c, 0xc5, 0x36, 0xfe, 0x81, 0xea,
		0x6e, 0x03, 0x57, 0x94, 0x38, 0x67, 0x1c, 0x87, 0x50, 0x7d, 0x67, 0xef, 0xd1, 0x7b, 0xbb, 0xe0,
		0xbd, 0x1e, 0x30, 0xc6, 0xe3, 0x8d, 0x24, 0x13, 0x19, 0x7a, 0x9a, 0x66, 0x25, 0xf3, 0x44, 0xe7,
		0x53, 0x11, 0x4c, 0x3c, 0x93, 0xc9, 0x8c, 0xcb, 0xae, 0x3c, 0xe3, 0x8f, 0xee, 0xfb, 0xf3, 0x25,
		0xe7, 0xfd, 0x23, 0xd1, 0x2e, 0x3f, 0xec, 0xff, 0x3f, 0xe1, 0xa4, 0x96, 0xf2, 0x37, 0x00, 0x00,
		0xff, 0xff, 0x09, 0x27, 0x03, 0x1c, 0xee, 0x03, 0x00, 0x00,
	}),
	"/index.html": embedded.NewFile("index.html", time.Now(), 1836, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x8c, 0x55, 0xdf, 0x6f, 0xdb, 0x36,
		0x10, 0x7e, 0x6e, 0xfe, 0x0a, 0x8e, 0x4f, 0x2d, 0x30, 0x49, 0xed, 0xb2, 0x0e, 0xee, 0x20, 0x19,
		0xc8, 0x92, 0x22, 0x1b, 0x0a, 0x24, 0x5b, 0x90, 0x2d, 0x5d, 0x86, 0x3d, 0xd0, 0xe4, 0xc9, 0x62,
		0x72, 0x22, 0x35, 0xde, 0x49, 0x8e, 0x20, 0xf8, 0x7f, 0x1f, 0x24, 0x5b, 0xf1, 0xaf, 0xb4, 0xf5,
		0x93, 0xcc, 0xe3, 0x77, 0xdf, 0x7d, 0xfc, 0x78, 0x47, 0xa7, 0xdf, 0x5d, 0x5c, 0x9f, 0xdf, 0xfe,
		0xfd, 0xfb, 0x47, 0x51, 0x70, 0x89, 0xd3, 0x93, 0x74, 0xfc, 0x80, 0x32, 0xd3, 0x93, 0x57, 0x69,
		0x09, 0xac, 0x84, 0x2e, 0x54, 0x20, 0xe0, 0x4c, 0xd6, 0x9c, 0x47, 0x13, 0x29, 0x92, 0x7e, 0x87,
		0x2d, 0x23, 0x4c, 0xbb, 0x2e, 0xbe, 0xed, 0x7f, 0x2c, 0x97, 0x69, 0xb2, 0x8a, 0x8c, 0x49, 0x4e,
		0x95, 0x90, 0xc9, 0xc6, 0xc2, 0xa2, 0xf2, 0x81, 0xa5, 0xd0, 0xde, 0x31, 0x38, 0xce, 0xe4, 0xc2,
		0x1a, 0x2e, 0x32, 0x03, 0x8d, 0xd5, 0x10, 0x0d, 0x8b, 0xef, 0x85, 0x75, 0x96, 0xad, 0xc2, 0x88,
		0xb4, 0x42, 0xc8, 0xde, 0xc9, 0x9e, 0x05, 0xad, 0x7b, 0x14, 0x01, 0x30, 0x93, 0xc4, 0x2d, 0x02,
		0x15, 0x00, 0x2c, 0x45, 0x11, 0x20, 0xcf, 0x64, 0xc1, 0x5c, 0xd1, 0xcf, 0x49, 0x52, 0x13, 0xc4,
		0xb9, 0x77, 0xac, 0x16, 0x40, 0xbe, 0x84, 0x58, 0xfb, 0x32, 0x09, 0x80, 0xa0, 0x08, 0x28, 0x69,
		0xde, 0xc7, 0x6f, 0xe3, 0x77, 0xa7, 0x89, 0x26, 0x4a, 0x14, 0x62, 0xac, 0x89, 0xa4, 0xb0, 0x8e,
		0x61, 0x1e, 0x2c, 0xb7, 0x99, 0xa4, 0x42, 0x9d, 0x4e, 0x7e, 0x8c, 0x2e, 0xae, 0xae, 0x7f, 0xbd,
		0xff, 0x69, 0xf2, 0xe7, 0xa4, 0xb8, 0xcf, 0x3f, 0x7d, 0xbe, 0x0e, 0xfc, 0x70, 0xd7, 0x3c, 0x3c,
		0xd5, 0x74, 0xe9, 0x3f, 0xdc, 0xfd, 0xe1, 0xc2, 0xd5, 0xd3, 0x0f, 0xf4, 0xdf, 0xe5, 0x5b, 0xce,
		0x69, 0x5e, 0x9c, 0x35, 0xfc, 0x17, 0xde, 0xdc, 0x9d, 0x72, 0xf3, 0xf8, 0xf9, 0xee, 0xbe, 0x78,
		0x3f, 0xb9, 0xfa, 0xf0, 0x50, 0x49, 0xa1, 0x83, 0x27, 0xf2, 0xc1, 0xce, 0xad, 0xcb, 0xa4, 0x72,
		0xde, 0xb5, 0xa5, 0xaf, 0xe9, 0x9b, 0x87, 0x48, 0xb8, 0x80, 0x5e, 0x33, 0x1d, 0x01, 0x05, 0xa7,
		0x7d, 0xed, 0x18, 0xc2, 0x71, 0x70, 0xb4, 0xb3, 0xa0, 0x42, 0x7b, 0x1c, 0xd8, 0x00, 0x2b, 0x8b,
		0xc7, 0x61, 0x4b, 0x6f, 0xd4, 0x06, 0x4a, 0x3a, 0xd8, 0x8a, 0x05, 0x05, 0xbd, 0x23, 0xf1, 0x81,
		0xe4, 0x34, 0x4d, 0x56, 0x9b, 0xd3, 0x93, 0x34, 0x59, 0x35, 0x53, 0x3a, 0xf3, 0xa6, 0x15, 0xde,
		0x3d, 0x42, 0x6b, 0xfc, 0xc2, 0x65, 0xb2, 0x50, 0xce, 0x20, 0x5c, 0xa2, 0x9f, 0x29, 0xfc, 0x04,
		0xed, 0x85, 0x5f, 0xb8, 0xd7, 0xd0, 0x80, 0xe3, 0x37, 0x03, 0xb9, 0xb1, 0x8d, 0xb0, 0x26, 0x93,
		0x85, 0x09, 0xfd, 0x7a, 0x27, 0x10, 0x21, 0xe4, 0x2c, 0x77, 0xba, 0xcf, 0xd8, 0xe6, 0x00, 0x55,
		0x5a, 0x63, 0x10, 0x86, 0xec, 0x57, 0xa9, 0x15, 0x1a, 0x15, 0x51, 0x26, 0x73, 0x45, 0x22, 0x57,
		0x91, 0x2e, 0xa0, 0x09, 0xde, 0x45, 0xda, 0x06, 0x8d, 0x10, 0x05, 0x3b, 0x2f, 0x58, 0xcc, 0x6a,
		0x66, 0xef, 0xa4, 0x18, 0x7a, 0x39, 0x93, 0x57, 0xf0, 0xc4, 0xe2, 0x1f, 0xf7, 0xaf, 0x14, 0xde,
		0x69, 0xb4, 0xfa, 0x31, 0x93, 0x0e, 0x9e, 0xf8, 0xb6, 0x0e, 0xee, 0xf5, 0x9b, 0xfe, 0x90, 0x76,
		0xc5, 0x4d, 0x95, 0x72, 0x43, 0x55, 0xae, 0x83, 0xeb, 0x75, 0xd9, 0x5c, 0xc4, 0xbf, 0x78, 0x15,
		0x4c, 0x7c, 0xe3, 0x6b, 0x67, 0x96, 0xcb, 0xe1, 0x23, 0xba, 0x6e, 0x37, 0xda, 0x75, 0x80, 0x04,
		0xcb, 0xe5, 0xd9, 0x42, 0x59, 0xb6, 0x6e, 0x2e, 0x7e, 0x1b, 0x06, 0x81, 0x6d, 0x03, 0x5d, 0x07,
		0x3d, 0x22, 0x4d, 0x7a, 0xee, 0x97, 0x4f, 0x50, 0x3b, 0xe3, 0xf7, 0x15, 0xdf, 0x78, 0xc4, 0xf5,
		0x38, 0xf5, 0x2c, 0x5b, 0xc2, 0x83, 0x47, 0xdc, 0xd0, 0x6f, 0xc9, 0x7f, 0xd9, 0xbb, 0xc1, 0x8f,
		0x2f, 0x58, 0x57, 0x61, 0x4d, 0x6b, 0xdf, 0x0e, 0x1d, 0x5b, 0x08, 0xed, 0xcb, 0x99, 0x62, 0xe5,
		0x78, 0xc7, 0xb6, 0xc5, 0xf9, 0x18, 0xde, 0xb1, 0x6e, 0x9f, 0x7b, 0x16, 0xbc, 0x2f, 0xf7, 0x59,
		0x2f, 0x00, 0x81, 0x41, 0x28, 0x44, 0x01, 0x0e, 0xca, 0x76, 0x53, 0x82, 0xb6, 0x6a, 0x98, 0x01,
		0x75, 0x86, 0xf8, 0xd1, 0x41, 0x69, 0x81, 0xbe, 0x5a, 0x47, 0xfb, 0xf9, 0x7e, 0x95, 0xeb, 0x8a,
		0xad, 0x77, 0xdb, 0x8c, 0xf3, 0xa1, 0x39, 0xd7, 0xf1, 0x17, 0x2c, 0x7b, 0xfe, 0x8e, 0xce, 0xad,
		0x5f, 0x39, 0x39, 0x3d, 0xe9, 0x3a, 0x86, 0xb2, 0x42, 0xc5, 0x20, 0x36, 0xf3, 0xd8, 0xbf, 0xae,
		0x52, 0xc4, 0xcb, 0xe5, 0xb6, 0xdb, 0xb3, 0xbe, 0x23, 0x22, 0x15, 0x40, 0xed, 0xa7, 0x0d, 0x3b,
		0x63, 0xd2, 0xd0, 0x38, 0xab, 0xcc, 0x83, 0x0b, 0x5b, 0x8d, 0xf0, 0x7e, 0xfa, 0x7a, 0xb0, 0xbf,
		0x98, 0x7f, 0xa0, 0x3e, 0xe7, 0xbd, 0x51, 0xd3, 0xbe, 0x6a, 0xd7, 0x8d, 0xd0, 0x75, 0xf1, 0xf9,
		0xb8, 0xda, 0xcc, 0xdb, 0x01, 0x05, 0xd9, 0xb2, 0x42, 0x88, 0x86, 0x77, 0x22, 0x32, 0x56, 0xa1,
		0x9f, 0xcb, 0xd1, 0xf9, 0x21, 0x28, 0x34, 0x7a, 0x02, 0xb3, 0xa9, 0xb3, 0xbd, 0x19, 0x0d, 0x37,
		0xd1, 0xdb, 0xbc, 0x7d, 0xc2, 0x1d, 0xc4, 0xb3, 0xc3, 0x5f, 0xc1, 0xac, 0xee, 0x95, 0x36, 0x98,
		0x03, 0x99, 0x2b, 0x9c, 0x6f, 0x20, 0xa0, 0x6a, 0x9f, 0x05, 0x8e, 0xd2, 0xd6, 0xf0, 0x34, 0xe9,
		0x9f, 0xac, 0xe1, 0x05, 0xeb, 0xff, 0x15, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x55, 0xfb, 0x2f,
		0x1e, 0x2c, 0x07, 0x00, 0x00,
	}),
	"/library.html": embedded.NewFile("library.html", time.Now(), 242, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x44, 0x8e, 0x31, 0x4e, 0xc6, 0x30,
		0x0c, 0x85, 0x67, 0xfe, 0x53, 0x58, 0x11, 0x03, 0x0c, 0x0d, 0x3b, 0x6a, 0x3b, 0x20, 0x16, 0x24,
		0x60, 0xe9, 0x09, 0x42, 0x62, 0x5a, 0x8b, 0x34, 0x11, 0x71, 0x54, 0x84, 0x2c, 0xdf, 0x1d, 0xa5,
		0x85, 0xb2, 0xbe, 0xf7, 0xfc, 0xf9, 0xeb, 0x03, 0x6d, 0x40, 0x61, 0x30, 0x91, 0xde, 0x8a, 0x2b,
		0xdf, 0x66, 0x14, 0xb9, 0x66, 0x8c, 0x70, 0x3f, 0x80, 0x7d, 0xc8, 0xae, 0x04, 0xfb, 0x7c, 0x34,
		0x13, 0x46, 0xf4, 0x95, 0x72, 0x52, 0xbd, 0x5c, 0x89, 0x74, 0x50, 0x5c, 0x9a, 0x11, 0xec, 0x4b,
		0x4e, 0x5c, 0xb1, 0x70, 0x8b, 0x77, 0x9a, 0x8f, 0x8e, 0xf9, 0x04, 0x76, 0x54, 0x71, 0x15, 0xa1,
		0x77, 0xc0, 0xcf, 0x73, 0xfc, 0xf4, 0x08, 0xed, 0x89, 0x2a, 0xfc, 0xad, 0x78, 0xa7, 0x63, 0x10,
		0xc1, 0x14, 0x54, 0x0d, 0xac, 0x4d, 0x4a, 0xe4, 0xff, 0xa2, 0x85, 0x39, 0xf9, 0x48, 0xfe, 0x63,
		0x30, 0xbc, 0xe4, 0xaf, 0xdf, 0xe6, 0xa6, 0x2e, 0xc4, 0xb7, 0xcd, 0xdb, 0x4e, 0xb9, 0x54, 0x4a,
		0xf3, 0xab, 0x5b, 0x51, 0xb5, 0xbf, 0x0b, 0xb4, 0x8d, 0x87, 0xea, 0xce, 0xbc, 0x1c, 0xc9, 0x4f,
		0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x51, 0x33, 0x4f, 0xf2, 0x00, 0x00, 0x00,
	}),
	"/options.html": embedded.NewFile("options.html", time.Now(), 302, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x8c, 0xcf, 0x31, 0x4e, 0x2c, 0x31,
		0x0c, 0xc6, 0xf1, 0x7a, 0xe7, 0x14, 0x91, 0xab, 0xf7, 0x1a, 0x40, 0x94, 0x68, 0x92, 0x02, 0x6d,
		0x43, 0xc5, 0x11, 0x50, 0x48, 0xbe, 0xb0, 0x16, 0x1e, 0x67, 0xa4, 0x38, 0x61, 0x57, 0xab, 0xbd,
		0x3b, 0x62, 0x00, 0x41, 0x49, 0xff, 0xb7, 0x7f, 0xf6, 0x9c, 0x79, 0xb8, 0x24, 0xb1, 0x35, 0x4f,
		0xc6, 0x26, 0xa0, 0xf0, 0xb8, 0x1a, 0x57, 0x6d, 0xf3, 0x75, 0xe6, 0x11, 0xa6, 0x2d, 0xe0, 0xec,
		0xa9, 0x30, 0x24, 0x37, 0x72, 0x55, 0x5f, 0x71, 0xca, 0xf5, 0x4d, 0x3d, 0x1d, 0xa2, 0x66, 0xc1,
		0x1e, 0x25, 0x76, 0xb1, 0xfb, 0x6e, 0x56, 0xf5, 0x1f, 0x06, 0xd4, 0xfe, 0x53, 0x98, 0x76, 0xbf,
		0x57, 0x4b, 0x7c, 0x86, 0x50, 0x78, 0x50, 0x36, 0x8e, 0xc6, 0x03, 0x6e, 0xcf, 0x09, 0x77, 0x5f,
		0xc6, 0x6e, 0x66, 0x5d, 0xbb, 0x7d, 0xc7, 0x9b, 0x44, 0xce, 0x4e, 0x2b, 0x3c, 0x19, 0x8e, 0x46,
		0x4e, 0xe3, 0x02, 0x4f, 0xac, 0x6c, 0x4f, 0x99, 0x13, 0xc8, 0x2d, 0xf1, 0x28, 0xd0, 0x17, 0x3b,
		0x78, 0xba, 0xbd, 0x21, 0x37, 0xa2, 0x74, 0x78, 0x3a, 0x9f, 0xaf, 0x7e, 0x88, 0x0f, 0xe1, 0x72,
		0x21, 0x17, 0xbb, 0xd5, 0x54, 0x97, 0x55, 0x60, 0xf0, 0x54, 0x4b, 0xf9, 0xeb, 0x0f, 0xdb, 0x64,
		0xa9, 0xa9, 0xb7, 0x30, 0x7d, 0x5e, 0xfa, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xce, 0x4a, 0x2d, 0x24,
		0x2e, 0x01, 0x00, 0x00,
	}),
	"/roll_initiative.html": embedded.NewFile("roll_initiative.html", time.Now(), 327, true, []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x8c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
		0x10, 0x86, 0xe7, 0xe6, 0x29, 0x4e, 0xa7, 0x0c, 0x20, 0x41, 0x61, 0x60, 0xaa, 0xe2, 0x0c, 0xa8,
		0x0b, 0x0b, 0xef, 0x70, 0xc4, 0xe7, 0xd6, 0xc2, 0x39, 0x87, 0xe4, 0x1c, 0x52, 0x59, 0x7e, 0x77,
		0x84, 0x29, 0xcc, 0x6c, 0x37, 0xdc, 0xf7, 0xff, 0xff, 0xd7, 0x59, 0xbf, 0x82, 0xb7, 0x06, 0x9d,
		0xe7, 0x60, 0x17, 0x84, 0x28, 0xef, 0x7c, 0xb1, 0xf1, 0x53, 0x0c, 0x9e, 0x49, 0x6c, 0xe0, 0x23,
		0x3b, 0x4a, 0x41, 0x9f, 0x93, 0x6a, 0x94, 0x1b, 0x5e, 0x59, 0xf4, 0x16, 0xfb, 0x66, 0x97, 0xf3,
		0x3d, 0xcc, 0x24, 0x27, 0x86, 0xd6, 0x8b, 0xe5, 0xed, 0x0e, 0xda, 0x19, 0x0e, 0x06, 0xf6, 0xa5,
		0x34, 0xbb, 0x1a, 0x3b, 0x04, 0x5a, 0x16, 0x83, 0x81, 0xde, 0x38, 0x60, 0x9f, 0x73, 0x3b, 0xef,
		0x5f, 0x69, 0xe4, 0x52, 0x0e, 0xdd, 0x83, 0xf5, 0x6b, 0xdf, 0xec, 0x3a, 0x2f, 0x53, 0xd2, 0xdf,
		0xc7, 0x3a, 0x01, 0x41, 0x2f, 0x13, 0x1b, 0x54, 0xde, 0x14, 0x41, 0x68, 0x64, 0x83, 0x15, 0x7d,
		0x39, 0x96, 0x82, 0x30, 0xd2, 0x16, 0x58, 0x4e, 0x7a, 0x36, 0xf8, 0x84, 0x40, 0x49, 0xe3, 0x10,
		0xc7, 0x29, 0xb0, 0xb2, 0xc1, 0xe8, 0x1c, 0xc2, 0x4a, 0x21, 0xfd, 0x21, 0xe2, 0xd5, 0x93, 0xfa,
		0x95, 0xbf, 0xd1, 0xff, 0x99, 0xe5, 0xec, 0x1d, 0xf0, 0xc7, 0x55, 0x0a, 0x1e, 0x4b, 0xa9, 0x2d,
		0x2e, 0x0e, 0x69, 0xc9, 0x99, 0xc5, 0x96, 0x72, 0x95, 0xaf, 0x77, 0xf3, 0xa3, 0xf2, 0x15, 0x00,
		0x00, 0xff, 0xff, 0xc8, 0x58, 0x5a, 0xf3, 0x47, 0x01, 0x00, 0x00,
	}),
})
