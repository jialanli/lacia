package lacia

import "testing"

func TestReadFile(t *testing.T) {
	src := "./example/t2.txt"
	data, err := ReadFile(src, 1024*5)
	if err != nil {
		t.Fatalf("ReadFile err: %v", err)
	}

	t.Logf("ReadFile end: %s", string(data))

	data, err = Read(src)
	if err != nil {
		t.Fatalf("Read err: %v", err)
	}

	t.Logf("Read end: %s", string(data))
}

func TestReadWrite(t *testing.T) {
	src := "./example/t2.txt"
	dst := "./example/t.txt"
	err := WriteFileAfterRead(dst, src, true, 0666)
	t.Logf("WriteFileAfterRead end: %v", err)
}

func TestFileMD5(t *testing.T) {
	src := "./example/t2.txt"
	md5, err := GetFileMD5(src)
	if err != nil {
		t.Fatalf("GetFileMD5 err: %v", err)
	}

	t.Logf("md5 is: %s", md5)

	md5, err = GetFileMd5(src)
	if err != nil {
		t.Fatalf("GetFileMd5 err: %v", err)
	}

	t.Logf("md5 is: %s", md5)
}
