package decode

import (
	"encoding/pem"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	privateKey := `-----BEGIN PRIVATE KEY-----MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC8n269SHL52Skw +1TMiuAlRGGzvccqFcSqP/tPefQPLVLBxloR6UMS5FNzS1JoD8o+KZxHQNQbsMOr 4KnVpzf9JQo75P0c7LlXImINX2ytZfOPdoKyWQDN8xqIMdhHKeGJxXUWkFo+R9sK TURsj+CBqBpxnFafNuC58mONE1bKtbD640UIIj00dmnB3F+hAj3p45EkoojvKCQg w4TJBdqzeVVHrbLDD0Nz8+YUgV+U7pfhX5nDUJrDlwU8QX5bRkOmG6uZp8giqzRe CI42sEq2/NGzHeUyS/hjzx9vfURFZD1htm8sbi15o9TygmhlBRyg77jelEAqk5oV hx1FSicdAgMBAAECggEAZLa4WWcCxTdX0CiSqbiJYDy38lIlsun2nXUWSeRzFVt8 8axoUWH8h2zUxh30vU2ZDkIHiOAroTFr+S03YbYspgUOBtcI81XOzKC2PC+0ho2G VBbSSEpPrKfehdQfiDfrjjWZhFfFZ7XUVhbVryXPmUtZ+0mf0tlper/auewWjps0 TUJQCGn5kGBjdSVmQ+SrQgn8N5GTfQvO+KkE7nrSOcFa3Wze2gcuj2tgEbWrjgOp ZX6ijsuLC1Lvy0teKFOfemkgTrYuw8RLM2y6UII2pRsjxH6iotFRfduwwod2x/mx NsfzqR7AH8Nm4vpsjqrUfA/iV2bUttHgCZTzj9/nQQKBgQDhLoM+DDtJzl5TQgHr hDdQNy5dLv593DxpOZT6yV620Re405O24PGfKAhQcuWJkGHbJdD2wjNHnhRFn+9z InbqdGwatY5LAKnOIpGIGTeY6gYazW/tKzTj9uIGwpgIBwW1LLUBWpuiMrkCALEK r3ut3kOCHGfyI/Q/bh4Q5QS1cQKBgQDWcAneGShj63XmIY/LGDaovBH4295xyFAJ RBPGEmmLMW5X5FWXMoXfTQ3+hDznqIm6OsGAtHlZzCnj+bvtvd7SmUPbt7nYtjUd VxjXxKPIfeOuBp6TWFT2ITMTBqh2DKmgeE1MIUw3wbVRtYimocCSAi2TSTaVIzNn aqVUPWxGbQKBgQDgIjvf9qzMueyJUpo2olOsNECh37TP8fbvhbEULntn3JeQEIhm 71+Q23frG3H9iwjb6Lgl5QbFskLbSmY8Y4GZvGgd7bjETtS1qRLSNMdyHP4Wj4Vo vCjY/1PbBrXsx5CW1PkVb+qZd4NOywFtcV3ZofHv//Y2vjLPL/wdNNeLAQKBgCC3 gVnofrgYAan17ulurVA7OYa9jQfRJxzyzQ2IW6AID6zLQp/J7pDuQbdWV7CEj1bh hgpr/qPLuYx12s1gDbIV/jt0oZ9F1X6/fiPVdaOuSWGQ0vrqHDRdRlydxHAWaH76 O9jkr6tGG7L2vn2l6iqcuHVjUzeX4kziAY2DvZRVAoGBAJatYWPhE+2n1zZYRsFx mBV1TFjEs8nfN/ePyQ8/bCRTTVdKutB5eiuU5SE5q/g9JGYQ0XWaunmg+WXppI+6 LC39ctyRgppc1WSBzlKyn7yiJWWoWy/Ik0jd7rR1sCneAbFw7i+X/WIxVjTihsiy fYfGl5QvZp4HoI5qOkoYl8Tj 
-----END PRIVATE KEY-----`
	block, err := pem.Decode([]byte(privateKey))
	fmt.Println(block, string(err))
}

func TestStr(t *testing.T) {
	str := "Go爱好者"
	//for i, c := range str {
	//	fmt.Printf("%d: %q [% x] % x\n", i, c, []byte(string(c)), []rune(string(c)))
	//}

	fmt.Println(str[0], str[0:1])
}
