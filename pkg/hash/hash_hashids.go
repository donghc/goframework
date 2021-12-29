package hash

import (
	"github.com/speps/go-hashids"
)

func (h *hash) HashidsEncode(params []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	hashID, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	return hashID.Encode(params)
}

func (h *hash) HashidsDecode(hash string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	ids, err := hashids.NewWithData(hd)

	if err != nil {
		return nil, err
	}
	return ids.DecodeWithError(hash)
}
