package util

import (
	"strconv"
)

// sBox
var S_Box = [][]int{
	{0x63, 0x7C, 0x77, 0x7B, 0xF2, 0x6B, 0x6F, 0xC5,
		0x30, 0x01, 0x67, 0x2B, 0xFE, 0xD7, 0xAB, 0x76},
	{0xCA, 0x82, 0xC9, 0x7D, 0xFA, 0x59, 0x47, 0xF0,
		0xAD, 0xD4, 0xA2, 0xAF, 0x9C, 0xA4, 0x72, 0xC0},
	{0xB7, 0xFD, 0x93, 0x26, 0x36, 0x3F, 0xF7, 0xCC,
		0x34, 0xA5, 0xE5, 0xF1, 0x71, 0xD8, 0x31, 0x15},
	{0x04, 0xC7, 0x23, 0xC3, 0x18, 0x96, 0x05, 0x9A,
		0x07, 0x12, 0x80, 0xE2, 0xEB, 0x27, 0xB2, 0x75},
	{0x09, 0x83, 0x2C, 0x1A, 0x1B, 0x6E, 0x5A, 0xA0,
		0x52, 0x3B, 0xD6, 0xB3, 0x29, 0xE3, 0x2F, 0x84},
	{0x53, 0xD1, 0x00, 0xED, 0x20, 0xFC, 0xB1, 0x5B,
		0x6A, 0xCB, 0xBE, 0x39, 0x4A, 0x4C, 0x58, 0xCF},
	{0xD0, 0xEF, 0xAA, 0xFB, 0x43, 0x4D, 0x33, 0x85,
		0x45, 0xF9, 0x02, 0x7F, 0x50, 0x3C, 0x9F, 0xA8},
	{0x51, 0xA3, 0x40, 0x8F, 0x92, 0x9D, 0x38, 0xF5,
		0xBC, 0xB6, 0xDA, 0x21, 0x10, 0xFF, 0xF3, 0xD2},
	{0xCD, 0x0C, 0x13, 0xEC, 0x5F, 0x97, 0x44, 0x17,
		0xC4, 0xA7, 0x7E, 0x3D, 0x64, 0x5D, 0x19, 0x73},
	{0x60, 0x81, 0x4F, 0xDC, 0x22, 0x2A, 0x90, 0x88,
		0x46, 0xEE, 0xB8, 0x14, 0xDE, 0x5E, 0x0B, 0xDB},
	{0xE0, 0x32, 0x3A, 0x0A, 0x49, 0x06, 0x24, 0x5C,
		0xC2, 0xD3, 0xAC, 0x62, 0x91, 0x95, 0xE4, 0x79},
	{0xE7, 0xC8, 0x37, 0x6D, 0x8D, 0xD5, 0x4E, 0xA9,
		0x6C, 0x56, 0xF4, 0xEA, 0x65, 0x7A, 0xAE, 0x08},
	{0xBA, 0x78, 0x25, 0x2E, 0x1C, 0xA6, 0xB4, 0xC6,
		0xE8, 0xDD, 0x74, 0x1F, 0x4B, 0xBD, 0x8B, 0x8A},
	{0x70, 0x3E, 0xB5, 0x66, 0x48, 0x03, 0xF6, 0x0E,
		0x61, 0x35, 0x57, 0xB9, 0x86, 0xC1, 0x1D, 0x9E},
	{0xE1, 0xF8, 0x98, 0x11, 0x69, 0xD9, 0x8E, 0x94,
		0x9B, 0x1E, 0x87, 0xE9, 0xCE, 0x55, 0x28, 0xDF},
	{0x8C, 0xA1, 0x89, 0x0D, 0xBF, 0xE6, 0x42, 0x68,
		0x41, 0x99, 0x2D, 0x0F, 0xB0, 0x54, 0xBB, 0x16}}

