
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}

}

func zero(ptr *[32]byte) {
*ptr = [32]byte{}
}