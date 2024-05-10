package watch

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io"
	"os"
	"time"
)

func watch(watchFile string, process func(data []byte)) {

	go func() {
		file, err := os.Open(watchFile)
		if err != nil {
			panic(fmt.Sprintf("open file err(%v)", err))
		}
		defer file.Close()

		fileInfo, _ := os.Stat(watchFile)
		lastModifyUnix := fileInfo.ModTime().Unix()
		lastFileSize := fileInfo.Size()

		ticker := time.NewTicker(time.Second * 5)
		for {
			<-ticker.C

			newFileInfo, _ := os.Stat(watchFile)
			newModifyUnix := newFileInfo.ModTime().Unix()
			newFileSize := newFileInfo.Size()
			if newModifyUnix > lastModifyUnix && newFileSize > lastFileSize {
				_, err = file.Seek(lastFileSize, io.SeekStart)
				if err != nil {
					spew.Dump("err1", err)
					return
				}
				buffer := make([]byte, newFileSize-lastFileSize)
				_, err = file.Read(buffer)
				if err != nil {
					spew.Dump("err2", err, newFileSize, lastFileSize)
				}
				// todo notice
				process(buffer)

				lastFileSize = newFileSize
				lastModifyUnix = newModifyUnix
			}
			fmt.Println(newModifyUnix, newFileSize, "watching...")
		}
	}()
}
