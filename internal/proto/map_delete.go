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
	"github.com/hazelcast/hazelcast-go-client/internal/serialization"

	"github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
)

func mapDeleteCalculateSize(name string, key *serialization.Data, threadId int64) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += stringCalculateSize(name)
	dataSize += dataCalculateSize(key)
	dataSize += bufutil.Int64SizeInBytes
	return dataSize
}

// MapDeleteEncodeRequest creates and encodes a client message
// with the given parameters.
// It returns the encoded client message.
func MapDeleteEncodeRequest(name string, key *serialization.Data, threadId int64) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, mapDeleteCalculateSize(name, key, threadId))
	clientMessage.SetMessageType(mapDelete)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendData(key)
	clientMessage.AppendInt64(threadId)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

// MapDeleteDecodeResponse(clientMessage *ClientMessage), this message has no parameters to decode
