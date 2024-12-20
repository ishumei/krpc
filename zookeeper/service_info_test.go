package zookeeper

import (
	"testing"
)

func TestRegistryInfo_Path(t *testing.T) {
	type fields struct {
		ServiceName string
		Port        string
		NodeInfo    *NodeInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				ServiceName: "test",
				Port:        "80",
				NodeInfo: &NodeInfo{
					Host:   "127.0.0.1",
					Port:   80,
					Weight: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &serviceInfo{
				ServiceName: tt.fields.ServiceName,
				Port:        tt.fields.Port,
				NodeInfo:    tt.fields.NodeInfo,
			}
			if got := n.Path(); got != tt.want {
				t.Logf("RegistryInfo.Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
