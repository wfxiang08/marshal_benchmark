package cachedata

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type CacheRecord struct {
	ID   int64
	Rank int64
}

func (d *CacheRecord) Size() (s uint64) {

	s += 16
	return
}
func (d *CacheRecord) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.ID >> 0)

		buf[1+0] = byte(d.ID >> 8)

		buf[2+0] = byte(d.ID >> 16)

		buf[3+0] = byte(d.ID >> 24)

		buf[4+0] = byte(d.ID >> 32)

		buf[5+0] = byte(d.ID >> 40)

		buf[6+0] = byte(d.ID >> 48)

		buf[7+0] = byte(d.ID >> 56)

	}
	{

		buf[0+8] = byte(d.Rank >> 0)

		buf[1+8] = byte(d.Rank >> 8)

		buf[2+8] = byte(d.Rank >> 16)

		buf[3+8] = byte(d.Rank >> 24)

		buf[4+8] = byte(d.Rank >> 32)

		buf[5+8] = byte(d.Rank >> 40)

		buf[6+8] = byte(d.Rank >> 48)

		buf[7+8] = byte(d.Rank >> 56)

	}
	return buf[:i+16], nil
}

func (d *CacheRecord) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = 0 | (int64(buf[0+0]) << 0) | (int64(buf[1+0]) << 8) | (int64(buf[2+0]) << 16) | (int64(buf[3+0]) << 24) | (int64(buf[4+0]) << 32) | (int64(buf[5+0]) << 40) | (int64(buf[6+0]) << 48) | (int64(buf[7+0]) << 56)

	}
	{

		d.Rank = 0 | (int64(buf[0+8]) << 0) | (int64(buf[1+8]) << 8) | (int64(buf[2+8]) << 16) | (int64(buf[3+8]) << 24) | (int64(buf[4+8]) << 32) | (int64(buf[5+8]) << 40) | (int64(buf[6+8]) << 48) | (int64(buf[7+8]) << 56)

	}
	return i + 16, nil
}

type CacheList struct {
	data []*CacheRecord
}

func (d *CacheList) Size() (s uint64) {

	{
		l := uint64(len(d.data))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.data {

			{
				if d.data[k0] != nil {

					{
						s += (*d.data[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	return
}
func (d *CacheList) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.data))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.data {

			{
				if d.data[k0] == nil {
					buf[i+0] = 0
				} else {
					buf[i+0] = 1

					{
						nbuf, err := (*d.data[k0]).Marshal(buf[i+1:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	return buf[:i+0], nil
}

func (d *CacheList) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.data)) >= l {
			d.data = d.data[:l]
		} else {
			d.data = make([]*CacheRecord, l)
		}
		for k0 := range d.data {

			{
				if buf[i+0] == 1 {
					if d.data[k0] == nil {
						d.data[k0] = new(CacheRecord)
					}

					{
						ni, err := (*d.data[k0]).Unmarshal(buf[i+1:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.data[k0] = nil
				}
			}

			i += 1

		}
	}
	return i + 0, nil
}
