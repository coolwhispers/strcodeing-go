package strcoding

import (
	"github.com/google/uuid"
)

//NewBase58 : Create a base58 string from uuid
func NewBase58() (str string, err error) {

	newID, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	strID := newID.String()

	str, err = ByteToBase58([]byte(strID))

	if err != nil {
		return "", err
	}

	return str, nil
}

//ByteToBase58 : Byte array to base58 string
func ByteToBase58(b []byte) (str string, err error) {
	//TODO
	return str, nil
}

//Base58ToByte : Base58 string to byte array
func Base58ToByte(str string) (b []byte, err error) {
	//TODO
	return b, nil
}
