https://github.com/OJ/gobuster
```
Gobuster is a tool used to brute-force:

- URIs (directories and files) in web sites.
- DNS subdomains (with wildcard support).
- Virtual Host names on target web servers.
- Open Amazon S3 buckets
- Open Google Cloud buckets
- TFTP servers
```


# Install
```
sudo apt-get -yq install gobuster seclists
```

# Quick Help

```
gobuster dir -u https://deadcell.dev -w /usr/share/wordlists/words.txt -x php,html,aspx,txt,etc

gobuster dns -d https://deadcell.dev -w /usr/share/wordlists/words.txt -i --wildcard

gobuster vhost -v -u https://deadcell.dev -w /usr/share/wordlists/words.txt -o deadcell.dev.gobuster.txt

gobuster s3 -w /usr/share/wordlists/words.txt

...


-e : completes printing the URL from enumerated directories

-n : will print results without the status code

-k : skip SSL verification

-t : assign the number of threads that will be used during enumeration

-r : allows redirection from one HTTP request to another

-p : allows proxy URLs to be used for requests on port 1080. This port can be changed in the URL.

-timeout : allows a timeout parameter to be set

-U byinarie -P P@SSw0rd1111!!! : define a username and password for basic HTTP authentication mechanisms
```

# Common Wordlists
	
	You should take some time to learn the contents of the below wordlists, or at the very least give yourself a general idea
	In general is is recommended to take some time to browse all of SecLists and understand the structure. 
	Common wordlists are also provided in offensive/wordlists/web
	
```
cd /usr/share/seclists/Discovery/Web-Content/
realpath * | xargs wc -l 
	20476 /usr/share/seclists/Discovery/Web-Content/big.txt
	141708 /usr/share/seclists/Discovery/Web-Content/directory-list-1.0.txt
	1273833 /usr/share/seclists/Discovery/Web-Content/directory-list-2.3-big.txt
	220560 /usr/share/seclists/Discovery/Web-Content/directory-list-2.3-medium.txt
	87664 /usr/share/seclists/Discovery/Web-Content/directory-list-2.3-small.txt
	1185254 /usr/share/seclists/Discovery/Web-Content/directory-list-lowercase-2.3-big.txt
	207643 /usr/share/seclists/Discovery/Web-Content/directory-list-lowercase-2.3-medium.txt
	81643 /usr/share/seclists/Discovery/Web-Content/directory-list-lowercase-2.3-small.txt
	12939 /usr/share/seclists/Discovery/Web-Content/dirsearch.txt
	1377715 /usr/share/seclists/Discovery/Web-Content/combined_directories.txt
	4726 /usr/share/seclists/Discovery/Web-Content/common.txt
	56164 /usr/share/seclists/Discovery/Web-Content/raft-large-directories-lowercase.txt
	62284 /usr/share/seclists/Discovery/Web-Content/raft-large-directories.txt
	2367 /usr/share/seclists/Discovery/Web-Content/raft-large-extensions-lowercase.txt
	2450 /usr/share/seclists/Discovery/Web-Content/raft-large-extensions.txt
	35325 /usr/share/seclists/Discovery/Web-Content/raft-large-files-lowercase.txt
	37050 /usr/share/seclists/Discovery/Web-Content/raft-large-files.txt
	107982 /usr/share/seclists/Discovery/Web-Content/raft-large-words-lowercase.txt
	119600 /usr/share/seclists/Discovery/Web-Content/raft-large-words.txt
	26584 /usr/share/seclists/Discovery/Web-Content/raft-medium-directories-lowercase.txt
	30000 /usr/share/seclists/Discovery/Web-Content/raft-medium-directories.txt
	1234 /usr/share/seclists/Discovery/Web-Content/raft-medium-extensions-lowercase.txt
	1290 /usr/share/seclists/Discovery/Web-Content/raft-medium-extensions.txt
	16244 /usr/share/seclists/Discovery/Web-Content/raft-medium-files-lowercase.txt
	17129 /usr/share/seclists/Discovery/Web-Content/raft-medium-files.txt
	56293 /usr/share/seclists/Discovery/Web-Content/raft-medium-words-lowercase.txt
	63088 /usr/share/seclists/Discovery/Web-Content/raft-medium-words.txt
	17770 /usr/share/seclists/Discovery/Web-Content/raft-small-directories-lowercase.txt
	20116 /usr/share/seclists/Discovery/Web-Content/raft-small-directories.txt
	914 /usr/share/seclists/Discovery/Web-Content/raft-small-extensions-lowercase.txt
	963 /usr/share/seclists/Discovery/Web-Content/raft-small-extensions.txt
	10848 /usr/share/seclists/Discovery/Web-Content/raft-small-files-lowercase.txt
	11424 /usr/share/seclists/Discovery/Web-Content/raft-small-files.txt
	38267 /usr/share/seclists/Discovery/Web-Content/raft-small-words-lowercase.txt
	43007 /usr/share/seclists/Discovery/Web-Content/raft-small-words.txt
```

