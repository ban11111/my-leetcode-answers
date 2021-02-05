package red_packet

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	pool, _ := Wrap(100, 5)
	size := unsafe.Sizeof(*pool)
	fmt.Println(size)
	size = unsafe.Sizeof(pool.total)
	fmt.Println(size)
	size = unsafe.Sizeof(pool.remain)
	fmt.Println(size)
	size = unsafe.Sizeof(pool.records)
	fmt.Println(size)
}

func TestMoneyPool_GrabBy(t *testing.T) {
	num := 10
	pool, _ := Wrap(1000, num)
	for i := 0; i < num; i++ {
		earned := pool.GrabBy(fmt.Sprintf("%d", i))
		fmt.Print(earned, "\t")
	}
	assert.Zero(t, pool.remain)
}
