package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {
	_file := flag.String("file", "", "Identify the X509 certificate")
	_altNames := flag.Bool("alt-names", false, "Show subject alternative names")
	_rem := flag.Bool("rem", false, "Show any remaining text")
	_crl := flag.Bool("crl", false, "Show any CRL distribution points")
	flag.Parse()

	if *_file == "" {
		fmt.Fprintf(os.Stderr, "-file required\n")
		os.Exit(1)
	}

	data, err := os.ReadFile(*_file)
	if err != nil {
		fmt.Printf("error reading file %s:  %v\n", *_file, err)
		os.Exit(1)
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

	fmt.Println("Certificate Details:")
	fmt.Printf("   Subject: %s\n", cert.Subject.String())
	fmt.Printf("   Version: %d\n", cert.Version)
	fmt.Printf("    Serial: %s\n", cert.SerialNumber.String())
	fmt.Printf("    Issuer: %s\n", cert.Issuer.String())
	fmt.Printf("     Start: %s\n", cert.NotBefore.String())
	fmt.Printf("   Expires: %s\n", cert.NotAfter.String())
	fmt.Printf("   SigAlgo: %s\n", cert.SignatureAlgorithm.String())
	fmt.Printf("PubKeyAlgo: %s\n", cert.PublicKeyAlgorithm.String())
	fmt.Printf(" Key Usage: %v\n", cert.KeyUsage)
	if *_altNames {
		fmt.Printf("      SAN:\n%v\n", cert.DNSNames)
	}

	if *_rem {
		fmt.Printf("  Remnant: %q\n", rest)
	}

	if *_crl {
		for _, v := range cert.CRLDistributionPoints {
			fmt.Println(v)
	}
		fmt.Printf("  Remnant: %q\n", rest)
	}
}
