package conf

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

var confPath *string

// Files 配置文件
type file struct {
	lock          sync.Mutex
	fileNameValue map[string]string
	base          string
}

var client *file

func init() {
	confPath = flag.String("conf", "", "path of the config files")
}

// Init init
func Init() error {
	// 读取文件
	base := *confPath
	base = filepath.FromSlash(base)
	fi, err := os.Stat(base)
	if err != nil {
		return err
	}
	var paths []string
	if fi.IsDir() {
		files, err := ioutil.ReadDir(base)
		if err != nil {
			return err
		}
		for _, f := range files {
			if !f.IsDir() && !strings.HasPrefix(filepath.Base(f.Name()), ".") {
				paths = append(paths, path.Join(base, f.Name()))
			}
		}
	} else {
		paths = append(paths, base)
	}
	// 读取内容
	raw := make(map[string]string)
	for _, fpath := range paths {
		k := path.Base(fpath)
		vb, err := ioutil.ReadFile(fpath)
		if err != nil {
			continue
		}
		v := string(vb)
		raw[k] = v

	}
	f := &file{
		fileNameValue: raw,
		base:          base,
	}
	client = f
	//
	go f.daemon()
	//监听
	//
	return nil
}

func (f *file) daemon() {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}
	if err = fswatcher.Add(f.base); err != nil {
		return
	}

	for event := range fswatcher.Events {
		switch event.Op {
		// use vim edit config will trigger rename
		case fsnotify.Write, fsnotify.Create:
			f.reloadFile(event.Name)
		case fsnotify.Chmod:
		default:
			log.Printf("unsupport event %s ingored", event)
		}
	}
}
func (f *file) reloadFile(file string) {
	k := filepath.Base(file)
	vb, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	f.fileNameValue[k] = string(vb)

}
