package ipLoc

import "testing"

// test

func TestIpToUint64(t *testing.T) {
	ip := "1.1.1.1"
	ipInt, err := IpToUint64(ip)
	if err != nil {
		t.Fatalf("ipInt: %d, err: %v", ipInt, err)
		t.Fatal(err)
	}
}

func TestUnint64ToIp(t *testing.T) {
	ip, err := Uint64ToIp(16843009)
	if err != nil {
		t.Fatal(err)
	}
	if ip != "1.1.1.1" {
		t.Fatalf("ip: %s", ip)
	}
}
