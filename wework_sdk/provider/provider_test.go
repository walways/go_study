package provider

import (
	"fmt"
	"testing"
)

func TestProvider_GetProviderToken(t *testing.T) {
	provider := NewProvider("wx2b03c6e1ce1d241c", "eYLUiM2V5bx4vzG9iaQzug8Q8", "", "", "")
	token, err := provider.GetToken()
	fmt.Printf("token:%+v, err:%+v", token, err)
}
