package tsct

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"sync"
	"time"
)

type Data struct {
	ID             int64
	Secret         string
	Time           *time.Time
	old, now, next string
	wg             sync.WaitGroup
}

var m = make(map[int64]*Data)

func LoadSecret(id int64, secret string) {
	m[id] = &Data{
		ID:     id,
		Secret: secret,
	}
}

func Find(id int64) *Data {
	d, ok := m[id]
	if !ok {
		return nil
	}
	return d
}

func (d *Data) GET() *Data {
	d.Time = getTime()

	d.wg.Add(3)
	go d.Old()
	go d.Now()
	go d.Next()
	return d
}

func (d *Data) Old() {
	d.old = d.gen(-30 * time.Second)
	d.wg.Done()
}

func (d *Data) Now() {
	d.now = d.gen(0)
	d.wg.Done()
}

func (d *Data) Next() {
	d.next = d.gen(30 * time.Second)
	d.wg.Done()
}

func (d *Data) gen(t time.Duration) (code string) {
	i := strconv.Itoa(gens(d.Secret, d.Time.Add(t)))
	switch len(i) {
	case 0:
		code = "000000"
	case 1:
		code = "00000" + i
	case 2:
		code = "0000" + i
	case 3:
		code = "000" + i
	case 4:
		code = "00" + i
	case 5:
		code = "0" + i
	default:
		code = i
	}
	return
}

func (d *Data) String() string {
	d.wg.Wait()
	return `_` + d.Time.Format("2006/01/02 15:04:05 MST") + `_

~` + d.old + `~  \|  ` +
		"`" + d.now + "`" + `  \|  ` + "`" + d.next + "`"
}

func gens(secret string, t time.Time) int {
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
	return c
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
