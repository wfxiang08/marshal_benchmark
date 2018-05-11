package cachedata

import (
	"github.com/vmihailenco/msgpack"
	"log"
	"testing"
	"time"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// 参考： https://github.com/andyleap/gencode/blob/master/bench/gencode_test.go

func TestMarshalDataLength(t *testing.T) {
	data := make([]*CacheRecord, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = &CacheRecord{
			ID:   int64(MaxInt),
			Rank: int64(MaxInt),
		}
	}
	encoded, _ := msgpack.Marshal(data)
	log.Printf("msgpack len: %d", len(encoded))

	list := CacheList{
		data: data,
	}
	encoded, _ = list.Marshal(nil)
	log.Printf("gencode: %d", len(encoded))

}

func BenchmarkMsgpackSerialize(b *testing.B) {
	data := make([]*CacheRecord, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = &CacheRecord{
			ID:   time.Now().Unix(),
			Rank: time.Now().Unix(),
		}
	}
	//encoded, err := msgpack.Marshal(data)
	//log.Printf("Len: %d, err: %v", len(encoded), err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(data)
	}
}

func BenchmarkGencodeSerialize(b *testing.B) {
	data := make([]*CacheRecord, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = &CacheRecord{
			ID:   time.Now().Unix(), //
			Rank: time.Now().Unix(),
		}
	}

	list := CacheList{
		data: data,
	}
	//encoded, err := list.Marshal(nil)
	//log.Printf("Len: %d, err: %v", len(encoded), err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = list.Marshal(nil)
	}
}

func BenchmarkGencodeDeSerialize(b *testing.B) {
	data := make([]*CacheRecord, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = &CacheRecord{
			ID:   time.Now().Unix(), //
			Rank: time.Now().Unix(),
		}
	}

	list := CacheList{
		data: data,
	}
	encoded, _ := list.Marshal(nil)
	//log.Printf("Len: %d, err: %v", len(encoded), err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n := &CacheList{}
		_, _ = n.Unmarshal(encoded)
	}
}

func BenchmarkMsgpackDeSerialize(b *testing.B) {
	data := make([]*CacheRecord, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = &CacheRecord{
			ID:   time.Now().Unix(), //
			Rank: time.Now().Unix(),
		}
	}

	encoded, _ := msgpack.Marshal(data)
	//encoded, err := msgpack.Marshal(data)
	//log.Printf("Len: %d, err: %v", len(encoded), err)
	//
	//var items []*CacheRecord
	//err = msgpack.Unmarshal(encoded, &items)
	//log.Printf("items: %d, err: %v", len(items), err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var items []*CacheRecord
		_ = msgpack.Unmarshal(encoded, &items)
	}
}
