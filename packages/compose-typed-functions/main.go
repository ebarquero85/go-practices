package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransFormFunc func(string) string

type Server struct {
	fileNameTransformFunc TransFormFunc
}

func (s *Server) handleRequest(filename string) error {
	newFileName := s.fileNameTransformFunc(filename)
	fmt.Println("new filename: ", newFileName)
	return nil
}

// sha1
// prefix GG_
// hmac
func hashFileName(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func prefixFileName(prefix string) TransFormFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func main() {

	s := &Server{
		//fileNameTransformFunc: hashFileName,
		fileNameTransformFunc: prefixFileName("HMC_"),
	}

	s.handleRequest("cool_picture.jpg")

}
