// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /home/ubuntu/go_1.11.11/go/src/cmd/compile/internal/gc/op_string.go

//line /home/ubuntu/go_1.11.11/go/src/cmd/compile/internal/gc/op_string.go:1
// Code generated by "stringer -type=Op -trimprefix=O"; DO NOT EDIT.

package gc

import "strconv"

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALADDSUBORXORADDSTRADDRANDANDAPPENDARRAYBYTESTRARRAYBYTESTRTMPARRAYRUNESTRSTRARRAYBYTESTRARRAYBYTETMPSTRARRAYRUNEASAS2AS2FUNCAS2RECVAS2MAPRAS2DOTTYPEASOPCALLCALLFUNCCALLMETHCALLINTERCALLPARTCAPCLOSECLOSURECMPIFACECMPSTRCOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDCLFIELDDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTINDINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMULDIVMODLSHRSHANDANDNOTNEWNOTCOMPLUSMINUSORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRRECOVERRECVRUNESTRSELRECVSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFBLOCKBREAKCASEXCASECONTINUEDEFEREMPTYFALLFORFORUNTILGOTOIFLABELPROCRANGERETURNSELECTSWITCHTYPESWTCHANTMAPTSTRUCTTINTERTFUNCTARRAYDDDDDDARGINLCALLEFACEITABIDATASPTRCLOSUREVARCFUNCCHECKNILVARKILLVARLIVEINDREGSPRETJMPGETGEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 36, 39, 45, 49, 55, 61, 73, 88, 100, 112, 127, 139, 141, 144, 151, 158, 165, 175, 179, 183, 191, 199, 208, 216, 219, 224, 231, 239, 245, 252, 258, 267, 275, 283, 289, 293, 302, 309, 313, 316, 323, 331, 339, 346, 352, 355, 361, 368, 376, 380, 387, 395, 397, 399, 401, 403, 405, 407, 410, 415, 423, 426, 435, 438, 442, 450, 457, 466, 469, 472, 475, 478, 481, 484, 490, 493, 496, 499, 503, 508, 512, 517, 522, 528, 533, 537, 542, 550, 558, 564, 573, 580, 584, 591, 598, 606, 610, 614, 618, 625, 632, 640, 646, 651, 656, 660, 665, 673, 678, 683, 687, 690, 698, 702, 704, 709, 713, 718, 724, 730, 736, 742, 747, 751, 758, 764, 769, 775, 778, 784, 791, 796, 800, 805, 809, 819, 824, 832, 839, 846, 854, 860, 864, 867}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
