package ipfs

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestIsFile(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", false},
		{"https", "https://something", false},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", true},
		{"file2", "../ipfs/getters.go", true},
		{"dir", "../ipfs", false},
		{"notfile", "not-a-file", false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
	}

	for _, v := range tests {
		res := isFile(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", true},
		{"https", "https://something", true},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", false},
		{"notfile", "not-a-file", false},
		{"dir", filepath.Join("..", "ipfs"), false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
	}

	for _, v := range tests {
		res := isURL(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}

func TestIsPathIPNS(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", false},
		{"https", "https://something", false},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", false},
		{"notfile", "not-a-file", false},
		{"dir", filepath.Join("..", "ipfs"), false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"ipns path", "/ipns/k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", true},
	}

	for _, v := range tests {
		res := isPathIPNS(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}

func TestIsKeyIPNS(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", false},
		{"https", "https://something", false},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", false},
		{"notfile", "not-a-file", false},
		{"dir", filepath.Join("..", "ipfs"), false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", true},
		{"ipns path", "/ipns/k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
	}

	for _, v := range tests {
		res := isKeyIPNS(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}

func TestIsPathCID(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", false},
		{"https", "https://something", false},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", false},
		{"notfile", "not-a-file", false},
		{"dir", filepath.Join("..", "ipfs"), false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"ipns path", "/ipns/k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"ipfs path", "/ipfs/QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", true},
	}

	for _, v := range tests {
		res := isPathIPFS(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}

func TestIsCID(t *testing.T) {
	tests := []struct {
		msg, in string
		out     bool
	}{
		{"http", "http://something", false},
		{"https", "https://something", false},
		{"ftp", "ftp://something", false},
		{"file", "getters.go", false},
		{"notfile", "not-a-file", false},
		{"dir", filepath.Join("..", "ipfs"), false},
		{"cid", "QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", true},
		{"key", "k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"ipns path", "/ipns/k5p1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
		{"ipfs path", "/ipfs/QmYp1dP9eyhsnmzjXtwZAZ1vKTYhQkLmhYoZfNjjqPdsd8", false},
	}

	for _, v := range tests {
		res := isCID(v.in)
		assert.Equal(t, res, v.out, v.msg)
	}
}
