package main

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	xades "github.com/artemkunich/goxades"
	"github.com/beevik/etree"
	dsig "github.com/russellhaering/goxmldsig"
)

var sampleXml = `<element id="signedData" xmlns="namespace">text</element>`

func main() {

	doc := etree.NewDocument()
	err := doc.ReadFromString(strings.ReplaceAll(sampleXml, "\n", ""))
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	keyStore, err := loadCert("key.pem", "cert.crt")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	root := removeComments(doc.Root())
	fmt.Printf("%+v", root)
	canonicalizer := dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList("")
	signContext := xades.SigningContext{
		DataContext: xades.SignedDataContext{
			Canonicalizer: canonicalizer,
			Hash:          crypto.SHA256,
			ReferenceURI:  "#signedData",
			IsEnveloped:   true,
		},
		PropertiesContext: xades.SignedPropertiesContext{
			Canonicalizer: canonicalizer,
			Hash:          crypto.SHA256,
		},
		Canonicalizer: canonicalizer,
		Hash:          crypto.SHA256,
		KeyStore:      *keyStore,
	}
	signature, err := xades.CreateSignature(root, &signContext)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}

	b, err := canonicalSerialize(signature)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Println(string(b))
}

func removeComments(elem *etree.Element) *etree.Element {
	copy := elem.Copy()
	for _, token := range copy.Child {
		_, ok := token.(*etree.Comment)
		if ok {
			copy.RemoveChild(token)
		}
	}
	for i, child := range elem.ChildElements() {
		copy.ChildElements()[i] = removeComments(child)
	}
	return copy
}

func canonicalSerialize(el *etree.Element) ([]byte, error) {
	doc := etree.NewDocument()
	doc.SetRoot(el.Copy())

	doc.WriteSettings = etree.WriteSettings{
		CanonicalAttrVal: true,
		CanonicalEndTags: false,
		CanonicalText:    true,
	}

	return doc.WriteToBytes()
}

func loadCert(keyPath string, certPath string) (*xades.MemoryX509KeyStore, error) {
	buffer, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	blockc, _ := pem.Decode(buffer)
	cert, err := x509.ParseCertificate(blockc.Bytes)

	buffer, err = os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	blockp, _ := pem.Decode(buffer)
	key, err := x509.ParsePKCS1PrivateKey(blockp.Bytes)

	return &xades.MemoryX509KeyStore{
		PrivateKey: key,
		Cert:       cert,
		CertBinary: blockc.Bytes,
	}, nil
}
