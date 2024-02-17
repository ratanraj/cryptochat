package cryptochat

import (
	"strings"
)

import (
	"golang.org/x/crypto/openpgp"
)

func VerifyPublicKey(Key string) (bool, error) {
	keyFile := strings.NewReader(Key)
	entityList, err := openpgp.ReadArmoredKeyRing(keyFile)
	if err != nil {
		return false, err
	}

	//for _, entity := range entityList {
	//	fmt.Printf("Identity: %s", entity.Identities)
	//	fmt.Printf("PublicKey: %s", entity.PrimaryKey)
	//}

	if len(entityList) > 0 {
		return true, nil
	}

	return true, nil
}
