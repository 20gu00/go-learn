package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    /*
       ioutil包：
           ReadFile()
           WriteFile()
           ReadDir()
           ..
    */

    //最方便就是不用os.Open打开文件或者目录,可以很方便的重复操作
    //1.读取文件中的所有的数据
    fileName1 := "/root/go-learn/io/ioutil/test"
    data, err := ioutil.ReadFile(fileName1)
    fmt.Println(err)
    fmt.Println(string(data))

    //2.写出数据
    fileName2 := "/root/go-learn/io/ioutil/test"
    s1 := "helloworld面朝大海春暖花开"
    err = ioutil.WriteFile(fileName2, []byte(s1), 0777)
    fmt.Println(err)

    //3.
    s2 := "qwertyuiopsdfghjklzxcvbnm"
    r1 := strings.NewReader(s2)
    data, _ = ioutil.ReadAll(r1)
    fmt.Println(data)

    //4.ReadDir(),读取一个目录下的子内容：子文件和子目录，但是仅有一层
    dirName := "/root/go-learn/io/ioutil/a"
    fileInfos, _ := ioutil.ReadDir(dirName)
    fmt.Println(len(fileInfos))
    //for i:=0;i
}
