package utils

import (
	"archive/zip"
	"encoding/xml"
	"frame/consts"
	"frame/internal/entity"
	"io/ioutil"
)

func GetZipInfo(url string, ManifestWhole entity.ManifestInterface) error {
	fr, err := zip.OpenReader(url)
	if err != nil {
		return err
	}
	defer fr.Close()

	for _, f := range fr.Reader.File {
		// 读取压缩包中 xml 文件内容, 可以自定义更改
		if f.Name == consts.XMLNAME {
			rc, err := f.Open()
			defer rc.Close()

			if err != nil {
				return err
			}
			data, err := ioutil.ReadAll(rc)
			err = xml.Unmarshal(data, &ManifestWhole)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

