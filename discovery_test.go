package onvif

import (
	"fmt"
	"testing"
)

func TestDiscoveryUnionCamQ5CamKeeper(t *testing.T) {
	const resp = `
	<Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope" xmlns:soapenc="http://www.w3.org/2003/05/soap-encoding" xmlns:enc="http://www.w3.org/2003/05/soap-encoding" xmlns:hcsx1="http://www.hcsx.com/ver10/private/wsdl" xmlns:rpc="http://www.w3.org/2003/05/soap-rpc" xmlns:xs="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:tt="http://www.onvif.org/ver10/schema" xmlns:tds="http://www.onvif.org/ver10/device/wsdl" xmlns:trt="http://www.onvif.org/ver10/media/wsdl" xmlns:timg="http://www.onvif.org/ver20/imaging/wsdl" xmlns:tev="http://www.onvif.org/ver10/events/wsdl" xmlns:tptz="http://www.onvif.org/ver20/ptz/wsdl" xmlns:tan="http://www.onvif.org/ver20/analytics/wsdl" xmlns:tst="http://www.onvif.org/ver10/storage/wsdl" xmlns:ter="http://www.onvif.org/ver10/error" xmlns:dn="http://www.onvif.org/ver10/network/wsdl" xmlns:tns1="http://www.onvif.org/ver10/topics" xmlns:ns1="http://www.placeholder.org/ver10/tmp/schema" xmlns:trc="http://www.onvif.org/ver10/recording/wsdl" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl" xmlns:wsoap12="http://schemas.xmlsoap.org/wsdl/soap12" xmlns:http="http://schemas.xmlsoap.org/wsdl/http" xmlns:d="http://schemas.xmlsoap.org/ws/2005/04/discovery" xmlns:wsadis="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:xop="http://www.w3.org/2004/08/xop/include" xmlns:wsnt="http://docs.oasis-open.org/wsn/b-2" xmlns:wsnt5="http://www.w3.org/2005/08/addressing" xmlns:wsa="http://www.w3.org/2005/08/addressing" xmlns:wsa5="http://www.w3.org/2005/08/addressing" xmlns:wstop="http://docs.oasis-open.org/wsn/t-1" xmlns:wsrf-bf="http://docs.oasis-open.org/wsrf/bf-2" xmlns:wsntw="http://docs.oasis-open.org/wsn/bw-2" xmlns:wsrf-rw="http://docs.oasis-open.org/wsrf/rw-2" xmlns:wsaw="http://www.w3.org/2006/05/addressing/wsdl" xmlns:wsrf-r="http://docs.oasis-open.org/wsrf/r-2" xmlns:tnsn="http://www.eventextension.com/2011/event/topics">
	<Header>
		<MessageID>
			urn:uuid:A000734
		</MessageID>
		<RelatesTo>
			uuid:39df5eb2-2376-42f1-bb8f-6edb3baa502c
		</RelatesTo>
		<To>
			http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous
		</To>
		<Action>
			http://schemas.xmlsoap.org/ws/2005/04/discovery/ProbeMatches
		</Action>
		<AppSequence InstanceId="1544111766" MessageNumber="119">
		</AppSequence>
	</Header>
	<Body>
		<ProbeMatches>
			<ProbeMatch>
				<EndpointReference>
					<Address>
						urn:uuid:A000734
					</Address>
				</EndpointReference>
				<Types>
					dn:NetworkVideoTransmitter
				</Types>
				<Scopes>
					onvif://www.onvif.org/type/video_encoder onvif://www.onvif.org/type/ptz onvif://www.onvif.org/type/audio_encoder onvif://www.onvif.org/type/network_video_transmitter onvif://www.onvif.org/hardware/ipc onvif://www.onvif.org/location/Germany/Germany onvif://www.onvif.org/name/CamKeeper onvif://www.onvif.org/Profile/Streaming
				</Scopes>
				<XAddrs>
					http://10.0.101.213:8090/onvif/device_service
				</XAddrs>
				<MetadataVersion>
					10
				</MetadataVersion>
			</ProbeMatch>
		</ProbeMatches>
	</Body>
</Envelope>`

	dev, err := readDiscoveryResponse("uuid:39df5eb2-2376-42f1-bb8f-6edb3baa502c", []byte(resp))

	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	if dev.XAddr != "http://10.0.101.213:8090/onvif/device_service" {
		fmt.Println("XAddr does not match")
		t.FailNow()
	}
}