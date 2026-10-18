   ## auth-basic  
    
    >  auth-basic -u rpcox
    service account password:
    user exists @ uid=rpcox,ou=people,dc=cox,dc=lan
    rpcox password:
    user auth sucess
    >
    
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 fd=10 ACCEPT from IP=10.10.10.10:52320 (IP=0.0.0.0:389)
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=0 BIND dn="uid=minion,ou=service,dc=cox,dc=lan" method=128
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=0 BIND dn="uid=minion,ou=service,dc=cox,dc=lan" mech=SIMPLE bind_ssf=0 ssf=0
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=0 RESULT tag=97 err=0 qtime=0.000026 etime=0.000250 text=
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=1 SRCH base="ou=people,dc=cox,dc=lan" scope=2 deref=0 filter="(uid=rpcox)"
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=1 SRCH attr=dn
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 op=1 SEARCH RESULT tag=101 err=0 qtime=0.000027 etime=0.000214 nentries=1 text=
    Oct 18 21:23:21 tequila slapd[88592]: conn=1033 fd=10 closed (connection lost)
    Oct 18 21:23:27 tequila slapd[88592]: conn=1034 fd=10 ACCEPT from IP=10.10.10.10:52321 (IP=0.0.0.0:389)
    Oct 18 21:23:27 tequila slapd[88592]: conn=1034 op=0 BIND dn="uid=rpcox,ou=people,dc=cox,dc=lan" method=128
    Oct 18 21:23:27 tequila slapd[88592]: conn=1034 op=0 BIND dn="uid=rpcox,ou=people,dc=cox,dc=lan" mech=SIMPLE bind_ssf=0 ssf=0
    Oct 18 21:23:27 tequila slapd[88592]: conn=1034 op=0 RESULT tag=97 err=0 qtime=0.000025 etime=0.000299 text=
    Oct 18 21:23:27 tequila slapd[88592]: conn=1034 fd=10 closed (connection lost)
    
    
    >  auth-basic -u spud
    service account password:
    user does not exist in directory information tree
    >
    
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 fd=10 ACCEPT from IP=10.10.10.10:52334 (IP=0.0.0.0:389)
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=0 BIND dn="uid=minion,ou=service,dc=cox,dc=lan" method=128
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=0 BIND dn="uid=minion,ou=service,dc=cox,dc=lan" mech=SIMPLE bind_ssf=0 ssf=0
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=0 RESULT tag=97 err=0 qtime=0.000023 etime=0.000237 text=
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=1 SRCH base="ou=people,dc=cox,dc=lan" scope=2 deref=0 filter="(uid=spud)"
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=1 SRCH attr=dn
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 op=1 SEARCH RESULT tag=101 err=0 qtime=0.000029 etime=0.000202 nentries=0 text=
    Oct 18 21:32:29 tequila slapd[88592]: conn=1041 fd=10 closed (connection lost)
    
    
    >  auth-basic -u rpcox
    service account password:
    %err: bind: LDAP Result Code 49 "Invalid Credentials":  for user 'uid=minion,ou=service,dc=cox,dc=lan'
    >
    
    Oct 18 21:25:01 tequila slapd[88592]: conn=1035 fd=10 ACCEPT from IP=10.10.10.10:52325 (IP=0.0.0.0:389)
    Oct 18 21:25:01 tequila slapd[88592]: conn=1035 op=0 BIND dn="uid=minion,ou=service,dc=cox,dc=lan" method=128
    Oct 18 21:25:01 tequila slapd[88592]: conn=1035 op=0 RESULT tag=97 err=49 qtime=0.000019 etime=0.000252 text=
    Oct 18 21:25:01 tequila slapd[88592]: conn=1035 fd=10 closed (connection lost)
    
    
