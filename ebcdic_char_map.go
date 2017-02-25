package main
var (
	/**
	 * Map EBCDIC byte to ASCII equivalent. If there is no ASCII equivalent we map it to 0x00 (NUL)
	 */
	ebcdic_to_ascii = map[byte]byte {
		0x00: 0x00,  // NUL -> NUL
		0x01: 0x01,  // SOH -> SOH
		0x02: 0x02,  // STX -> STX
		0x03: 0x03,  // ETX -> ETX
		0x04: 0x00,  // PF -> 
		0x05: 0x09,  // HT -> HT
		0x06: 0x00,  // LC -> 
		0x07: 0x7F,  // DEL -> DEL
		0x08: 0x00,  // GE -> 
		0x09: 0x00,  // RPT -> 
		0x0A: 0x00,  // SMM -> 
		0x0B: 0x0B,  // VT -> VT
		0x0C: 0x0C,  // FF -> FF
		0x0D: 0x0D,  // CR -> CR
		0x0E: 0x0E,  // SO -> SO
		0x0F: 0x0F,  // SI -> OF
		0x10: 0x10,  // DLE -> DLE
		0x11: 0x11,  // DC1 -> DC1
		0x12: 0x12,  // DC2 -> DC2
		0x13: 0x13,  // DCE -> DC3
		0x14: 0x00,  // RES -> 
		0x15: 0x0A,  // NL -> LF
		0x16: 0x0B,  // BS -> BS
		0x17: 0x00,  // IL -> 
		0x18: 0x18,  // CAN -> CAN
		0x19: 0x19,  // EM -> EM
		0x1A: 0x00,  // UBS -> 
		0x1B: 0x00,  // CU1 -> 
		0x1C: 0x1C,  // FS -> FS
		0x1D: 0x1D,  // GS -> GS
		0x1E: 0x1E,  // RS -> RS
		0x1F: 0x1F,  // US -> US
		0x20: 0x00,  // DS -> 
		0x21: 0x00,  // SOS -> 
		0x22: 0x1C,  // FS -> FS
		0x23: 0x00,  // WUS -> 
		0x24: 0x00,  // BYP -> 
		0x25: 0x0A,  // LF -> LF
		0x26: 0x17,  // ETB -> ETB
		0x27: 0x1B,  // ESC -> ESC
		0x28: 0x00,  //  -> 
		0x29: 0x00,  //  -> 
		0x2A: 0x00,  // SM -> 
		0x2B: 0x00,  // CU2 -> 
		0x2C: 0x00,  //  -> 
		0x2D: 0x05,  // ENQ -> ENQ
		0x2E: 0x06,  // ACK -> ACK
		0x2F: 0x07,  // BEL -> BEL
		0x30: 0x00,  //  -> 
		0x31: 0x00,  //  -> 
		0x32: 0x16,  // SYN -> SYN
		0x33: 0x00,  //  -> 
		0x34: 0x00,  // PP -> 
		0x35: 0x00,  // RS -> 
		0x36: 0x00,  // UC -> 
		0x37: 0x04,  // EOT -> EOT
		0x38: 0x00,  // SBS -> 
		0x39: 0x00,  // IT -> 
		0x3A: 0x00,  // RFF -> 
		0x3B: 0x00,  // CU3 -> 
		0x3C: 0x14,  // DC4 -> DC4
		0x3D: 0x15,  // NAK -> NAK
		0x3E: 0x00,  //  -> 
		0x3F: 0x1A,  // SUB -> SUB
		0x40: 0x20,  //  -> 
		0x41: 0x00,  //  -> 
		0x42: 0xE2,  // â -> â
		0x43: 0xE4,  // ä -> ä
		0x44: 0xE0,  // à -> à
		0x45: 0xE1,  // á -> á
		0x46: 0xE3,  // ã -> ã
		0x47: 0xE5,  // å -> å
		0x48: 0xE7,  // ç -> ç
		0x49: 0xF0,  // ñ -> ñ
		0x4A: 0xA2,  // ¢ -> ¢
		0x4B: 0x2E,  // . -> .
		0x4C: 0x3C,  // < -> <
		0x4D: 0x28,  // ( -> (
		0x4E: 0x2B,  // + -> +
		0x4F: 0x7C,  // | -> |
		0x50: 0x26,  // & -> &
		0x51: 0xE9,  // é -> é
		0x52: 0xEA,  // ê -> ê
		0x53: 0xEB,  // ë -> ë
		0x54: 0xE8,  // è -> è
		0x55: 0xED,  // í -> í
		0x56: 0xEE,  // î -> î
		0x57: 0xEF,  // ï -> ï
		0x58: 0xEC,  // ì -> ì
		0x59: 0xDF,  // ß -> ß
		0x5A: 0x21,  // ! -> !
		0x5B: 0x24,  // $ -> $
		0x5C: 0x2A,  // * -> *
		0x5D: 0x29,  // ) -> )
		0x5E: 0x3B,  // ; -> ;
		0x5F: 0x5E,  // ^ -> ^
		0x60: 0x2D,  // - -> -
		0x61: 0x2F,  // / -> /
		0x62: 0xC2,  // Â -> Â
		0x63: 0xC4,  // Ä -> Ä
		0x64: 0xC0,  // À -> À
		0x65: 0xC1,  // Á -> Á
		0x66: 0xC3,  // Ã -> Ã
		0x67: 0xC5,  // Å -> Å
		0x68: 0xC7,  // Ç -> Ç
		0x69: 0xD1,  // Ñ -> Ñ
		0x6A: 0xA6,  // ¦ -> ¦
		0x6B: 0x2C,  // , -> ,
		0x6C: 0x25,  // % -> %
		0x6D: 0x5F,  // _ -> _
		0x6E: 0x3E,  // > -> >
		0x6F: 0x3F,  // ? -> ?
		0x70: 0xF8,  // ø -> ø
		0x71: 0xC9,  // É -> É
		0x72: 0xCA,  // Ê -> Ê
		0x73: 0xCB,  // Ë -> Ë
		0x74: 0xC8,  // È -> È
		0x75: 0xCD,  // Í -> Í
		0x76: 0xCE,  // Î -> Î
		0x77: 0xCF,  // Ï -> Ï
		0x78: 0xCC,  // Ì -> Ì
		0x79: 0x5F,  // ` -> `
		0x7A: 0x3A,  // : -> :
		0x7B: 0x23,  // # -> #
		0x7C: 0x40,  // @ -> @
		0x7D: 0x27,  //  -> 
		0x7E: 0x3D,  // = -> =
		0x7F: 0x22,  // " -> "
		0x80: 0xD8,  // Ø -> Ø
		0x81: 0x61,  // a -> a
		0x82: 0x62,  // b -> b
		0x83: 0x63,  // c -> c
		0x84: 0x64,  // d -> d
		0x85: 0x65,  // e -> e
		0x86: 0x66,  // f -> f
		0x87: 0x67,  // g -> g
		0x88: 0x68,  // h -> h
		0x89: 0x69,  // i -> i
		0x8A: 0xAB,  // « -> «
		0x8B: 0xBB,  // » -> »
		0x8C: 0xF0,  // ð -> ð
		0x8D: 0x0D,  // ý -> ý
		0x8E: 0xDE,  // Þ -> Þ
		0x8F: 0xB1,  // ± -> ±
		0x90: 0xB0,  // ° -> °
		0x91: 0x6A,  // j -> j
		0x92: 0x6B,  // k -> k
		0x93: 0x6C,  // l -> l
		0x94: 0x6D,  // m -> m
		0x95: 0x6E,  // n -> n
		0x96: 0x6F,  // o -> o
		0x97: 0x70,  // p -> p
		0x98: 0x71,  // q -> q
		0x99: 0x72,  // r -> r
		0x9A: 0xAA,  // ª -> ª
		0x9B: 0xBA,  // º -> º
		0x9C: 0xE6,  // æ -> æ
		0x9D: 0xB8,  // ¸ -> ¸
		0x9E: 0xC6,  // Æ -> Æ
		0x9F: 0xA4,  // ¤ -> ¤
		0xA0: 0xB5,  // µ -> µ
		0xA1: 0x7E,  // ~ -> ~
		0xA2: 0x73,  // s -> s
		0xA3: 0x74,  // t -> t
		0xA4: 0x75,  // u -> u
		0xA5: 0x76,  // v -> v
		0xA6: 0x77,  // w -> w
		0xA7: 0x78,  // x -> x
		0xA8: 0x79,  // y -> y
		0xA9: 0x7A,  // z -> z
		0xAA: 0xA1,  // ¡ -> ¡
		0xAB: 0xBF,  // ¿ -> ¿
		0xAC: 0xD0,  // Ð -> Ð
		0xAD: 0x5B,  // [ -> [
		0xAE: 0xFE,  // þ -> þ
		0xAF: 0xAE,  // ® -> ®
		0xB0: 0xAC,  // ¬ -> ¬
		0xB1: 0xA3,  // £ -> £
		0xB2: 0xA5,  // ¥ -> ¥
		0xB3: 0x95,  // • -> •
		0xB4: 0xA9,  // © -> ©
		0xB5: 0xA7,  // § -> §
		0xB6: 0xB6,  // ¶ -> ¶
		0xB7: 0xBC,  // ¼ -> ¼
		0xB8: 0xBD,  // ½ -> ½
		0xB9: 0xBE,  // ¾ -> ¾
		0xBA: 0xDD,  // Ý -> Ý
		0xBB: 0xA8,  // ¨ -> ¨
		0xBC: 0xAF,  // ¯ -> ¯
		0xBD: 0x5D,  // ] -> ]
		0xBE: 0x92,  // ’ -> ’
		0xBF: 0xD7,  // × -> ×
		0xC0: 0x7B,  // { -> {
		0xC1: 0x41,  // A -> A
		0xC2: 0x42,  // B -> B
		0xC3: 0x43,  // C -> C
		0xC4: 0x44,  // D -> D
		0xC5: 0x45,  // E -> E
		0xC6: 0x46,  // F -> F
		0xC7: 0x47,  // G -> G
		0xC8: 0x48,  // H -> H
		0xC9: 0x49,  // I -> I
		0xCA: 0x9B,  // – -> –
		0xCB: 0xF4,  // ô -> ô
		0xCC: 0xF6,  // ö -> ö
		0xCD: 0xF2,  // ò -> ò
		0xCE: 0xF3,  // ó -> ó
		0xCF: 0xF5,  // õ -> õ
		0xD0: 0x7D,  // } -> }
		0xD1: 0x4A,  // J -> J
		0xD2: 0x4B,  // K -> K
		0xD3: 0x4C,  // L -> L
		0xD4: 0x4D,  // M -> M
		0xD5: 0x4E,  // N -> N
		0xD6: 0x4F,  // O -> O
		0xD7: 0x50,  // P -> P
		0xD8: 0x51,  // Q -> Q
		0xD9: 0x52,  // R -> R
		0xDA: 0xB9,  // ¹ -> ¹
		0xDB: 0xFB,  // û -> û
		0xDC: 0xFC,  // ü -> ü
		0xDD: 0xF9,  // ù -> ù
		0xDE: 0xFA,  // ú -> ú
		0xDF: 0xFF,  // ÿ -> ÿ
		0xE0: 0x5C,  // \ -> \
		0xE1: 0xF7,  // ÷ -> ÷
		0xE2: 0x53,  // S -> S
		0xE3: 0x54,  // T -> T
		0xE4: 0x55,  // U -> U
		0xE5: 0x56,  // V -> V
		0xE6: 0x57,  // W -> W
		0xE7: 0x58,  // X -> X
		0xE8: 0x59,  // Y -> Y
		0xE9: 0x5A,  // Z -> Z
		0xEA: 0xB2,  // ² -> ²
		0xEB: 0xD4,  // Ô -> Ô
		0xEC: 0xD6,  // Ö -> Ö
		0xED: 0xD2,  // Ò -> Ò
		0xEE: 0xD3,  // Ó -> Ó
		0xEF: 0xD5,  // Õ -> Õ
		0xF0: 0x30,  // 0 -> 0
		0xF1: 0x31,  // 1 -> 1
		0xF2: 0x32,  // 2 -> 2
		0xF3: 0x33,  // 3 -> 3
		0xF4: 0x34,  // 4 -> 4
		0xF5: 0x35,  // 5 -> 5
		0xF6: 0x36,  // 6 -> 6
		0xF7: 0x37,  // 7 -> 7
		0xF8: 0x38,  // 8 -> 8
		0xF9: 0x39,  // 9 -> 9
		0xFA: 0xB3,  // ³ -> ³
		0xFB: 0xDB,  // Û -> Û
		0xFC: 0xDC,  // Ü -> Ü
		0xFD: 0xD9,  // Ù -> Ù
		0xFE: 0xDA,  // Ú -> Ú
		0xFF: 0x00,  //  -> 
	}
)