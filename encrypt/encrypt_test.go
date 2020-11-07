package encrypt

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestEncrypt(t *testing.T) {
	originalText := "test_test_test"
	originalBytes := []byte(originalText)

	keys := [][]byte{
		{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC},
		[]byte("test_key"),
	}

	for _, key := range keys {
		cryptoText, _ := DesEncryption(key, originalBytes)
		fmt.Println(string(cryptoText))
		decryptedText, _ := DesDecryption(key, cryptoText)
		assert.Equal(t, originalText, string(decryptedText))
	}
}
