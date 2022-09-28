package exiencode0_1

import (
	"errors"
	"os"
	"strings"
)

/////////////////////HeaderOptionsOutputType.class///////////////
type HeaderOptionsOutputType struct {
	none         string
	lessSchemaId string
	all          string
}

var VARHeaderOptionsOutputType HeaderOptionsOutputType = HeaderOptionsOutputType{"none", "lessSchemaId", "all"}

/////////////////////////end HeaderOptionsOutputType.class//////
//////////EXIOptions.class////////////////////////////////////////////////////////////////////////////////////////////
type EXIOptions struct {
	ADD_LESSCOMMON                        int
	ADD_UNCOMMON                          int
	ADD_ALIGNMENT                         int
	ADD_PRESERVE                          int
	ADD_COMMON                            int
	ADD_VALUE_MAX_LENGTH                  int
	ADD_VALUE_PARTITION_CAPACITY          int
	ADD_FRAGMENT                          int
	ADD_DTRM                              int
	BLOCKSIZE_DEFAULT                     int
	VALUE_MAX_LENGTH_UNBOUNDED            int
	VALUE_PARTITION_CAPACITY_UNBOUNDED    int
	m_alignmentType                       AlignmentType
	m_isFragment                          bool
	m_isStrict                            bool
	m_preserveComments                    bool
	m_preservePIs                         bool
	m_preserveDTD                         bool
	m_preserveNS                          bool
	m_infuseSC                            bool
	m_schemaId                            SchemaId
	m_blockSize                           int
	m_valueMaxLength                      int
	m_valuePartitionCapacity              int
	m_preserveLexicalValues               bool
	m_n_datatypeRepresentationMapBindings int
	m_datatypeRepresentationMap           []QName
}

func (e *EXIOptions) EXIOptions() {
	//e.init()//init() используется для выполнения инициализации сервлета. Создает и загружает объекты, которые используются сервлетом при обработке его запроса.
	e.m_datatypeRepresentationMap = make([]QName, 16, 16)
	/*//думаю не нужно
	for var1 := 0; var1 < len(q.m_datatypeRepresentationMap); var1++ {
	   q.m_datatypeRepresentationMap[var1] = new QName();
	}
	*/
}

var VAREXIOptions EXIOptions = EXIOptions{
	ADD_LESSCOMMON:                     1,
	ADD_UNCOMMON:                       2,
	ADD_ALIGNMENT:                      4,
	ADD_PRESERVE:                       8,
	ADD_COMMON:                         16,
	ADD_VALUE_MAX_LENGTH:               32,
	ADD_VALUE_PARTITION_CAPACITY:       64,
	ADD_FRAGMENT:                       128,
	ADD_DTRM:                           256,
	BLOCKSIZE_DEFAULT:                  1000000,
	VALUE_MAX_LENGTH_UNBOUNDED:         -1,
	VALUE_PARTITION_CAPACITY_UNBOUNDED: -1,
}

func (e *EXIOptions) getAlignmentType() AlignmentType {
	return e.m_alignmentType
}

func (e *EXIOptions) isFragment() bool {
	return e.m_isFragment
}
func (e *EXIOptions) isStrict() bool {
	return e.m_isStrict
}
func (e *EXIOptions) getPreserveComments() bool {
	return e.m_preserveComments
}
func (e *EXIOptions) getPreservePIs() bool {
	return e.m_preservePIs
}

func (e *EXIOptions) getPreserveDTD() bool {
	return e.m_preserveDTD
}

func (e *EXIOptions) getPreserveNS() bool {
	return e.m_preserveNS
}

func (e *EXIOptions) getInfuseSC() bool {
	return e.m_infuseSC
}

func (e *EXIOptions) getSchemaId() SchemaId {
	return e.m_schemaId
}

func (e *EXIOptions) getBlockSize() int {
	return e.m_blockSize
}

func (e *EXIOptions) getValueMaxLength() int {
	return e.m_valueMaxLength
}

func (e *EXIOptions) getValuePartitionCapacity() int {
	return e.m_valuePartitionCapacity
}

func (e *EXIOptions) getPreserveLexicalValues() bool {
	return e.m_preserveLexicalValues
}

func (e *EXIOptions) getDatatypeRepresentationMapBindingsCount() int {
	return e.m_n_datatypeRepresentationMapBindings
}

func (e *EXIOptions) getDatatypeRepresentationMap() []QName {
	return e.m_datatypeRepresentationMap
}
func (e *EXIOptions) setAlignmentType(VARAlignmentType AlignmentType) (*EXIOptions, error) { ////throws EXIOptionsException
	var err error
	var var2 AlignmentType
	if e.m_infuseSC {
		err = EXIOptionsException1(e, VARAlignmentType)
	}
	if err != nil {
		e.m_alignmentType = var2
		return e, err //e.m_alignmentType,err
	}
	e.m_alignmentType = VARAlignmentType
	return e, err ////e.m_alignmentType,err
}
func EXIOptionsException1(e *EXIOptions, VARAlignmentType AlignmentType) error { //my custom
	var err error
	if e.m_alignmentType.compress == VARAlignmentType.compress { //!!!??как переменная типа AlignmentType сравниваеться с частью самого AlignmentType!!!????думаю так....
		err = errors.New("selfContained option and compression option cannot be used together")
	}
	if e.m_alignmentType.preCompress == VARAlignmentType.preCompress { //!!!??как переменная типа AlignmentType сравниваеться с частью самого AlignmentType!!!???думаю так....
		err = errors.New("selfContained option and pre-compression option cannot be used together")
	}
	return err
}

