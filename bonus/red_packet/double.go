package red_packet

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	maxMoney = 20000 // 单位分
)

type record struct {
	earned uint16
	name   string
}

type MoneyPool struct {
	lock    sync.Mutex
	rand    *rand.Rand
	total   uint16
	remain  uint16
	records []*record //记录抢到的结果
}

func Wrap(amount uint16, num int) (*MoneyPool, error) {
	if amount > maxMoney {
		return nil, errors.New("呵呵")
	}
	return &MoneyPool{
		rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
		total:   amount,
		remain:  amount,
		records: make([]*record, 0, num),
	}, nil
}

func (p *MoneyPool) GrabBy(name string) uint16 {
	p.lock.Lock()
	defer p.lock.Unlock()

	remainNum := cap(p.records) - len(p.records)
	earned := p.remain
	if remainNum > 1 {
		randMax := (int(p.remain) / remainNum) * 2
		if randMax <= 0 {
			randMax = 1
		}
		earned = uint16(p.rand.Intn(randMax))
		if earned == 0 {
			earned++
		}
	}
	p.queuedCalculate(earned, name)
	return earned
}

func (p *MoneyPool) queuedCalculate(earned uint16, name string) {
	p.remain -= earned
	p.records = append(p.records, &record{
		earned: earned,
		name:   name,
	})
}
