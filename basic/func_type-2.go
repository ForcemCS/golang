package main

import "fmt"

// 定义一种函数类型
// 规定：这类函数接收 文件名、文件大小，并返回 error
type FinishHandler func(filename string, size int64) error

// Backup 负责备份
// handler 是备份完成后要执行的函数
func Backup(source string, handler FinishHandler) error {

	fmt.Println("开始备份：", source)

	// 假装这里真的执行了备份
	fmt.Println("正在压缩文件...")
	fmt.Println("正在写入备份目录...")

	// 假装最终生成了这个文件
	filename := "backup-20260825.tar.gz"
	var size int64 = 1024 * 1024 * 500

	fmt.Println("备份完成")

	// 备份完成之后
	// 调用外部传进来的处理函数
	if handler != nil {
		err := handler(filename, size)
		if err != nil {
			return err
		}
	}

	return nil
}

func main35() {

	// 这里完全按照你说的方式：
	// var 变量 类型 = 一整个匿名函数
	var finishHandler FinishHandler = func(filename string, size int64) error {

		fmt.Println("===== 执行备份完成后的处理 =====")

		fmt.Println("备份文件：", filename)
		fmt.Println("文件大小：", size)

		// 判断备份文件是否太大
		if size > 1024*1024*1024 {
			return fmt.Errorf("备份文件超过 1GB")
		}

		fmt.Println("记录备份日志")
		fmt.Println("发送备份成功通知")
		fmt.Println("更新备份状态")

		return nil
	}

	// 把函数变量传给 Backup
	err := Backup(
		"/var/lib/mysql",
		finishHandler,
	)

	if err != nil {
		fmt.Println("备份失败：", err)
	}
}
