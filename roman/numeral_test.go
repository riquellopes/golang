package roman

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestArabic_1_should_roman_I(t *testing.T) {

    arabicToRoman := ArabicToRoman(1)
    assert.Equal(t, arabicToRoman, "I")
}

func TestArabic_2_should_roman_II(t *testing.T) {

    arabicToRoman := ArabicToRoman(2)
    assert.Equal(t, arabicToRoman, "II")
}
