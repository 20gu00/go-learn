package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/*
	   文件操作：
	   1.路径：
	       相对路径：relative
	           ab.txt
	           相对于当前工程
	       绝对路径：absolute
	           /Users/ruby/Documents/pro/a/aa.txt

	       .当前目录
	       ..上一层
	   2.创建文件夹，如果文件夹存在，创建失败
	       os.MkDir()，创建一层
	       os.MkDirAll()，可以创建多层

	   3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	       os.Create()，创建文件

	   4.打开文件：让当前的程序，和指定的文件之间建立一个连接
	       os.Open(filename)
	       os.OpenFile(filename,mode,perm)

	   5.关闭文件：程序和文件之间的链接断开。
	       file.Close()

	   5.删除文件或目录：慎用，慎用，再慎用
	       os.Remove()，删除文件和空目录
	       os.RemoveAll()，删除所有
	*/

	//删除文件或目录
	//os.Remove("abc/d/e/f");
	//
	////删除指定目录下所有文件
	//os.RemoveAll("abc");
	//
	////重命名文件
	//os.Rename("./2.txt", "./2_new.txt");
	//
	////判断是否为同一文件
	////unix下通过底层结构的设备和索引节点是否相同来判断
	////其他系统可能是通过文件绝对路径来判断
	//fs1, _ := f1.Stat();
	//f2, _ := os.Open("./1.txt");
	//fs2, _ := f2.Stat();
	//fmt.Println(os.SameFile(fs1, fs2));

	//返回临时目录
	fmt.Println(os.TempDir())
	//1.路径
	fileName1 := "/root/go-learn/io/file_operator_os/file_operator_os.go"
	fileName2 := "bb.txt"
	fmt.Println(filepath.IsAbs(fileName1)) //true absolute绝对的
	fmt.Println(filepath.IsAbs(fileName2)) //false
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2)) // /Users/ruby/go/src/l_file/bb.txt

	fmt.Println("获取父目录：", path.Join(fileName1, ".."))

	//2.创建目录
	err := os.Mkdir("/root/go-learn/io/file_operator_os/a/bb", os.ModePerm) //0777
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("文件夹创建成功。。")
	err = os.MkdirAll("/root/go-learn/io/file_operator_os/a/cc/dd/ee", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("多层文件夹创建成功")

	//3.创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	file1, err := os.Create("/root/go-learn/io/file_operator_os/a/ab.txt")
	if err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println(file1)

	dir, _ := os.Open("/root/go-learn/io/file_operator_os/nowdir/")
	if err = dir.Chdir(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getwd()) //打印当前的工作目录

	file2, err := os.Create(fileName2) //创建相对路径的文件，是以当前工程为参照的
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	fmt.Println(file2)

	//4.打开文件：
	file3, err := os.Open(fileName1) //默认只读的
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file3)
	/*
	   第一个参数：文件名称
	   第二个参数：文件的打开方式
	       const (
	   // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	       O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	       O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	       O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	       // The remaining values may be or'ed in to control behavior.
	       O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	       O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	       O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	       O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	       O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	   )
	   第三个参数：文件的权限：文件不存在创建文件，需要指定权限
	*/
	file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)

	//5关闭文件，
	file4.Close()

	//6.删除文件或文件夹：
	//删除文件和空文件夹
	err = os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除文件成功。。")
	//删除目录
	err = os.RemoveAll("/Users/ruby/Documents/pro/a/cc")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除目录成功。。")
}

////获取主机名
//fmt.Println(os.Hostname());
//
////获取当前目录
//fmt.Println(os.Getwd());
//
////获取用户ID
//fmt.Println(os.Getuid());
//
////获取有效用户ID
//fmt.Println(os.Geteuid());
//
////获取组ID
//fmt.Println(os.Getgid());
//
////获取有效组ID
//fmt.Println(os.Getegid());
//
////获取进程ID
//fmt.Println(os.Getpid());
//
////获取父进程ID
//fmt.Println(os.Getppid());
//
////获取环境变量的值
//fmt.Println(os.Getenv("GOPATH"));
//
////设置环境变量的值
//os.Setenv("TEST", "test");
//
////改变当前工作目录
//os.Chdir("C:/");
//fmt.Println(os.Getwd());
//
////创建文件
//f1, _ := os.Create("./1.txt");
//defer f1.Close();
//
////修改文件权限
//if err := os.Chmod("./1.txt", 0777); err != nil {
//fmt.Println(err);
//}
//
////修改文件所有者
//if err := os.Chown("./1.txt", 0, 0); err != nil {
//fmt.Println(err);
//}
//
////修改文件的访问时间和修改时间
//os.Chtimes("./1.txt", time.Now().Add(time.Hour), time.Now().Add(time.Hour));
//
////获取所有环境变量
//fmt.Println(strings.Join(os.Environ(), "
