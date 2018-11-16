# LogServer
a log server for neo-cli with LogBackendPlugin
* save logs from your cn-node into files 
* transfer your logs to partner, another LogServer
## requisite
install go at least v1.10  
use with LogBackendPlugin v1.0.0
## installation
* git clone https://github.com/neo-ngd/LogServer.git
## configuration
after installation, configure your this application in **config.json**
```
{
    "name": "ngd",//the name of your node, let partners to recognize you
    "cnport": 8080, //the port to receive log from cn-node and partners
    "webport": 8081,//the port to serve websocket
    "logpath": "./persist",//directory to store the log
    "logname": "concensus.log",//log file name
    "logfileexpire": 30,//day, how many days you want to store the log
    "logfilesplit": 1,//hour, how many hours to split the file
    "sendto":[//partners you want to send your log
           "http://ngd1.logbackend:8081"
    ]
}
```

> if you want to send your cn-node logs to partners, set your partners' LogServer url in "__sendto__"

> set your node name to let partners recognize you, you can not use "" and "~~local~~", becase local is reserved for your cn-node, LogServer use name to set in Http Header "From" when sending to partner.

> "__cnport__ " is for serving cn-node LogBackendPlugin, see [LogBackendPlugin](https://github.com/KickSeason/LogBackendPlugin)
## start

```cd log-backend && go run main.go```
