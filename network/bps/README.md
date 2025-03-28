## bps


    tcpdump -i <device> -l -e -n

    -e Print the link-level header
    -l Make stdout line buffered
    -n Don't convert addreses to names



    $ sudo tcpdump -i eth0 -l -e -n | bps
    tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
    listening on eth0, link-type EN10MB (Ethernet), snapshot length 262144 bytes
             131.24 Bps
            4394.31 Bps
             104.09 Bps
             106.52 Bps
             158.01 Bps
             150.48 Bps
             131.84 Bps
          269825.62 Bps
             103.85 Bps
             200.52 Bps
             297.13 Bps
    ^C9939 packets captured
    9942 packets received by filter
    0 packets dropped by kernel
    $
