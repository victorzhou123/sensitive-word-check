package main

import (
	"testing"
)

var sensitiveWord sensitiveWords

func init() {
	sensitiveWord, _ = newSensitiveWords()
}

func Test_sensitiveWords_read(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"test1",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sensitiveWord.read(); (err != nil) != tt.wantErr {
				t.Errorf("sensitiveWords.read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
