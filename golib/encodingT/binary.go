package encodingT

import (
	"encoding/binary"
	"fmt"
)

func Test_binary() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//#1 LittleEndian 解析
	literV := []byte{0xF1, 0xF2, 0x00, 0x00}
	iliterV := binary.LittleEndian.Uint32(literV)
	fmt.Println("#1 LittleEndian:", iliterV)
	//i := literV[4] //边界检查
	//#1 LittleEndian 打包
	binary.LittleEndian.PutUint32(literV, 255)
	fmt.Println("#1 LittleEndian:", literV)

}
