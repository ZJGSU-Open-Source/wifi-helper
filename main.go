package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    "time"
)

type Auth struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func readConf() (*Auth, error) {
    b, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Println(err)
        return nil, err
    }

    auth := new(Auth)
    err = json.Unmarshal(b, auth)
    return auth, err
}

func connect(auth *Auth) error {
    postForm := fmt.Sprintf(
        "buttonClicked=4&redirect_url=&err_flag=0&username=%s&password=%s",
        auth.Username,
        auth.Password,
    )

    postForm = strings.Replace(postForm, " ", "%20", -1)

    // log.Println(postForm)

    reader := strings.NewReader(postForm)

    req, err := http.NewRequest("POST", "https://2.2.2.2/login.html", reader)
    if err != nil {
        return err
    }

    req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    req.Header.Set("Accept-Encoding", "gzip, deflate")
    req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
    req.Header.Set("Connection", "keep-alive")
    req.Header.Set("Host", "2.2.2.2")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.11; rv:41.0) Gecko/20100101 Firefox/41.0")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr,
        Timeout: 20 * time.Second,
    }

    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    data, _ := ioutil.ReadAll(resp.Body)
    idx := strings.Index(string(data), "Login Successful")
    if idx < 0 {
        // log.Println(string(data))
        return fmt.Errorf("login failed")
    }

    return nil
}

func Check() bool {
    req, err := http.NewRequest("GET", "https://baidu.com", nil)
    if err != nil {
        return false
    }

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr,
        Timeout: 20 * time.Second,
    }

    resp, err := client.Do(req)
    if err != nil {
        return false
    }

    data, _ := ioutil.ReadAll(resp.Body)
    idx := strings.Index(string(data), "百度一下，你就知道")
    if idx < 0 {
        return false
        // fmt.Errorf("connect failed")
    }

    return true
}

func main() {
    auth, err := readConf()
    if err != nil {
        log.Println("读取配置文件出错：", err)
        os.Exit(1)
    }

    if err := connect(auth); err != nil {
        log.Println("登陆失败：", err)
    } else {
        log.Println("登陆成功")
    }

    for {

        if !Check() {
            log.Println("重新连接 Wi-Fi")
            if err := connect(auth); err != nil {
                log.Println("登陆失败：", err)
            } else {
                log.Println("登陆成功")
            }
        } else {
            log.Println("Wi-Fi 连接正常")
        }

        time.Sleep(10 * time.Second)
    }

    os.Exit(0)
}
