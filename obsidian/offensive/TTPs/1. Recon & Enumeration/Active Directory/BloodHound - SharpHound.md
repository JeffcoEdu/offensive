# Install BloodHound CE
```
cd offensive/ad/BloodHound/examples/docker-compose

docker compose -f docker-compose.yml up -d


```

# Old BloodHound

```
sudo apt-get -y install bloodhound
sudo neo4j start
```
# Grab the admin password for BloodHound CE
```

{"level":"info","time":"2024-03-08T00:27:04.171968449Z","message":"#                                                                 #"}
{"level":"info","time":"2024-03-08T00:27:04.171970089Z","message":"# Initial Password Set To:    {password here}    #"}
{"level":"info","time":"2024-03-08T00:27:04.171971409Z","message":"#                                                                 #"}
{"level":"info","time":"2024-03-08T00:27:04.171972809Z","message":"###################################################################"}
```


add 127.0.0.1 bloodhound to /etc/hosts

vi /etc/hosts
```
127.0.0.1 bloodhound
```


# Run SharpHound
```
SharpHound.exe -c all -s
```