package compatibility

import (
	"fmt"
	"github.com/hazelcast/go-client/config"
	. "github.com/hazelcast/go-client/internal/common"
	"github.com/hazelcast/go-client/internal/serialization"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

type BinaryCompatibilityTest struct {
	version   byte
	object    interface{}
	byteOrder bool
}

type i interface {
	Available() int32
	ReadInt32() (int32, error)
	ReadUTF() (string, error)
	GetPosition() int32
	SetPosition(pos int32)
}

func TestBinaryCompatibility(t *testing.T) {
	var supporteds []string = []string{
		"1-NULL-BIG_ENDIAN",
		"1-NULL-LITTLE_ENDIAN",
		"1-Boolean-BIG_ENDIAN",
		"1-Boolean-LITTLE_ENDIAN",
		"1-Byte-BIG_ENDIAN",
		"1-Byte-LITTLE_ENDIAN",
		"1-Character-BIG_ENDIAN",
		"1-Character-LITTLE_ENDIAN",
		"1-Double-BIG_ENDIAN",
		"1-Double-LITTLE_ENDIAN",
		"1-Short-BIG_ENDIAN",
		"1-Short-LITTLE_ENDIAN",
		"1-Float-BIG_ENDIAN",
		"1-Float-LITTLE_ENDIAN",
		"1-Integer-BIG_ENDIAN",
		"1-Integer-LITTLE_ENDIAN",
		"1-Long-BIG_ENDIAN",
		"1-Long-LITTLE_ENDIAN",
		"1-String-BIG_ENDIAN",
		"1-String-LITTLE_ENDIAN",
		"1-AnInnerPortable-BIG_ENDIAN",
		"1-AnInnerPortable-LITTLE_ENDIAN",
		"1-boolean[]-BIG_ENDIAN",
		"1-boolean[]-LITTLE_ENDIAN",
		"1-byte[]-BIG_ENDIAN",
		"1-byte[]-LITTLE_ENDIAN",
		"1-char[]-BIG_ENDIAN",
		"1-char[]-LITTLE_ENDIAN",
		"1-double[]-BIG_ENDIAN",
		"1-double[]-LITTLE_ENDIAN",
		"1-short[]-BIG_ENDIAN",
		"1-short[]-LITTLE_ENDIAN",
		"1-float[]-BIG_ENDIAN",
		"1-float[]-LITTLE_ENDIAN",
		"1-int[]-BIG_ENDIAN",
		"1-int[]-LITTLE_ENDIAN",
		"1-long[]-BIG_ENDIAN",
		"1-long[]-LITTLE_ENDIAN",
		"1-String[]-BIG_ENDIAN",
		"1-String[]-LITTLE_ENDIAN",
		"1-AnIdentifiedDataSerializable-BIG_ENDIAN",
		"1-AnIdentifiedDataSerializable-LITTLE_ENDIAN",
		"1-APortable-BIG_ENDIAN",
		"1-APortable-LITTLE_ENDIAN",
	}

	var dataMap map[string]*serialization.Data = make(map[string]*serialization.Data)

	dat, _ := ioutil.ReadFile("1.serialization.compatibility.binary")

	i := serialization.NewObjectDataInput(dat, 0, nil, true)
	var index int
	for i.Available() != 0 {
		objectKey, _ := i.ReadUTF()
		length, _ := i.ReadInt32()
		if length != NULL_ARRAY_LENGTH {
			payload := dat[i.GetPosition() : i.GetPosition()+length]
			i.SetPosition(i.GetPosition() + length)
			if supporteds[index] == *objectKey {
				dataMap[*objectKey] = &serialization.Data{payload}
			}
		}
		index++
		if index == len(supporteds) {
			break
		}
	}

	serviceLE := createSerializationService(false)
	serviceBE := createSerializationService(true)

	objects := allTestObjects{}.getAllTestObjects()

	var retObjects []interface{} = make([]interface{}, len(supporteds)/2)

	var temp interface{}
	var temp2 interface{}
	for i := 0; i < len(supporteds); i++ {

		if strings.HasSuffix(supporteds[i], "BIG_ENDIAN") {
			temp, _ = serviceBE.ToObject(dataMap[supporteds[i]])
		} else {
			temp2, _ = serviceLE.ToObject(dataMap[supporteds[i]])
			if !reflect.DeepEqual(temp, temp2) {
				t.Errorf("compatibility test is incorrectly coded!")
			}
		}
		retObjects[i/2] = temp
	}

	if !reflect.DeepEqual(objects, retObjects) {
		fmt.Println(objects)
		fmt.Println(retObjects)
		t.Errorf("Go Serialization is not compatible with Java!")
	}

}

func createSerializationService(byteOrder bool) *serialization.SerializationService {
	serConfing := config.NewSerializationConfig()
	pf := &aPortableFactory{}
	idf := &aDataSerializableFactory{}
	serConfing.AddPortableFactory(PORTABLE_FACTORY_ID, pf)
	serConfing.AddDataSerializableFactory(IDENTIFIED_DATA_SERIALIZABLE_FACTORY_ID, idf)

	if byteOrder {
		return serialization.NewSerializationService(serConfing)
	}
	serConfing.SetByteOrder(false)
	return serialization.NewSerializationService(serConfing)
}
