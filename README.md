# Go today, go tomorrow, go forever

- Load balancer (go): https://github.com/geffersonFerraz/lb-http-2-quic

- Database: Mongo



## If the console displays any warnings regarding buffer size, execute the following command:
```
sysctl -w net.core.rmem_max=2500000 
sysctl -w net.core.wmem_max=2500000


source( https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes )
```

