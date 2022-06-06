/**
 * @date: 2022/6/4
 * @desc:
 */

package utils

import (
	"net/http"
	"os"
)

func GetFileContentType(out *os.File) (string, error) {
	// 需要前 512 个字节
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
