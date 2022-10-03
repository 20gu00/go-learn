package main

import (
	"io"
	"os"
)

//Copy（dst,src） 为复制src 全部到 dst 中。
//
//CopyN(dst,src,n) 为复制src 中 n 个字节到 dst。
//
//CopyBuffer（dst,src,buf）为指定一个buf缓存区，以这个大小完全复制。

func main() {
	copyFile2("/root/go-learn/io/file_copy/test1", "/root/go-learn/io/file_copy/test2")
}
func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}

//无论是哪个copy方法最终都是由copyBuffer（）这个私有方法实现的
//func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
//    // If the reader has a WriteTo method, use it to do the copy.
//    // Avoids an allocation and a copy.
//    if wt, ok := src.(WriterTo); ok {
//        return wt.WriteTo(dst)
//    }
//    // Similarly, if the writer has a ReadFrom method, use it to do the copy.
//    if rt, ok := dst.(ReaderFrom); ok {
//        return rt.ReadFrom(src)
//    }
//    if buf == nil {
//        size := 32 * 1024
//        if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
//            if l.N < 1 {
//                size = 1
//            } else {
//                size = int(l.N)
//            }
//        }
//        buf = make([]byte, size)
//    }
//    for {
//        nr, er := src.Read(buf)
//        if nr > 0 {
//            nw, ew := dst.Write(buf[0:nr])
//            if nw > 0 {
//                written += int64(nw)
//            }
//            if ew != nil {
//                err = ew
//                break
//            }
//            if nr != nw {
//                err = ErrShortWrite
//                break
//            }
//        }
//        if er != nil {
//            if er != EOF {
//                err = er
//            }
//            break
//        }
//    }
//    return written, err
//}
