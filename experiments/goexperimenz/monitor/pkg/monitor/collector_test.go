package monitor

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockHttpCli struct {
}

func (m mockHttpCli) Do(req *http.Request) (*http.Response, error) {
	json := `[{"systemId":"02e4bb65-9d61-4988-b064-8ca4ed4e19fa","systemName":"DVGW-AZEDOCAPP02","customerId":29,"customerName":"DVGW Deutscher Verein des Gas- und Wasserfaches e.V. Technisch-Wissenschaftlicher Verein","internal":false,"diskUsagePercent":18,"ramUsagePercent":53,"lastContact":1594993442},{"systemId":"04388272-d60c-4514-a36d-329896a354c8","systemName":"Leipheim (Neu)","customerId":7,"customerName":"BRITAX R\u00d6MER Kindersicherheit GmbH","internal":false,"diskUsagePercent":28,"ramUsagePercent":48,"lastContact":1594993442},{"systemId":"06218070-aad3-4340-8cb9-5b66c70de4e2","systemName":"invoice app server","customerId":19,"customerName":"ystral gmbh maschinenbau + processtechnik","internal":false,"diskUsagePercent":27,"ramUsagePercent":63,"lastContact":1594993442},{"systemId":"0668ed5a-69b4-4f04-9799-b3c71fbf765c","systemName":"invoiceapp.haenchen.de","customerId":13,"customerName":"Herbert H\u00e4nchen GmbH & Co. KG","internal":false,"diskUsagePercent":38,"ramUsagePercent":52,"lastContact":1594993442},{"systemId":"0fc8445b-f468-41ff-9848-9c409a974dba","systemName":"edocapp3test5 (192.168.207.7)","customerId":6,"customerName":"edoc solutions ag - Team App Development Project","internal":true,"diskUsagePercent":40,"ramUsagePercent":39,"lastContact":1593446162}]`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestCollector_CollectSystemData(t *testing.T) {
	c := &Collector{httpClient: mockHttpCli{}}
	monitorData, err := c.CollectSystemData()
	if err != nil {
		t.Fatal(err)
	}
	if len(monitorData) != 5 {
		t.Fail()
	}
}
