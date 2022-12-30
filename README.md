# masscan-online

Preparation
~~~
$ sudo iptables -A INPUT -p tcp --dport 61000 -j DROP  
$ sudo iptables-save  
$ sudo chmod +s /usr/bin/masscan  
~~~
