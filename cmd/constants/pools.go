package constants

// Uniswap V2 Pool Addresses

var UniV2Pools = map[string]string{
	"eEUR_eAUD_Pool": "0x47167006b08358292bc99eb1be24124e7363ba50",
    "eEUR_eCAD_Pool": "0x37a0b08cfa766fb3ef9cf1792be438022f46e515",
    "eEUR_eCHF_Pool": "0x850718e01057785eeb4cdb8650cc0359a020af09",
    "eEUR_eGBP_Pool": "0x3f42d3c827bd9344c4d87ae94effae836736c714",
    "eEUR_eHKD_Pool": "0x1fa3171bc411f3871da26a15d2808feabfc72846",
    "eEUR_eJPY_Pool": "0x591903689c23863e1ae3a449765c2e5ab27121a1",
    "eEUR_eMXN_Pool": "0xbbeb15978deff140b5dda965c465255c63afee37",
    "eEUR_eNOK_Pool": "0x3b0dd9f3f8a873d012e49ecf391d4cbf50475ebb",
    "eEUR_eNZD_Pool": "0x08a15828cde24297ec0b6270f6abc1609f536840",
    "eEUR_eSDR_Pool": "0x3c00d0936c04694bef7018172b3929fb820df8bd",
    "eEUR_eSGD_Pool": "0xe1686ad2bee3ecae5c92ed00a46b6ae1f0291786",
    "eEUR_eZAR_Pool": "0x93de9c487cdbf3cbcff2e2bdf9f5d005d980a1b4",
    "eUSD_eAUD_Pool": "0x2ba6cadcfe56c665c3e5b73fbdfb46808fa1d2b7",
    "eUSD_eCAD_Pool": "0x6f6709bb8ba052906cfd5b728fb86de2fb2b2315",
    "eUSD_eCHF_Pool": "0xcbe821513e253e6ea68ec4816c491365ecea209e",
    "eUSD_eEUR_Pool": "0x7cd9e0a93cac1729369c5d428daed0b1f9d07fe7",
    "eUSD_eGBP_Pool": "0xa8a4b9f80d42654dd260e964533cebae00781e4c",
    "eUSD_eHKD_Pool": "0x6687312904abd15a8d1b258443298706726bdc01",
    "eUSD_eJPY_Pool": "0x4ccbb5a234224f1dfe998cb07a38706e2872dd48",
    "eUSD_eMXN_Pool": "0x5e1ff572196e3b20f21d8f6cb43e8b3b36f14765",
    "eUSD_eNOK_Pool": "0x972c2d5a94411b996c9c5f34156b58ba5fb59bc8",
    "eUSD_eNZD_Pool": "0x64f6cefc2a2446db2ffe89318f1ceb7a25146caf",
    "eUSD_eSDR_Pool": "0x93f12e4ab16e7c112cfcba1167c7caa892a3246e",
    "eUSD_eSGD_Pool": "0x8ec58cbb1197313a9f2061e9bb5a1335e1719a7a",
    "eUSD_eZAR_Pool": "0x8c173d881318a0f4ff327b79fdb38f4a20c28f05",
    "wTEL_eUSD_Pool": "0x403663b51c85d628ecc39d6ce480dafbefd238d2",
    "wTEL_eEUR_Pool": "0x0b3fe394a0faeb9011bfc1f0893080645e4266fa",
}

// Target ratios for Pools (for demo purposes)
// These would be the "ideal" ratios for each pool
var TargetRatios = map[string]float64{
	"eEUR_eAUD_Pool": 1.64, // Example: 1 eEUR = 1.64 eAUD
    "eEUR_eCAD_Pool": 1.46, // Example: 1 eEUR = 1.46 eCAD
    "eUSD_eEUR_Pool": 0.92, // Example: 1 eUSD = 0.92 eEUR
    // Add more as needed
}