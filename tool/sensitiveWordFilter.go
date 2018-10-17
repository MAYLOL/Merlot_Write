package tool

import (
	"merlot_write/model"
	"os"
	"bufio"
	"strings"
	"io"
	"fmt"
)

type sensitiveWord struct {
	//词库版本
	version   string
	//词数量
	amount    int
}

func init() {
	ViewVersion()
}

//读取文件词库版本，如果版本有变化，那么更新敏感词库，否则不变
func ViewVersion(){
	//读取服务器version


	//  if version ！= newVersion
	//   UpdateSensitiveWord()
	//update

     UpdateSensitiveWord()

	//else    使用之前读取的敏感词库
	//
}


func UpdateSensitiveWord(){
	//1.将增加的词文件读取出来，然后加入到服务器当前使用的词库中来

	//2.或者选择直接替换词库文件到最新版本

	//3.读取词库
	ReadLine("CensorWords.txt", Print)
}


//读取词库
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

func Print(line string) {
	fmt.Println(line)
}

//

func HasSensitiveWords(article model.Article) bool{

  //传进来article
  //返回有没有敏感词，以及哪个词是敏感词
  //




  return  true

}
