package helper

import ( 
    "fmt"
    "log"
    "os"
) 

func GetByteCount(filename string) (int64, error) {
    // Get the fileinfo 
    fileInfo, err := os.Stat(filename) 
  
    // Checks for the error 
    if err != nil { 
        log.Fatal(err) 
    }  
  
    // Gives the size of the file in bytes 
    fileSize := fileInfo.Size() 
    fmt.Println("Size of the file:", fileSize)
    return fileSize, nil
}