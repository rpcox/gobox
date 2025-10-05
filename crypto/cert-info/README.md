## cert-info

    > cert-info -key-usage -algo *.pem
          File: google.pem
       Subject: CN=*.google.com
       Version: 3
        Serial: 275681765554378825529392251565661853316
        Issuer: CN=WE2,O=Google Trust Services,C=US
         Start: 2025-07-07 08:34:14 +0000 UTC
       Expires: 2025-09-29 08:34:13 +0000 UTC
    PubKeyAlgo: ECDSA
       SigAlgo: ECDSA-SHA256
     Key Usage: Digital Signature

          File: security.stackexchange.com.pem
       Subject: CN=security.stackexchange.com
       Version: 3
        Serial: 500221542290708826615611861855328612605033
        Issuer: CN=E5,O=Let's Encrypt,C=US
         Start: 2025-07-28 15:50:16 +0000 UTC
       Expires: 2025-10-26 15:50:15 +0000 UTC
    PubKeyAlgo: ECDSA
       SigAlgo: ECDSA-SHA384
     Key Usage: Digital Signature



    > cert-info -h
    Usage of ./cert-info:
      -algo
        Show public and signature algorithms
      -crl
        Show any CRL distribution points
      -key-usage
        Show the list of valid usage
      -rem
        Show any remaining text
      -san
        Show subject alternative names
      -ski
        Show subject key identifer