// inv_sBox
var Inv_S_Box = [][]int{
	{0x52, 0x09, 0x6A, 0xD5, 0x30, 0x36, 0xA5, 0x38,
		0xBF, 0x40, 0xA3, 0x9E, 0x81, 0xF3, 0xD7, 0xFB},
	{0x7C, 0xE3, 0x39, 0x82, 0x9B, 0x2F, 0xFF, 0x87,
		0x34, 0x8E, 0x43, 0x44, 0xC4, 0xDE, 0xE9, 0xCB},
	{0x54, 0x7B, 0x94, 0x32, 0xA6, 0xC2, 0x23, 0x3D,
		0xEE, 0x4C, 0x95, 0x0B, 0x42, 0xFA, 0xC3, 0x4E},
	{0x08, 0x2E, 0xA1, 0x66, 0x28, 0xD9, 0x24, 0xB2,
		0x76, 0x5B, 0xA2, 0x49, 0x6D, 0x8B, 0xD1, 0x25},
	{0x72, 0xF8, 0xF6, 0x64, 0x86, 0x68, 0x98, 0x16,
		0xD4, 0xA4, 0x5C, 0xCC, 0x5D, 0x65, 0xB6, 0x92},
	{0x6C, 0x70, 0x48, 0x50, 0xFD, 0xED, 0xB9, 0xDA,
		0x5E, 0x15, 0x46, 0x57, 0xA7, 0x8D, 0x9D, 0x84},
	{0x90, 0xD8, 0xAB, 0x00, 0x8C, 0xBC, 0xD3, 0x0A,
		0xF7, 0xE4, 0x58, 0x05, 0xB8, 0xB3, 0x45, 0x06},
	{0xD0, 0x2C, 0x1E, 0x8F, 0xCA, 0x3F, 0x0F, 0x02,
		0xC1, 0xAF, 0xBD, 0x03, 0x01, 0x13, 0x8A, 0x6B},
	{0x3A, 0x91, 0x11, 0x41, 0x4F, 0x67, 0xDC, 0xEA,
		0x97, 0xF2, 0xCF, 0xCE, 0xF0, 0xB4, 0xE6, 0x73},
	{0x96, 0xAC, 0x74, 0x22, 0xE7, 0xAD, 0x35, 0x85,
		0xE2, 0xF9, 0x37, 0xE8, 0x1C, 0x75, 0xDF, 0x6E},
	{0x47, 0xF1, 0x1A, 0x71, 0x1D, 0x29, 0xC5, 0x89,
		0x6F, 0xB7, 0x62, 0x0E, 0xAA, 0x18, 0xBE, 0x1B},
	{0xFC, 0x56, 0x3E, 0x4B, 0xC6, 0xD2, 0x79, 0x20,
		0x9A, 0xDB, 0xC0, 0xFE, 0x78, 0xCD, 0x5A, 0xF4},
	{0x1F, 0xDD, 0xA8, 0x33, 0x88, 0x07, 0xC7, 0x31,
		0xB1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xEC, 0x5F},
	{0x60, 0x51, 0x7F, 0xA9, 0x19, 0xB5, 0x4A, 0x0D,
		0x2D, 0xE5, 0x7A, 0x9F, 0x93, 0xC9, 0x9C, 0xEF},
	{0xA0, 0xE0, 0x3B, 0x4D, 0xAE, 0x2A, 0xF5, 0xB0,
		0xC8, 0xEB, 0xBB, 0x3C, 0x83, 0x53, 0x99, 0x61},
	{0x17, 0x2B, 0x04, 0x7E, 0xBA, 0x77, 0xD6, 0x26,
		0xE1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0C, 0x7D},
}

// column_box
var column_box = [][]string{
	{"02", "03", "01", "01"},
	{"01", "02", "03", "01"},
	{"01", "01", "02", "03"},
	{"03", "01", "01", "02"},
}

var inv_columnBox = [][]string{
	{"0e", "0b", "0d", "09"},
	{"09", "0e", "0b", "0d"},
	{"0d", "09", "0e", "0b"},
	{"0b", "0d", "09", "0e"},
} //逆列混合

var Rcon = [][]string{
	{"01", "00", "00", "00"},
	{"02", "00", "00", "00"},
	{"04", "00", "00", "00"},
	{"08", "00", "00", "00"},
	{"10", "00", "00", "00"},
	{"20", "00", "00", "00"},
	{"40", "00", "00", "00"},
	{"80", "00", "00", "00"},
	{"1b", "00", "00", "00"},
	{"36", "00", "00", "00"},
}

// 子密钥扩展
func keyExpension(key string) (data [40][4]string) {
	keyToByte := []byte(key)
	tempList := [44][4]int{}
	// 循环插入数组
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			tempList[i][j] = int(keyToByte[(i*4)+j])
		}
	}
	for i := 4; i < 44; i++ {
		roundConList := Rcon[i/4-1] //代入轮常量
		if i%4 != 0 {
			wi4 := tempList[i-4]
			wi1 := tempList[i-1]
			for j := 0; j < 4; j++ {
				tempList[i][j] = wi1[j] ^ wi4[j]
				data[i-4][j] = strconv.FormatInt(int64(wi1[j]^wi4[j]), 16)
			}
		} else {
			wi4 := tempList[i-4]
			// 字循环
			tempWi1 := append(tempList[i-1][1:], tempList[i-1][:1]...)
			for j := 0; j < 4; j++ {
				// 字节代换Sbox
				fixNum := string(strconv.FormatInt(int64(tempWi1[j]), 16))
				if len(fixNum) < 2 {
					fixNum = "0" + fixNum
				}
				x, _ := strconv.ParseInt(string(fixNum[0]), 16, 32)
				y, _ := strconv.ParseInt(string(fixNum[1]), 16, 32)
				tempWi1[j] = S_Box[x][y]
				// 轮常量异或
				tempNum, _ := strconv.ParseInt(roundConList[j], 16, 32)
				tempList[i][j] = wi4[j] ^ (int(tempNum) ^ tempWi1[j])
				data[i-4][j] = strconv.FormatInt(int64(wi4[j])^(tempNum^int64(tempWi1[j])), 16)
			}
		}
	}
	return data
}

