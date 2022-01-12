# NFCLI
free5GC Network Function (NF) Command Line Interface (CLI) based on Thrift 

# Thrift Installation Steps
1) Donwload Thrift 0.15.0
2) ./bootstrap.sh
3) ./configure --with-go --without-cpp  --with-boost=/usr/local --without-python --without-csharp --without-java --without-erlang --without-perl --without-php --without-php_extension --without-ruby --without-haskell
4) Modify the following lines in lib/go/Makefile
    > check-local:
      * #$(GO) test -mod=mod -race ./thrift
      * $(GO) test -race ./thrift
        
    > all-local:
      * #$(GO) build -mod=mod ./thrift
      * $(GO) build ./thrift
5) make
6) sudo make check
7) sudo make install 
8) thrift --version
    > Thrift version 0.15.0

References
* https://downloads.apache.org/thrift/
* https://thrift-tutorial.readthedocs.io/en/latest/installation.html

# NF Thrift Compilation
* cd nfcli 
* thrift -r --gen go thrift/UpfService.thrift

# Supported
| NF Name                |   Thrift Service  |
| ---------------------- | -----------------:|
| AMF                    | TODO              |  
| SMF                    | TODO              |   
| UPF                    | TODO              | 
| UPF CP  (Proprietary)  | TODO              |   
| UPF DP  (Proprietary)  | In-Progress       | 
| AUSF                   | TODO              | 
| N3IWF                  | TODO              | 
| NRF                    | TODO              | 
| NSSF                   | TODO              | 
| PCF                    | TODO              | 
| UDM                    | TODO              | 
| UDR                    | TODO              | 
| NEF                    | TODO              | 
