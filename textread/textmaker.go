package textread

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func mkdirifnotexist(path string) {
	_, err := os.ReadDir(path)
	if !os.IsExist(err) {

		os.Mkdir(path, 0644)

	}

}
func isdirexist(path string) bool {
	_, err := os.ReadFile(path)
	return !os.IsExist(err)
}
func ReadText(dirpath string, filename string) ([]byte, error) {
	filepath := path.Join(dirpath, filename)
	if isdirexist(filepath) {
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {

		err := fmt.Errorf("nofile")
		return nil, err
	}

}

func SaveText(dirpath string, filename string, data []byte) error {

	filepath := path.Join(dirpath, filename)
	mkdirifnotexist(dirpath)
	err := ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil

}

func SaveTmp(data []byte) {

	SaveText("./tmp", "~tmp.txt", data)

}

func StartRead() []byte {
	data, err := ReadText("./tmp", "~tmp.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data

}
