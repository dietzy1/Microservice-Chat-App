package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	cache := New(2)

	// Test adding a value to the cache
	cache.Set("key1", "value1")
	val, ok := cache.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val)

	// Test updating a value in the cache
	cache.Set("key1", "value2")
	val, ok = cache.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value2", val)

	// Test adding a second value to the cache
	cache.Set("key2", "value2")
	val, ok = cache.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "value2", val)

	// Test removing the least recently used value from the cache
	cache.Set("key3", "value3")
	val, ok = cache.Get("key1")
	assert.False(t, ok)
	assert.Equal(t, "", val)
}