# Usage

## dir - Discover open directories and files
	First identity the tech stack using whatwebb or wappalyzer so we can determine the correct file extensions we want to add
	Check the 200 status codes first
	Browse everything, take note of errors
	Download all lootable files
	Use -k to ignore ssl errors caused by self signed certs
	Always output your results to a file using -o
	Always begin by limiting your threads to avoid crashing the application
	Not adding -x to specify exentions will result in missing loot.
	Use the --delay flag if you experience any lag or errors, you may adjust the threads as well
	--delay 1s
	-t 10
	-t 5
```
whatweb http://10.10.12.201:83/

http://10.10.12.201:83/ [200 OK] Apache[2.4.7], Country[RESERVED][ZZ], HTML5, HTTPServer[Ubuntu Linux][Apache/2.4.7 (Ubuntu)], IP[10.10.12.201], JQuery[1.12.4], MetaGenerator[WordPress 4.6], PHP[5.5.9-1ubuntu4.21], PoweredBy[WordPress,WordPress,], Script[text/javascript], Title[Acme Corp &#8211; Just another WordPress site], UncommonHeaders[link], WordPress[4.6], X-Powered-By[PHP/5.5.9-1ubuntu4.21]

...

gobuster dir -u http://10.10.12.201:83/ -w /usr/share/seclists/Discovery/Web-Content/raft-medium-directories-lowercase.txt -t 10 -o deadcell.dev.dirs.txt -x php,txt,html,js
===============================================================
Gobuster v3.6
by OJ Reeves (@TheColonial) & Christian Mehlmauer (@firefart)
===============================================================
[+] Url:                     http://10.10.12.201:83/
[+] Method:                  GET
[+] Threads:                 10
[+] Wordlist:                /usr/share/seclists/Discovery/Web-Content/raft-medium-directories-lowercase.txt
[+] Negative Status codes:   404
[+] User Agent:              gobuster/3.6
[+] Extensions:              txt,html,js,php
[+] Timeout:                 10s
===============================================================
Starting gobuster in directory enumeration mode
===============================================================
/wp-content           (Status: 301) [Size: 319] [--> http://10.10.12.201:83/wp-content/]
/wp-admin             (Status: 301) [Size: 317] [--> http://10.10.12.201:83/wp-admin/]
/wp-includes          (Status: 301) [Size: 320] [--> http://10.10.12.201:83/wp-includes/]
/xmlrpc.php           (Status: 405) [Size: 42]
/index.php            (Status: 301) [Size: 0] [--> http://10.10.12.201:83/]
/wp-trackback.php     (Status: 200) [Size: 135]
/wp-login.php         (Status: 200) [Size: 2166]
/readme.html          (Status: 200) [Size: 7342]
/license.txt          (Status: 200) [Size: 19935]
/server-status        (Status: 403) [Size: 292]
/.php                 (Status: 403) [Size: 283]
/.html                (Status: 403) [Size: 284]
/wp-config.php        (Status: 200) [Size: 0]
```

## dns - subdomain discovery
	Identity the wordlist(s) you want to use
	/usr/share/wordlists/
	/usr/share/seclists/Discovery/Web-Content/
	If needed, add the initial domain to your /etc/hosts
	Remove -i if you don't want to view the IP addresses
	-r to specify the resolver you want to use
```
gobuster dns -r 10.10.100.200 -d google.com -w /usr/share/seclists/Discovery/Web-Content/raft-small-words.txt -t 10 --delay 1s -i -o google.com.dns.txt
===============================================================
Gobuster v3.6
by OJ Reeves (@TheColonial) & Christian Mehlmauer (@firefart)
===============================================================
[+] Domain:     google.com
[+] Threads:    10
[+] Delay:      1s
[+] Resolver:   10.10.100.200
[+] Show IPs:   true
[+] Timeout:    1s
[+] Wordlist:   /usr/share/seclists/Discovery/Web-Content/raft-small-words.txt
===============================================================
Starting gobuster in DNS enumeration mode
===============================================================
Found: admin.google.com [142.250.190.14,2607:f8b0:4009:818::200e]

Found: search.google.com [172.217.4.78,2607:f8b0:4009:81a::200e]

Found: images.google.com [142.250.191.110,2607:f8b0:4009:819::200e]

```


## fuzz - fuzzing for parameters
```
/usr/share/seclists/Discovery/Web-Content/burp-parameter-names.txt
/usr/share/seclists/Discovery/Web-Content/url-params_from-top-55-most-popular-apps.txt


# "FUZZ" - replaces the `FUZZ` keyword. 
# You may want to use ffuf or wfuzz rather than gobuster for fuzzing.

gobuster fuzz -u https://example.com?FUZZ=test -w parameter-names.txt

```