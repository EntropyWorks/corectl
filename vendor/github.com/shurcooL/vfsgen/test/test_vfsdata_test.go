// Code generated by vfsgen; DO NOT EDIT

package test_test

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	mustUnmarshalTextTime := func(text string) time.Time {
		var t time.Time
		err := t.UnmarshalText([]byte(text))
		if err != nil {
			panic(err)
		}
		return t
	}

	fs := _vfsgen_fs{
		"/": &_vfsgen_dirInfo{
			name:    "/",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderA": &_vfsgen_dirInfo{
			name:    "folderA",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderA/file1.txt": &_vfsgen_fileInfo{
			name:    "file1.txt",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
			content: []byte("\x53\x74\x75\x66\x66\x20\x69\x6e\x20\x2f\x66\x6f\x6c\x64\x65\x72\x41\x2f\x66\x69\x6c\x65\x31\x2e\x74\x78\x74\x2e"),
		},
		"/folderA/file2.txt": &_vfsgen_fileInfo{
			name:    "file2.txt",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
			content: []byte("\x53\x74\x75\x66\x66\x20\x69\x6e\x20\x2f\x66\x6f\x6c\x64\x65\x72\x41\x2f\x66\x69\x6c\x65\x32\x2e\x74\x78\x74\x2e"),
		},
		"/folderB": &_vfsgen_dirInfo{
			name:    "folderB",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderB/folderC": &_vfsgen_dirInfo{
			name:    "folderC",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
		},
		"/folderB/folderC/file3.txt": &_vfsgen_fileInfo{
			name:    "file3.txt",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
			content: []byte("\x53\x74\x75\x66\x66\x20\x69\x6e\x20\x2f\x66\x6f\x6c\x64\x65\x72\x42\x2f\x66\x6f\x6c\x64\x65\x72\x43\x2f\x66\x69\x6c\x65\x33\x2e\x74\x78\x74\x2e"),
		},
		"/not-worth-compressing-file.txt": &_vfsgen_fileInfo{
			name:    "not-worth-compressing-file.txt",
			modTime: mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
			content: []byte("\x49\x74\x73\x20\x6e\x6f\x72\x6d\x61\x6c\x20\x63\x6f\x6e\x74\x65\x6e\x74\x73\x20\x61\x72\x65\x20\x68\x65\x72\x65\x2e"),
		},
		"/sample-file.txt": &_vfsgen_compressedFileInfo{
			name:             "sample-file.txt",
			modTime:          mustUnmarshalTextTime("0001-01-01T00:00:00Z"),
			uncompressedSize: 189,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\xc9\xc8\x2c\x56\x48\xcb\xcc\x49\x55\x48\xce\xcf\x2d\x28\x4a\x2d\x2e\x4e\x2d\x56\x28\x4f\xcd\xc9\xd1\x53\x70\xca\x49\x1c\xd4\x20\x43\x11\x10\x00\x00\xff\xff\xe7\x47\x81\x3a\xbd\x00\x00\x00"),
		},
	}

	fs["/"].(*_vfsgen_dirInfo).entries = []os.FileInfo{
		fs["/folderA"].(os.FileInfo),
		fs["/folderB"].(os.FileInfo),
		fs["/not-worth-compressing-file.txt"].(os.FileInfo),
		fs["/sample-file.txt"].(os.FileInfo),
	}
	fs["/folderA"].(*_vfsgen_dirInfo).entries = []os.FileInfo{
		fs["/folderA/file1.txt"].(os.FileInfo),
		fs["/folderA/file2.txt"].(os.FileInfo),
	}
	fs["/folderB"].(*_vfsgen_dirInfo).entries = []os.FileInfo{
		fs["/folderB/folderC"].(os.FileInfo),
	}
	fs["/folderB/folderC"].(*_vfsgen_dirInfo).entries = []os.FileInfo{
		fs["/folderB/folderC/file3.txt"].(os.FileInfo),
	}

	return fs
}()

type _vfsgen_fs map[string]interface{}

func (fs _vfsgen_fs) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *_vfsgen_compressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &_vfsgen_compressedFile{
			_vfsgen_compressedFileInfo: f,
			gr: gr,
		}, nil
	case *_vfsgen_fileInfo:
		return &_vfsgen_file{
			_vfsgen_fileInfo: f,
			Reader:           bytes.NewReader(f.content),
		}, nil
	case *_vfsgen_dirInfo:
		return &_vfsgen_dir{
			_vfsgen_dirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// _vfsgen_compressedFileInfo is a static definition of a gzip compressed file.
type _vfsgen_compressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *_vfsgen_compressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *_vfsgen_compressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *_vfsgen_compressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *_vfsgen_compressedFileInfo) Name() string       { return f.name }
func (f *_vfsgen_compressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *_vfsgen_compressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *_vfsgen_compressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *_vfsgen_compressedFileInfo) IsDir() bool        { return false }
func (f *_vfsgen_compressedFileInfo) Sys() interface{}   { return nil }

// _vfsgen_compressedFile is an opened compressedFile instance.
type _vfsgen_compressedFile struct {
	*_vfsgen_compressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *_vfsgen_compressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f._vfsgen_compressedFileInfo.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *_vfsgen_compressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case os.SEEK_SET:
		f.seekPos = 0 + offset
	case os.SEEK_CUR:
		f.seekPos += offset
	case os.SEEK_END:
		f.seekPos = f._vfsgen_compressedFileInfo.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *_vfsgen_compressedFile) Close() error {
	return f.gr.Close()
}

// _vfsgen_fileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type _vfsgen_fileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *_vfsgen_fileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *_vfsgen_fileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *_vfsgen_fileInfo) NotWorthGzipCompressing() {}

func (f *_vfsgen_fileInfo) Name() string       { return f.name }
func (f *_vfsgen_fileInfo) Size() int64        { return int64(len(f.content)) }
func (f *_vfsgen_fileInfo) Mode() os.FileMode  { return 0444 }
func (f *_vfsgen_fileInfo) ModTime() time.Time { return f.modTime }
func (f *_vfsgen_fileInfo) IsDir() bool        { return false }
func (f *_vfsgen_fileInfo) Sys() interface{}   { return nil }

// _vfsgen_file is an opened file instance.
type _vfsgen_file struct {
	*_vfsgen_fileInfo
	*bytes.Reader
}

func (f *_vfsgen_file) Close() error {
	return nil
}

// _vfsgen_dirInfo is a static definition of a directory.
type _vfsgen_dirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *_vfsgen_dirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *_vfsgen_dirInfo) Close() error               { return nil }
func (d *_vfsgen_dirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *_vfsgen_dirInfo) Name() string       { return d.name }
func (d *_vfsgen_dirInfo) Size() int64        { return 0 }
func (d *_vfsgen_dirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *_vfsgen_dirInfo) ModTime() time.Time { return d.modTime }
func (d *_vfsgen_dirInfo) IsDir() bool        { return true }
func (d *_vfsgen_dirInfo) Sys() interface{}   { return nil }

// _vfsgen_dir is an opened dir instance.
type _vfsgen_dir struct {
	*_vfsgen_dirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *_vfsgen_dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == os.SEEK_SET {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d._vfsgen_dirInfo.name)
}

func (d *_vfsgen_dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d._vfsgen_dirInfo.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d._vfsgen_dirInfo.entries)-d.pos {
		count = len(d._vfsgen_dirInfo.entries) - d.pos
	}
	e := d._vfsgen_dirInfo.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