// 字符串转bit
func strToBit(str string) map[int]string {
	// 定义一个二维数组去存放128bit
	data := make(map[int]string, 128)
	newData := []byte(str)
	// tempStr := ""
	for _, v := range newData {
		conv := strconv.FormatInt(int64(int(v)), 2)
		for len(conv) < 8 {
			conv = "0" + conv
		}
	}
	return data
}

// 初始变换
func xorKeys(plainText string, cipherText string) []string {
	plainList := []byte(plainText)
	cipherList := []uint8(cipherText)
	var data []string = []string{"00", "00", "00", "00", "00", "00", "00",
		"00", "00", "00", "00", "00", "00", "00", "00", "00"}
	for i := 0; i < len(plainList); i++ {
		tempData := strconv.FormatInt(int64(plainList[i]^cipherList[i]), 16)
		if len(tempData) < 2 {
			tempData = "0" + tempData
		}
		data[i] = tempData
	}
	return data
}

// 字节代换
func subBytes(subList [4][4]string, method string) [4][4]string {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// 这里拿到左右两个参数的byte,之后string打印, 然后再转为int, 不得不说有点麻烦
			x, _ := strconv.ParseInt(string(subList[i][j][0]), 16, 64)
			y, _ := strconv.ParseInt(string(subList[i][j][1]), 16, 64)
			if method == "encode" {
				subList[i][j] = strconv.FormatInt((int64(S_Box[x][y])), 16)
			} else {
				subList[i][j] = strconv.FormatInt((int64(Inv_S_Box[x][y])), 16)
			}
		}
	}
	return subList
}

// 旋转数组
func turnList(dataList [4][4]string) (data [4][4]string) {
	// 向右旋转数组 将第i列等于第i行
	for k1, v1 := range dataList {
		for k2, _ := range v1 {
			data[k1][k2] = dataList[k2][k1]
		}
	}
	return
}

// 行移位
func movRows(dataList [4][4]string, direction string) [4][4]string {
	data := [4][4]string{}
	// 旋转数组
	tempList := turnList(dataList)
	// 每行移动i位
	if direction == "left" {
		for j := 0; j < 4; j++ {
			for i := 0; i < 4; i++ {
				data[j][i] = append(tempList[j][j:], tempList[j][:j]...)[i]
			}
		}
	} else {
		for j := 0; j < 4; j++ {
			for i := 0; i < 4; i++ {
				data[j][i] = append(tempList[j][4-j:], tempList[j][:4-j]...)[i]
			}
		}
	}
	data = turnList(data)
	return data
}

// 列混合在有限域F^2上运算
func mixcolumns(columnData [4][4]string) (data [4][4]string) {
	for k1, v1 := range columnData {
		for k2, v2 := range column_box {
			var res int64
			for i := 0; i < len(v2); i++ {
				tempData, _ := strconv.ParseInt(v1[i], 16, 64)
				switch v2[i] {
				case "01":
					res = res ^ tempData
				case "02":
					res = res ^ GFMul2(tempData)
				case "03":
					res = res ^ GFMul3(tempData)
				}
				strData := strconv.FormatInt(res, 16)
				for len(strData) < 2 {
					strData = "0" + strData
				}
				data[k1][k2] = strData
			}
		}
	}
	return
}

// 轮秘钥加
func addRoundKey(oriData [4][4]string, keyList [40][4]string, round int) (data [4][4]string) {
	xorList := keyList[4*round : 4*round+4][:]
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			lData, _ := strconv.ParseInt(oriData[i][j], 16, 32)
			rData, _ := strconv.ParseInt(xorList[i][j], 16, 32) //处理两个数组元素
			tempData := strconv.FormatInt(lData^rData, 16)
			if len(tempData) < 2 {
				tempData = "0" + tempData
			}
			data[i][j] = tempData
		}
	}
	return
}

// 9轮迭代
func goround(xorList [4][4]string, keys [40][4]string, round int) [4][4]string {
	data := subBytes(xorList, "encode")
	data = movRows(data, "left")
	data = mixcolumns(data)
	data = addRoundKey(data, keys, round)
	return data
}

// 按行分组
func sortByColumns(oriData []string) (data [4][4]string) {
	// 四组对应成行
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			data[i][j] = oriData[i*4+j]
		}
	}
	return
}

