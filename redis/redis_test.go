package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func setupClient() Client {
	conf := Config{
		Host:               "localhost",
		Port:               6379,
		ConnectionTimeout:  5000,
		MaxIdleConnections: 10,
	}

	client, _ := NewClient(conf)
	return client
}

func TestClient(t *testing.T) {
	client := setupClient()
	assert.NotNil(t, client.Client())
}

func TestSet(t *testing.T) {
	client := setupClient()
	defer client.Close()

	t.Run("Set String", func(t *testing.T) {
		err := client.Set("test_key", "test_value", time.Minute)
		assert.NoError(t, err)
	})

	t.Run("Set Struct", func(t *testing.T) {
		type TestStruct struct {
			Field1 string
			Field2 int
		}

		testStruct := TestStruct{Field1: "value1", Field2: 42}
		err := client.SetStruct("struct_key", testStruct, time.Minute)
		assert.NoError(t, err)
	})

	t.Run("Set Bytes", func(t *testing.T) {
		err := client.Set("bytes_key", []byte("bytes_value"), time.Minute)
		assert.NoError(t, err)
	})
}

func TestGet(t *testing.T) {
	client := setupClient()
	defer client.Close()

	t.Run("get a string present in cache", func(t *testing.T) {
		err := client.Set("test_key", "test_value", time.Minute)
		assert.NoError(t, err)

		value, err := client.GetString("test_key")
		assert.NoError(t, err)
		assert.Equal(t, "test_value", value)
	})

	t.Run("return error on cache miss", func(t *testing.T) {
		_, err := client.GetString("invalid _key")
		assert.Error(t, err)
	})

	t.Run("Get Struct", func(t *testing.T) {
		type TestStruct struct {
			Field1 string
			Field2 int
		}

		testStruct := TestStruct{Field1: "value1", Field2: 42}
		err := client.SetStruct("struct_key", testStruct, time.Minute)
		assert.NoError(t, err)

		var result TestStruct
		err = client.GetStruct("struct_key", &result)
		assert.NoError(t, err)
		assert.Equal(t, testStruct, result)
	})

	t.Run("Get Bytes", func(t *testing.T) {
		err := client.Set("bytes_key", []byte("bytes_value"), time.Minute)
		assert.NoError(t, err)

		value, err := client.GetBytes("bytes_key")
		assert.NoError(t, err)
		assert.Equal(t, []byte("bytes_value"), value)
	})

	t.Run("Get Non-Existent Key", func(t *testing.T) {
		_, err := client.GetString("non_existent_key")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cache_miss")
	})
}

func TestDel(t *testing.T) {
	client := setupClient()
	defer client.Close()

	t.Run("Delete Key", func(t *testing.T) {
		err := client.Set("delete_key", "to_be_deleted", time.Minute)
		assert.NoError(t, err)

		err = client.Del("delete_key")
		assert.NoError(t, err)

		_, err = client.GetString("delete_key")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cache_miss")
	})
}
