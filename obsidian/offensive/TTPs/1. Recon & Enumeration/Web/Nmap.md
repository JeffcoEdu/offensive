## **Network Discovery**

### Common Nmap Scan

```
nmap -sV -sT -sC -T5 -v -A $targetip
```

* **-sV** (Version detection)
* \-**sT** (TCP connect scan)
* \-**sC** (Performs a script scan using the default set of scripts)
* **-v** (Increase verbosity level)


#### **All TCP port scan**

```
nmap -p- -sT -v $targetip
```

#### **All UDP Port Scan (With Service and Script Scan)**

```
nmap -p- -sV -sU -sC $targetip
```

#### **100 most** **common ports**

```
nmap $targetip -F 100
```

### **Nmap Script Scan**

#### **Find the nse scripts**

```
locate *.nse | grep <script.nse>
```

For example, the below command will find smb scripts.

```
locate *.nse | grep smb
```

#### **Scan using a specific NSE script**

```
nmap -sV -p 443 --script=ssl-heartbleed.nse $targetip
```

For example, below command will find any smb scripts running on port 139 & 445.

```
nmap -p 139,445 --script=smb-* $targetip
```

All smb scripts:&#x20;

```
nmap --script smb-vuln* -p 139,445 [ip] 
```

SMB share paths enumeration:

```
nmap --script smb-enum-shares -p 139,445 [ip]
```

&#x20;Using Metasploit portscan

```
use auxiliary/scanner/portscan/
```

## **Service Discovery**

### Port 80 & 443 - Web Discovery

#### Find SSL Heartbleed Vulnerablity

```
sslscan $targetip:443
```

```
nmap -sV --script=ssl-heartbleed 192.168.101.8
```
