package main

/*
func diskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range direnst(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func direnst(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64

	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	diskUsage(nfiles, nbytes)
}
*/
