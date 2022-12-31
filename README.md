# masscan-online

Preparation
~~~
$ sudo iptables -A INPUT -p tcp --dport 61000 -j DROP  
$ sudo iptables-save  
$ sudo chmod +s /usr/bin/masscan  
~~~

Build
~~~
$ make
~~~

CrossCompile
- Linux
    - amd64
        - ~~~
          $ make GOOS=linux GOARCH=amd64
    - arm
        - ~~~
          $ make GOOS=linux GOARCH=arm
    - arm64
        - ~~~
          $ make GOOS=linux GOARCH=arm64