package lacia

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
file opts contains:
	count/list
	file create
	file perm
	file dir/path
	file md5
	file convert
	file copy/move/remove
	file zip/unzip
	file readwrite
*/

// *************************** file count/list ***************************

func WalkDirBySuffix(dirPth, suffix string) (files []string, err error) {
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func GetFiles(folder string) (filesList []string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GetFiles(folder + "/" + file.Name())
		} else {
			filesList = append(filesList, file.Name())
		}
	}

	return
}

// 统计目录下的文件数量
func FilesCountAndFiles(dirPath string, recursion bool) (files []string, err error) {
	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		log.Printf("WalkDir: %s ---> %s isDir: %v\n", dirPath, path, d.IsDir())

		idDir := d.IsDir()
		if (idDir && recursion) || !idDir {
			files = append(files, path)
		}

		return nil
	})

	return
}

func FilesCount(dirPath string, recursion bool) (count int, err error) {
	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		log.Printf("WalkDir: %s ---> %s isDir: %v\n", dirPath, path, d.IsDir())

		idDir := d.IsDir()
		if (idDir && recursion) || !idDir {
			count++
		}

		return nil
	})

	return
}

// *************************** file create ***************************

func CreateDir(dir string, perm os.FileMode) (bool, error) {
	if err := os.MkdirAll(dir, perm); err != nil {
		return false, err
	}

	return true, nil
}

func SafeCreateFile(file string, perm os.FileMode) (err error) {
	dir, _ := filepath.Split(file)
	_, err = CreateDir(dir, perm)
	if err != nil {
		return
	}

	_, err = CreateFile(file)
	return
}

func CreateFile(fileName string) (*os.File, error) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return dstFile, nil
}

// *************************** file perm ***************************

func ChmodFile(fileName string, mode os.FileMode) error {
	err := os.Chmod(fileName, mode)
	if err != nil {
		return err
	}

	return nil
}

func FileIsExecutable(fileName string) (isExecutable bool) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return
	}

	file_mode := fileInfo.Mode()

	perm := file_mode.Perm()

	// 73: 000 001 001 001
	flag := perm & os.FileMode(73)

	if uint32(flag) == uint32(73) {
		isExecutable = true
	}

	return
}

// *************************** file dir/path ***************************

func IsDir(path string) (isDir bool, err error) {
	f, err := os.Stat(path)
	if err != nil {
		return
	}

	if f.IsDir() {
		isDir = true
	}

	return
}

func IsFile(path string) (isFile bool, err error) {
	f, err := os.Stat(path)
	if err != nil {
		return
	}

	if !f.IsDir() {
		isFile = true
	}

	return
}

func CurrentDir() string {
	cwdPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return strings.Replace(cwdPath, "\\", "/", -1)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	return false, err
}

// *************************** file md5 ***************************

func GetFileMd5(filename string) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var result bytes.Buffer
	_, err = io.Copy(&result, file)
	if err != nil {
		return "", err
	}

	checksum := md5.Sum(result.Bytes())
	return fmt.Sprintf("%x", checksum), nil
}

func GetFileMD5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	md5Hash := md5.New()
	if _, err = io.Copy(md5Hash, file); err != nil {
		return "", err
	}

	hashInBytes := md5Hash.Sum(nil)
	return fmt.Sprintf("%x", hashInBytes), nil
}

func GetMD5ByFile(file *os.File) (string, error) {
	file.Seek(0, 0)
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", err
	}

	hashInBytes := md5Hash.Sum(nil)
	return fmt.Sprintf("%x", hashInBytes), nil
}

func GetMD5ByFileBody(body string) (result string, err error) {
	var buf bytes.Buffer
	_, err = buf.Write([]byte(body))
	checksum := md5.Sum(buf.Bytes())
	result = fmt.Sprintf("%x", checksum)
	return
}

