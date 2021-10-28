package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const pwd_db = "../../.password.db"

func StorePassword(plataform, username, password string) (bool, error) {
	entry := fmt.Sprintf("%s,%s,%s\n", plataform, username, password)
	f, err := os.OpenFile(pwd_db, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("The file doesnt exist: %v", err)
		return false, err
	}
	defer f.Close()
	content, err := f.WriteString(entry)
	if err != nil {
		log.Fatalf("it hadn't writed in the file: %v", err)
		return false, err
	}
	log.Println(content, "bytes written")
	return true, nil
}

func RetrivePassword(plataform string) (string, error) {
	f, err := os.Open(pwd_db)
	if err != nil {
		log.Fatalf("The file doesnt exist: %v", err)
		return "", err
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == plataform {
			response := strings.Join(entry[1:], "|")
			return response, nil
		}
	}
	log.Fatalf("Platform %s not know\n", plataform)
	return "", fmt.Errorf("plataform %s not know", plataform)
}
