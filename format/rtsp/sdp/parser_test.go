package sdp

import (
	"testing"
)

func TestParse(t *testing.T) {
	s, infos := Parse(`
v=0
o=- 1459325504777324 1 IN IP4 192.168.0.123
s=RTSP/RTP stream from Network Video Server
i=mpeg4cif
t=0 0
a=tool:LIVE555 Streaming Media v2009.09.28
a=type:broadcast
a=control:*
a=range:npt=0-
a=x-qt-text-nam:RTSP/RTP stream from Network Video Server
a=x-qt-text-inf:mpeg4cif
m=video 0 RTP/AVP 96
c=IN IP4 0.0.0.0
b=AS:300
a=rtpmap:96 H264/90000
a=fmtp:96 profile-level-id=420029; packetization-mode=1; sprop-parameter-sets=Z00AHpWoKA9k,aO48gA==
a=x-dimensions: 720, 480
a=x-framerate: 15
a=control:track1
m=audio 0 RTP/AVP 96
c=IN IP4 0.0.0.0
b=AS:256
a=rtpmap:96 MPEG4-GENERIC/16000/2
a=fmtp:96 streamtype=5;profile-level-id=1;mode=AAC-hbr;sizelength=13;indexlength=3;indexdeltalength=3;config=1408
a=control:track2
m=audio 0 RTP/AVP 0
c=IN IP4 0.0.0.0
b=AS:50
a=recvonly
a=control:rtsp://109.195.127.207:554/mpeg4cif/trackID=2
a=rtpmap:0 PCMU/8000
a=Media_header:MEDIAINFO=494D4B48010100000400010010710110401F000000FA000000000000000000000000000000000000;
a=appversion:1.0
`)
	t.Logf("%s,%+v\n", s, infos)

	s, infos = Parse(`
v=0
o=- 5783354 0 IN IP4 192.168.100.3
s=000-001_594008_61010102031320000028_ffcs.root.77d97fa4fa3bb25e64eb62e995e8eb75
u=www.ffcs.cn
e=ffcs@ffcs.cn
c=IN IP4 192.168.100.3
b=AS:5100
t=0 0
a=control:*
m=video 0 RTP/AVP 96
c=IN IP4 0.0.0.0
b=AS:5000
a=recvonly
a=control:trackID=1
a=rtpmap:96 H265/90000
a=fmtp:98 profile-id=1;sprop-sps=QgEBAUAAAAMAAAMAAAMAAAMAmaACgIAuHxOWu5Gwa5VB;sprop-pps=RAHAc8BMkA==;sprop-vps=QAEMAf//AUAAAAMAAAMAAAMAAAMAmawJ
m=audio 0 RTP/AVP 8
a=recvonly
b=AS:50
a=rtpmap:8 PCMA/8000
a=control:trackID=2
a=Media_header:MEDIAINFO=494D4B48010100000400050000000000000000000000000000000000000000000000000000000000;
a=appversion:1.0
`)
	t.Logf("%s,%+v\n", s, infos)

	s, infos = Parse(`
v=0
o=- 2229913047 2229913047 IN IP4 0.0.0.0
s=Media Server
c=IN IP4 0.0.0.0
t=0 0
a=control:*
a=packetization-supported:DH
a=rtppayload-supported:DH
a=range:npt=now-
m=video 0 RTP/AVP 96
a=control:trackID=0
a=framerate:25.000000
a=rtpmap:96 H264/90000
a=fmtp:96 packetization-mode=1;profile-level-id=64002A;sprop-parameter-sets=Z2QAKqwsaoHgCJ+WbgoCCoAAAAMAgAAAGUIA,aO4xshsA
a=recvonly
`)
	t.Logf("%s,%+v\n", s, infos)
}
