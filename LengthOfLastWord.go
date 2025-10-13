package main

func LengthOfLastWord(s string) int {

	var char_flag bool // digunakan untuk trigger jika sudah menemukan karakter bukan spasi
	var length int = 0 // panjang karakter

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' { // validasi per karakter tidak boleh spasi
			char_flag = true //  deklarasi sudah menemukan karakter
			length++         //menambhkan value dari variable length satu per satu
		} else {
			if char_flag { // ketika triggernya benar maka langsung return length
				return length
			}
		}
	}
	return length
}
