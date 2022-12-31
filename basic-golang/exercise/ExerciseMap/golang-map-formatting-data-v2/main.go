package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	header := make(map[string]int)
	formatData := make(map[string][]string)
	sort.Strings(data)
	for i := 0; i < len(data); i++ {
		for j := i+1; j < len(data); j++ {
			tmp1 := strings.Split(data[i], "-")
			tmp2 := strings.Split(data[j], "-")
			if tmp1[0] == tmp2[0] {
				ind1, _ := strconv.Atoi(tmp1[1])
				ind2, _ := strconv.Atoi(tmp2[1])
				if ind1 > ind2 {
					data[i], data[j] = data[j], data[i]
				}
			}
		}
	}
	for _, val := range data {
		dataDetail := strings.Split(val, "-")
		if dataDetail[2] == "first" {
			formatData[dataDetail[0]] = append(formatData[dataDetail[0]], dataDetail[3])
			header[dataDetail[0]]++
		} else {
			formatData[dataDetail[0]][header[dataDetail[0]]-1] = fmt.Sprintf("%s %s", formatData[dataDetail[0]][header[dataDetail[0]]-1], dataDetail[3])
		}
	}
	return formatData // TODO: replace this
}

func main() {
	data := []string{"aaaa-0-first-QiFTiTbFXr", "aaaa-0-last-mfZepTbbEO", "bbbb-0-first-SVdJeXaPKs", "bbbb-0-last-VxZcudCSWm", "cccc-0-first-HLaDBGhIGf", "cccc-0-last-KurIFSiTFE", "aaaa-1-first-ZuvHUUAMKv", "aaaa-1-last-pGplWTCATk", "bbbb-1-first-VeaAqklJzV", "bbbb-1-last-MDlzOFhTzp", "cccc-1-first-ylkSyPpRpi", "cccc-1-last-YzYUnkGshv", "aaaa-2-first-EPrbsZSCZo", "aaaa-2-last-oNtRuYQZdC", "bbbb-2-first-DcrwxlNPDv", "bbbb-2-last-FCBzlKpJYl", "cccc-2-first-ASndktHELU", "cccc-2-last-LsiihAuVSK", "aaaa-3-first-mUOBYmsLTg", "aaaa-3-last-nmNvGdyiwL", "bbbb-3-first-ykoHGWdlJq", "bbbb-3-last-DkmlsDXDIT", "cccc-3-first-GXHAtQOBJs", "cccc-3-last-nPZaLPRVFJ", "aaaa-4-first-yvsNJQQnTG", "aaaa-4-last-epYPcDmiUx", "bbbb-4-first-dFoSFpVUIX", "bbbb-4-last-tAfpnyfOes", "cccc-4-first-VQmIGPhEpq", "cccc-4-last-hqdWCkYrxW", "aaaa-5-first-FkKFJvOCab", "aaaa-5-last-VickAoCQin", "bbbb-5-first-OsuSbnfmVY", "bbbb-5-last-sPfNccFksy", "cccc-5-first-hjwjMDntdL", "cccc-5-last-sTJTQpPGPm", "aaaa-6-first-dUNzitkIuN", "aaaa-6-last-YlbwpORVIO", "bbbb-6-first-jnJEOqHypu", "bbbb-6-last-NHvLdHETvz", "cccc-6-first-QnywNlvsVV", "cccc-6-last-fTCwVWIlEm", "aaaa-7-first-HnpoBXeBXM", "aaaa-7-last-UfYoBXqXHc", "bbbb-7-first-nHVMZzhHim", "bbbb-7-last-vDHgFQWFEn", "cccc-7-first-pWqcpUAWCV", "cccc-7-last-oaTlqWNWbu", "aaaa-8-first-GilIfjifph", "aaaa-8-last-LMSEEFqKEo", "bbbb-8-first-aEPtgRpuun", "bbbb-8-last-KtxJIteXIS", "cccc-8-first-NhIkRqeUhE", "cccc-8-last-uJFJIbRVLA", "aaaa-9-first-OcjYAcJuOP", "aaaa-9-last-eXuRcGUzwo", "bbbb-9-first-vDWJwKAcrX", "bbbb-9-last-WNRdYLHZnB", "cccc-9-first-pxOVIPunFw", "cccc-9-last-OpDZYZIAQT", "aaaa-10-first-RhGtkhdkuC", "aaaa-10-last-odQZAqziIW", "bbbb-10-first-WLVWcMsPVO", "bbbb-10-last-djPNgFjIQO", "cccc-10-first-xPmiwVzMCk", "cccc-10-last-ZfiySJTEYY", "aaaa-11-first-YYPVCWgZod", "aaaa-11-last-fMRWBXdCLn", "bbbb-11-first-RgwuNIvDZc", "bbbb-11-last-JgaqNQwsHJ", "cccc-11-first-rcbCOIbVCy", "cccc-11-last-sKIjfmKTNr", "aaaa-12-first-QGdDLkgwtH", "aaaa-12-last-nEffnEwtvy", "bbbb-12-first-oBueXNbyqT", "bbbb-12-last-UUDzTFQPpt", "cccc-12-first-LEdzjWusAi", "cccc-12-last-dRQTTOzofT", "aaaa-13-first-DBinnzZpNL", "aaaa-13-last-PtzDDlWGIW", "bbbb-13-first-GhaPvaVTWX", "bbbb-13-last-TTFvXGPqtY", "cccc-13-first-eePfOHqWlN", "cccc-13-last-YKWLCIjMUD", "aaaa-14-first-etLIWpFlmY", "aaaa-14-last-HWFzieKBox", "bbbb-14-first-uqIHmTUfhN", "bbbb-14-last-fGWsuEtgMo", "cccc-14-first-SeUgqBrAtd", "cccc-14-last-eqWUCaMhwD", "aaaa-15-first-EeZQuZBGTY", "aaaa-15-last-SMGkCGknTB", "bbbb-15-first-djCAytQFuV", "bbbb-15-last-ievUiyVyQL", "cccc-15-first-hKVUKrdDnP", "cccc-15-last-RqEMWHIgdU", "aaaa-16-first-wbvgaidlta", "aaaa-16-last-tXHcFDvCHH", "bbbb-16-first-wySYxRkfWk", "bbbb-16-last-QmIpqlSIdw", "cccc-16-first-CRqdhSXljb", "cccc-16-last-tRwZdPKivP", "aaaa-17-first-EhuNlRmfSn", "aaaa-17-last-flAJlaNMYD", "bbbb-17-first-ELRQpanNOJ", "bbbb-17-last-cUrVRieXlB", "cccc-17-first-NTkwVgnHRi", "cccc-17-last-PBkNpcnbAI", "aaaa-18-first-MJmmdxzwWn", "aaaa-18-last-HggllNKhUT", "bbbb-18-first-LwBthWHOpW", "bbbb-18-last-iAIGFXYbQI", "cccc-18-first-PsGmZwOcdr", "cccc-18-last-fNYWNVTdov", "aaaa-19-first-XZZCtChuhv", "aaaa-19-last-jrjCxGIEPy", "bbbb-19-first-GDAuUAXhmw", "bbbb-19-last-CPSmElnyQx", "cccc-19-first-DZpmoGGizS", "cccc-19-last-etqqErSfNf", "aaaa-20-first-xtsqwGpdpd", "aaaa-20-last-nxDSEDlpba", "bbbb-20-first-SLzKIQYyWS", "bbbb-20-last-ebKrQxYMtw", "cccc-20-first-hrcKASDkIX", "cccc-20-last-OPfPgrotfe", "aaaa-21-first-WaoLMYSaya", "aaaa-21-last-dBCyQNLDEh", "bbbb-21-first-YZZnIqVcyK", "bbbb-21-last-uFKrMvaprF", "cccc-21-first-lqwTEqzKBR", "cccc-21-last-sGfrYzieUG", "aaaa-22-first-HyEQHkSQiO", "aaaa-22-last-tOOTTZQgdJ", "bbbb-22-first-IKVsLhFNNy", "bbbb-22-last-sIVMEfrmcB", "cccc-22-first-USSITEjpKM", "cccc-22-last-QXRQfwlewr", "aaaa-23-first-naicPBiqXa", "aaaa-23-last-kHoEjyQTnI", "bbbb-23-first-xdQMpiGUpy", "bbbb-23-last-IRjJNfNQrd", "cccc-23-first-DZYEJUGhWV", "cccc-23-last-THBQvfbMjD", "aaaa-24-first-QxtlrPbMRm", "aaaa-24-last-DekcrKFNwk", "bbbb-24-first-suLgnMuyCE", "bbbb-24-last-LVNtDBxKJl", "cccc-24-first-PmmtmZhgIc", "cccc-24-last-pKZFymAIAI", "aaaa-25-first-bCPdzXlaxh", "aaaa-25-last-rCstLRNHQJ", "bbbb-25-first-sEhtzlaliv", "bbbb-25-last-XjMZyqUefw", "cccc-25-first-amQWeKBaWW", "cccc-25-last-BmrBCbOIwL", "aaaa-26-first-TaVeGVwhKN", "aaaa-26-last-imyxeiBOpe", "bbbb-26-first-lStSCzoLpn", "bbbb-26-last-fWHqiXowMU", "cccc-26-first-makWRXTxaC", "cccc-26-last-ZFVyaRpyol", "aaaa-27-first-mjQZhjATPI", "aaaa-27-last-JXcuNDNYvl", "bbbb-27-first-xPrZtczpZs", "bbbb-27-last-enNmGJoyJo", "cccc-27-first-VbOelBCjnp", "cccc-27-last-VoBalZXodi", "aaaa-28-first-AEXqeKKRVg", "aaaa-28-last-xnEedruqgJ", "bbbb-28-first-ZTPELgSaDn", "bbbb-28-last-goskLxnrHM", "cccc-28-first-cqaUWhNLFE", "cccc-28-last-FawEabyZQb", "aaaa-29-first-QaPIajbRKH", "aaaa-29-last-faLkekNvBr", "bbbb-29-first-LPFlWCKzBr", "bbbb-29-last-QQzcWFupTQ", "cccc-29-first-qIjhGXbKyp", "cccc-29-last-QxyxegOPFO"}
	fmt.Println(ChangeOutput(data))
}
