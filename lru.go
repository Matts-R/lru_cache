package main

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string)
}

type CacheMapElement struct {
	el    *list.Element
	value string
}

type LRUCache struct {
	cache    map[string]*CacheMapElement
	capacity int
	l        list.List
}

type KeyNotFoundError struct {
	Key string
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("chave não encontrada: %s", e.Key)
}

func NewLRU(capacity int) LRUCache {
	return LRUCache{
		cache:    map[string]*CacheMapElement{},
		capacity: capacity,
		l:        list.List{},
	}
}

func (c *LRUCache) Get(key string) (string, error) {
	v, ok := c.cache[key]

	if !ok {
		return "", &KeyNotFoundError{Key: key}
	}

	c.l.MoveToFront(v.el)

	return v.value, nil
}

func (c *LRUCache) Set(key, value string) {
	v, ok := c.cache[key]

	if !ok {
		el := c.l.PushFront(key)
		c.cache[key] = &CacheMapElement{
			el:    el,
			value: value,
		}

		if c.l.Len() > c.capacity {
			back := c.l.Back()
			bk := back.Value.(string)
			c.l.Remove(back)
			delete(c.cache, bk)
		}
	} else {
		v.value = value
		c.l.MoveToFront(v.el)
	}
}

func main() {

}
