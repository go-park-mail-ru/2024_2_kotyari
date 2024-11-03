package utils

import (
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/errs"
	"strconv"
)

const (
	convertBase = 10
	bitSize     = 32
)

func StrToUint32(arg string) (uint32, error) {
	argUint64, err := strconv.ParseUint(arg, convertBase, bitSize)
	if err != nil {
		return 0, errs.ParsingURLArg
	}

	return uint32(argUint64), nil
}
