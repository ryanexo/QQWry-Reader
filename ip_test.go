package ip

import (
    `fmt`
    `testing`
)

func TestData_Query(t *testing.T) {
    q, err := New(`./qqwry.dat`)
    if err != nil {
        panic(err)
    }
    r, err := q.Query("1.0.0.1")
    if err != nil {
        panic(err)
    }
    if string(r.Country) != "美国" {
        t.Error("unexpected country")
    }
    if string(r.Region) != "APNIC&CloudFlare公共DNS服务器" {
        t.Error("unexpected region")
    }

    fmt.Println(string(r.Location))
}
