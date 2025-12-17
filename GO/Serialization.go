package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type Serializer struct{}

func NewSerializer() *Serializer {
	return &Serializer{}
}

// Бинарная сериализация Stack
func (s *Serializer) BinarySerializeStack(obj *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %v", err)
	}
	defer file.Close()

	data := obj.ToVector()
	size := uint64(len(data))

	// Записываем размер
	if err := binary.Write(file, binary.LittleEndian, size); err != nil {
		return err
	}

	// Записываем каждую строку
	for _, str := range data {
		strBytes := []byte(str)
		strSize := uint64(len(strBytes))

		if err := binary.Write(file, binary.LittleEndian, strSize); err != nil {
			return err
		}

		if _, err := file.Write(strBytes); err != nil {
			return err
		}
	}
	return nil
}

func (s *Serializer) BinaryDeserializeStack(obj *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %v", err)
	}
	defer file.Close()

	var size uint64
	if err := binary.Read(file, binary.LittleEndian, &size); err != nil {
		return err
	}

	data := []string{}
	for i := uint64(0); i < size; i++ {
		var strSize uint64
		if err := binary.Read(file, binary.LittleEndian, &strSize); err != nil {
			return err
		}

		strBytes := make([]byte, strSize)
		if _, err := file.Read(strBytes); err != nil {
			return err
		}

		data = append(data, string(strBytes))
	}

	obj.FromVector(data)
	return nil
}

// Текстовая сериализация Stack
func (s *Serializer) TextSerializeStack(obj *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %v", err)
	}
	defer file.Close()

	data := obj.ToVector()
	if _, err := fmt.Fprintf(file, "%d\n", len(data)); err != nil {
		return err
	}

	for _, str := range data {
		if _, err := fmt.Fprintln(file, str); err != nil {
			return err
		}
	}
	return nil
}

func (s *Serializer) TextDeserializeStack(obj *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %v", err)
	}
	defer file.Close()

	var size int
	if _, err := fmt.Fscanf(file, "%d\n", &size); err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	data := []string{}
	for i := 0; i < size; i++ {
		if scanner.Scan() {
			data = append(data, scanner.Text())
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	obj.FromVector(data)
	return nil
}
