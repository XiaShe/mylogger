package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里写日志相关代码

// FileLogger 文件日志结构体
type FileLogger struct {
	Level       Loglevel // 日志级别
	filePath    string   // 文件保存路径
	fileName    string   // 日志文件保存的文件名
	fileObj     *os.File // 所有日志文件
	errFileObj  *os.File // error级别以上日志文件
	maxFileSize int64    // 每个日志的大小
}

//######################################构造函数#########################################//

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	// 按照文件路径和文件名将将要保存日志的文件打开
	fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile() error {
	// 路径+文件名
	fullFileName := path.Join(f.filePath, f.fileName)
	// 读写文件
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//#####################################日志功能方法#########################################//

// 根据日志级别（loglevel），判断是否需要记录该日志
func (f *FileLogger) enable(loglevel Loglevel) bool {
	return loglevel >= f.Level
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	// 返回文件信息 ---> 一个结构体，包括文件各种信息，如下
	/*
	type FileInfo interface {
		Name() string       // base name of the file 文件名.扩展名 aa.txt
		Size() int64        // 文件大小，字节数 12540
		Mode() FileMode     // 文件权限 -rw-rw-rw-
		ModTime() time.Time // 修改时间 2018-04-13 16:30:53 +0800 CST
		IsDir() bool        // 是否文件夹
		Sys() interface{}   // 基础数据源接口(can return nil)
	}
	 */
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	// fileInfo.Size()返回文件大小，如果当前文件大小 > 等于日志文件设定的最大值 就应该返回true，进行分割
	return fileInfo.Size() >= f.maxFileSize
}

// 切割日志文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	// 拿到当前的日志文件的完整路径
	logName := path.Join(f.filePath, fileInfo.Name())
	// 拼接一个日志文件备份的名字
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr)

	//1. 关闭当前的日志文件
	file.Close()

	//2. 备份一下 rename xx.log -> xx.log.bak202005061422
	os.Rename(logName, newLogName)
	//3. 打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err%v\n", err)
		return nil, err
	}
	//4. 将打开的新日志文件对象赋值给 f.fileObj
	return fileObj, nil
}

// 记录日志的方法，将日志信息写入文件
func (f *FileLogger) log(lv Loglevel, format string, a ...interface{}) {
	if f.enable(lv) {
		// 将传入参数格式化为字符串
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		// 检查日志文件大小，一旦达到设定的文件大小，即创建新的日志文件 ---> 切割日志文件
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj) // 切割日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		// 函数名，文件名，行号
		funcName, fileName, lineNo := getInfo(3)
		// 将日志传入文件
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)

		// 如果要记录的日志大于等于ERROR级别（即error与fatal级别），还要在err日志文件中再记录一遍
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj) // 切割，同上
				if err != nil {
					return
				}
				f.errFileObj = newFile
			}
			// 将error与fatal日志再记录一边，生成日志文件
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getlogString(lv), fileName, funcName, lineNo, msg)

		}
	}
}

//#########################下面是写入文件中不同级别日志内容的方法###############################//

// 传入： 级别  描述字符串 自定义参数
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
