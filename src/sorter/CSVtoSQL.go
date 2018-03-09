package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("check file ....")
	var inputByte bytes.Buffer
	var flag bool = true
	skillfolder := `/Users/King/Documents/AG001_CSV`
	//skillfolder := `./`
	_, err := os.Stat(skillfolder + "pre_static_text.csv")
	if err != nil || os.IsExist(err) {
		fmt.Println("ERROE: file pre_static_text.csv is not exist")
		return
	}
	// 获取所有文件
	files, _ := ioutil.ReadDir(skillfolder)
	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".csv") || file.Name() == "pre_static_text.csv" {
			continue
		} else {
			var index []string
			fmt.Println(file.Name() + " :")
			res_file, err := os.Open(skillfolder + "/" + file.Name())
			if err != nil {
				fmt.Println("ERROE: Failed to open the input file", err)
			}
			defer res_file.Close()
			inputReader := bufio.NewReader(res_file)
			var primaryKey []string
			var unKey []string
			for n := 0; ; n++ {
				//开始读取文件
				inputContent, isPrefix, readerError := inputReader.ReadLine()
				pkStr := ""
				unStr := ""
				if inputContent != nil {
					inputByte.Write(inputContent)
					//判断一行是否超出预设缓冲区
					if !isPrefix {
						dataString := mahonia.NewDecoder("gbk").ConvertString(inputByte.String())
						if n == 3 {
							index = strings.Split(dataString, ",")
						} else {
							dataArray := strings.Split(dataString, ",")
							if n > 4 {
								for key, value := range index {
									if (value == "pk" || value == "unique") && len(dataArray[key]) == 0 {
										fmt.Sprintf("ERROE: %x is null line: %d", value, n)
									} else if value == "pk" {
										pkStr += dataArray[key]
									} else if value == "unique" {
										unStr += dataArray[key]
									}
								}
								if len(pkStr) > 0 {
									primaryKey = append(primaryKey, pkStr)
								}
								if len(unStr) > 0 {
									unKey = append(unKey, unStr)
								}
							}
						}
						inputByte.Reset()
					}
				}
				if readerError == io.EOF {
					break
				}
			}
			pkErr := BubbleSort(primaryKey, "pk")
			if len(unKey) > 0 {
				uniqueErr := BubbleSort(unKey, "unique")
				fmt.Println(uniqueErr)
			}
			if pkErr != "" {
				flag = false
				fmt.Println(pkErr)
			} else {
				fmt.Println("OK")
			}
		}
	}
	if flag == true {
		//创建4个sql文件
		sqlitefile, err := os.Create(skillfolder + "sqlite.sql")
		if err != nil {
			fmt.Println(err)
		}
		mysqlfile, err := os.Create(skillfolder + "mysql.sql")
		if err != nil {
			fmt.Println(err)
		}
		defer sqlitefile.Close()
		defer mysqlfile.Close()
		m := crateMap(skillfolder)
		fmt.Println("file create doing ...... ")
		for _, file := range files {
			var (
				sqlite      string
				mysql       string
				primarykey  string
				keymysql    string
				keysqlite   string
				unkeysqlite string
				unkeymysql  string
				headstr     string
				insert      []string
				insertsqlit string
				insertmysql string
				IDtotext    string
				head        []string
			)
			if file.IsDir() || !strings.Contains(file.Name(), ".csv") || file.Name() == "pre_static_text.csv" {
				continue
			} else {
				fmt.Println(file.Name() + " :")
				res_file, err := os.Open(skillfolder + file.Name())
				if err != nil {
					fmt.Println("ERROE: Failed to open the input file", err)
				}
				defer res_file.Close()
				mysql_table := strings.Replace(file.Name(), ".csv", "", -1)
				sqlite_table := strings.Replace(mysql_table, "pre_static_", "", -1)
				inputReader := bufio.NewReader(res_file)
				sqlite += "\nDROP TABLE IF EXISTS `" + sqlite_table + "`;\nCREATE TABLE " + sqlite_table + "(\n"
				mysql += "\nDROP TABLE IF EXISTS `" + mysql_table + "`;\nCREATE TABLE " + mysql_table + "(\n"
				primarykey += "PRIMARY KEY ("
				for n := 0; ; n++ {
					inputContent, isPrefix, readerError := inputReader.ReadLine()
					if inputContent != nil {
						inputByte.Write(inputContent)
						if !isPrefix {
							dataString := mahonia.NewDecoder("gbk").ConvertString(inputByte.String())
							dataArray := strings.Split(dataString, ",")
							if n < 5 {
								if n == 0 {
									for _, value := range dataArray {
										if len(value) > 0 {
											head = append(head, value+",")
											headstr += "`" + value + "`,"
										}
									}
								} else {
									for i := 0; i < len(head); i++ {
										head[i] += dataArray[i] + ","
									}
								}
							} else {
								insert = append(insert, dataString)
							}
							inputByte.Reset()
						}
					}
					if readerError == io.EOF {
						break
					}
				}
				for key, value := range head {
					headArray := strings.Split(value, ",")
					if len(headArray[4]) > 0 {
						IDtotext += fmt.Sprintf("%d,", key)
					}
					sqlite += fmt.Sprintf("`%s` ", headArray[0])
					mysql += fmt.Sprintf("`%s` %s", headArray[0], headArray[1])
					if strings.Index(headArray[1], "int") != -1 {
						sqlite += "INTEGER"
					} else if strings.Index(headArray[1], "float") != -1 {
						sqlite += "FLOAT"
					} else {
						sqlite += "TEXT"
					}
					if len(headArray[2]) != 0 {
						mysql += headArray[2]
					}
					sqlite += " NOTE NULL"
					mysql += " NOTE NULL"
					if len(headArray[3]) != 0 {
						switch headArray[3] {
						case "pk":
							primarykey += fmt.Sprintf("`%s`,", headArray[0])
						case "index":
							if keysqlite == "" {
								keysqlite = "CREATE INDEX `" + sqlite_table + "_index` ON " + sqlite_table + "("
								keymysql = "CREATE INDEX `" + sqlite_table + "_index` ON " + mysql_table + "("
							}
							keysqlite += fmt.Sprintf("`%s`,", headArray[0])
							keymysql += fmt.Sprintf("`%s`,", headArray[0])
						case "unique":
							if unkeysqlite == "" {
								unkeysqlite = "CREATE UNIQUE INDEX `" + sqlite_table + "_unique` ON " + sqlite_table + " ("
								unkeymysql = "CREATE UNIQUE INDEX `" + sqlite_table + "_unique` ON " + mysql_table + " ("
							}
							unkeysqlite += fmt.Sprintf("`%s`,", headArray[0])
							unkeymysql += fmt.Sprintf("`%s`,", headArray[0])
						default:
							fmt.Printf("ERROE: Index exceeds the specified type! Cell:[5, %d]", key)
						}

					}
					if headArray[3] == "pk" {
						sqlite += ",\n"
						mysql += ",\n"
					} else if strings.Index(headArray[1], "char") != -1 {
						sqlite += " DEFAULT '',\n"
						mysql += " DEFAULT '',\n"
					} else {
						sqlite += " DEFAULT '0',\n"
						sqlite += " DEFAULT '0',\n"
					}
				}
				if primarykey != "" {
					sqlite += strings.TrimRight(primarykey, ",") + ")"
					mysql += strings.TrimRight(primarykey, ",") + ")"
				}
				mysql += "\n)ENGINE=InnoDB;\n"
				sqlite += "\n);\n"
				if keysqlite != "" {
					sqlite += strings.TrimRight(keysqlite, ",") + ");\n"
					mysql += strings.TrimRight(keymysql, ",") + ");\nGOOS=windows GOARCH=amd64 go build"
				}
				if unkeysqlite != "" {
					sqlite += strings.TrimRight(unkeysqlite, ",") + ");\n"
					mysql += strings.TrimRight(unkeymysql, ",") + ");\n"
				}
				headstr = strings.TrimRight(headstr, ",")
				insertsqlit = "INSERT INTO " + sqlite_table + " (" + headstr + ") VALUES "
				insertmysql = "INSERT INTO " + mysql_table + " (" + headstr + ") VALUES "
				for row, value := range insert {
					item := strings.Split(value, ",")
					insertItem := "("
					for key, val := range item {
						val = strings.TrimRight(val, "#")
						if strings.Index(IDtotext, fmt.Sprintf("%d", key)) != -1 {
							Tag := ""
							if strings.Index(val, "#") != -1 {
								countTag := strings.Split(val, "#")
								countLen := len(countTag)
								for i := 0; i < countLen; i++ {
									if str, ok := m[countTag[i]]; ok {
										Tag += fmt.Sprintf("%s#", str)
									} else {
										Tag += ""
										fmt.Printf("ERROE: There is no dictionary line %d : %d\n", row+6, key+1)
									}
								}
							} else {
								if str, ok := m[val]; ok {
									Tag = fmt.Sprintf("%s#", str)
								} else {
									Tag = ""
									fmt.Printf("ERROE: There is no dictionary line %d : %d\n", row+6, key+1)
								}
							}
							insertItem += fmt.Sprintf("'%s',", strings.TrimRight(Tag, "#"))
						} else {
							insertItem += fmt.Sprintf("'%s',", val)
						}
						//fmt.Println(key)
						//fmt.Println(IDtotext)
					}
					insertsqlit += fmt.Sprintf("%s),", strings.TrimRight(insertItem, ","))
					insertmysql += fmt.Sprintf("%s),", strings.TrimRight(insertItem, ","))
				}
			}
			sqlitefile.WriteString(sqlite)
			sqlitefile.WriteString(strings.TrimRight(insertsqlit, ","))
			mysqlfile.WriteString(mysql)
			mysqlfile.WriteString(strings.TrimRight(insertmysql, ","))
		}
	}
	for {

	}
}

