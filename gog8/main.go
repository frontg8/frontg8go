package main

import (
	"fmt"
	"github.com/frontg8/frontg8go/frontg8"
)

func main() {
	err, message := frontg8.NewEncryptedMessage("Hello")
	err, content := message.GetContent()

	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	} else {
		fmt.Printf("Content: '%s'\nLength: %d\nBytes: % #x\n", *content, len(*content), *content)
	}

	println(" -- Changing content -- ")

	text := []byte{0x66, 0x67, 0x00, 0x38, 0x21}
	err = message.SetContent(string(text))

	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	}

	err, content = message.GetContent()
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	} else {
		fmt.Printf("Content: '%s'\nLength: %d\nBytes: % #x\n", *content, len(*content), *content)
	}

	println(" -- Serializing message -- ")

	err, serialized := message.Serialize()
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	} else {
		fmt.Printf("Serialized: '%s'\nLength: %d\nBytes: % #x\n", *serialized, len(*serialized), *serialized)
	}

	println(" -- Deserializing message -- ")

	err, deserialized := frontg8.DeserializeEncryptedMessage(*serialized)
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	}

	err, content = deserialized.GetContent()
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	} else {
		fmt.Printf("Content: '%s'\nLength: %d\nBytes: % #x\n", *content, len(*content), *content)
	}

	println(" -- Clearing message -- ")

	err = deserialized.Clear()
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	}

	err, content = deserialized.GetContent()
	if err.HasError() {
		fmt.Printf("Error: %s\n", err.Message())
	} else {
		fmt.Printf("Content: '%s'\nLength: %d\nBytes: % #x\n", *content, len(*content), *content)
	}

	println(" -- Comparing to original -- ")

	if deserialized.EqualTo(message) {
		println("Both are equal")
	} else {
		println("The messages differ")
	}

	println(" -- Comparing to itself -- ")

	if deserialized.EqualTo(deserialized) {
		println("Both are equal")
	} else {
		println("The messages differ")
	}

	println(" -- Checking content -- ")

	if deserialized.HasContent() {
		println("Cleared message has content")
	} else {
		println("Cleared message has no content")
	}

}