func ReaderToBytes(reader io.Reader) ([]byte, error) {
	var result bytes.Buffer

	_, err := io.Copy(&result, reader)
	if err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

// *************************** file convert ***************************

func FileToBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// *************************** file copy/move/remove ***************************

func CommandCP(src, dst string) (err error) {
	if src == dst {
		return nil
	}

	_, err = ExecCommand("cp", src, dst)
	return
}

func CommandMV(src, dst string) (err error) {
	if src == dst {
		return nil
	}

	_, err = ExecCommand("mv", src, dst)
	return
}

func CopyFile(src, dst string) (size int64, err error) {
	log.Printf("ready copy '%s' to '%s'\n", src, dst)

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0644)
	if err != nil {
		err = fmt.Errorf("failed to open source file: %v", err)
		return
	}
	defer srcFile.Close()

	info, err := srcFile.Stat()
	if err != nil {
		err = fmt.Errorf("failed to get file info: %v\n", err)
		return
	}

	size = info.Size()

	if info.IsDir() {
		err = fmt.Errorf("source is a directory, does not support copying directories recursively.")
		return
	}

	dstFile, err := createDstFile(dst, info)
	if err != nil {
		err = errors.New(fmt.Sprintf("failed to create destination file: %v", err))
		return
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("failed to copy file: %v\n", err))
		return
	}

	log.Printf("successfully copied '%s' to '%s'\n", src, dst)
	return
}

func createDstFile(dst string, info os.FileInfo) (*os.File, error) {
	dir := filepath.Dir(dst)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination directory: %w", err)
	}

	dstFile, err := CreateFile(dst)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %w", err)
	}

	err = dstFile.Chmod(info.Mode())
	if err != nil {
		dstFile.Close()
		return nil, fmt.Errorf("failed to set destination file permissions: %w", err)
	}

	return dstFile, nil
}

func MoveFile(src, dst string) error {
	return os.Rename(src, dst)
}

