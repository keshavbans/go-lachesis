package poset

import (
	"fmt"
	"testing"

	"github.com/Fantom-foundation/go-lachesis/lachesis/src/crypto"
)

func TestSignBlock(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()

	block := NewBlock(0, 1,
		[]byte("framehash"),
		[][]byte{
			[]byte("abc"),
			[]byte("def"),
			[]byte("ghi"),
		})

	sig, err := block.Sign(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	res, err := block.Verify(sig)
	if err != nil {
		t.Fatalf("Error verifying signature: %v", err)
	}
	if !res {
		t.Fatal("Verify returned false")
	}
}

func TestAppendSignature(t *testing.T) {
	privateKey, _ := crypto.GenerateKey()
	pubKeyBytes := privateKey.Public().Bytes()

	block := NewBlock(0, 1,
		[]byte("framehash"),
		[][]byte{
			[]byte("abc"),
			[]byte("def"),
			[]byte("ghi"),
		})

	sig, err := block.Sign(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	err = block.SetSignature(sig)
	if err != nil {
		t.Fatal(err)
	}

	blockSignature, err := block.GetSignature(fmt.Sprintf("0x%X", pubKeyBytes))
	if err != nil {
		t.Fatal(err)
	}

	res, err := block.Verify(blockSignature)
	if err != nil {
		t.Fatalf("Error verifying signature: %v", err)
	}
	if !res {
		t.Fatal("Verify returned false")
	}

}