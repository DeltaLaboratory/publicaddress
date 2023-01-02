package publicaddress

import "testing"

func TestGetInformation(t *testing.T) {
	res, err := GetInformation()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("IP : %s\nCountry: %s", res.IP, res.Country)
	return
}
