package utils

import (
	"archive/tar"
	"compress/gzip"
	"crypto_webtool/global"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Tar(src, dst string) (err error) {
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if !fi.Mode().IsRegular() {
			return nil
		}
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}
		global.CW_LOG.Info("成功打包文件：", zap.String(fileName, "共写入："+strconv.FormatInt(n, 10)+" 字节的数据"))
		return nil
	})
}
