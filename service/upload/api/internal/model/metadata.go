/**
 * @date: 2022/5/21
 * @desc:
 */

package model

// FileMetadata : 文件元信息结构
type FileMetadata struct {
	FileHash        string
	FileName        string
	FileContentType string
	FileSize        int64
	Location        string
	UploadAt        string
}
