package concurrency

import (
	"crypto/sha1"
	"sync"
)

type Shard struct {
	sync.RWMutex
	m map[string]interface{}
}

type ShardedMap []*Shard

func NewShardedMap(nshards uint) ShardedMap {
	shards := make([]*Shard, nshards)

	var i uint
	for ; i < nshards; i++ {
		shardMap := make(map[string]interface{})
		shards[i] = &Shard{m: shardMap}
	}

	return shards
}

func (shardedMap ShardedMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key))
	// For more than 255 shards use:
	// hash := int(sum[13]) << 8 | int(sum[17])
	hash := int(checksum[17])
	return hash % len(shardedMap)
}

func (shardedMap ShardedMap) getShard(key string) *Shard {
	index := shardedMap.getShardIndex(key)
	return shardedMap[index]
}

func (shardedMap ShardedMap) Get(key string) interface{} {
	shard := shardedMap.getShard(key)
	shard.RLock()
	defer shard.RUnlock()
	return shard.m[key]
}

func (shardedMap ShardedMap) Set(key string, value interface{}) {
	shard := shardedMap.getShard(key)
	shard.Lock()
	defer shard.Unlock()
	shard.m[key] = value
}

func (shardedMap ShardedMap) Keys() []string {
	keys := make([]string, 0)
	mutex := sync.Mutex{}

	var wg sync.WaitGroup
	wg.Add(len(shardedMap))

	for _, shard := range shardedMap {
		go func(s *Shard) {
			s.RLock()
			defer s.RUnlock()
			defer wg.Done()

			for key := range s.m {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}
		}(shard)
	}

	wg.Wait()

	return keys
}
