# log-backend
a log backend for neo-cli with logbackendplugin
## requisite
install go at least v1.10
## installation
* git clone https://github.com/KickSeason/log-backend.git 
## configuration
after installation, configure your this application in **config.json**
```
{
    "name": "ngd",//the name of your node, let partners to recognize you
    "cnport": 8080, //the port to receive log from cn-node and partners
    "webport": 8081,//the port to solve web api request
    "logpath": "./persist",//directory to store the log
    "logname": "concensus.log",//log file name
    "logfileexpire": 30,//day, how many days you want to store the log
    "logfilesplit": 1,//hour, how many hours to split the file
    "sendto":[//partners you want to send your log
           "http://ngd1.logbackend:8081"
    ]
}
```

> if you want to send your cn-node logs to partners, set your partners log-backend url in "__*sendto*__"
>set your node name to let partners recognize you, you can not use "" and "~~local~~", becase local is reserved for your cn-node, partners' logs using their names setting in Http Header "From"
## start
```go run main.go```