func BubbleSort(value []string, indexType string) (err string) {
	len := len(value)
	for i := 0; i < len; i++ {
		if strings.Contains(value[i], "-") {
			err += fmt.Sprintf("ERROE: %s not unsigned line: %d \n", indexType, i+6)
		}
		for j := i + 1; j < len; j++ {
			if value[i] == value[j] {
				err += fmt.Sprintf("ERROE: %s not only line: %d\n", indexType, i+6)
			}
		}
	}
	return
}

func crateMap(path string) (m map[string]string) {
	m = make(map[string]string)
	var inputByte bytes.Buffer
	reslut, _ := os.Open(path + "pre_static_text.csv")
	defer reslut.Close()
	inputReader := bufio.NewReader(reslut)
	n := 1
	for {
		n++
		if n <= 5 {
			inputReader.ReadLine()
			continue
		} else {
			inputContent, isPrefix, readerError := inputReader.ReadLine()
			if inputContent != nil {
				inputByte.Write(inputContent)
				if !isPrefix {
					dataString := inputByte.String()
					//dataString := mahonia.NewDecoder("gbk").ConvertString(inputByte.String())
					if strings.Index(dataString, ",") != 0 {
						dataArray := strings.Split(dataString, ",")
						m[dataArray[0]] = dataArray[1]
					}
					inputByte.Reset()
				}
			}
			if readerError == io.EOF {
				break
			}
		}
	}
	return
}
