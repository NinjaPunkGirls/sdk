package docs

import (
	"encoding/hex"
	"path"
	"time"
	"strconv"
	"os"
	"fmt"
)


type Document struct {
	client *Client
	Time  string
	Place Place
	Parent string
	Class string
	Data  interface{}
}

func (self *Document) Serialise() (string, error) {
	s, err := self.client.CompactSerial(self)
	if err != nil {
		return "", err
	}
	return s, nil
}

func (self *Document) Save() error {

	s, err := self.Serialise()
	if err != nil {
		return err
	}

	for _, parentNodeID := range self.Place.ParentHashes() {

		folders := path.Join(parentNodeID, self.TimePrefix())
		if err := os.MkdirAll(folders, 0777); err != nil {
			return err
		}

		filename := path.Join(folders, self.Class)
		println("SAVING", filename)

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}

		defer f.Close()

		_, err = f.WriteString(self.Time + " " + hex.EncodeToString(Hash([]byte(s))) + "\n")
		if err != nil {
			panic(err)
		}
	
	}
	return nil
}


func (self *Document) TimePrefix() string {

	i, err := strconv.ParseInt(self.Time, 10, 64)
    if err != nil {
        panic(err)
    }
    t := time.Unix(i, 0)

	return fmt.Sprintf("%d/%02d/%02d/%02d/%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute())
}