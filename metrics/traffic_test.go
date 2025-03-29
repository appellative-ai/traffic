package metrics

import (
	"encoding/json"
	"fmt"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/iox"
	//"github.com/behavioral-ai/domain/testrsc"
)

func _ExampleMarshalTrafficProfile() {
	p := NewTrafficProfile()
	buf, err := json.Marshal(&p)

	fmt.Printf("test: json.Marshal() -> [err:%v] [buf:%v]\n", err, string(buf))

	//Output:
	//fail
}

func ExampleUnmarshalTrafficProfile() {
	p := NewTrafficProfile()
	buf, err := iox.ReadFile(testrsc.ResiliencyTrafficProfile1)
	fmt.Printf("test: iox.ReadFile() -> [err:%v] [buf:%v]\n", err, len(buf) > 0)

	err = json.Unmarshal(buf, &p)
	fmt.Printf("test: json.Unmarshal() -> [err:%v] [profile:%v]\n", err, p)

	traffic := p.Now()
	fmt.Printf("test: Profile.Now() -> [traffic:%v]\n", traffic)

	//Output:
	//fail
}

/*
func ExampleResolver() {
	buf, err := iox.ReadFile(testrsc.ResiliencyTrafficProfile1)
	fmt.Printf("test: iox.ReadFile() -> [err:%v] [buf:%v]\n", err, len(buf) > 0)

	r := collective.NewEphemeralResolver()
	status := r.PutContent(ProfileName, "author", buf, 1)
	fmt.Printf("test: PutContent() -> [status:%v]\n", status)

	p, status1 := collective.Resolve[TrafficProfile](ProfileName, 1, r)
	fmt.Printf("test: Resolve[TrafficProfile]() -> [status:%v]\n", status1)

	traffic := p.Now()
	fmt.Printf("test: Profile.Now() -> [traffic:%v]\n", traffic)

	//Output:
	//test: iox.ReadFile() -> [err:<nil>] [buf:true]
	//test: PutContent() -> [status:OK]
	//test: Resolve[TrafficProfile]() -> [status:OK]
	//test: Profile.Now() -> [traffic:low]

}


*/
