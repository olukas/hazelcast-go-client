// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proto

import (
	"github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
	"github.com/hazelcast/hazelcast-go-client/internal/serialization"
)

func listRemoveWithIndexCalculateSize(name string, index int32) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	dataSize += bufutil.Int32SizeInBytes
	return dataSize
}

// ListRemoveWithIndexEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func ListRemoveWithIndexEncodeRequest(name string, index int32) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, listRemoveWithIndexCalculateSize(name, index))
	clientMessage.SetMessageType(listRemoveWithIndex)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendInt32(index)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// ListRemoveWithIndexDecodeResponse decodes the given client message.
// It returns a function which returns the response parameters.
func ListRemoveWithIndexDecodeResponse(clientMessage *ClientMessage) func() (response *serialization.Data) {
	// Decode response from client message
	return func() (response *serialization.Data) {

		if !clientMessage.ReadBool() {
			response = clientMessage.ReadData()
		}
		return
	}
}
