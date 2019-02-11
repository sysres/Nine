package fs_test

import (
	"testing"

	"github.com/madlambda/Nine/sys/fs"
	"github.com/madlambda/spells/assert"
)

func basicFileTest(t *testing.T, name string, data []byte, expectedErr error) {
	memfs := fs.NewInMem()
	assert.EqualErrs(t, memfs.Mkfile(name, data), expectedErr, "mkfile")
	fdata, err := memfs.Data(name)
	assert.EqualErrs(t, err, expectedErr, "data")
	assert.EqualStrings(t, name, fdata.Name(), "name mismatch")
}

func TestInMemFiles(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []byte
		err  error
	}{
		{"some", []byte(`thing`), nil},
		{"secrets", []byte{0, 0, 0, 0, 0, 0, 0, 0}, nil},
		{"", []byte("anything"), fs.ErrInvalidFilename},
	} {
		basicFileTest(t, tc.name, tc.data, tc.err)
	}
}

func testBasicDirs(t *testing.T, dirname string, files map[string]string, expectedErr error) {
	mem := fs.NewInMem()
	err := mem.Mkdir(dirname)
	assert.EqualErrs(t, err, expectedErr, "mkdir")

	if err != nil {
		return // directory not created
	}

	testdir, err := mem.Data(dirname)
	assert.EqualErrs(t, err, expectedErr, "data")

	for fname, fvalue := range files {
		assert.NoError(t, testdir.Dir().Mkfile(fname, []byte(fvalue)))
	}

	counter := 0
	err = mem.Walk(func(f *fs.FileData) error {
		counter++
		return nil
	})
	assert.NoError(t, err)
	assert.EqualInts(t, counter, 1)

	counter = 0
	err = testdir.Dir().Walk(func(f *fs.FileData) error {
		counter++
		return nil
	})

	assert.NoError(t, err)
	assert.EqualInts(t, counter, len(files))
}

func TestInMemDirs(t *testing.T) {
	for _, tc := range []struct {
		dirname string
		files   map[string]string
		err     error
	}{
		{
			dirname: "",
			err:     fs.ErrInvalidFilename,
		},
		{
			dirname: "test",
			files: map[string]string{
				"a": "abc",
				"b": "bcd",
				"c": "cde",
			},
		},
		{
			dirname: "opt",
			files:   map[string]string{},
		},
		{
			dirname: "etc",
			files: map[string]string{
				"passwd": "root",
				"shadow": "root::::::",
				"groups": "root",
			},
		},
	} {
		testBasicDirs(t, tc.dirname, tc.files, tc.err)
	}
}
