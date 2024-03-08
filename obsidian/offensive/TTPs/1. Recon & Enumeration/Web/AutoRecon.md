# Installation

```python
cd offensive/AutoRecon
sudo apt -yq install python3 python3-pip seclists seclists curl dnsrecon enum4linux feroxbuster gobuster impacket-scripts nbtscan nikto nmap onesixtyone oscanner redis-tools smbclient smbmap snmp sslscan sipvicious tnscmd10g whatweb wkhtmltopdf
pip3 install -r requirements.txt
pipx install git+https://github.com/Tib3rius/AutoRecon.git

autorecon --help                                                                                                                      
usage: autorecon [-t TARGET_FILE] [-p PORTS] [-m MAX_SCANS]
```

## Path issues
Ensure .local/bin is exported.
```
echo $PATH | egrep ".local/bin"
export PATH=/home/byinarie/.npm-global/bin:/home/byinarie/.local/bin <snip>
```

If you do not see /home/<your-user>/.local/bin you will need to add it manually.


See: https://www.youtube.com/watch?v=hk0RwVC6uts 


