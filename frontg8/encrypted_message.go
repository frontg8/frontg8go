// Copyright (c) 2016, Felix Morgner
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//  * Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//  * Neither the name of The frontg8 Project nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package frontg8

/*
#cgo LDFLAGS: -lfrontg8 -lprotobuf

#include <frontg8/protocol/message/encrypted.h>
*/
import "C"

import (
	"runtime"
)

type fg8EncryptedMessage struct {
	c_message C.fg8_protocol_message_encrypted_t
}

func NewEncryptedMessage(content string) (*fg8Error, *fg8EncryptedMessage) {
	encrypted := new(fg8EncryptedMessage)
	err := libNewError()

	data := C.CString(content)
	length := C.size_t(len(content))
	encrypted.c_message = C.fg8_protocol_message_encrypted_create(data, length, &err.c_error)

	runtime.SetFinalizer(encrypted, destroyEncryptedMessage)
	return err, encrypted
}

func DeserializeEncryptedMessage(serialized string) (*fg8Error, *fg8EncryptedMessage) {
	encrypted := new(fg8EncryptedMessage)
	err := libNewError()

	c_data := C.CString(serialized)
	length := C.size_t(len(serialized))
	encrypted.c_message = C.fg8_protocol_message_encrypted_deserialize(c_data, length, &err.c_error)

	return err, encrypted
}

func (obj *fg8EncryptedMessage) Copy() (*fg8Error, *fg8EncryptedMessage) {
	encrypted := new(fg8EncryptedMessage)
	err := libNewError()

	encrypted.c_message = C.fg8_protocol_message_encrypted_copy(obj.c_message, &err.c_error)

	runtime.SetFinalizer(encrypted, destroyEncryptedMessage)
	return err, encrypted
}

func (obj *fg8EncryptedMessage) Serialize() (*fg8Error, *string) {
	length := C.size_t(0)
	err := libNewError()

	c_data := C.fg8_protocol_message_encrypted_serialize(obj.c_message, &length, &err.c_error)
	data := C.GoStringN(c_data, C.int(length))

	return err, &data
}

func (obj *fg8EncryptedMessage) GetContent() (*fg8Error, *string) {
	length := C.size_t(0)
	err := libNewError()

	data := C.fg8_protocol_message_encrypted_get_content(obj.c_message, &length, &err.c_error)
	content := C.GoStringN(data, C.int(length))

	return err, &content
}

func (obj *fg8EncryptedMessage) SetContent(content string) *fg8Error {
	err := libNewError()

	data := C.CString(content)
	length := C.size_t(len(content))
	C.fg8_protocol_message_encrypted_set_content(obj.c_message, data, length, &err.c_error)

	return err
}

func (obj *fg8EncryptedMessage) Clear() *fg8Error {
	err := libNewError()

	C.fg8_protocol_message_encrypted_set_content(obj.c_message, nil, 0, &err.c_error)

	return err
}

func (obj *fg8EncryptedMessage) EqualTo(other *fg8EncryptedMessage) bool {
	return bool(C.fg8_protocol_message_encrypted_compare_equal(obj.c_message, other.c_message))
}

func (obj *fg8EncryptedMessage) HasContent() bool {
	return bool(C.fg8_protocol_message_encrypted_is_valid(obj.c_message))
}

func destroyEncryptedMessage(obj *fg8EncryptedMessage) {
	C.fg8_protocol_message_encrypted_destroy(obj.c_message)
	obj.c_message = nil
}