func (e *EXIOptions) setFragment(var1 bool) *EXIOptions {
	e.m_isFragment = var1
	return e
}

func (e *EXIOptions) setStrict(var1 bool) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_preserveComments {
		err = errors.New("preserve.comments option and strict option cannot be used together")
		e.m_isStrict = var2
		return e, err
	} else if e.m_preservePIs {
		err = errors.New("preserve.pis option and strict option cannot be used together")
		e.m_isStrict = var2
		return e, err
	} else if e.m_preserveDTD {
		err = errors.New("preserve.dtd option and strict option cannot be used together")
		e.m_isStrict = var2
		return e, err
	} else if e.m_preserveNS {
		err = errors.New("preserve.prefixes option and strict option cannot be used together")
		e.m_isStrict = var2
		return e, err
	} else if e.m_infuseSC {
		err = errors.New("selfContained option and strict option cannot be used together")
		e.m_isStrict = var2
		return e, err
	}
	e.m_isStrict = var1
	return e, err
}

func (e *EXIOptions) setPreserveComments(var1 bool) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_isStrict && var1 {
		err = errors.New("preserve.comments option and strict option cannot be used together")
		e.m_preserveComments = var2
		return e, err
	}
	e.m_preserveComments = var1
	return e, err
}

func (e *EXIOptions) setPreservePIs(var1 bool) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_isStrict && var1 {
		err = errors.New("preserve.pis option and strict option cannot be used together")
		e.m_preservePIs = var2
		return e, err
	}
	e.m_preservePIs = var1
	return e, err
}

func (e *EXIOptions) setPreserveDTD(var1 bool) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_isStrict && var1 {
		err = errors.New("preserve.dtd option and strict option cannot be used together")
		e.m_preserveDTD = var2
		return e, err
	}
	e.m_preserveDTD = var1
	return e, err
}

func (e *EXIOptions) setPreserveNS(var1 bool) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_isStrict && var1 {
		err = errors.New("preserve.prefixes option and strict option cannot be used together")
		e.m_preserveNS = var2
		return e, err
	}
	e.m_preserveNS = var1
	return e, err
}