// 第十轮
func finalEncodeRound(xorList [4][4]string, keys [40][4]string) [4][4]string {
	data := subBytes(xorList, "encode")
	data = movRows(data, "right")
	data = addRoundKey(data, keys, 9)
	return data
}

// 加密
func EnAES(plainText string, key string) (data string) {
	keys := keyExpension(key)
	xorList := xorKeys(plainText, key)
	tempList := sortByColumns(xorList)
	for i := 0; i < 9; i++ {
		tempList = goround(tempList, keys, i)
	}
	tempList = finalEncodeRound(tempList, keys)
	for _, v1 := range tempList {
		for _, v2 := range v1 {
			data += v2
		}
	}
	return // 打印密文
}

// 解密第一轮
func reFirstDecodeRound(cipherData [4][4]string, keyList [40][4]string) (data [4][4]string) {
	data = addRoundKey(cipherData, keyList, 9) // 轮密钥加
	data = movRows(data, "right")              // 逆行移位
	data = subBytes(data, "decode")            // 逆字节替换
	return
}

// 函数思想 -> 嵌套
func GFMul2(cipherData int64) (data int64) {
	tempData := strconv.FormatInt(int64(cipherData), 2)
	if len(tempData) < 8 { // 当tempData长度小于7 即a7不等于1
		for len(tempData) < 8 {
			tempData = "0" + tempData
		}
		tempData = tempData[1:] + "0"
		data, _ = strconv.ParseInt(tempData, 2, 64)
	} else { //长度等于8
		tempData = tempData[1:] + "0"
		data, _ := strconv.ParseInt(tempData, 2, 64)
		data = data ^ 27
	}
	return data
}

func GFMul3(cipher int64) (data int64) {
	return cipher ^ GFMul2(cipher)
}

func GFMul4(cipherData int64) (data int64) {
	return GFMul2(GFMul2(cipherData))
}

func GFMul8(cipherData int64) (data int64) {
	return GFMul2(GFMul4(cipherData))
}

func GFMul9(cipherData int64) (data int64) {
	return cipherData ^ GFMul8(cipherData)
}

func GFMulb(cipherData int64) (data int64) {
	return GFMul2(cipherData) ^ GFMul9(cipherData)
}

func GFMuld(cipherData int64) (data int64) { // 13 = 11 + 2
	return GFMul2(cipherData) ^ GFMulb(cipherData)
}

func GFMule(cipherData int64) (data int64) { // 14 = 13 + 1
	return cipherData ^ GFMuld(cipherData)
}

// 逆列混合
func invMixColumns(cipherData [4][4]string) (data [4][4]string) {
	for k1, v1 := range cipherData {
		for k2, v2 := range inv_columnBox {
			var res int64
			for i := 0; i < len(v2); i++ {
				tempData, _ := strconv.ParseInt(v1[i], 16, 64)
				switch v2[i] {
				case "09":
					res = res ^ GFMul9(tempData)
				case "0b":
					res = res ^ GFMulb(tempData)
				case "0d":
					res = res ^ GFMuld(tempData)
				case "0e":
					res = res ^ GFMule(tempData)
				}
			}
			tempData := strconv.FormatInt(res, 16)
			for len(tempData) < 2 {
				tempData = "0" + tempData
			}

			data[k1][k2] = tempData
		}
	}
	return
}

// 解密最后异或初始秘钥
func xorReginKey(cipherData [4][4]string, key string) (data []uint8) {
	keyData := []uint8(key)
	for k1, v1 := range cipherData {
		for k2, v2 := range v1 {
			tempData, _ := strconv.ParseInt(v2, 16, 64)
			final_data := (int64(keyData[4*k1+k2]) ^ tempData)
			data = append(data, uint8(final_data))
		}
	}
	return
}

// 解密
func deAES(cipherText string, key string) (data string) {
	tempList := [32]string{}
	for k, v := range cipherText {
		tempList[k] = string(v)
	}
	cipherData := [4][4]string{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			cipherData[i][j] = tempList[i*8+j*2] + tempList[i*8+j*2+1]
		}
	}
	keyList := keyExpension(key)                         // 扩展后的密钥
	cipherData = reFirstDecodeRound(cipherData, keyList) // 单独四组中每项解密 第一轮

	for i := 8; i >= 0; i-- {
		// 这里拿到第一轮处理后的cipher 进行逆9轮迭代
		cipherData = addRoundKey(cipherData, keyList, i) // 轮密钥加
		cipherData = invMixColumns(cipherData)           //逆列混合
		cipherData = movRows(cipherData, "right")        // 逆行移位
		cipherData = subBytes(cipherData, "decode")      // 逆字节替换

	}

	// 初始秘钥加
	finalData := xorReginKey(cipherData, key)
	data = string(finalData)
	return
}
