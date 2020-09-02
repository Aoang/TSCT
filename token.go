package tsct

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

type data struct {
	ID             int64
	Secret         string
	Time           *time.Time
	old, now, next string
	wg             sync.WaitGroup
}

var m = make(map[int64]*data)

func LoadSecret(id int64, secret string) {
	m[id] = &data{
		ID:     id,
		Secret: secret,
	}
}

func Find(id int64) *data {
	d, ok := m[id]
	if !ok {
		return nil
	}
	return d
}

func (d *data) GET() *data {
	d.Time = getTime()

	d.wg.Add(3)
	go d.oldCode()
	go d.nowCode()
	go d.nextCode()
	return d
}

func (d *data) oldCode() {
	d.old = Token(d.Secret, d.Time.Add(-30*time.Second))
	d.wg.Done()
}

func (d *data) nowCode() {
	d.now = Token(d.Secret, d.Time.Add(0))
	d.wg.Done()
}

func (d *data) nextCode() {
	d.next = Token(d.Secret, d.Time.Add(30*time.Second))
	d.wg.Done()
}

func (d *data) String() string {
	d.wg.Wait()
	return `_` + d.Time.Format("2006/01/02 15:04:05 MST") + `_

~` + d.old + `~  \|  ` +
		"`" + d.now + "`" + `  \|  ` + "`" + d.next + "`"
}

func Token(secret string, t time.Time) string {
	ss, _ := hex.DecodeString(secret)
	b := sha256.Sum256(append(ss, times(t)...))

	d := make([]byte, 0, 64)
	for i := range b {
		d = append(d, b[i]>>4, b[i]&0x0f)
	}

	c := 0
	for i := 0; i < 6; i++ {
		t := d[i+1]
		t += d[i+1+1*7]
		t += d[i+1+2*7]
		t += d[i+1+3*7]
		t += d[i+1+4*7]
		t += d[i+1+5*7]
		t += d[i+1+6*7]
		t += d[i+1+7*7]
		t += d[i+1+8*7]
		c = c*10 + int(t%10)
	}
	return zeroPadding(c)
}

func zeroPadding(i int) string {
	s := strconv.Itoa(i)
	switch len(s) {
	case 1:
		s = "00000" + s
	case 2:
		s = "0000" + s
	case 3:
		s = "000" + s
	case 4:
		s = "00" + s
	case 5:
		s = "0" + s
	}
	return s
}

func times(t time.Time) []byte {
	b := []byte(":00")
	if t.Second() > 30 {
		b = []byte(":30")
	}

	return append([]byte(t.Format("2006-01-02 15:04")), b...)
}

func getTime() *time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t := time.Now().In(loc)
	return &t
}
