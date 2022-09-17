package main

func removeDuplicates(inputStream chan string, outputStream chan string) {
	workMap := make(map[string]bool, len(inputStream))

	for i := range inputStream {
		if ok, _ := workMap[i]; ok {
			continue
		} else {
			workMap[i] = true
			outputStream <- i
		}
	}

	close(outputStream)

}
