package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/nbright/nomadcoin/utils"
)

/*
	절차 /이론코드

1) we hash the msg
"i love you" -> hash(x) -> "hashed_message"

2) generate key pair
KeyPair (private Key, public Key) (save priv to a file)

3) sign the hash
("hashed_message" + private Key) -> "signature"

4) verify
("hashed_message" + "signature" + public Key) -> true / false

이론코드

	//2) generate key pair : create private key (Elliptic Curve Digital Signature Algorithem) y2=x3+ax+b
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//x509 key 를 parse 하는 표준
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)

	fmt.Printf("privateKey AsBytes: %x\n\n\n", keyAsBytes)

	utils.HandleErr(err)
	//1) we hash the msg

	fmt.Println("hashedMessage: ", hashedMessage)

	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	//3) sign the hash : signature
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("signature: %x", signature)

	utils.HandleErr(err)

	//4) verify
	ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)
	fmt.Println(ok)
*/
const (
	signature  string = "abd2abf373997128e6208ae0a34dfb9d76f412ee34aa77e84122f31fd8445e2076392189d818101628fdeef07cc199cf96a3597ae76936f4b01eedddee42ef17"
	privateKey string = "3077020101042082c0876db49cade9c968891441ce78a24332f976e838ca8b9edc62304045ccf7a00a06082a8648ce3d030107a14403420004551e6a5ead3b1f90003619b3ecac809e1e6f35f12f60dc64079b74d471829aa07b8e798b91d09b685a02db0d396d712fb0b8108282b8fc11587e722ea9d032c2"
	//&{{0x84bec0 38500331679112089423409746125272053511493861677303916564936039130802912664224 55886211533036942996220913519515284038369672374467810070347726616707787141826} 59140839645673750371203875215128068884123494416857918168582949937637667228919}
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	// 복구 과정
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)
	restoredKey, err := x509.ParseECPrivateKey(privByte)
	utils.HandleErr(err)
	fmt.Println(restoredKey)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)

}
