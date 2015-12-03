# wifi-helper
ZJGSU Wi-Fi helper

## Feature

+ Support multi-user login using same account at the same time.
+ Support auto-login when connection closed.
+ Keep your connection alive.

## Usage

### OS X / Linux

Copy and paste the following code into your terminal:
```bash
curl 'https://2.2.2.2/login.html' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3' -H 'Connection: keep-alive' -H 'Host: 2.2.2.2' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:41.0) Gecko/20100101 Firefox/41.0' -H 'Content-Type: application/x-www-form-urlencoded' --data 'buttonClicked=4&redirect_url=&err_flag=0&username=帐号%20%20%20%20%20%20%20%20%20%20%20&password=密码' -k
```

Please remember to modify your 帐号 and 密码.

### iOS

+ Download iCurlHTTP app via APP Store.
+ Open iCurlHTTP and click `User` button, then copy below into `POST DATA`:  
```bash
buttonClicked=4&redirect_url=&err_flag=0&username=帐号%20%20%20%20%20%20%20%20%20%20%20&password=密码
```  
Please remember modify 帐号 and 密码.  
+ Click `Done` back to main screen, enter below url:  
```bash
https://2.2.2.2/login.html
```  
Click `POST`.  
+ Finally if you see `Login Successful` in the response, you can surf the Internet now.

### Android

+ Download [Terminal emulator](https://play.google.com/store/apps/details?id=jackpal.androidterm) via Google Play or other third-party Android market.
+ Copy and paste the following code into the terminal.
```bash
curl 'https://2.2.2.2/login.html' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3' -H 'Connection: keep-alive' -H 'Host: 2.2.2.2' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:41.0) Gecko/20100101 Firefox/41.0' -H 'Content-Type: application/x-www-form-urlencoded' --data 'buttonClicked=4&redirect_url=&err_flag=0&username=帐号%20%20%20%20%20%20%20%20%20%20%20&password=密码' -k
```
Please remember to modify your 帐号 and 密码.

+ Finally if you see `Logged In` in the response, you can surf the Internet now.

![Screenshot-Android](https://dn-jinwei.qbox.me/Screenshot_2015-12-03-13-03-05.png)

### Windows
+ If you are using Windows, you can either [download](https://github.com/ZJGSU-Open-Source/wifi-helper/releases/download/v1.1/wifi-helper.zip) our pre-built binary release to login or if you are using Cygwin/MSYS2, you can still use the `curl` way.
+ Please remember to modify your username and password in `config.json`.

## Note

You must connect to `ZJGSU` or `ZJGSU-STU` Wi-Fi first.
