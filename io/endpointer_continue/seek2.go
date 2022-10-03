package main

import (
    "fmt"
    "io"
    "os"
    "strconv"
)

func main() {
    /*
       断点续传：
           文件传递：文件复制
               /Users/ruby/Documents/pro/a/guliang.jpeg

           复制到
               guliang4.jpeg

       思路：
           边复制，边记录复制的总量
    */

    //首先思考几个问题
    //Q1：如果你要传的文件，比较大，那么是否有方法可以缩短耗时？
    //Q2：如果在文件传递过程中，程序因各种原因被迫中断了，那么下次再重启时，文件是否还需要重头开始？
    //Q3：传递文件的时候，支持暂停和恢复么？即使这两个操作分布在程序进程被杀前后。
    //
    //通过断点续传可以实现，不同的语言有不同的实现方式。我们看看Go语言中，通过Seek()方法如何实现：
    //
    //先说一下思路：想实现断点续传，主要就是记住上一次已经传递了多少数据，那我们可以创建一个临时文件，记录已经传递的数据量，当恢复传递的时候，先从临时文件中读取上次已经传递的数据量，然后通过Seek()方法，设置到该读和该写的位置，再继续传递数据。

    srcFile := "/Users/ruby/Documents/pro/a/guliang.jpeg"
    destFile := "guliang4.jpeg"
    tempFile := destFile + "temp.txt"
    //fmt.Println(tempFile)
    file1, _ := os.Open(srcFile)
    file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
    file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

    defer file1.Close()
    defer file2.Close()
    //1.读取临时文件中的数据，根据seek
    file3.Seek(0, io.SeekStart)
    bs := make([]byte, 100, 100)
    n1, err := file3.Read(bs)
    fmt.Println(n1)
    countStr := string(bs[:n1])
    fmt.Println(countStr)
    //count,_:=strconv.Atoi(countStr)
    count, _ := strconv.ParseInt(countStr, 10, 64)
    fmt.Println(count)

    //2. 设置读，写的偏移量
    file1.Seek(count, 0)
    file2.Seek(count, 0)
    data := make([]byte, 1024, 1024)
    n2 := -1            // 读取的数据量
    n3 := -1            //写出的数据量
    total := int(count) //读取的总量

    for {
        //3.读取数据
        n2, err = file1.Read(data)
        if err == io.EOF {
            fmt.Println("文件复制完毕。。")
            file3.Close()
            os.Remove(tempFile)
            break
        }
        //将数据写入到目标文件
        n3, _ = file2.Write(data[:n2])
        total += n3
        //将复制总量，存储到临时文件中
        file3.Seek(0, io.SeekStart)
        file3.WriteString(strconv.Itoa(total))

        //假装断电
        //if total>8000{
        //  panic("假装断电了。。。，假装的。。。")
        //}
    }

}

