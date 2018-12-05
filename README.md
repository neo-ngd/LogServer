# LogServer
A log server for neo-cli with LogPlugin
* save logs from your consensus node into files 
* transfer your logs to partner, another LogServer  

**warning**: LogServer2.0 integrated with LogMonitor you dont need to deploy logmonitor anymore.
## 1. requisite
install go at least v1.10
use with LogPlugin v1.0.0
## 2. installation
#### 2.1 source code
* git clone https://github.com/neo-ngd/LogServer.git into your $GOPATH/github.com/neo-ngd
#### 2.2 download application directly
* https://github.com/neo-ngd/LogServer/releases
## 3. configuration
after installation, configure your this application in **config.json**
```
{
    "name": "ngd",//the name of your node, let partners to recognize you
    "port": 8080, //server port
    "logpath": "./persist",//directory to store the log
    "logname": "concensus.log",//log file name
    "logfileexpire": 30,//day, how many days you want to store the log 
    "logfilesplit": 1,//hour, how many hours to split the file into
    "sendto":[//partners you want to send your log
           "http://ngd1.logbackend:8080"
    ]
}
```

> if you want to send your consensus node logs to partners, set your partners' LogServer url in "__sendto__"

> set your node name to let partners recognize you, LogServer use name to set in Http Header "From" when sending to partner.

> "__port__ " is for serving cn-node LogPlugin using "__http://logserverip:port/log__" and LogMonitor using "__http://logserverip:port__".

## 4. start

```cd log-backend && go run main.go```

## 5. usage
LogServer can both receive logs from cn and provide LogMonitor 
__LogPlugin__: put "http://LogServerIP:port/log" into config.json of LogPlugin.
__LogMonitor__: use "http://LogServerIP:port" to visit web service directly.
