# make_Folder_from_list  

## Description  
Make folder from text file.  
**Note**  
An error will occur if the Windows folder contains characters that cannot be used.  

## Usage  
```
$ go run main.go <A text file with the folder name>
```

## Example  
```
D:\tool\Go\make_Folder_from_list>go run main.go folder_list.txt
Error:
mkdir あｓふぁ\ｆさｄ: The system cannot find the path specified.
Error:
mkdir /あああああああ: Cannot create a file when that file already exists.

Done!
```

## Requirements  
- Windows
- The folder list is a \*.txt file in UTF8 (without BOM)