func (e *EXIOptions) setInfuseSC(var1 bool /*VARAlignmentType AlignmentType*/) (*EXIOptions, error) { //добавил VARAlignmentType AlignmentType//throws EXIOptionsException
	var err error
	var var2 bool
	if e.m_infuseSC != var1 {
		if var1 {
			if e.m_alignmentType.compress == VARAlignmentType.compress { //опять таже штука с AlignmentType
				err = errors.New("selfContained option and compression option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
			if e.m_alignmentType.preCompress == VARAlignmentType.preCompress { //опять таже штука с AlignmentType
				err = errors.New("selfContained option and pre-compression option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
			if e.m_isStrict {
				err = errors.New("selfContained option and strict option cannot be used together")
				e.m_infuseSC = var2
				return e, err
			}
		}
		e.m_infuseSC = var1
		return e, err
	}
	err = errors.New("metod setInfuseSC: e.m_infuseSC == var1")
	e.m_infuseSC = var2
	return e, err
}

func (e *EXIOptions) setSchemaId(var1 SchemaId) *EXIOptions {
	e.m_schemaId = var1
	return e
}

func (e *EXIOptions) setBlockSize(var1 int) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	if var1 <= 0 {
		err = errors.New("blockSize option value cannot be a negative number")
		e.m_blockSize = 0
		return e, err
	}
	e.m_blockSize = var1
	return e, err
}

func (e *EXIOptions) setValueMaxLength(var1 int) *EXIOptions {
	e.m_valueMaxLength = var1
	return e
}

func (e *EXIOptions) setValuePartitionCapacity(var1 int) *EXIOptions {
	e.m_valuePartitionCapacity = var1
	return e
}

func (e *EXIOptions) setPreserveLexicalValues(var1 bool) *EXIOptions { //было еще throws EXIOptionsException (исключение, но не понятно к чему оно)
	e.m_preserveLexicalValues = var1
	return e
}

func (e *EXIOptions) setDatatypeRepresentationMap(var1 []QName, var2 int) *EXIOptions {
	//var err error
	var var3 int = (2 * var2)

	var var4 int
	for var4 = 0; var4 < var2; var4++ {
		//var var5 int = var4 << 1  //var5 declared but not used
		/*invalid operation: cannot compare var1[var5 + 1] == nil.....
		if var1[var5] == nil || var1[var5+1] == nil {
			err = errors.New("a qname in datatypeRepresentationMap cannot be null")
		}
		*/
	}

	if len(e.m_datatypeRepresentationMap) < var3 {
		var var6 []QName = make([]QName, var3, var3)

		for var4 = 0; var4 < len(e.m_datatypeRepresentationMap); var4++ {
			var6[var4] = e.m_datatypeRepresentationMap[var4]
		}
		/*я считаю, что это не нужно
		for var4; var4<len(var6); var4++{
			var6[var4]=new QName
		}
		*/
		e.m_datatypeRepresentationMap = var6
	}
	for var4 = 0; var4 < 2*var2; var4++ {
		var var7 QName = var1[var4]
		e.m_datatypeRepresentationMap[var4].setValue2(var7.namespaceName, var7.localName, var7.prefix, var7.qName)
	}
	e.m_n_datatypeRepresentationMapBindings = var2
	return e
}

func (e *EXIOptions) appendDatatypeRepresentationMap(var1, var2 EventDescription) (*EXIOptions, error) { //throws EXIOptionsException
	var err error
	if var1 != nil && var2 != nil {
		var var3 int = e.m_n_datatypeRepresentationMapBindings + 1
		var var4 int = 2 * var3
		if len(e.m_datatypeRepresentationMap) < var4 {
			var var6 []QName = make([]QName, var4, var4)
			var var5 int
			for var5 = 0; var5 < len(e.m_datatypeRepresentationMap); var5++ {
				var6[var5] = e.m_datatypeRepresentationMap[var5]
			}
			//я считаю что это лишнее
			//for _; var5 < len(var6); var5++ {
			//	var6[var5] = new QName();
			//}
			e.m_datatypeRepresentationMap = var6
		}
		var var7 int = 2 * e.m_n_datatypeRepresentationMapBindings
		e.m_datatypeRepresentationMap[var7+1].setValue2(var1.getURI(), var1.getName(), "", "") //[var7++]
		e.m_datatypeRepresentationMap[var7+2].setValue2(var2.getURI(), var2.getName(), "", "") //[var7++]

		if var7 != var4 {
			err = errors.New("class EXIOptions, method appendDatatypeRepresentationMap, must -> var7 == var4")
		}
		if err != nil {
			return e, err
		}
		e.m_n_datatypeRepresentationMapBindings = var3
		return e, err
	}
	err = errors.New("A qname in datatypeRepresentationMap cannot be null.")
	return e, err
}

func (e *EXIOptions) toGrammarOptions() int {
	var var1 int = 2
	if e.m_isStrict {
		return 1
	}
	if e.m_preserveComments {
		var1 = VARGrammarOptions.addCM(var1) //было без VAR...
	}

	if e.m_preservePIs {
		var1 = VARGrammarOptions.addPI(var1) //было без VAR...
	}

	if e.m_preserveDTD {
		var1 = VARGrammarOptions.addDTD(var1) //было без VAR...
	}

	if e.m_preserveNS {
		var1 = VARGrammarOptions.addNS(var1) //было без VAR...
	}

	if e.m_infuseSC {
		var1 = VARGrammarOptions.addSC(var1) //было без VAR...
	}

	return var1

}

func (e *EXIOptions) setGrammarOptions(var1 int) { //throws EXIOptionsException
	e.m_isStrict = (var1 == 1) //if (this.m_isStrict = var1 == 1) {
	if e.m_isStrict {
		e.m_preserveComments = false
		e.m_preservePIs = false
		e.m_preserveDTD = false
		e.m_preserveNS = false
		e.m_infuseSC = false
	} else {
		e.m_preserveComments = VARGrammarOptions.hasCM(var1) //было без VAR...
		e.m_preservePIs = VARGrammarOptions.hasPI(var1)      //было без VAR...
		e.m_preserveDTD = VARGrammarOptions.hasDTD(var1)     //было без VAR...
		e.m_preserveNS = VARGrammarOptions.hasNS(var1)       //было без VAR...
		e.setInfuseSC(VARGrammarOptions.hasSC(var1))         //было без VAR...
	}
}

func (e *EXIOptions) getOutline(var1 bool) int {
	var var2 int = 0
	var var3 bool = false
	if e.m_alignmentType.byteAligned == VARAlignmentType.byteAligned || e.m_alignmentType.preCompress == VARAlignmentType.preCompress { //как всю структуру можно сравнить с чем - то конкретным?
		//if e.m_alignmentType == AlignmentType.byteAligned || e.m_alignmentType == AlignmentType.preCompress {!!???!!!???!!!???!!!???!!!???!!!???!!!
		var3 = true
	}

	var var4 bool = (e.m_valuePartitionCapacity != -1)
	var var5 bool = (e.m_valueMaxLength != -1)
	var var6 bool = (e.m_n_datatypeRepresentationMapBindings != 0)
	var var7 bool = false

	if var3 {
		var7 = true
	} else if e.m_infuseSC {
		var7 = true
	} else if var5 {
		var7 = true
	} else if var4 {
		var7 = true
	} else if var6 {
		var7 = true
	}

	var var8 bool = false
	if e.m_preserveComments || e.m_preservePIs || e.m_preserveDTD || e.m_preserveNS || e.m_preserveLexicalValues {
		var8 = true
	}

	var var9 bool = false
	if e.m_blockSize != 1000000 {
		var9 = true
	}

	var var10 bool = false
	if var7 || var8 || var9 {
		var10 = true
	}

	var var11 bool = false
	if e.m_alignmentType.compress == VARAlignmentType.compress {
		//if (e.m_alignmentType == AlignmentType.compress) {!!??!!??!!??!!??!?!?
		var11 = true
	} else if e.m_isFragment {
		var11 = true
	} else if var1 && e.m_schemaId.m_value != "" { //e.m_schemaId != nil
		var11 = true
	}

	if var10 {
		var2 |= 1 //не знаю что это
	}

	if var7 {
		var2 |= 2
		if var3 {
			var2 |= 4
		}

		if var5 {
			var2 |= 32
		}

		if var4 {
			var2 |= 64
		}

		if var6 {
			var2 |= 256
		}
	}

	if var8 {
		var2 |= 8
	}

	if var11 {
		var2 |= 16
		if e.m_isFragment {
			var2 |= 128
		}
	}

	return var2

}

//////GrammarOptions.class/////////////////////////
type GrammarOptions struct {
	RESTRICT_XSI_NIL_TYPE_MASK int
	ADD_UNDECLARED_EA_MASK     int
	ADD_NS                     int
	ADD_SC                     int
	ADD_DTD                    int
	ADD_CM                     int
	ADD_PI                     int
	OPTIONS_UNUSED             int
	DEFAULT_OPTIONS            int
	STRICT_OPTIONS             int
}

var VARGrammarOptions GrammarOptions = GrammarOptions{
	RESTRICT_XSI_NIL_TYPE_MASK: 1,
	ADD_UNDECLARED_EA_MASK:     2,
	ADD_NS:                     4,
	ADD_SC:                     8,
	ADD_DTD:                    16,
	ADD_CM:                     32,
	ADD_PI:                     64,
	OPTIONS_UNUSED:             0,
	DEFAULT_OPTIONS:            2,
	STRICT_OPTIONS:             1,
}

func (g *GrammarOptions) GrammarOptions() {}

func (g *GrammarOptions) restrictXsiNilType(var0 int, var1 bool) int {
	if var1 {
		return (var0 | 1)
	}
	return (var0 & -2)
}

func (g *GrammarOptions) isXsiNilTypeRestricted(var0 int) bool {
	return ((var0 & 1) != 0)
}

func (g *GrammarOptions) isPermitDeviation(var0 int) bool {
	return ((var0 & 2) != 0)
}

func (g *GrammarOptions) hasNS(var0 int) bool {
	return ((var0 & 4) != 0)
}

func (g *GrammarOptions) hasSC(var0 int) bool {
	return ((var0 & 8) != 0)
}

func (g *GrammarOptions) hasDTD(var0 int) bool {
	return ((var0 & 16) != 0)
}

func (g *GrammarOptions) hasCM(var0 int) bool {
	return ((var0 & 32) != 0)
}

func (g *GrammarOptions) hasPI(var0 int) bool {
	return ((var0 & 64) != 0)
}

func (g *GrammarOptions) addNS(var0 int) int {
	return (var0 | 4)
}

func (g *GrammarOptions) addSC(var0 int) int {
	return (var0 | 8)
}

func (g *GrammarOptions) addDTD(var0 int) int {
	return (var0 | 16)
}

func (g *GrammarOptions) addCM(var0 int) int {
	return (var0 | 32)
}

func (g *GrammarOptions) addPI(var0 int) int {
	return (var0 | 64)
}

//////end GrammarOptions.class/////////////////////
///////////////EventDescription.class///////////////////////////////////////////
type EventDescription interface {
	MetodAddVARVAREventDescription() *STRUCTEventDescription //my metod

	getEventKind() byte

	getURI() string

	getName() string

	getPrefix() string

	getURIId() int

	getNameId() int

	getCharacters() Characters

	getBinaryDataSource() BinaryDataSource

	getEventType() EventType
}
type STRUCTEventDescription struct {
	NOT_AN_EVENT int //byte
	EVENT_SD     byte
	EVENT_ED     byte
	EVENT_SE     byte
	EVENT_AT     byte
	EVENT_TP     byte
	EVENT_NL     byte
	EVENT_CH     byte
	EVENT_EE     byte
	EVENT_NS     byte
	EVENT_PI     byte
	EVENT_CM     byte
	EVENT_ER     byte
	EVENT_DTD    byte
	EVENT_BLOB   byte
}

var VAREventDescription STRUCTEventDescription = STRUCTEventDescription{
	NOT_AN_EVENT: -1,
	EVENT_SD:     0,
	EVENT_ED:     1,
	EVENT_SE:     2,
	EVENT_AT:     3,
	EVENT_TP:     4,
	EVENT_NL:     5,
	EVENT_CH:     6,
	EVENT_EE:     7,
	EVENT_NS:     8,
	EVENT_PI:     9,
	EVENT_CM:     10,
	EVENT_ER:     11,
	EVENT_DTD:    12,
	EVENT_BLOB:   13,
}

func (s *STRUCTEventDescription) MetodAddVARVAREventDescription(VAREventDescription STRUCTEventDescription) *STRUCTEventDescription {
	s.NOT_AN_EVENT = VAREventDescription.NOT_AN_EVENT
	s.EVENT_SD = VAREventDescription.EVENT_SD
	s.EVENT_ED = VAREventDescription.EVENT_ED
	s.EVENT_SE = VAREventDescription.EVENT_SE
	s.EVENT_AT = VAREventDescription.EVENT_AT
	s.EVENT_TP = VAREventDescription.EVENT_TP
	s.EVENT_NL = VAREventDescription.EVENT_NL
	s.EVENT_CH = VAREventDescription.EVENT_CH
	s.EVENT_EE = VAREventDescription.EVENT_EE
	s.EVENT_NS = VAREventDescription.EVENT_NS
	s.EVENT_PI = VAREventDescription.EVENT_PI
	s.EVENT_CM = VAREventDescription.EVENT_CM
	s.EVENT_ER = VAREventDescription.EVENT_ER
	s.EVENT_DTD = VAREventDescription.EVENT_DTD
	s.EVENT_BLOB = VAREventDescription.EVENT_BLOB
	return s
}

/////////////Characters.class/////////////////////////////
//////////////end Characters.class////////////////////////
/////////////BinaryDataSource.class///////////////////////
type BinaryDataSource struct {
	m_byteArray        []byte
	m_startIndex       int
	m_length           int
	m_n_remainingBytes int
	m_scanner          IBinaryValueScanner
}

func (b *BinaryDataSource) getByteArray() []byte {
	return b.m_byteArray
}

func (b *BinaryDataSource) getStartIndex() int {
	return b.m_startIndex
}

func (b *BinaryDataSource) getLength() int {
	return b.m_length
}

func (b *BinaryDataSource) getRemainingBytesCount() int {
	return b.m_n_remainingBytes
}

func (b *BinaryDataSource) setValues(var1 []byte, var2, var3 int, var4 IBinaryValueScanner, var5 int) *BinaryDataSource {
	b.m_byteArray = var1
	b.m_startIndex = var2
	b.m_length = var3
	b.m_n_remainingBytes = var5
	b.m_scanner = var4
	return b
}

func (b *BinaryDataSource) hasNext() bool {
	return b.m_n_remainingBytes > 0
}

//todo scan()
func (b *BinaryDataSource) next() int { //throws IOException
	if b.m_n_remainingBytes > 0 {
		b.m_scanner.scan(b.m_n_remainingBytes, b)
		return b.m_length
	}
	return -1
}

/////////IBinaryValueScanner.interface///////////////////////
type IBinaryValueScanner interface {
	//BinaryDataSource scan(long var1, BinaryDataSource var3) throws IOException;
}

/////end IBinaryValueScanner.interface///////////////////////

//////////end BinaryDataSource.class//////////////////////
/*////////////////////////////////////////////////////////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//////////EventType.class/////////////////////////////////

type EventType struct {
	ITEM_PI                      byte
	ITEM_CM                      byte
	ITEM_ER                      byte
	ITEM_CH                      byte
	ITEM_ED                      byte
	ITEM_SE_WC                   byte
	ITEM_SC                      byte
	ITEM_NS                      byte
	ITEM_AT_WC_ANY_UNTYPED       byte
	ITEM_EE                      byte
	ITEM_DTD                     byte
	ITEM_SE                      byte
	ITEM_AT                      byte
	ITEM_SD                      byte
	ITEM_SCHEMA_WC_ANY           byte
	ITEM_SCHEMA_WC_NS            byte
	ITEM_SCHEMA_AT               byte
	ITEM_SCHEMA_AT_WC_ANY        byte
	ITEM_SCHEMA_AT_WC_NS         byte
	ITEM_SCHEMA_CH               byte
	ITEM_SCHEMA_CH_MIXED         byte
	ITEM_SCHEMA_NIL              byte
	ITEM_SCHEMA_TYPE             byte
	ITEM_SCHEMA_AT_INVALID_VALUE byte
	depth                        byte
	uri                          string
	name                         string
	subsequentGrammar            IGrammar
	m_path                       []EventCode
	m_uriId                      int
	m_nameId                     int
	m_eventKind                  byte
	m_index                      int
	m_ownerList                  EventTypeList
}

var VAREventType EventType = EventType{
	ITEM_PI:                      0,
	ITEM_CM:                      1,
	ITEM_ER:                      2,
	ITEM_CH:                      3,
	ITEM_ED:                      4,
	ITEM_SE_WC:                   5,
	ITEM_SC:                      6,
	ITEM_NS:                      7,
	ITEM_AT_WC_ANY_UNTYPED:       8,
	ITEM_EE:                      9,
	ITEM_DTD:                     10,
	ITEM_SE:                      11,
	ITEM_AT:                      12,
	ITEM_SD:                      13,
	ITEM_SCHEMA_WC_ANY:           14,
	ITEM_SCHEMA_WC_NS:            15,
	ITEM_SCHEMA_AT:               16,
	ITEM_SCHEMA_AT_WC_ANY:        17,
	ITEM_SCHEMA_AT_WC_NS:         18,
	ITEM_SCHEMA_CH:               19,
	ITEM_SCHEMA_CH_MIXED:         20,
	ITEM_SCHEMA_NIL:              21,
	ITEM_SCHEMA_TYPE:             22,
	ITEM_SCHEMA_AT_INVALID_VALUE: 23,
}

/*
public EventType(byte var1, EventTypeList var2, byte var3, IGrammar var4) {
	this((String)null, (String)null, -1, -1, var1, var2, var3, (byte)-1, var4);
}*/

/*
public EventType(String var1, String var2, int var3, int var4, byte var5, EventTypeList var6, byte var7, IGrammar var8) {
    this(var1, var2, var3, var4, var5, var6, var7, (byte)-1, var8);
}
*/ /*eventType.class!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (e *EventType) EventType3(var1, var2 string, var3, var4 int, var5 byte, var6 EventTypeList, var7, var8 byte, var9 IGrammar) *EventType {
	VAREventCode.SUPEREventCode(var7) // super(var7)
	e.depth = var5
	e.m_path = make([]EventCode, int(var5), int(var5))
	for i := 0; i < int(var5); i++ {
		e.m_path[i] = VAREventCode
	}
	e.uri = var1
	e.name = var2
	e.m_uriId = var3
	e.m_nameId = var4
	e.m_index = -1
	e.m_ownerList = var6
	e.subsequentGrammar = var9
	e.m_eventKind = var8
	return e
}

func (e *EventType) getDepth() int {
	return int(e.depth) //хотя тип байт...
}

func (e *EventType) getItemPath() []EventCode {
	return e.m_path
}

func (e *EventType) getURIId() int {
	return e.m_uriId
}

func (e *EventType) getNameId() int {
	return e.m_nameId
}

func (e *EventType) getIndex() (int,error){
	var err error
	if e.m_index == -1{
		err = errors.New("EventType.class getIndex - e.m_index == -1")
	}
	/////////////////////////////////////////////////////////////////вернуться позже///////////////////////////////////////////////////////
	return this.m_ownerList.isReverse ? this.m_ownerList.getLength() - (this.m_index + 1) : this.m_index;
 }
///////EventTypeList.class/////
type EventTypeList bool{
	isReverse bool
	//public static final EventTypeList EMPTY = new 1(false);
}
func (e *EventTypeList)MetodEventTypeList(var1 bool) {
	e.isReverse = var1
}
func (e *EventType) getLength() int{}//скорее всего - длинна масива EventCode(ибо масивов рядом нет....)

func item(int var1) EventType{}//or metod?

func getSD() EventType{}//or metod?

func (e *EventType) getED() *EventType{
	var var1 int = e.getLength();

	for(int var2 = 0; var2 < var1; ++var2) {
	   EventType var3 = this.item(var2);
	   if (var3.itemType == 4) {
		  return var3;
	   }
	}

	return null;
 }
////end EventTypeList.class////

///////EventCode.class/////////
type EventCode struct {
	ITEM_TUPLE             int //byte
	EVENT_CODE_DEPTH_ONE   byte
	EVENT_CODE_DEPTH_TWO   byte
	EVENT_CODE_DEPTH_THREE byte
	//EventCode parent = null;
	position int
	itemType byte
}

var VAREventCode EventCode = EventCode{
	ITEM_TUPLE:             -1,
	EVENT_CODE_DEPTH_ONE:   1,
	EVENT_CODE_DEPTH_TWO:   2,
	EVENT_CODE_DEPTH_THREE: 3,
	//parent = null;
	position: -1,
}

func (e *EventCode) SUPEREventCode(var1 byte) {
	e.itemType = var1
}

/* i d`no
   public final void setParentalContext(int var1, EventCode var2) {
      this.position = var1;
      this.parent = var2;
   }
*/
//замена тому что выше
/*eventtype.Class!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (e *EventCode) setParentalContext(var1 int) {
	e.position = var1
}

///////end EventCode.class/////

//////////end EventType.class/////////////////////////////
*/ ////////////////////////////////////////////////////////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
////////////end EventDescription.class//////////////////////////////////////////

///QName.class////////////////////////
type QName struct {
	qName         string
	namespaceName string
	localName     string
	prefix        string
	uriId         int
	localNameId   int
}

func (q *QName) QName1() {}
func (q *QName) QName2(var1, var2 string) {
	q.qName = var1
	var var3 int = strings.Index(q.qName, ":")
	if var3 != -1 {
		var var4 int = len(q.qName)
		var var5 int = var4 - 1
		//label40:
		for var5 > 0 {
			if string(var1[var5]) == string('\t') {
			} else if string(var1[var5]) == string('\n') {
			} else if string(var1[var5]) == string('\r') {
			} else if string(var1[var5]) == string(' ') {
				var5 = (var5 - 1)
				break
			} else {
				break
			}
		}
		q.localName = q.qName[var3 : var5+1]
		var5 = 0

		//label31:
		for var5 < var4 {
			if string(var1[var5]) == string('\t') {
			} else if string(var1[var5]) == string('\n') {
			} else if string(var1[var5]) == string('\r') {
			} else if string(var1[var5]) == string(' ') {
				var5 = (var5 + 1)
				break
			} else {
				break
			}
		}
		q.prefix = q.qName[var5-1 : var3]
		q.namespaceName = var2

	} else {
		q.localName = q.qName
		if var2 != "" {
			q.namespaceName = var2
		} else {
			q.namespaceName = ""
		}
		q.prefix = ""
	}
}

func (q *QName) setValue1(var1, var2, var3 string) *QName {
	var var4 string = ""
	q.qName = var4
	q.namespaceName = var1
	q.localName = var2
	q.prefix = var3
	return q
}

func (q *QName) setValue2(var1, var2, var3, var4 string) *QName { //(*string, *string, *string, *string) {
	q.qName = var4
	q.namespaceName = var1
	q.localName = var2
	q.prefix = var3
	return q //&q.qName, &q.namespaceName, &q.localName, &q.prefix
}

/*нужно подумать нужно ли оно в принципе
public boolean equals(Object var1) {
	if (var1 != null && var1 instanceof QName) {
	   QName var2 = (QName)var1;
	   if (this.namespaceName != null) {
		  return this.namespaceName.equals(var2.namespaceName) && this.localName.equals(var2.localName);
	   }
	}

	return false;
}
*/
func (q *QName) toString() string {
	if q.namespaceName != "" {
		return ("{ " + q.namespaceName + " }" + q.localName)
	}
	return ("")
}
func (q *QName) isSame(var0 []QName, var1 int, var2 []QName, var3 int) bool {
	var var4 int = 2 * var3
	if var0 == nil {
		panic("met isSame QNane struct - var0 is nil")
	}
	if var1 == 0 {
		if var4 == 0 {
			return true
		}
	} else {
		var lenVar0 int = len(var0) //my var
		var lenVar2 int = len(var2) //my var
		var var5 int = 2 * var1

		if var5 == 0 {
			panic("met isSame QNane struct - var5 is 0")
		}
		if var4 != 0 && var4 == var5 {
			var var6 int
			for var6 = 0; var6 < var4; var6++ {
				if lenVar0 == var6 {
					if lenVar2 < var6 {
						break
					}
				} else if (lenVar2 == var6) || (!((var2[var6+1]) == (var0[var6+1]))) {
					break
				}
				/*
					var var7 QName = var0[var6]
					var var8 QName = var2[var6]
				*/ //убрал, ибо не вижу пока смысла от переменных
			}

			if var6 == var4 {
				return true
			}
		}
	}
	return false
}

//////////end QName.class/////////////
////SchemaId.class////////////////////
type SchemaId struct {
	m_value string
}

var VARSchemaId SchemaId = SchemaId{""}

func (s *SchemaId) SchemaId1(var1 string) {
	s.m_value = var1
}

func (s *SchemaId) getValue() string {
	return s.m_value
}

////end SchemaId.class/////////////////

////////////AlignmentType.class////////
type AlignmentType struct {
	bitPacked   string //наверное
	byteAligned string //наверное
	preCompress string //наверное
	compress    string //наверное
}

var VARAlignmentType AlignmentType = AlignmentType{"bitPacked", "byteAligned", "preCompress", "compress"}

////////////end AlignmentType.class//////
///////////////////////////////////////////end EXIOptions.class//////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////Transmogrifier.clas////////////////////////////////////////////////////////////////////////////
type Transmogrifier struct {
	m_xmlReader           *os.File //возможно string
	m_saxHandler          string   // обработка SAX событий(событий XML)
	m_outputOptions       HeaderOptionsOutputType
	m_exiOptions          EXIOptions
	SCHEMAID_NO_SCHEMA    SchemaId
	SCHEMAID_EMPTY_SCHEMA SchemaId
}

var VARTransmogrifier Transmogrifier = Transmogrifier{
	SCHEMAID_NO_SCHEMA:    VARSchemaId,
	SCHEMAID_EMPTY_SCHEMA: VARSchemaId,
}

func Transmogrifier1() { //throws TransmogrifierRuntimeException
	Transmogrifier3(false)
}

func Transmogrifier2(var1 SAXParserFactory) { //throws TransmogrifierRuntimeException
	VARTransmogrifier.Transmogrifier4(var1, false)
}

func Transmogrifier3(var1 bool) { //throws TransmogrifierRuntimeException
	VARTransmogrifier.Transmogrifier4(createSAXParserFactory(), var1)
}

//to do
/*
func (t *Transmogrifier) Transmogrifier4(var1 SAXParserFactory, var2 bool) {
	var err error
	if (!var1.isNamespaceAware()){
		err = errors.New("Transmogrifier class - !var1(SAXParserFactory).isNamespaceAware()==true")
	} else {
		t.m_saxHandler=""//// обработка SAX событий(событий XML)!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	}
	var var3 SAXParser
}
*/

//Transmogrifier(SAXParserFactory var1, boolean var2) throws TransmogrifierRuntimeException {
//	if (!var1.isNamespaceAware()) {
//	   throw new TransmogrifierRuntimeException(3, (String[])null);
//	} else {
//	   this.m_saxHandler = new org.openexi.sax.Transmogrifier.SAXEventHandler(this);
//
//	   try {
//		  SAXParser var3 = var1.newSAXParser();
//		  this.m_xmlReader = var3.getXMLReader();
//		  this.m_xmlReader.setFeature("http://xml.org/sax/features/namespace-prefixes", var2);
//	   } catch (Exception var6) {
//		  throw new TransmogrifierRuntimeException(2, (String[])null);
//	   }
//
//	   this.m_xmlReader.setContentHandler(this.m_saxHandler);
//
//	   try {
//		  this.m_xmlReader.setProperty("http://xml.org/sax/properties/lexical-handler", this.m_saxHandler);
//	   } catch (SAXException var5) {
//		  TransmogrifierRuntimeException var4 = new TransmogrifierRuntimeException(1, new String[]{"http://xml.org/sax/properties/lexical-handler"});
//		  var4.setException(var5);
//		  throw var4;
//	   }
//
//	   this.m_outputOptions = HeaderOptionsOutputType.none;
//	   this.m_exiOptions = new EXIOptions();
//	}
// }

func createSAXParserFactory() SAXParserFactory {
	var var0 SAXParserFactory
	var0 = VARSAXParserFactory
	var0.namespaceAware = true
	return var0
}

///////my SAXParserFactory/////
type SAXParserFactory struct {
	validating     bool
	namespaceAware bool
}

var VARSAXParserFactory SAXParserFactory = SAXParserFactory{
	validating:     false,
	namespaceAware: false,
}

func (s *SAXParserFactory) isNamespaceAware() bool {
	return s.namespaceAware
}

//public boolean isNamespaceAware() {
//	return this.namespaceAware;
// }

//////end my SAXParserFactory//

///////////////////////////////////////end Transmogrifier.class///////////////////////////////////////////////////////////////////////////

//////////////////////////////////////GrammarCache.class//////////////////////////////////////////////////////////////////////////////////
type GrammarCache struct {
	m_schema EXISchema
}

////////////////////EXISchema.class/////////////////////////
type EXISchema struct {
	COOKIE                      []byte
	NIL_NODE                    int
	NIL_VALUE                   int
	EMPTY_STRING                int
	NIL_GRAM                    int
	EVENT_AT_WILDCARD           int
	EVENT_SE_WILDCARD           int
	EVENT_CH_UNTYPED            int
	EVENT_CH_TYPED              int
	MIN_EVENT_ID                int
	EVENT_TYPE_AT               byte
	EVENT_TYPE_SE               byte
	EVENT_TYPE_AT_WILDCARD_NS   byte
	EVENT_TYPE_SE_WILDCARD_NS   byte
	TRUE_VALUE                  int
	FALSE_VALUE                 int
	UNBOUNDED_OCCURS            int
	CONSTRAINT_NONE             int
	CONSTRAINT_DEFAULT          int
	CONSTRAINT_FIXED            int
	WHITESPACE_PRESERVE         int
	WHITESPACE_REPLACE          int
	WHITESPACE_COLLAPSE         int
	VARIANT_STRING              byte
	VARIANT_FLOAT               byte
	VARIANT_DECIMAL             byte
	VARIANT_INTEGER             byte
	VARIANT_INT                 byte
	VARIANT_LONG                byte
	VARIANT_DATETIME            byte
	VARIANT_DURATION            byte
	VARIANT_BASE64              byte
	VARIANT_BOOLEAN             byte
	VARIANT_HEXBIN              byte
	INTEGER_CODEC_DEFAULT       int
	INTEGER_CODEC_NONNEGATIVE   int
	ANCESTRY_IDS                []int //[]byte
	ELEMENT_NAMES               []string
	DEFAULT_TYPABLES            []bool
	UR_SIMPLE_TYPE              byte
	ATOMIC_SIMPLE_TYPE          byte
	LIST_SIMPLE_TYPE            byte
	UNION_SIMPLE_TYPE           byte
	m_elems                     []int
	m_attrs                     []int
	m_types                     []int
	uris                        []string
	localNames                  [][]string
	m_localNames                [][]int
	m_names                     []string
	m_strings                   []string
	m_ints                      []int
	m_mantissas                 []int
	m_exponents                 []int
	m_signs                     []bool
	m_integralDigits            []string
	m_reverseFractionalDigits   []string
	m_integers                  []int //big.int
	m_longs                     []int
	m_datetimes                 []string //XSDateTime[]
	m_durations                 []int64  //[]Duration
	m_binaries                  [][]byte
	m_variantTypes              []byte
	m_variants                  []int
	m_computedDatetimes         []string //XSDateTime[]
	m_variantCharacters         []string //[]Characters
	m_n_stypes                  int
	ancestryIds                 []byte
	m_stypes_end                int
	m_grammars                  []int
	m_productions               []int
	m_eventTypes                []byte
	m_eventData                 []int
	m_grammarCount              int
	m_fragmentINodes            []int
	m_n_fragmentElems           int
	m_globalElementsDirectory   map[string]int
	m_globalAttributesDirectory map[string]int
	m_globalTypesDirectory      map[string]int
	m_buitinTypes               []int
	m_globalElems               []int
	m_globalAttrs               []int
	datatypeFactory             DatatypeFactory
}

var VAREXISchema = EXISchema{
	COOKIE:                    []byte{36, 51, 43, 45},
	NIL_NODE:                  -1,
	NIL_VALUE:                 -1,
	EMPTY_STRING:              0,
	NIL_GRAM:                  -1,
	EVENT_AT_WILDCARD:         -1,
	EVENT_SE_WILDCARD:         -2,
	EVENT_CH_UNTYPED:          -3,
	EVENT_CH_TYPED:            -4,
	MIN_EVENT_ID:              -4,
	EVENT_TYPE_AT:             0,
	EVENT_TYPE_SE:             1,
	EVENT_TYPE_AT_WILDCARD_NS: 2,
	EVENT_TYPE_SE_WILDCARD_NS: 3,
	TRUE_VALUE:                1,
	FALSE_VALUE:               0,
	UNBOUNDED_OCCURS:          -1,
	CONSTRAINT_NONE:           0,
	CONSTRAINT_DEFAULT:        1,
	CONSTRAINT_FIXED:          2,
	WHITESPACE_PRESERVE:       0,
	WHITESPACE_REPLACE:        1,
	WHITESPACE_COLLAPSE:       2,
	VARIANT_STRING:            0,
	VARIANT_FLOAT:             1,
	VARIANT_DECIMAL:           2,
	VARIANT_INTEGER:           3,
	VARIANT_INT:               4,
	VARIANT_LONG:              5,
	VARIANT_DATETIME:          6,
	VARIANT_DURATION:          7,
	VARIANT_BASE64:            8,
	VARIANT_BOOLEAN:           9,
	VARIANT_HEXBIN:            10,
	INTEGER_CODEC_DEFAULT:     255,
	INTEGER_CODEC_NONNEGATIVE: 254,
	ANCESTRY_IDS:              []int{-1, -1, 2, 3, 4, 5, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, -1, -1, 21}, //[]byte
	ELEMENT_NAMES:             []string{"", "", "StringType", "BooleanType", "DecimalType", "FloatType", "FloatType", "DurationType", "DateTimeType", "TimeType", "DateType", "GYearMonthType", "GYearType", "GMonthDayType", "GDayType", "GMonthType", "HexBinaryType", "Base64BinaryType", "AnyURIType", "QNameType", "QNameType", "IntegerType"},
	DEFAULT_TYPABLES:          []bool{true, true, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, false, true, false, true, true, false, true, false, true, true, true, true, false, true, true, false, false, false, false, false, false, false},
	UR_SIMPLE_TYPE:            0,
	ATOMIC_SIMPLE_TYPE:        1,
	LIST_SIMPLE_TYPE:          2,
	UNION_SIMPLE_TYPE:         3,
}

/////////////////end EXISchema.class////////////////////////
////////////////////////////////////end GrammarCache.class////////////////////////////////////////////////////////////////////////////////
func EncodeEXI(sourceFile, destinationFile string) {

}
