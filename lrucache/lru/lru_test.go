package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int { return len(s) }

func TestCache_Get(t *testing.T) {
	type argument struct {
		maxBytes  int64
		onEvicted func(string, Value)
		key       string
		getKey    string
		value     Value
	}
	type result struct {
		value Value
		ok    bool
	}
	tests := []struct {
		name   string
		arg    argument
		result result
	}{
		{name: "test1", arg: argument{maxBytes: 0, onEvicted: nil, key: "key1", value: String("value1"), getKey: "key1"}, result: result{String("value1"), true}},
		{name: "test2", arg: argument{maxBytes: 100, onEvicted: nil, key: "key2", value: String("value2"), getKey: "nokey2"}, result: result{nil, false}},
		{name: "test3", arg: argument{maxBytes: 1, onEvicted: func(key string, value Value) {}, key: "key3", value: String("value3"), getKey: "nokey3"}, result: result{nil, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.arg.maxBytes, tt.arg.onEvicted)
			c.Add(tt.arg.key, tt.arg.value)
			if tt.name == "test1" {
				c.Add(tt.arg.key, tt.arg.value)
			}

			if gotValue, gotOk := c.Get(tt.arg.getKey); gotOk == tt.result.ok && !reflect.DeepEqual(gotValue, tt.result.value) {
				t.Errorf("Cache.Get() gotValue = %v, want %v", gotValue, tt.result.value)
			}
		})
	}
}
