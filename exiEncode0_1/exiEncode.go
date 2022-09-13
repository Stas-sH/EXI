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

func (e *EXIOptions) setInfuseSC(var1 bool, VARAlignmentType AlignmentType) (*EXIOptions, error) { //добавил VARAlignmentType AlignmentType//throws EXIOptionsException
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

/*
public void appendDatatypeRepresentationMap(EventDescription var1, EventDescription var2) throws EXIOptionsException {
	if (var1 != null && var2 != null) {
	   int var3 = this.m_n_datatypeRepresentationMapBindings + 1;
	   int var4 = 2 * var3;
	   if (this.m_datatypeRepresentationMap.length < var4) {
		  QName[] var6 = new QName[var4];

		  int var5;
		  for(var5 = 0; var5 < this.m_datatypeRepresentationMap.length; ++var5) {
			 var6[var5] = this.m_datatypeRepresentationMap[var5];
		  }

		  while(var5 < var6.length) {
			 var6[var5] = new QName();
			 ++var5;
		  }

		  this.m_datatypeRepresentationMap = var6;
	   }

	   int var7 = 2 * this.m_n_datatypeRepresentationMapBindings;
	   this.m_datatypeRepresentationMap[var7++].setValue(var1.getURI(), var1.getName(), (String)null, (String)null);
	   this.m_datatypeRepresentationMap[var7++].setValue(var2.getURI(), var2.getName(), (String)null, (String)null);

	   assert var7 == var4;

	   this.m_n_datatypeRepresentationMapBindings = var3;
	} else {
	   throw new EXIOptionsException("A qname in datatypeRepresentationMap cannot be null.");
	}
 }
*/
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
	im_startIndex      int
	m_length           int
	m_n_remainingBytes int
	m_scanner          IBinaryValueScanner
}

func (b *BinaryDataSource) getByteArray() []byte {
	return b.m_byteArray
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
type Transmogrifier struct {
	m_xmlReader     *os.File //возможно string
	m_saxHandler    string   // обработка SAX событий(событий XML)
	m_outputOptions HeaderOptionsOutputType
}

func EncodeEXI(sourceFile, destinationFile string) {

}
