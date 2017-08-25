package godirwalk

import (
	"os"
)

// Dirent stores the name and file system mode type of discovered file system
// entries.
type Dirent struct {
	name     string
	modeType os.FileMode
}

// Name returns the basename of the file system entry.
func (de Dirent) Name() string { return de.name }

// ModeType returns the mode bits that specify the file system node type. We
// could make our own enum-like data type for encoding the file type, but Go's
// runtime already gives us architecture independent file modes, as discussed in
// `os/types.go`:
//
//    Go's runtime FileMode type has same definition on all systems, so that
//    information about files can be moved from one system to another portably.
func (de Dirent) ModeType() os.FileMode { return de.modeType }

// IsDir returns true if and only if the Dirent represents a file system
// directory. Not that some operating systems, more than one file mode bit may
// be set for a node. For instance, on some operating systems, a symbolic link
// that points to a directory will have both the directory and the symbolic link
// bits set.
func (de Dirent) IsDir() bool { return de.modeType&os.ModeDir != 0 }

// Dirents represents a slice of Dirent pointers, which are sortable by
// name. This type satisfies the `sort.Interface` interface.
type Dirents []*Dirent

// Len returns the count of Dirent structures in the slice.
func (l Dirents) Len() int { return len(l) }

// Less returns true if and only if the Name of the element specified by the
// first index is lexicographically less than that of the second index.
func (l Dirents) Less(i, j int) bool { return l[i].name < l[j].name }

// Swap exchanges the two Dirent entries specified by the two provided indexes.
func (l Dirents) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

// ReadDirents returns a sortable slice of pointers to Dirent structures,
// representing the immediate descendants of the specified directory. If the
// specified directory is a symbolic link, it will be resolved.
//
//    children, err := godirwalk.ReadDirents(osPathname, 0)
//    if err != nil {
//    	return nil, errors.Wrap(err, "cannot get list of directory children")
//    }
//    sort.Sort(children)
//    for _, child := range children {
//        fmt.Printf("%s %s\n", child.ModeType, child.Name)
//    }
func ReadDirents(osDirname string, max int) (Dirents, error) {
	return readdirents(osDirname, max)
}

// ReadDirnames returns a slice of strings, representing the immediate
// descendants of the specified directory. If the specified directory is a
// symbolic link, it will be resolved.
//
//    children, err := godirwalk.ReadDirnames(osPathname, 0)
//    if err != nil {
//    	return nil, errors.Wrap(err, "cannot get list of directory children")
//    }
//    sort.Strings(children)
//    for _, child := range children {
//        fmt.Printf("%s\n", child)
//    }
func ReadDirnames(osDirname string, max int) ([]string, error) {
	return readdirnames(osDirname, max)
}