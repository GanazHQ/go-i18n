// This file is generated by i18n/plural/codegen/generate.sh; DO NOT EDIT

package plural

// DefaultRules returns a map of Rules generated from CLDR language data.
func DefaultRules() Rules {
	rules := Rules{}

	addPluralRules(rules, []string{"bm", "bo", "dz", "id", "ig", "ii", "in", "ja", "jbo", "jv", "jw", "kde", "kea", "km", "ko", "lkt", "lo", "ms", "my", "nqo", "root", "sah", "ses", "sg", "th", "to", "vi", "wo", "yo", "yue", "zh"}, &Rule{
		PluralForms: newPluralFormSet(Other),
		PluralFormFunc: func(ops *Operands) Form {
			return Other
		},
	})
	addPluralRules(rules, []string{"am", "as", "bn", "fa", "gu", "hi", "kn", "zu"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 0 or n = 1
			if intEqualsAny(ops.I, 0) ||
				ops.NEqualsAny(1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ff", "fr", "hy", "kab"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 0,1
			if intEqualsAny(ops.I, 0, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"pt"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 0..1
			if intInRange(ops.I, 0, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ast", "ca", "de", "en", "et", "fi", "fy", "gl", "ia", "io", "it", "ji", "nl", "pt_PT", "sc", "scn", "sv", "sw", "ur", "yi", "ht"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"si"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0,1 or i = 0 and f = 1
			if ops.NEqualsAny(0, 1) ||
				intEqualsAny(ops.I, 0) && intEqualsAny(ops.F, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ak", "bh", "guw", "ln", "mg", "nso", "pa", "ti", "wa"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0..1
			if ops.NInRange(0, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"tzm"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0..1 or n = 11..99
			if ops.NInRange(0, 1) ||
				ops.NInRange(11, 99) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"af", "asa", "az", "bem", "bez", "bg", "brx", "ce", "cgg", "chr", "ckb", "dv", "ee", "el", "eo", "es", "eu", "fo", "fur", "gsw", "ha", "haw", "hu", "jgo", "jmc", "ka", "kaj", "kcg", "kk", "kkj", "kl", "ks", "ksb", "ku", "ky", "lb", "lg", "mas", "mgo", "ml", "mn", "mr", "nah", "nb", "nd", "ne", "nn", "nnh", "no", "nr", "ny", "nyn", "om", "or", "os", "pap", "ps", "rm", "rof", "rwk", "saq", "sd", "sdh", "seh", "sn", "so", "sq", "ss", "ssy", "st", "syr", "ta", "te", "teo", "tig", "tk", "tn", "tr", "ts", "ug", "uz", "ve", "vo", "vun", "wae", "xh", "xog"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"da"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1 or t != 0 and i = 0,1
			if ops.NEqualsAny(1) ||
				!intEqualsAny(ops.T, 0) && intEqualsAny(ops.I, 0, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"is"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// t = 0 and i % 10 = 1 and i % 100 != 11 or t != 0
			if intEqualsAny(ops.T, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
				!intEqualsAny(ops.T, 0) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"mk"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 10 = 1 and i % 100 != 11 or f % 10 = 1 and f % 100 != 11
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
				intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ceb", "fil", "tl"}, &Rule{
		PluralForms: newPluralFormSet(One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i = 1,2,3 or v = 0 and i % 10 != 4,6,9 or v != 0 and f % 10 != 4,6,9
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I, 1, 2, 3) ||
				intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I%10, 4, 6, 9) ||
				!intEqualsAny(ops.V, 0) && !intEqualsAny(ops.F%10, 4, 6, 9) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"lv", "prg"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n % 10 = 0 or n % 100 = 11..19 or v = 2 and f % 100 = 11..19
			if ops.NModEqualsAny(10, 0) ||
				ops.NModInRange(100, 11, 19) ||
				intEqualsAny(ops.V, 2) && intInRange(ops.F%100, 11, 19) {
				return Zero
			}
			// n % 10 = 1 and n % 100 != 11 or v = 2 and f % 10 = 1 and f % 100 != 11 or v != 2 and f % 10 = 1
			if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11) ||
				intEqualsAny(ops.V, 2) && intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) ||
				!intEqualsAny(ops.V, 2) && intEqualsAny(ops.F%10, 1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"lag"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// i = 0,1 and n != 0
			if intEqualsAny(ops.I, 0, 1) && !ops.NEqualsAny(0) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ksh"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"iu", "naq", "se", "sma", "smi", "smj", "smn", "sms"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 2
			if ops.NEqualsAny(2) {
				return Two
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"shi"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 0 or n = 1
			if intEqualsAny(ops.I, 0) ||
				ops.NEqualsAny(1) {
				return One
			}
			// n = 2..10
			if ops.NInRange(2, 10) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"mo", "ro"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// v != 0 or n = 0 or n % 100 = 2..19
			if !intEqualsAny(ops.V, 0) ||
				ops.NEqualsAny(0) ||
				ops.NModInRange(100, 2, 19) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"bs", "hr", "sh", "sr"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 10 = 1 and i % 100 != 11 or f % 10 = 1 and f % 100 != 11
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) ||
				intEqualsAny(ops.F%10, 1) && !intEqualsAny(ops.F%100, 11) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14 or f % 10 = 2..4 and f % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) ||
				intInRange(ops.F%10, 2, 4) && !intInRange(ops.F%100, 12, 14) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"gd"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1,11
			if ops.NEqualsAny(1, 11) {
				return One
			}
			// n = 2,12
			if ops.NEqualsAny(2, 12) {
				return Two
			}
			// n = 3..10,13..19
			if ops.NInRange(3, 10) || ops.NInRange(13, 19) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"sl"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 100 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 1) {
				return One
			}
			// v = 0 and i % 100 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 2) {
				return Two
			}
			// v = 0 and i % 100 = 3..4 or v != 0
			if intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 3, 4) ||
				!intEqualsAny(ops.V, 0) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"dsb", "hsb"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 100 = 1 or f % 100 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 1) ||
				intEqualsAny(ops.F%100, 1) {
				return One
			}
			// v = 0 and i % 100 = 2 or f % 100 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 2) ||
				intEqualsAny(ops.F%100, 2) {
				return Two
			}
			// v = 0 and i % 100 = 3..4 or f % 100 = 3..4
			if intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 3, 4) ||
				intInRange(ops.F%100, 3, 4) {
				return Few
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"he", "iw"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// i = 2 and v = 0
			if intEqualsAny(ops.I, 2) && intEqualsAny(ops.V, 0) {
				return Two
			}
			// v = 0 and n != 0..10 and n % 10 = 0
			if intEqualsAny(ops.V, 0) && !ops.NInRange(0, 10) && ops.NModEqualsAny(10, 0) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"cs", "sk"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// i = 2..4 and v = 0
			if intInRange(ops.I, 2, 4) && intEqualsAny(ops.V, 0) {
				return Few
			}
			// v != 0
			if !intEqualsAny(ops.V, 0) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"pl"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// i = 1 and v = 0
			if intEqualsAny(ops.I, 1) && intEqualsAny(ops.V, 0) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) {
				return Few
			}
			// v = 0 and i != 1 and i % 10 = 0..1 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 12..14
			if intEqualsAny(ops.V, 0) && !intEqualsAny(ops.I, 1) && intInRange(ops.I%10, 0, 1) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 5, 9) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 12, 14) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"be"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n % 10 = 1 and n % 100 != 11
			if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11) {
				return One
			}
			// n % 10 = 2..4 and n % 100 != 12..14
			if ops.NModInRange(10, 2, 4) && !ops.NModInRange(100, 12, 14) {
				return Few
			}
			// n % 10 = 0 or n % 10 = 5..9 or n % 100 = 11..14
			if ops.NModEqualsAny(10, 0) ||
				ops.NModInRange(10, 5, 9) ||
				ops.NModInRange(100, 11, 14) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"lt"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n % 10 = 1 and n % 100 != 11..19
			if ops.NModEqualsAny(10, 1) && !ops.NModInRange(100, 11, 19) {
				return One
			}
			// n % 10 = 2..9 and n % 100 != 11..19
			if ops.NModInRange(10, 2, 9) && !ops.NModInRange(100, 11, 19) {
				return Few
			}
			// f != 0
			if !intEqualsAny(ops.F, 0) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"mt"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 0 or n % 100 = 2..10
			if ops.NEqualsAny(0) ||
				ops.NModInRange(100, 2, 10) {
				return Few
			}
			// n % 100 = 11..19
			if ops.NModInRange(100, 11, 19) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ru", "uk"}, &Rule{
		PluralForms: newPluralFormSet(One, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 10 = 1 and i % 100 != 11
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) && !intEqualsAny(ops.I%100, 11) {
				return One
			}
			// v = 0 and i % 10 = 2..4 and i % 100 != 12..14
			if intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 2, 4) && !intInRange(ops.I%100, 12, 14) {
				return Few
			}
			// v = 0 and i % 10 = 0 or v = 0 and i % 10 = 5..9 or v = 0 and i % 100 = 11..14
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 0) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I%10, 5, 9) ||
				intEqualsAny(ops.V, 0) && intInRange(ops.I%100, 11, 14) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"br"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n % 10 = 1 and n % 100 != 11,71,91
			if ops.NModEqualsAny(10, 1) && !ops.NModEqualsAny(100, 11, 71, 91) {
				return One
			}
			// n % 10 = 2 and n % 100 != 12,72,92
			if ops.NModEqualsAny(10, 2) && !ops.NModEqualsAny(100, 12, 72, 92) {
				return Two
			}
			// n % 10 = 3..4,9 and n % 100 != 10..19,70..79,90..99
			if (ops.NModInRange(10, 3, 4) || ops.NModEqualsAny(10, 9)) && !(ops.NModInRange(100, 10, 19) || ops.NModInRange(100, 70, 79) || ops.NModInRange(100, 90, 99)) {
				return Few
			}
			// n != 0 and n % 1000000 = 0
			if !ops.NEqualsAny(0) && ops.NModEqualsAny(1000000, 0) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ga"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 2
			if ops.NEqualsAny(2) {
				return Two
			}
			// n = 3..6
			if ops.NInRange(3, 6) {
				return Few
			}
			// n = 7..10
			if ops.NInRange(7, 10) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"gv"}, &Rule{
		PluralForms: newPluralFormSet(One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// v = 0 and i % 10 = 1
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 1) {
				return One
			}
			// v = 0 and i % 10 = 2
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%10, 2) {
				return Two
			}
			// v = 0 and i % 100 = 0,20,40,60,80
			if intEqualsAny(ops.V, 0) && intEqualsAny(ops.I%100, 0, 20, 40, 60, 80) {
				return Few
			}
			// v != 0
			if !intEqualsAny(ops.V, 0) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"ar", "ars"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 2
			if ops.NEqualsAny(2) {
				return Two
			}
			// n % 100 = 3..10
			if ops.NModInRange(100, 3, 10) {
				return Few
			}
			// n % 100 = 11..99
			if ops.NModInRange(100, 11, 99) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"cy"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n = 2
			if ops.NEqualsAny(2) {
				return Two
			}
			// n = 3
			if ops.NEqualsAny(3) {
				return Few
			}
			// n = 6
			if ops.NEqualsAny(6) {
				return Many
			}
			return Other
		},
	})
	addPluralRules(rules, []string{"kw"}, &Rule{
		PluralForms: newPluralFormSet(Zero, One, Two, Few, Many, Other),
		PluralFormFunc: func(ops *Operands) Form {
			// n = 0
			if ops.NEqualsAny(0) {
				return Zero
			}
			// n = 1
			if ops.NEqualsAny(1) {
				return One
			}
			// n % 100 = 2,22,42,62,82 or n%1000 = 0 and n%100000=1000..20000,40000,60000,80000 or n!=0 and n%1000000=100000
			if ops.NModEqualsAny(100, 2, 22, 42, 62, 82) ||
				ops.NModEqualsAny(1000, 0) && (ops.NModInRange(100000, 1000, 20000) || ops.NModEqualsAny(100000, 40000, 60000, 80000)) ||
				!ops.NEqualsAny(0) && ops.NModEqualsAny(1000000, 100000) {
				return Two
			}
			// n % 100 = 3,23,43,63,83
			if ops.NModEqualsAny(100, 3, 23, 43, 63, 83) {
				return Few
			}
			// n != 1 and n % 100 = 1,21,41,61,81
			if !ops.NEqualsAny(1) && ops.NModEqualsAny(100, 1, 21, 41, 61, 81) {
				return Many
			}
			return Other
		},
	})

	return rules
}
