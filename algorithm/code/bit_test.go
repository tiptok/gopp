package code

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber(t *testing.T) {
	var table []struct {
		in     []int
		except int
	} = []struct {
		in     []int
		except int
	}{
		{in: []int{2, 2, 1}, except: 1},
		{in: []int{4, 1, 2, 1, 2}, except: 4},
	}

	for i := range table {
		out := singleNumber(table[i].in)
		if out != table[i].except {
			t.Fatalf("in:%v except:%v out:%v", table[i].in, table[i].except, out)
		}
	}
}

func Test_singleNumber2(t *testing.T) {
	var table []struct {
		in     []int
		except int
	} = []struct {
		in     []int
		except int
	}{
		{in: []int{2, 2, 2, 1}, except: 1},
		{in: []int{4, 1, 1, 2, 1, 2, 2}, except: 4},
	}

	for i := range table {
		out := singleNumber2(table[i].in)
		if out != table[i].except {
			t.Fatalf("in:%v except:%v out:%v", table[i].in, table[i].except, out)
		}
	}
}

func Test_singleNumber3(t *testing.T) {
	var table []struct {
		in     []int
		except []int
	} = []struct {
		in     []int
		except []int
	}{
		{in: []int{2, 2, 4, 1}, except: []int{4,1}},
		{in: []int{4, 1, 1, 2, 1, 2, 3}, except:[]int{4,3}},
	}

	for i := range table {
		out := singleNumber3(table[i].in)
		if reflect.DeepEqual(out ,table[i].except) {
			t.Fatalf("in:%v except:%v out:%v", table[i].in, table[i].except, out)
		}
	}
}


func Test_hammingWeight(t *testing.T){
	var table []struct {
		in     uint32
		except int
	} = []struct {
		in     uint32
		except int
	}{
		{in: 11, except:3},
		{in: 1024, except: 1},
	} 
	for i := range table {
		out := hammingWeight(table[i].in)
		if out!=table[i].except{
			t.Fatalf("in:%v except:%v out:%v", table[i].in, table[i].except, out)
		}
	}
}


func Test_countBits(t *testing.T) {
	var table []struct {
		in     int
		except []int
	} = []struct {
		in     int
		except []int
	}{
		{in: 5, except: []int{0,1,1,2,1,2}},
		{in: 4, except:[]int{0,1,1,2,1}},
	}

	for i := range table {
		out := countBits(table[i].in)
		if !assert.ElementsMatch(t,out,table[i].except){
			t.Fatalf("in:%v except:%v out:%v", table[i].in, table[i].except, out)
		}
	}
}


// 136. Single Number I
func singleNumber(nums []int) int {
	var res int
	for i := range nums {
		res ^= nums[i]
	}
	return res
}
// 137. Single Number II
func singleNumber2(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := range nums {
			sum += (nums[j] >> i) & 1
		}
		res |= (sum % 3) << i
	}
	return res
}

// 260. Single Number III
func singleNumber3(nums []int) []int {
	diff :=0
	for j := range nums {
		diff ^= nums[j]
	}
	res :=[]int{diff,diff}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
    diff=(diff&(diff-1))^diff
	for j := range nums {
		if diff & nums[j] ==0{
			res[0] ^=nums[j]
		}else{
			res[1] ^=nums[j]
		}
	}
	return res
}

// 191. Number of 1 Bits (汉明重量)
func hammingWeight(num uint32) int {
    var res = 0
	for num>0{
		if num & 1 >0{
			res ++
		}
		num = num >> 1
	}
	return res
}

// 338: countBits
func countBits(num int) []int {
	res :=make([]int,num+1)
	for i:=0;i<=num;i++{
		res[i] = hammingWeight(uint32(i))
	}
	return res
}

// 190. Reverse Bits
func reverseBits(num uint32) uint32 {
	var res uint32
	var bitIndex int=31
	for num !=0{
		res += (num & 1) << uint32(bitIndex)
		num = num >> 1
		bitIndex --
	}
	return res
}
// 201.bitwise-and-of-numbers-range
func rangeBitwiseAnd(m int, n int) int {
	// m 5 1 0 1
    //   6 1 1 0
    // n 7 1 1 1
    // 把可能包含0的全部右移变成
    // m 5 1 0 0
    //   6 1 0 0
    // n 7 1 0 0
    // 所以最后结果就是m<<count
	var bitIndex int
	for m!=n{
		m = m >> 1
		n = n >> 1
		bitIndex ++
	}
	return m << bitIndex
}