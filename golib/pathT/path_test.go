package pathT

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	spath := os.Getenv("GOPATH")
	pathList := strings.Split(spath, string(filepath.ListSeparator))
	if len(pathList) <= 0 {
		t.Fatal("gopath not exits")
	}
	log.Println(pathList)
	gopath := pathList[0]
	log.Println("GOPATH:", gopath)

	gopath = filepath.ToSlash(gopath)
	log.Println("GOPATH ToSlash:", gopath)

	//路径分割符 常量
	log.Println("-----------------路径分割符 常量--------------")
	log.Println("os.PathSeparator:", string(filepath.Separator))
	log.Println("os.PathListSeparator", string(filepath.ListSeparator))

	//返回路径的最后一个元素
	log.Println("-----------------返回路径的最后一个元素--------------")
	log.Println(path.Base(gopath))
	log.Println(path.Base(`./Go/GoWorkSpace/dd`))

	//返回等价的最短路径
	//1.用一个斜线替换多个斜线
	//2.清除当前路径.
	//3.清除内部的..和他前面的元素
	//4.以/..开头的，变成/
	log.Println(path.Clean(`./a/b//c`))

	//返回路径最后一个元素的目录
	log.Println("-----------------返回路径最后一个元素的目录--------------")
	log.Println("目录", path.Dir(`./a/b/e/c.img`))

	//返回路径中的扩展名
	log.Println("扩展名", path.Ext(`./a/b/e/c.img`))

	//返回路径最后一个元素的目录
	log.Println("-----------------判断路径是不是绝对路径--------------")
	log.Println("绝对路径", path.IsAbs(`./a/b/e/c.img`))

	log.Println("-----------------匹配文件名，完全匹配则返回true--------------")
	log.Println(path.Match("*", "a"))

	log.Println("-----------------分割路径中的目录与文件--------------")
	log.Println(path.Split(`./a/b/e/c.img`))

	//https://www.cnblogs.com/jkko123/p/6923962.html
}

func TestFilePath(t *testing.T) {
	var spath string = `./a/b/e/c.img`
	log.Println("path:", spath)
	abs, _ := filepath.Abs(spath)
	log.Println("路径的绝对路径", abs)

	log.Println("Base:", filepath.Base(`./a/b/e/c.img`))

	log.Println("Dir:", filepath.Dir(spath))
	log.Println("Ext:", filepath.Ext(spath))
	log.Println("FromSlash:", filepath.FromSlash(spath))
	log.Println("VolumeName:", filepath.VolumeName("\\\\server\\path\\file"))
	log.Println("VolumeName:", filepath.VolumeName("D://a/b"))

	dir, file := filepath.Split(spath)
	log.Println("Split:", dir, file)

	log.Println(filepath.EvalSymlinks(`1.lnk`))

	filepath.Walk("D:\\Hearthstone", func(path string, info os.FileInfo, err error) error {
		log.Println("遍历", path)
		return nil
	})

	//test ssh
}

type MatchTest struct {
	pattern, s string
	match      bool
	err        error
}

var matchTests = []MatchTest{
	{"**[1-9]bc", "5bc", true, nil},
	{"abc", "abc", true, nil},
	{"*", "abc", true, nil},
	{"*c", "abc", true, nil},
}

func TestMatch(t *testing.T) {
	for _, tt := range matchTests {
		ok, err := path.Match(tt.pattern, tt.s)
		if ok != tt.match || err != tt.err {
			t.Errorf("Match(%#q, %#q) = %v, %v want %v, %v", tt.pattern, tt.s, ok, err, tt.match, tt.err)
		}
	}
}