func RemoveFile(file string) error {
	return os.Remove(file)
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// *************************** file zip/unzip/compress ***************************

func CompressData(data []byte) (cpsData string, err error) {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	_, err = writer.Write(data)
	defer func() {
		err = writer.Close()
	}()

	if err != nil {
		return
	}

	cpsData = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

func Unzip(zipFile, unzipDir string, infoOut bool) (paths []string, err error) {
	zipReader, _ := zip.OpenReader(zipFile)
	for i, file := range zipReader.Reader.File {

		if infoOut {
			log.Printf("decompressing %d/%d. name=%s, isDir=%v, size=%d.\n",
				i+1, len(zipReader.Reader.File), file.Name, file.FileInfo().IsDir(), file.FileInfo().Size())
		}

		err = func(i int, file *zip.File) error {
			zippedFile, zErr := file.Open()
			if zErr != nil {
				return zErr
			}

			defer zippedFile.Close()

			extractedFilePath := filepath.Join(unzipDir, file.Name)

			if file.FileInfo().IsDir() {
				mkErr := os.MkdirAll(extractedFilePath, file.Mode())
				if infoOut {
					log.Printf("mkdir %s %v\n", file.Name, mkErr)
				}
				return mkErr
			}

			return func() error {
				outputFile, fErr := os.OpenFile(extractedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

				if fErr != nil {
					return fErr
				}
				defer outputFile.Close()

				paths = append(paths, outputFile.Name())

				_, fErr = io.Copy(outputFile, zippedFile)
				if fErr != nil {
					return fErr
				}

				return nil
			}()

		}(i, file)

		if err != nil {
			return
		}
	}

	return
}

func zipFile(srcPaths []string, outputPath string, useBasePathInZip, infoOut bool, outputPathPerm os.FileMode) error {
	if len(srcPaths) == 0 {
		return nil
	}

	file, openErr := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, outputPathPerm)
	if openErr != nil {
		return openErr
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	for _, path := range srcPaths {

		info, err := os.Stat(path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			err = addFilesToDirectory(zipWriter, path, "", useBasePathInZip, infoOut)
			if err != nil {
				return err
			}

			continue
		}

		if err = compressFile(zipWriter, path, useBasePathInZip); err != nil {
			return fmt.Errorf("add file %s to zip failed: %s", path, err)
		}
	}

	return nil
}

func addFilesToDirectory(zw *zip.Writer, dir, baseInZip string, useBasePathInZip, infoOut bool) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	var newBaseInZip string
	for _, fileInfo := range files {
		if useBasePathInZip {
			newBaseInZip = filepath.Join(baseInZip, fileInfo.Name())
		}

		newFullPath := filepath.Join(dir, fileInfo.Name())

		if fileInfo.IsDir() {
			if err = addFilesToDirectory(zw, newFullPath, newBaseInZip, useBasePathInZip, infoOut); err != nil {
				return err
			}

			if infoOut {
				log.Printf("addToZip success: dir=%s, newFullPath=%s, newBaseInZip=%s\n", fileInfo.Name(), newFullPath, newBaseInZip)
			}
			continue
		}

		// 处理单个文件
		if err = compressFile(zw, newFullPath, useBasePathInZip); err != nil {
			return err
		}
		if infoOut {
			log.Printf("addToZip success: file=%s, newFullPath=%s, newBaseInZip=%s\n", fileInfo.Name(), newFullPath, newBaseInZip)
		}
	}

	return nil
}

func compressFile(zw *zip.Writer, srcFile string, useBasePathInZip bool) error {
	fileToZip, err := os.Open(srcFile)
	if err != nil {
		log.Fatalf("compressFile failed to open %s: %v", srcFile, err)
		return err
	}
	defer fileToZip.Close()

	var zipFileWriter io.Writer

	if !useBasePathInZip {
		info, err := fileToZip.Stat()
		if err != nil {
			log.Printf("failed to open file %s: %v\n", srcFile, err)
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Printf("failed to create file header for %s: %v\n", srcFile, err)
			return err
		}

		header.Name = filepath.Base(srcFile)
		zipFileWriter, err = zw.CreateHeader(header)
		if err != nil {
			return err
		}
	} else {
		zipFileWriter, err = zw.Create(srcFile)
		if err != nil {
			return err
		}
	}

	_, err = io.Copy(zipFileWriter, fileToZip)
	return err
}

// *************************** file readwrite ***************************

func Read(filePath string) (bs []byte, err error) {
	return ioutil.ReadFile(filePath)
}

func ReadFile(filePath string, bufLen int64) (chunks []byte, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	if bufLen <= 0 {
		bufLen = 5 * 1024
	}

	reader := bufio.NewReader(f)
	dataByte := make([]byte, bufLen)
	for {
		var n int
		n, err = reader.Read(dataByte)
		if err != nil || 0 == n {
			break
		}

		chunks = append(chunks, dataByte[:n]...)
	}

	if err != nil {
		if isEOF := strings.Compare(err.Error(), io.EOF.Error()); isEOF == 0 {
			err = nil
		}
	}

	return
}

func WriteContentToFile(filePath string, content string, perm os.FileMode) (written int, md5Str string, err error) {
	dir := filepath.Dir(filePath)
	_, _ = CreateDir(dir, perm)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	written, err = writer.WriteString(content)
	if err != nil {
		return
	}

	err = writer.Flush()
	if err != nil {
		return
	}

	md5Str, err = GetFileMD5(filePath)
	return
}

func WriteFileAfterRead(writerFilePath, readFilePath string, append bool, perm os.FileMode) (err error) {
	srcBody, err := os.OpenFile(readFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer srcBody.Close()

	stat, err := srcBody.Stat()
	if err != nil {
		return
	}

	var flag = os.O_CREATE | os.O_WRONLY
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}

	f, err := os.OpenFile(writerFilePath, flag, perm)
	if err != nil {
		return
	}
	defer f.Close()

	//buf := &bytes.Buffer{}
	_, err = io.CopyN(f, srcBody, stat.Size())
	return
}

func WriteFileAfterReadBody(writerFilePath string, body io.Reader) (err error) {

	f, err := os.OpenFile(writerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}

	defer f.Close()
	bs := make([]byte, 5*1024)
	writer := bufio.NewWriter(f)
	for {
		var read int
		switch read, err = body.Read(bs[:]); true {
		case read < 0:
			return
		case read == 0, err == io.EOF:
			return writer.Flush()
		case read > 0:
			_, err = writer.Write(bs[:read])
			if err != nil {
				return
			}
		}
	}

	return
}
