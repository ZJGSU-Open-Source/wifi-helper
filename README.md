# wifi-helper
ZJGSU Wi-Fi helper

## Feature

1. Support multi-people use same account login at same time.  
2. Support auto login when connection closed.  
3. Keep connection alive.

## Yet Another Wifi Trick

### OS X / Linux

Copy and paste into your terminal:  
```bash
curl 'https://2.2.2.2/login.html' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3' -H 'Connection: keep-alive' -H 'Host: 2.2.2.2' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:41.0) Gecko/20100101 Firefox/41.0' -H 'Content-Type: application/x-www-form-urlencoded' --data 'buttonClicked=4&redirect_url=&err_flag=0&username=帐号%20%20%20%20%20%20%20%20%20%20%20&password=密码' -k
```

Please remember modify 帐号 and 密码.

### IOS

1. Download iCurlHTTP app vai APP Store.
2. Open iCurlHTTP and click `User` button, then copy below into `POST DATA`:  
```bash
buttonClicked=4&redirect_url=&err_flag=0&username=帐号%20%20%20%20%20%20%20%20%20%20%20&password=密码
```  
Please remember modify 帐号 and 密码.  
3. Click `Done` back to main screen, enter below url:  
```bash
https://2.2.2.2/login.html
```  
Click `POST`.  
4. Finally if you see `Login Successful` in the response, you can surf the Internet.

### Android

You can find same app on you android.

## Note

You must connect zjgsu wifi before you do those tricks.
