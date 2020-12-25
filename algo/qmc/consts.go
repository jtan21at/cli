package qmc

var oggPublicHeader1 = []byte{
	0x4f, 0x67, 0x67, 0x53, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0x1e, 0x01, 0x76, 0x6f, 0x72,
	0x62, 0x69, 0x73, 0x00, 0x00, 0x00, 0x00, 0x02, 0x44, 0xac, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xee, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0xb8, 0x01, 0x4f, 0x67, 0x67, 0x53, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xff, 0xff}

var oggPublicHeader2 = []byte{
	0x03, 0x76, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x2c, 0x00, 0x00, 0x00, 0x58, 0x69, 0x70, 0x68, 0x2e,
	0x4f, 0x72, 0x67, 0x20, 0x6c, 0x69, 0x62, 0x56, 0x6f, 0x72, 0x62, 0x69, 0x73, 0x20, 0x49, 0x20,
	0x32, 0x30, 0x31, 0x35, 0x30, 0x31, 0x30, 0x35, 0x20, 0x28, 0xe2, 0x9b, 0x84, 0xe2, 0x9b, 0x84,
	0xe2, 0x9b, 0x84, 0xe2, 0x9b, 0x84, 0x29, 0xff, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00, 0x54,
	0x49, 0x54, 0x4c, 0x45, 0x3d}

var oggPublicConfidence1 = []uint{
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 0, 0,
	0, 0, 9, 9, 9, 9, 0, 0, 0, 0, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 6, 3, 3, 3, 3, 6, 6, 6, 6,
	3, 3, 3, 3, 6, 6, 6, 6, 6, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 9, 9, 9, 9,
	0, 0, 0, 0}

var oggPublicConfidence2 = []uint{
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 0, 1, 3, 3, 0, 1, 3, 3, 3,
	3, 3, 3, 3, 3}

var (
	defaultKey256Mask44 = []byte{
		0xde, 0x51, 0xfa, 0xc3, 0x4a, 0xd6, 0xca, 0x90,
		0x7e, 0x67, 0x5e, 0xf7, 0xd5, 0x52, 0x84, 0xd8,
		0x47, 0x95, 0xbb, 0xa1, 0xaa, 0xc6, 0x66, 0x23,
		0x92, 0x62, 0xf3, 0x74, 0xa1, 0x9f, 0xf4, 0xa0,
		0x1d, 0x3f, 0x5b, 0xf0, 0x13, 0x0e, 0x09, 0x3d,
		0xf9, 0xbc, 0x00, 0x11}
	headerFlac = []byte{'f', 'L', 'a', 'C'}
	headerOgg  = []byte{'O', 'g', 'g', 'S'}
)
var key256MappingAll [][]int //[idx256][idx128]idx44
var key256Mapping128to44 map[int]int

func init() {
	{ // init all mapping
		key256MappingAll = make([][]int, 256)
		for i := 0; i < 128; i++ {
			realIdx := (i*i + 27) % 256
			if key256MappingAll[realIdx] == nil {
				key256MappingAll[realIdx] = []int{i}
			} else {
				key256MappingAll[realIdx] = append(key256MappingAll[realIdx], i)
			}
		}
	}
	{ // init
		key256Mapping128to44 = make(map[int]int, 128)
		idx44 := 0
		for _, all128 := range key256MappingAll {
			if all128 != nil {
				for _, _i128 := range all128 {
					key256Mapping128to44[_i128] = idx44
				}
				idx44++
			}
		}
	}

}
