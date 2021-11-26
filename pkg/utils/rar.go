package utils

import (
	"encoding/xml"
	"frame/consts"
	"frame/internal/entity"
	"io/ioutil"
	"strings"
	"github.com/mholt/archiver/v3"
)

func GetRarInfo(url string, ManifestWhole entity.ManifestInterface) error {
	z := archiver.NewRar()
	if err := z.Walk(url, func(f archiver.File) error {
		defer f.Close()

		Split := strings.Split(f.Name(), "/")
		if len(Split) == 2 && Split[1] == consts.XMLNAME {
			data, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}
			err = xml.Unmarshal(data, &ManifestWhole)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

