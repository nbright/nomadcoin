package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/nbright/nomadcoin/utils"
)

/*
	절차

1) we hash the msg
"i love you" -> hash(x) -> "hashed_message"

2) generate key pair
KeyPair (private Key, public Key) (save priv to a file)

3) sign the hash
("hashed_message" + private Key) -> "signature"

4) verify
("hashed_message" + "signature" + public Key) -> true / false
*/
func Start() {
	//2) generate key pair : create private key (Elliptic Curve Digital Signature Algorithem) y2=x3+ax+b
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	utils.HandleErr(err)
	//1) we hash the msg
	message := "i love you"
	hashedMessage := utils.Hash(message)
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	//3) sign the hash : signature
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleErr(err)

	//4) verify
	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)

	fmt.Println(ok)

	//fmt.Printf("R:%d\nS:%d", r, s)

}
