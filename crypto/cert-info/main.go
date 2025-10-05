package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	_algo := flag.Bool("algo", false, "Show public and signature algorithms")
	_crl := flag.Bool("crl", false, "Show any CRL distribution points")
	_rem := flag.Bool("rem", false, "Show any remaining text")
	_san := flag.Bool("san", false, "Show subject alternative names")
	_ski := flag.Bool("ski", false, "Show subject key identifer")
	_ku := flag.Bool("key-usage", false, "Show subject key identifer")
	flag.Parse()

	files := flag.Args()
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("error reading file %s:  %v\n", file, err)
			continue
		}

		block, rest := pem.Decode(data)
		if block == nil || block.Type != "CERTIFICATE" {
			fmt.Println("Failed to decode PEM block containing certificate")
			os.Exit(1)
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			fmt.Printf("Failed to parse certificate: %v\n", err)
			return
		}

		fmt.Printf("      File: %s\n", file)
		fmt.Printf("   Subject: %s\n", cert.Subject.String())
		fmt.Printf("   Version: %d\n", cert.Version)
		fmt.Printf("    Serial: %s\n", cert.SerialNumber.String())
		fmt.Printf("    Issuer: %s\n", cert.Issuer.String())
		fmt.Printf("     Start: %s\n", cert.NotBefore.String())
		fmt.Printf("   Expires: %s\n", cert.NotAfter.String())

		if *_algo {
			fmt.Printf("PubKeyAlgo: %s\n", cert.PublicKeyAlgorithm.String())
			fmt.Printf("   SigAlgo: %s\n", cert.SignatureAlgorithm.String())
		}

		if *_crl {
			fmt.Printf("       CRL:\n")
			for _, v := range cert.CRLDistributionPoints {
				fmt.Printf("          %v\n", v)
			}
		}

		if *_ku {
			ku := make(map[x509.KeyUsage]string)
			ku[x509.KeyUsageDigitalSignature] = "Digital Signature"
			ku[x509.KeyUsageContentCommitment] = "Content Commitment"
			/*ku[3] = "Key Encipherment"
			ku[4] = "Data Encipherment"
			ku[5] = "Key Agreement"
			ku[6] = "Cert Sign"
			ku[7] = "CRL Sign"
			ku[8] = "Encipher Only"
			ku[9] = "Decipher Only"*/

			fmt.Print(" Key Usage: ")
			var i x509.KeyUsage
			first := true
			for i = x509.KeyUsageDigitalSignature; i <= x509.KeyUsageDecipherOnly; i++ {
				if or := i | cert.KeyUsage; or == 1 {
					if first {
						first = false
						fmt.Printf("%s\n", ku[i])
					} else {
						fmt.Printf("           %s\n", ku[i])
					}
				}
			}
		}

		if *_rem {
			fmt.Printf("  Remnant: %q\n", rest)
		}

		if *_san {
			fmt.Printf("       SAN:\n%v\n", cert.DNSNames)
		}

		if *_ski {
			fmt.Printf(" Subject Key ID: %v\n", cert.SubjectKeyId)
		}

		fmt.Println()
	}
}
